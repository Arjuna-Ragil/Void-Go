package main

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
	"os"
	"strings"
	"sync"

	"github.com/miekg/dns"
)

var (
		list = make(map[string]bool)
		mu   sync.RWMutex
	)

func LoadList() {

	const listURL = "https://raw.githubusercontent.com/StevenBlack/hosts/master/hosts"

	resp, err := http.Get(listURL); if err != nil{
		fmt.Println("failed to fetch blocklist")
		return
	}
	defer resp.Body.Close()

	newList := make(map[string]bool)
	scanner := bufio.NewScanner(resp.Body)

	for scanner.Scan(){
		line  := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#"){
			continue
		}

		if strings.HasPrefix(line, "0.0.0.0 "){
			parts := strings.Fields(line)
			if len(parts) >= 2{
				domain := parts[1]
				fqdn := domain + "."
				newList[fqdn] = true
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("error reading input: %v", err)
	}

	mu.Lock()
	list = newList
	mu.Unlock()

	fmt.Println("Updated block list")
}

func main() {
	LoadList()

	AutoUpdater()

	addr := net.UDPAddr{
		IP: net.ParseIP("127.0.0.1"),
		Port: 53,
	}
	conn, err := net.ListenUDP("udp", &addr); if err != nil{
		fmt.Println("failed to connect")
		os.Exit(1)
	}
	defer conn.Close()

	fmt.Println("ad-blocker up")

	buffer := make([]byte, 1024)

	for {
		n, clientAddr, err := conn.ReadFromUDP(buffer); if err != nil{
			fmt.Println("failed to get data")
			continue
		}

		msg := new(dns.Msg)

		err = msg.Unpack(buffer[:n]); if err != nil{
			fmt.Println("Failed to unpack: ", err)
			continue
		}

		if len(msg.Question) > 0{
			domain := msg.Question[0].Name
			qtype  := msg.Question[0].Qtype

			reply := new(dns.Msg)
			reply.SetReply(msg)

			mu.RLock()
			isBlocked := list[domain]
			mu.RUnlock()

			if isBlocked{
				if qtype == dns.TypeA{
					answer := fmt.Sprintf("%s 60 IN A 0.0.0.0", domain)
					rr, err := dns.NewRR(answer); if err == nil{
						reply.Answer = append(reply.Answer, rr)
					}
				}
				fmt.Println("[Blocked] ", domain)
			} else {
				c := new(dns.Client)
				upstream, _, err := c.Exchange(msg, "1.1.1.1:53"); if err == nil && upstream != nil{
					reply.Answer = upstream.Answer
				} else{
					fmt.Println("Failed to get upstream response")
				}
			}

			replyBytes, err := reply.Pack(); if err != nil{
				fmt.Println("failed to pack: ", err)
				continue
			}

			_, err = conn.WriteToUDP(replyBytes, clientAddr); if err != nil{
				fmt.Println("failed to send response")
			}
		}
	}
}