package http

import "net/http"

// FileHandler type implements the service that can serve files over HTTP connection
type FileHandler interface {
	GetCaseSensitiveFileHandler(webDirectoryPath string) func(http.ResponseWriter, *http.Request)

	GetCaseInsensitiveFileHandler(serverMux *http.ServeMux, webDirectoryPath string) (func(http.ResponseWriter, *http.Request), http.Handler)
}
