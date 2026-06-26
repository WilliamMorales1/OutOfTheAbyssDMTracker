package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os/exec"
	"strconv"
	"os"
	"strings"

	"github.com/gorilla/websocket"
)

var wsUpgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func handleNotesLSP(w http.ResponseWriter, r *http.Request) {
	conn, err := wsUpgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("lsp ws upgrade:", err)
		return
	}
	defer conn.Close()

	libtorch := os.Getenv("LIBTORCH_PATH")
	if libtorch == "" {
		log.Println("lsp: LIBTORCH_PATH not set - natural-syntax-ls will likely fail to load; set LIBTORCH_PATH to your libtorch installation directory")
	}
	cmd := exec.Command("natural-syntax-ls")
	env := os.Environ()
	if libtorch != "" {
		env = append(env,
			`LIBTORCH=`+libtorch,
			`PATH=`+libtorch+`\lib;`+os.Getenv("PATH"),
		)
	}
	cmd.Env = env
	stdin, err := cmd.StdinPipe()
	if err != nil {
		log.Println("lsp stdin pipe:", err)
		return
	}
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Println("lsp stdout pipe:", err)
		return
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		log.Println("lsp stderr pipe:", err)
		return
	}
	if err := cmd.Start(); err != nil {
		log.Println("lsp start:", err)
		conn.WriteMessage(websocket.TextMessage, []byte(`{"jsonrpc":"2.0","method":"window/showMessage","params":{"type":1,"message":"failed to start natural-syntax-ls"}}`))
		return
	}
	defer cmd.Process.Kill()

	// drain stderr so the process doesn't block; log everything for debugging
	go func() {
		b, _ := io.ReadAll(stderr)
		if len(b) > 0 {
			log.Printf("lsp stderr: %s", b)
		}
	}()
	go func() {
		if err := cmd.Wait(); err != nil {
			log.Println("lsp process exited:", err)
		}
	}()

	done := make(chan struct{})

	// LSP stdout → WebSocket: strip Content-Length framing, forward raw JSON
	go func() {
		defer close(done)
		defer log.Println("lsp stdout reader exited")
		reader := bufio.NewReader(stdout)
		for {
			contentLength := 0
			for {
				line, err := reader.ReadString('\n')
				if err != nil {
					log.Println("lsp stdout read err:", err)
					return
				}
				line = strings.TrimRight(line, "\r\n")
				if line == "" {
					break
				}
				if strings.HasPrefix(line, "Content-Length: ") {
					n, _ := strconv.Atoi(strings.TrimPrefix(line, "Content-Length: "))
					contentLength = n
				}
			}
			if contentLength == 0 {
				continue
			}
			body := make([]byte, contentLength)
			if _, err := io.ReadFull(reader, body); err != nil {
				return
			}
			if err := conn.WriteMessage(websocket.TextMessage, body); err != nil {
				return
			}
		}
	}()

	// WebSocket → LSP stdin: add Content-Length framing
	for {
		select {
		case <-done:
			return
		default:
		}
		_, msg, err := conn.ReadMessage()
		if err != nil {
			return
		}
		header := fmt.Sprintf("Content-Length: %d\r\n\r\n", len(msg))
		if _, err := io.WriteString(stdin, header); err != nil {
			log.Println("lsp stdin write header:", err)
			return
		}
		if _, err := stdin.Write(msg); err != nil {
			log.Println("lsp stdin write body:", err)
			return
		}
	}
}
