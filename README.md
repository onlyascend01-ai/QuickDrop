# QuickDrop

QuickDrop is a lightning-fast CLI tool for sharing files over your local network. No cloud, no upload limits, just instant peer-to-peer transfer.

## Features
- **Instant Sharing**: Share files or folders immediately.
- **QR Code**: Automatically generates a QR code in your terminal to scan with your phone.
- **Zero Configuration**: Finds your local IP and starts an HTTP server automatically.
- **Cross-Platform**: Works on Windows, macOS, and Linux.

## Installation

**Prerequisites:** [Go](https://go.dev/dl/) installed.

1.  **Clone the repo:**
    ```bash
    git clone https://github.com/onlyascend01-ai/QuickDrop.git
    cd QuickDrop
    ```

2.  **Build the binary:**
    ```bash
    go build -o quickdrop.exe main.go
    ```

3.  **Run it:**
    ```bash
    ./quickdrop.exe <path-to-file-or-folder>
    ```

## Usage

**Share a specific file:**
```bash
./quickdrop.exe C:\Users\You\Documents\cool-video.mp4
```

**Share the current folder:**
```bash
./quickdrop.exe .
```

**Change the port (default: 8080):**
```bash
./quickdrop.exe -port 9090 .
```

## How it Works
1.  Finds your local IP address (e.g., `192.168.1.50`).
2.  Starts a lightweight HTTP server on the specified port.
3.  Generates a QR code pointing to `http://<YOUR-IP>:<PORT>/<FILE>`.
4.  Your phone scans the code and downloads the file directly from your computer!

## License
MIT License. Created by **Lux**.
