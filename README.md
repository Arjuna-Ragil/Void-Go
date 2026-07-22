```

██╗   ██╗ ██████╗ ██╗██████╗        ██████╗  ██████╗ 
██║   ██║██╔═══██╗██║██╔══██╗      ██╔════╝ ██╔═══██╗
██║   ██║██║   ██║██║██║  ██║█████╗██║  ███╗██║   ██║
╚██╗ ██╔╝██║   ██║██║██║  ██║╚════╝██║   ██║██║   ██║
 ╚████╔╝ ╚██████╔╝██║██████╔╝      ╚██████╔╝╚██████╔╝
  ╚═══╝   ╚═════╝ ╚═╝╚═════╝        ╚═════╝  ╚═════╝

```

> A lightning-fast, lightweight, headless DNS sinkhole and ad-blocker written in Go with DNS-over-HTTPS (DoH) support.
#

Void-Go is designed for privacy enthusiasts and developers who need a reliable, zero-bloat network filter. By leveraging the performance of Go and multi-stage Docker builds, this project delivers a fully functional DNS ad-blocker in an ultra-compact container.


## Core Feature
- **Headless Architecture**: No resource-heavy web UI. Configured purely via code and environment files for maximum efficiency.
- **DNS-over-HTTPS (DoH)**: Encrypts your DNS queries to upstream providers (e.g., Cloudflare, Quad9) to prevent ISP snooping.
- **Aggressive Ad & Tracker Blocking**: Routes malicious and telemetry domains to a literal void (0.0.0.0).
- **Minimal Footprint**: Compiled down to a static binary. The entire Docker image is under 10MB.


## Why Void-Go? (vs. Pi-hole)

While Pi-hole is an incredible and feature-rich tool, it comes with a web UI and multiple dependencies (PHP, lighttpd, etc.) which might be overkill for minimalist setups. **Void-Go** is built for people who just want a silent, ultra-lightweight filter without the frontend overhead.

| Feature | Void-Go | Pi-hole |
| :--- | :--- | :--- |
| **Architecture** | Headless / CLI-first | Web UI / Dashboard |
| **Docker Image Size** | **~9 MB** | ~120 MB |
| **Language** | Go (Compiled, Static) | FTL (C), PHP, Bash |
| **Configuration** | Environment / Code (`.env`) | Web Interface / CLI |
| **DoH (DNS over HTTPS)** | **Native (Out-of-the-box)** | Requires external proxy (e.g., `cloudflared`) |
| **Target Audience** | Developers, Minimalists | General users, Data visualizers |


## Deployment Setup
The easiest way to deploy Void-Go is via Docker. The image is officially hosted on the GHCR.

### Docker CLI
```
docker run -d \
  --name void-go \
  -p 53:53/tcp \
  -p 53:53/udp \
  --restart unless-stopped \
  ghcr.io/arjuna-ragil/void-go:v1
```

### Docker Compose
```
services:
  void-go:
    image: ghcr.io/arjuna-ragil/void-go:v1
    container_name: void-go
    ports:
      - "53:53/tcp"
      - "53:53/udp"
    restart: unless-stopped
```

## Thanks You For Using Void-Go :)
