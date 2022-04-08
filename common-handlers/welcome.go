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
	wb "github.com/NovusEdge/web-byte/src"
	"io"
	"net/http"
)

// Welcome specifies the default serving message for the server
func Welcome(w http.ResponseWriter, _ *http.Request) {
	io.WriteString(w, wb.BannerArt)
	io.WriteString(w, fmt.Sprintf("%s%s%s\n", wb.BoldColorCyan, wb.WelcomeMessage, wb.ColorReset))
}
