package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
	"tui-portfolio/cmd/server"
)

const (
	defaultHost       = "localhost"
	defaultPort       = 2222
	defaultSSHKeyPath = ".ssh/term_info_ed25519"
	defaultDataPath   = "data/portfolio.json"
)

func main() {
	// Command line flags
	var (
		host     = flag.String("host", defaultHost, "Host to bind the SSH server to")
		port     = flag.Uint("port", defaultPort, "Port to bind the SSH server to")
		dataPath = flag.String("data", defaultDataPath, "Path to portfolio data JSON file")
		help     = flag.Bool("help", false, "Show help message")
	)
	flag.Parse()

	if *help {
		printHelp()
		return
	}

	// Validate data file exists
	if _, err := os.Stat(*dataPath); os.IsNotExist(err) {
		log.Printf("Warning: Data file not found at %s", *dataPath)
		log.Printf("The application will start with fallback content.")
		log.Printf("Create the data file or use -data flag to specify a different path.")
	}

	// Create and start server
	srv, err := server.NewServer(*host, *port, defaultSSHKeyPath, *dataPath)
	if err != nil {
		log.Fatalln("Failed to create server:", err)
	}

	// Setup graceful shutdown
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// Start server in goroutine
	go func() {
		if err = srv.ListenAndServe(); err != nil {
			log.Fatalln("Server failed:", err)
		}
	}()

	log.Printf("Portfolio SSH server running on %s:%d", *host, *port)
	log.Printf("Data source: %s", *dataPath)
	log.Printf("Press Ctrl+C to stop")

	// Wait for shutdown signal
	<-done

	log.Println("Shutting down SSH server...")

	// Create shutdown context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Shutdown server gracefully
	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("Server shutdown error: %v", err)
	} else {
		log.Println("Server shutdown complete")
	}
}

func printHelp() {
	log.Printf(`Portfolio SSH Terminal Server

Usage: %s [options]

Options:
  -host string
        Host to bind the SSH server to (default "%s")
  -port uint
        Port to bind the SSH server to (default %d)
  -data string
        Path to portfolio data JSON file (default "%s")
  -help
        Show this help message

Examples:
  # Start with default settings
  %s

  # Start on different port with custom data file
  %s -port 3333 -data ./my-portfolio.json

  # Start on all interfaces
  %s -host 0.0.0.0

Data File:
  The data file should be a JSON file containing your portfolio information.
  See the included portfolio.json for the expected structure.

Connection:
  Once running, connect with: ssh %s -p %d

Controls (once connected):
  Tab/Shift+Tab  Navigate sections
  ?              Toggle help
  e              Toggle effects
  x              Trigger explosion
  q              Quit
`,
		os.Args[0],
		defaultHost, defaultPort, defaultDataPath,
		os.Args[0],
		os.Args[0],
		os.Args[0],
		defaultHost, defaultPort,
	)
}
