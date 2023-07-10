package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"

	"github.com/fatih/color"
)

func printBanner() {
	red := color.New(color.FgRed).SprintFunc()
	white := color.New(color.FgWhite).SprintFunc()
	fmt.Println("")
	fmt.Println(white("[Core |"), red("Threat]"), white(" PFORWARD"))
	fmt.Println("")

}

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
	printBanner()

	localPort := flag.Int("localPort", 8080, "Lokaler Port zum Lauschen")
	listenIP := flag.String("listenIP", "0.0.0.0", "IP-Adresse zum Lauschen")
	targetIP := flag.String("targetIP", "google.com", "Ziel-IP-Adresse")
	targetPort := flag.Int("targetPort", 443, "Ziel-Port")

	flag.Parse()

	listenAddr := fmt.Sprintf("%s:%d", *listenIP, *localPort)
	targetAddr := fmt.Sprintf("%s:%d", *targetIP, *targetPort)

	listener, err := net.Listen("tcp", listenAddr)
	if err != nil {
		log.Fatalf("Fehler beim Starten des Listeners: %s", err)
	}
	defer listener.Close()

	log.Printf("Portweiterleitung gestartet. Hörer aktiv auf %s. Zieladresse: %s\n", listenAddr, targetAddr)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Fehler bei der Annahme der Verbindung: %s", err)
			continue
		}

		go handleConnection(conn, targetAddr)
	}
}
