package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
	"tui-portfolio/cmd/server"
)

const (
	host       = "localhost"
	port       = 2222
	sshKeyPath = ".ssh/term_info_ed25519"
)

func main() {
	server, err := server.NewServer(host, port, sshKeyPath)
	if err != nil {
		log.Fatalln(err)
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err = server.ListenAndServe(); err != nil {
			log.Fatalln(err)
		}
	}()

	<-done

	log.Println("Stopping SSH server")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer func() { cancel() }()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalln(err)
	}
}
