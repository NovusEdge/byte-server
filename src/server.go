package wbyte

/*
// Author: Aliasgar Khimani (NovusEdge)
// Project: github.com/ARaChn3/web-byte
//
// Copyright: MPL 2.0
// See the LICENSE file for more info.
//
// All Rights Reserved
*/

import (
	"fmt"
	"io"
	"net/http"
)

// ByteServer represents a web-server.
type ByteServer struct {
	// Handlers holds a mapping of respective patterns and functions
	// used by the server.
	Handlers map[string]func(w http.ResponseWriter, _ *http.Request)

	// ServingPort specifies the port on which the server will serve.
	ServingPort uint32

	// Handler specifies the main handler for the server.
	Handler http.Handler
}

// MakeByteServer creates a new ByteServer and returns a pointer to it.
func MakeByteServer(servingPort uint32, serverHandler http.Handler) (*ByteServer, error) {
	InitLogging()

	handlers := make(map[string]func(w http.ResponseWriter, _ *http.Request))
	newBS := ByteServer{
		Handler:     serverHandler,
		ServingPort: servingPort,
		Handlers:    handlers,
	}

	return &newBS, nil
}

// Serve starts the serving process of the server.
func (bs *ByteServer) Serve() error {
	bs.LogInfo(fmt.Sprintf("Starting Server on port %d", bs.ServingPort))

	for pattern, handleFunc := range bs.Handlers {
		http.HandleFunc(pattern, handleFunc)
	}

	err := http.ListenAndServe(fmt.Sprintf(":%d", bs.ServingPort), bs.Handler)
	if err != nil {
		bs.LogError(err)
	}
	return err
}

// AddHandler adds a handler function to ByteServer.
func (bs *ByteServer) AddHandler(pattern string, handleFunc func(w http.ResponseWriter, _ *http.Request)) {
	bs.Handlers[pattern] = handleFunc
}

// DefaultServe specifies the default serving message for the server
func DefaultServe(w http.ResponseWriter, _ *http.Request) {
	io.WriteString(w, BannerArt)
	io.WriteString(w, fmt.Sprintf("%s%s%s\n", BoldColorCyan, DefaultServerWelcomeMessage, ColorReset))
}
