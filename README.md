# 🕒 clockwall: A Multi-Zone Wall Clock in Go

This project consists of two Go programs:

- **clock server**: A simple time server that listens on a given port and sends the current time every second.
- **clockwall**: A client that connects to multiple `clock` servers and displays their time output side-by-side — similar to a "wall of clocks" seen in international business offices.

## 🔧 Setup

Each `clock` server instance serves the time in the **local system timezone**. To simulate different geographical locations, you can launch multiple instances with different `TZ` environment variables.

### Requirements

- Go 1.18+
- Unix-based shell (Linux/macOS) or compatible terminal on Windows

## 📦 Installation

```bash
git clone https://github.com/yourusername/clockwall.git
cd clockwall
go build -o clock ./clock.go
go build -o clockwall ./clockwall.go

