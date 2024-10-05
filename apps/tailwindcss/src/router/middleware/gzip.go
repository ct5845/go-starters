package middleware

import (
	"compress/gzip"
	"io"
	"net/http"
	"strings"
)

// gzipHandler wraps the handler to enable GZIP compression based on Content-Type
func GZip(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Create a response writer that allows us to write GZIP responses
		gz := gzip.NewWriter(w)
		defer gz.Close()

		// Check if the request accepts GZIP encoding
		if strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			// Set the appropriate header
			w.Header().Set("Content-Encoding", "gzip")
			// Use our custom response writer
			w = gzipResponseWriter{ResponseWriter: w, Writer: gz}
		}

		// Call the original handler
		h.ServeHTTP(w, r)
	})
}

// gzipResponseWriter is a custom ResponseWriter that writes GZIP compressed data
type gzipResponseWriter struct {
	http.ResponseWriter
	Writer io.Writer
}

func (g gzipResponseWriter) Write(b []byte) (int, error) {
	return g.Writer.Write(b)
}
