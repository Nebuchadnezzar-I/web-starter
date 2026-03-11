package compression

import (
	"compress/gzip"
	"io"
	"net/http"
	"strings"
)

type gzipResponseWriter struct {
	http.ResponseWriter
	writer io.Writer
}

func (g gzipResponseWriter) Write(b []byte) (int, error) {
	return g.writer.Write(b)
}

func shouldCompress(path string, contentType string) bool {
	path = strings.ToLower(path)
	contentType = strings.ToLower(contentType)

	switch {
		case strings.HasSuffix(path, ".png"),
		strings.HasSuffix(path, ".jpg"),
		strings.HasSuffix(path, ".jpeg"),
		strings.HasSuffix(path, ".webp"),
		strings.HasSuffix(path, ".gif"),
		strings.HasSuffix(path, ".mp4"),
		strings.HasSuffix(path, ".webm"),
		strings.HasSuffix(path, ".woff2"),
		strings.HasSuffix(path, ".gz"),
		strings.HasSuffix(path, ".zip"),
		strings.HasSuffix(path, ".pdf"):
		return false
	}

	return true
}

func GzipMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			next.ServeHTTP(w, r)
			return
		}

		if !shouldCompress(r.URL.Path, w.Header().Get("Content-Type")) {
			next.ServeHTTP(w, r)
			return
		}

		w.Header().Add("Vary", "Accept-Encoding")
		w.Header().Set("Content-Encoding", "gzip")
		w.Header().Del("Content-Length")

		gz := gzip.NewWriter(w)
		defer gz.Close()

		gzw := gzipResponseWriter{
			ResponseWriter: w,
			writer:         gz,
		}

		next.ServeHTTP(gzw, r)
	})
}
