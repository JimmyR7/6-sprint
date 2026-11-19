package handlers

import (
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/JimmyR7/6-sprint/internal/service"
)

func Index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	http.ServeFile(w, r, "index.html")
}

func Upload(logger *log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.Header().Set("Allow", http.MethodPost)
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
		fail := func(err error) {
			logger.Printf("upload: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		if err := r.ParseMultipartForm(10 << 20); err != nil {
			fail(err)
			return
		}
		if r.MultipartForm != nil {
			defer r.MultipartForm.RemoveAll()
		}
		file, header, err := r.FormFile("myFile")
		if err != nil {
			fail(err)
			return
		}
		defer file.Close()
		data, err := io.ReadAll(file)
		if err != nil {
			fail(err)
			return
		}
		result, err := service.Convert(string(data))
		if err != nil {
			fail(err)
			return
		}
		name := "converted-" + strconv.FormatInt(time.Now().UTC().UnixNano(), 10) + filepath.Ext(header.Filename)
		output, err := os.OpenFile(name, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0600)
		if err != nil {
			fail(err)
			return
		}
		if _, err = io.WriteString(output, result); err != nil {
			output.Close()
			os.Remove(name)
			fail(err)
			return
		}
		if err = output.Close(); err != nil {
			os.Remove(name)
			fail(err)
			return
		}
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		_, _ = io.WriteString(w, result)
	}
}
