package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/mdp/qrterminal/v3"
)

func main() {
	// Parse CLI flags
	port := flag.Int("port", 8080, "Port to run the server on")
	flag.Parse()

	var input string

	// Check if argument provided, otherwise prompt user
	if len(flag.Args()) > 0 {
		input = flag.Args()[0]
	} else {
		// Interactive Mode
		fmt.Println("Welcome to QuickDrop!")
		fmt.Println("Generate a QR code for a Link or Share a File locally.")
		fmt.Print("Enter a URL or File Path (press Enter to share current folder): ")
		
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			input = scanner.Text()
		}
		input = strings.TrimSpace(input)
		if input == "" {
			input = "." // Default to current directory
		}
	}

	// 1. URL Mode (Just generate QR)
	if strings.HasPrefix(input, "http://") || strings.HasPrefix(input, "https://") || strings.HasPrefix(input, "www.") {
		if strings.HasPrefix(input, "www.") {
			input = "https://" + input
		}
		fmt.Print("\033[H\033[2J") // Clear screen
		fmt.Println("QR Code for URL:", input)
		fmt.Println("Scan to open in browser:")
		generateQR(input)
		return
	}

	// 2. File Server Mode
	fileInfo, err := os.Stat(input)
	if err != nil {
		log.Fatalf("Error: File or path not found: %v\n", err)
	}

	// Get local IP address
	ip, err := getLocalIP()
	if err != nil {
		log.Fatalf("Error getting local IP: %v\n", err)
	}

	// Construct URL
	url := fmt.Sprintf("http://%s:%d", ip, *port)
	if !fileInfo.IsDir() {
		url += "/" + filepath.Base(input)
	}

	// Clear terminal screen
	fmt.Print("\033[H\033[2J")

	// Print Banner
	fmt.Println("QuickDrop Server Started!")
	fmt.Println("Sharing:", input)
	fmt.Println("Local URL:", url)
	fmt.Println("\nScan with your phone to download:")
	
	generateQR(url)

	fmt.Println("\nPress Ctrl+C to stop sharing.")

	// Start HTTP Server
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request from: %s -> %s\n", r.RemoteAddr, r.URL.Path)
		if fileInfo.IsDir() {
			http.FileServer(http.Dir(input)).ServeHTTP(w, r)
		} else {
			http.ServeFile(w, r, input)
		}
	})

	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(*port), nil))
}

func generateQR(content string) {
	config := qrterminal.Config{
		Level:     qrterminal.L,
		Writer:    os.Stdout,
		BlackChar: qrterminal.WHITE,
		WhiteChar: qrterminal.BLACK,
		QuietZone: 1,
	}
	qrterminal.GenerateWithConfig(content, config)
}

// Helper to get local IP address
func getLocalIP() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}
	for _, address := range addrs {
		// Check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String(), nil
			}
		}
	}
	return "", fmt.Errorf("no local IP found")
}
