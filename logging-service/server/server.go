package server

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	"github.com/revandpratama/reflect/logging-service/adapter"
)

type Server struct {
	Shutdown chan os.Signal
	Done     chan bool

	LogFile *os.File
	Writer  *bufio.Writer
}

func NewServer() *Server {
	return &Server{
		Shutdown: make(chan os.Signal, 1),
		Done:     make(chan bool, 1),
	}
}

func (s *Server) Start() {
	signal.Notify(s.Shutdown, os.Interrupt, syscall.SIGTERM)
	k := adapter.NewKafka()
	k.InitKafka()

	if err := s.initLogFile(); err != nil {
		log.Fatalf("Failed to initialize log file: %v", err)
		s.Done <- true
	}
	defer s.closeLogFile()

	go func() {
		s.ReadMessageFromKafka()
	}()

	select {
	case sh := <-s.Shutdown:
		log.Println("Shutdown cause:", sh)
	case <-s.Done:
		log.Println("Server shutting down due to an error.")
	}

}

func (s *Server) initLogFile() error {
	now := time.Now()
	dir := "./log/"

	// Ensure the directory exists
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	fileName := filepath.Join(dir, fmt.Sprintf("log-%d-%s-%d.log", now.Year(), now.Month(), now.Day()))
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	// Assign file and writer
	s.LogFile = file
	s.Writer = bufio.NewWriter(file)
	return nil
}

func (s *Server) closeLogFile() {
	if s.Writer != nil {
		s.Writer.Flush()
	}
	if s.LogFile != nil {
		s.LogFile.Close()
	}
}

func (s *Server) writeToDocs(msg string) error {
	if _, err := s.Writer.WriteString(msg + "\n"); err != nil {
		return err
	}

	if err := s.Writer.Flush(); err != nil { // Flush after each write
		return err
	}

	return nil
}

type LogEntry struct {
	Service   string    `json:"service"`
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timestamp"`
}

func (s *Server) ReadMessageFromKafka() {
	// fmt.Println("2.1")
	reader := adapter.KafkaReader
	defer reader.Close()

	fmt.Println("waiting for message")
	for {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			break
		}
		// str := fmt.Sprintf("%s-%s-%v", m.Topic, string(m.Value), m.Time)
		logEntry := &LogEntry{
			Service:   string(m.Key),
			Message:   string(m.Value),
			Timestamp: time.Now(),
		}

		jsonMsg, err := json.Marshal(logEntry)
		if err != nil {
			fmt.Println("Error marshalling JSON:", err)
			return
		}
		s.writeToDocs(string(jsonMsg))
		fmt.Printf("Message on %s: %s\n", m.Topic, string(m.Value))
	}

	s.Done <- true
}
