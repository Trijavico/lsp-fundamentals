package main

import (
	"bufio"
	"encoding/json"
	"log"
	"os"

	"github.com/Trijavico/first-lsp/lsp"
	"github.com/Trijavico/first-lsp/rpc"
)

func main() {
	logger := getLogger("/home/javierpp/go/lsp/log.txt")
	logger.Println("Started!!")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(rpc.Split)

	for scanner.Scan() {
		msg := scanner.Bytes()
		method, contents, err := rpc.DecodeMessage(msg)
		if err != nil {
			logger.Printf("Got an error %s\n", err)
		}

		handleMessage(logger, method, contents)
	}
}

func handleMessage(logger *log.Logger, method string, contents []byte) {
	logger.Printf("Received method: %s\n", method)

	switch method {
	case "initialize":
		var request lsp.InitializeRequest
		if err := json.Unmarshal(contents, &request); err != nil {
			log.Printf("Couldnt parse it correctly %s\n", err)
		}

		logger.Printf("Connected to: %s %s\n", request.Params.ClientInfo.Name, request.Params.ClientInfo.Version)

		msg := lsp.NewInitializeResponse(request.ID)
		reply := rpc.EncodeMessage(msg)

		writer := os.Stdout
		_, err := writer.Write([]byte(reply))

		if err != nil {
			logger.Printf("Failed to sent a reply: %v\n", err)
			return
		}

		logger.Println("Sent the reply!!")
	}
}

func getLogger(filename string) *log.Logger {
	logFile, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file %s: %v\n", filename, err)
	}

	return log.New(logFile, "[educationalps]", log.Ldate|log.Ltime|log.Lshortfile)
}
