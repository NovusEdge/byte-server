package wbyte

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

// Init initializes the server.
func (bs *ByteServer) Init() {
	for pattern, handleFunc := range bs.Handlers {
		http.HandleFunc(pattern, handleFunc)
	}
	
	InitLogging()
}

// Serve starts the serving process of the server
func (bs *ByteServer) Serve(useTLS bool) error {
	var err error
	if useTLS {
		err = http.ListenAndServe(fmt.Sprintf(":%d", bs.ServingPort), bs.Handler)
	}
	// else {  // TODO: Implement TLS styled serving for this server
	// 	err = http.ListenAndServeTLS(fmt.Sprintf(":%d", bs.ServingPort), bs.Handler)
	// }

	if err != nil {
		bs.LogError(err)
	}

	return err
}

// AddHandler adds a handler function to ByteServer.
func (bs *ByteServer) AddHandler(pattern string, handleFunc func(w http.ResponseWriter, _ *http.Request)) {
	bs.Handlers[pattern] = handleFunc
}

// the default serving message for the server
func defaultServe(w http.ResponseWriter, _ *http.Request) {
	io.WriteString(w, DefaultServerWelcomeMessage)
}
