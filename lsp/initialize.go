package lsp

type InitializeRequest struct {
    Request
    Params InitializeRequestParams `json:"params"`
}

type InitializeRequestParams struct {
    ClientInfo *ClientInfo `json:"clientInfo"`

}

type ClientInfo struct {
    Name string `json:"name"`
    Version string `json:"version"`
}

type ServerInfo struct {
    Name string `json:"name"`
    Version string `json:"version"`
}


type InitializeResult struct{
    Capabilities ServerCapabilities `json:"capabilities"`
    ServerInfo ServerInfo `json:"serverInfo"`
}

type InitializeResponse struct {
    Response
    ResponseResult InitializeResult `json:"result"`
}

func NewInitializeResponse(id int) InitializeResponse{
    return InitializeResponse{
    	Response:       Response{
    		RPC: "2.0",
    		ID: &id, 
    	},
    	ResponseResult: InitializeResult{
    		Capabilities: ServerCapabilities{
    			TextDocumentSync: 1,
    		},
    		ServerInfo:   ServerInfo{
    			Name:    "educationalsp",
    			Version: "0.0.0.0.0.0-beta1.final",
    		},
    	},
    } 
}


type ServerCapabilities struct {
    TextDocumentSync int `json:"textDocumentSync"`
}
