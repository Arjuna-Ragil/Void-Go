package main

import (
	"fmt"
	"time"
)

func AutoUpdater() {
	ticker := time.NewTicker(7 * 24 * time.Hour)

	go func(){
		for{
			<-ticker.C
			fmt.Println("Update list")
			LoadList()
		}
	}()
}