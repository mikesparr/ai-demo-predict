package main

import (
	"context"
	"fmt"
	"github.com/mikesparr/ai-demo-predict/cache"
	"github.com/mikesparr/ai-demo-predict/handler"
	"github.com/mikesparr/ai-demo-predict/message"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// server
	addr := ":8080"
	/* #nosec */
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Error occurred: %s", err.Error())
	}

	// pubsub
	projectID, topicID :=
		os.Getenv("PROJECT_ID"),
		os.Getenv("TOPIC_ID")
	producer, err := message.Initialize(projectID, topicID)
	if err != nil {
		log.Fatalf("Could not set up messaging: %v", err)
	}
	defer producer.Topic.Stop()

	// cache
	redisHost, redisPort :=
		os.Getenv("REDISHOST"),
		os.Getenv("REDISPORT")
	redisAddr := fmt.Sprintf("%s:%s", redisHost, redisPort)
	const maxConnections = 100
	client, err := cache.Initialize(redisAddr, maxConnections)
	if err != nil {
		log.Fatalf("Could not set up cache: %v", err)
	}

	// inject cache client and pubsub producer
	httpHandler := handler.NewHandler(client, producer)
	server := &http.Server{
		Handler: httpHandler,
	}
	go func() {
		err := server.Serve(listener)
		if err != nil {
			log.Printf("Error starting the server %v\n", err)
		}
	}()
	defer Stop(server)
	log.Printf("Started server on %s", addr)
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	log.Println(fmt.Sprint(<-ch))
	log.Println("Stopping API server.")
}

// Stop safely shuts down server
func Stop(server *http.Server) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Printf("Could not shut down server correctly: %v\n", err)
		os.Exit(1)
	}
}
