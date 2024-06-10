package rpc

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
)

func EncodeMessage(msg any) string{
    content, err := json.Marshal(msg)
    if err != nil{
        panic(err)
    }

    return fmt.Sprintf("Content-Length: %d\r\n\r\n%s", len(content), content)
}

type BaseMesssage struct{
    Method string `json:"method"`
}

func DecodeMessage(msg []byte) (string, int, error){
    header, content, found := bytes.Cut(msg, []byte{'\r', '\n', '\r', '\n'})
    if !found {
        return "", 0, errors.New("Did not find separator")
    }

    contentLengthBytes := header[len("Content-Length: "):]
    contentLength, err := strconv.Atoi(string(contentLengthBytes))
    if err != nil{
        return "", 0, err
    }

    var baseMesssage BaseMesssage
    if err := json.Unmarshal(content[:contentLength], &baseMesssage); err != nil{
        return "", 0, err
    }

    return baseMesssage.Method, contentLength, nil
}