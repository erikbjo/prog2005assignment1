package server

import (
	"io"
	"net/http"
	"strings"
)

func DiagHandler(w http.ResponseWriter, r *http.Request) {
	linebreak := "\n"
	output := "Server is running" + linebreak
	output += "Request path: " + r.URL.Path + linebreak
	output += "Request method: " + r.Method + linebreak

	if r.URL.RawQuery != "" {
		output += "Request query(ies)" + linebreak
		for parKey, parVal := range r.URL.Query() {
			output += "- " + parKey + ", Value: " + strings.Join(parVal, "; ") + linebreak
		}
	}

	if r.Header != nil {
		output += "Request headers:" + linebreak
		for headerKey, headerVal := range r.Header {
			output += "- " + headerKey + ", Value: " + strings.Join(headerVal, "; ") + linebreak
		}
	}

	output += "Content type: " + r.Header.Get("Content-Type") + linebreak
	output += "Content length: " + r.Header.Get("Content-Length") + linebreak

	content, err1 := io.ReadAll(r.Body)
	if err1 != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}

	if len(content) > 0 {
		output += "Request body: " + string(content) + linebreak
	}

	_, err2 := w.Write([]byte(output))
	if err2 != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
	}
}
