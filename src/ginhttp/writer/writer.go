package writer

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

const DefaultBodySize = 100

type Writer struct {
	gin.ResponseWriter

	body    []byte
	written bool
}

func GinContextUseNewWriter(c *gin.Context) *Writer {
	writer := NewWriter(c.Writer)
	writer.BindGinContext(c)
	return writer
}

func NewWriter(writer gin.ResponseWriter) *Writer {
	return &Writer{
		ResponseWriter: writer,
		body:           make([]byte, 0, DefaultBodySize),
		written:        false,
	}
}

func (w *Writer) BindGinContext(c *gin.Context) {
	c.Writer = w
}

func (w *Writer) Status() int {
	return w.ResponseWriter.Status()
}

func (w *Writer) Write(b []byte) (int, error) {
	w.body = append(w.body, b...)
	return len(b), nil
}

func (w *Writer) WriteString(s string) (int, error) {
	b := []byte(s)
	w.body = append(w.body, b...)
	return len(b), nil
}

func (w *Writer) Size() int {
	return len(w.body)
}

func (w *Writer) Written() bool {
	return w.written
}

func (w *Writer) WriteHeaderNow() {
	w.ResponseWriter.WriteHeaderNow()
}

func (w *Writer) Pusher() http.Pusher {
	return w.ResponseWriter.Pusher()
}

func (w *Writer) WriteToHttp() (int, error) {
	if w.written {
		return len(w.body), nil
	}
	w.written = true
	return w.ResponseWriter.Write(w.body)
}

func (w *Writer) Reset() error {
	if w.written {
		return fmt.Errorf("body has been written")
	}
	w.body = make([]byte, 0, DefaultBodySize)
	return nil
}
