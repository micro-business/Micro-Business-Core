package http

import "net/http"

type HttpFileHandler interface {
	GetCaseSensitiveFileHandler(webDirectoryPath string) func(http.ResponseWriter, *http.Request)

	GetCaseInsensitiveFileHandler(serverMux *http.ServeMux, webDirectoryPath string) (func(http.ResponseWriter, *http.Request), http.Handler)
}
