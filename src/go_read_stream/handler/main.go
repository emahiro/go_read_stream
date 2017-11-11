package handler

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

var host = "localhost"
var port = "8080"

type RewriteWriter struct {
	reader io.Reader
}

// Readメソッドの再wrap
func (r *RewriteWriter) Read(p []byte) (int, error) {
	buf := make([]byte, len(p))
	n, err := r.reader.Read(buf) // ReaderのinterfaceとしてのReadを呼び出している
	if err != nil && err != io.EOF {
		return n, err
	}

	return copy(p, bytes.Replace(buf, []byte("0"), []byte("1"), -1)), nil
}

func Top(w http.ResponseWriter, r *http.Request) {
	res, err := http.Get(fmt.Sprintf("http://%v:%v/data", host, port))
	if err != nil {
		fmt.Printf("request get error. err: %v", err)
	}
	body := res.Body
	defer body.Close()

	// /data で0000と表示されるところを1111に置き換えている
	// あくまでwriteされているのは、Dataメソッドの中身
	io.Copy(w, &RewriteWriter{body})
}

func Data(w http.ResponseWriter, r *http.Request) {
	var str string
	// 10000行のコードを書き換えていく
	for i := 0; i < 10000; i++ {
		str = str + fmt.Sprintf("%v\n", "000")
	}

	w.Write([]byte(str))
}
