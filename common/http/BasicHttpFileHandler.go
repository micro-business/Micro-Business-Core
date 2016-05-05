package http

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/microbusinesses/Micro-Businesses-Core/common/diagnostics"
)

type BasicHttpFileHandler struct {
}

func (basicHttpFileHandler BasicHttpFileHandler) GetCaseSensitiveFileHandler(webDirectoryPath string) func(http.ResponseWriter, *http.Request) {
	diagnostics.IsNotNilOrEmptyOrWhitespace(webDirectoryPath, "webDirectoryPath", "webDirectoryPath cannot be nil, empty or contains whitespace only.")

	return getFileHandler(webDirectoryPath)
}

func (basicHttpFileHandler BasicHttpFileHandler) GetCaseInsensitiveFileHandler(serverMux *http.ServeMux, webDirectoryPath string) (func(http.ResponseWriter, *http.Request), http.Handler) {
	diagnostics.IsNotNil(serverMux, "serverMux", "serverMux cannot be nil.")
	diagnostics.IsNotNilOrEmptyOrWhitespace(webDirectoryPath, "webDirectoryPath", "webDirectoryPath cannot be nil, empty or contains whitespace only.")

	caseInsensitiveFilehandler := http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		request.URL.Path = strings.ToLower(request.URL.Path)

		serverMux.ServeHTTP(writer, request)
	})

	return getFileHandler(webDirectoryPath), caseInsensitiveFilehandler
}

func getFileHandler(webDirectoryPath string) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Access-Control-Allow-Origin", "*")
		writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")
		writer.Header().Set("Access-Control-Allow-Methods", "GET, POST")

		filePath := webDirectoryPath + request.URL.Path

		if request.URL.Path == "/" || len(request.URL.Path) == 0 {
			if doesFileExist(webDirectoryPath + "/index.html") {
				filePath = webDirectoryPath + "/index.html"
			} else if doesFileExist(webDirectoryPath + "/index.htm") {
				filePath = webDirectoryPath + "/index.htm"
			}
		}

		if doesFileExist(filePath) {
			if fileContent, err := loadPage(filePath); err != nil {
				http.Error(writer, err.Error(), http.StatusInternalServerError)
			} else {
				switch strings.ToLower(filepath.Ext(filePath)) {
				case ".css":
					writer.Header().Set("Content-Type", "text/css; charset=utf-8")
				case ".js", ".jsx":
					writer.Header().Set("Content-Type", "application/javascript; charset=utf-8")
				case ".html", ".htm":
					writer.Header().Set("Content-Type", "text/html; charset=utf-8")
				default:
					writer.Header().Set("Content-Type", "text/plain; charset=utf-8")
				}

				fmt.Fprintln(writer, fileContent)
			}
		} else {
			http.NotFound(writer, request)
		}
	}
}

func doesDirectoryExist(directoryPath string) bool {
	if _, err := os.Stat(directoryPath); os.IsNotExist(err) {
		return false
	} else {
		return true
	}
}

func doesFileExist(filePath string) bool {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return false
	} else {
		return true
	}
}

func loadPage(filePath string) (string, error) {
	if content, err := ioutil.ReadFile(filePath); err != nil {
		return "", err
	} else {
		return string(content[:len(content)]), nil
	}
}
