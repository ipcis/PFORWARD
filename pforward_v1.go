package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
)

func handleConnection(conn net.Conn, targetAddr string) {
	defer conn.Close()

	// Verbindung zum Ziel-Host herstellen
	targetConn, err := net.Dial("tcp", targetAddr)
	if err != nil {
		log.Printf("Fehler bei der Verbindung zum Ziel-Host: %s", err)
		return
	}
	defer targetConn.Close()

	// Daten zwischen den Verbindungen kopieren
	go func() {
		_, err := io.Copy(targetConn, conn)
		if err != nil {
			log.Printf("Fehler beim Kopieren der Daten von Client zu Ziel: %s", err)
		}
	}()

	_, err = io.Copy(conn, targetConn)
	if err != nil {
		log.Printf("Fehler beim Kopieren der Daten von Ziel zu Client: %s", err)
	}
}

func main() {
	localPort := flag.Int("localPort", 8080, "Lokaler Port zum Lauschen")
	targetAddr := flag.String("targetAddr", "google.com:443", "Zieladresse (Host:Port)")

	flag.Parse()

	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *localPort))
	if err != nil {
		log.Fatalf("Fehler beim Starten des Listeners: %s", err)
	}
	defer listener.Close()

	log.Printf("Portweiterleitung gestartet. Hörer aktiv auf localhost:%d. Zieladresse: %s\n", *localPort, *targetAddr)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Fehler bei der Annahme der Verbindung: %s", err)
			continue
		}

		go handleConnection(conn, *targetAddr)
	}
}
