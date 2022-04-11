package cmnhandles

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
	"io"
	"net/http"
)

// MakeBasicWebPage makes a simple web-page provided the pageBody
func MakeBasicWebPage(pageBody string) func(w http.ResponseWriter, _ *http.Request) {
    return func(w http.ResponseWriter, _ *http.Request) {
        io.WriteString(w, pageBody)
    }
}
