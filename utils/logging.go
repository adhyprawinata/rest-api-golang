package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

// Middleware untuk mencatat request dan respons ke dalam file log
func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Membuat buffer untuk menyimpan body request
		requestBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			// Jika terjadi kesalahan dalam membaca body request, lanjutkan ke handler selanjutnya
			next.ServeHTTP(w, r)
			return
		}

		// Mengatur ulang body request setelah dibaca agar bisa digunakan di handler selanjutnya
		r.Body = ioutil.NopCloser(bytes.NewBuffer(requestBody))

		// Menyimpan informasi request
		requestInfo := fmt.Sprintf("%s - %s %s\n", time.Now().Format(time.RFC3339), r.Method, r.RequestURI)
		requestInfo += fmt.Sprintf("Request Body: %s\n", string(requestBody))

		// Merekam request ke dalam file log
		file, err := os.OpenFile("request.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			// Jika terjadi kesalahan dalam membuka file, lanjutkan ke handler selanjutnya
			next.ServeHTTP(w, r)
			return
		}
		defer file.Close()

		if _, err := file.WriteString(requestInfo); err != nil {
			// Jika terjadi kesalahan dalam menulis ke file, lanjutkan ke handler selanjutnya
			next.ServeHTTP(w, r)
			return
		}

		// Menangkap respons yang dihasilkan oleh handler selanjutnya
		responseWriter := NewResponseWriter(w)
		next.ServeHTTP(responseWriter, r)

		// Menyimpan informasi respons
		responseInfo := fmt.Sprintf("Response Status: %d\n", responseWriter.StatusCode())
		responseInfo += "Response Body: \n"

		// Merender respons ke dalam format JSON
		if responseWriter.ContentType() == "application/json" {
			var jsonResponse interface{}
			if err := json.Unmarshal(responseWriter.Body(), &jsonResponse); err == nil {
				responseBody, _ := json.MarshalIndent(jsonResponse, "", "  ")
				responseInfo += string(responseBody)
			}
		} else {
			responseInfo += string(responseWriter.Body())
		}

		// Merekam respons ke dalam file log
		if _, err := file.WriteString(responseInfo); err != nil {
			// Jika terjadi kesalahan dalam menulis ke file, lanjutkan ke handler selanjutnya
			next.ServeHTTP(w, r)
			return
		}

		// Mengembalikan respons ke client
		responseWriter.Render()
	})
}

// Implementasi custom ResponseWriter untuk menangkap respons
type customResponseWriter struct {
	http.ResponseWriter
	statusCode   int
	body         []byte
	contentType  string
	wroteHeaders bool
}

func NewResponseWriter(w http.ResponseWriter) *customResponseWriter {
	return &customResponseWriter{
		ResponseWriter: w,
		statusCode:     http.StatusOK,
		contentType:    "text/plain",
	}
}

func (w *customResponseWriter) WriteHeader(statusCode int) {
	w.statusCode = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}

func (w *customResponseWriter) Write(body []byte) (int, error) {
	if !w.wroteHeaders {
		w.WriteHeader(http.StatusOK)
	}
	w.body = body
	return w.ResponseWriter.Write(body)
}

func (w *customResponseWriter) Header() http.Header {
	return w.ResponseWriter.Header()
}

func (w *customResponseWriter) Render() {
	if !w.wroteHeaders {
		w.ResponseWriter.WriteHeader(w.statusCode)
	}
}

func (w *customResponseWriter) StatusCode() int {
	return w.statusCode
}

func (w *customResponseWriter) Body() []byte {
	return w.body
}

func (w *customResponseWriter) ContentType() string {
	return w.contentType
}

func (w *customResponseWriter) WriteHeaderNow() {
	if !w.wroteHeaders {
		w.ResponseWriter.WriteHeader(w.statusCode)
		w.wroteHeaders = true
	}
}

func (w *customResponseWriter) Push(target string, opts *http.PushOptions) error {
	pusher, ok := w.ResponseWriter.(http.Pusher)
	if !ok {
		return http.ErrNotSupported
	}
	return pusher.Push(target, opts)
}
