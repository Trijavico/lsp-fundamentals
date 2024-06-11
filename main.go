package main

import (
	"bufio"
	"log"
	"os"

	"github.com/Trijavico/first-lsp/rpc"
)

func main(){
    logger := getLogger("/home/javierpp/go/lsp/log.txt")
    logger.Println("Started!!")

    scanner := bufio.NewScanner(os.Stdin)
    scanner.Split(rpc.Split)

    for scanner.Scan() {
        msg := scanner.Text()
        handleMessage(msg)
    }
}

func handleMessage(msg string){
}

func getLogger(filename string) *log.Logger{
    logFile, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
    if err != nil{
        log.Fatalf("Failed to open log file %s: %v", filename, err)
    }

    return log.New(logFile, "[educationalps]", log.Ldate|log.Ltime|log.Lshortfile)
}
