package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net"
	"strings"

	"nhooyr.io/websocket"
)

func RmMarshal(line string) interface{} {
	cutLine, ok := cutFixes(line)
	if !ok {
		log.Println("LINE MISSING CORRECT PREFIX OR SUFFIX")
		return nil
	}
	data := strings.Split(cutLine, ",")
	var info interface{}
	var err error
	fmt.Println(cutLine)
	switch data[0] {
	case "F":
		info, err = parseHeartbeat(data)
	case "A":
		info, err = parseCompetitorInfo(data)
	case "COMP":
		info, err = parseCompInfo(data)
	case "B":
		info, err = ParseRunInfo(data)
	case "C":
		info, err = paseClassInfo(data)
	case "E":
		info = parseSettingInfo(data)
	case "G":
		info, err = parseRaceInfo(data)
	case "H":
		info, err = ParsePQInfo(data)
	case "I":
		info, err = ParseInitRecord(data)
	case "J":
		info, err = parsePassingInfo(data)
	case "COR":
		info, err = ParseCorrectedFinish(data)
	default:
		info = nil
	}
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return TimingMessage{Type: data[0], Data: info}
}

func main() {
	// Connect to the TCP server
	conn, err := net.Dial("tcp", "localhost:5001")
	if err != nil {
		log.Fatalf("Failed to connect to TCP server: %v", err)
	}
	defer conn.Close()
	fmt.Println("Connected to TCP server")

	// Connect to the WebSocket server
	wsUrl := "ws://css-container-ztob2eeuta-uc.a.run.app/timing/publish/irwindale/1"
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	wsConn, _, err := websocket.Dial(ctx, wsUrl, &websocket.DialOptions{})
	if err != nil {
		log.Fatalf("Failed to connect to WebSocket server: %v", err)
	}
	defer wsConn.Close(websocket.StatusNormalClosure, "closing connection")

	// Read from the TCP connection and send to the WebSocket
	reader := bufio.NewReader(conn)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			log.Printf("Error reading from TCP connection: %v", err)
			break
		}

		err = wsConn.Write(ctx, websocket.MessageText, []byte(line))
		if err != nil {
			log.Printf("Error writing to WebSocket: %v", err)
			break
		}
	}

	log.Println("Connection closed, exiting program")
}
