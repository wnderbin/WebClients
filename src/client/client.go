package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s <url>", filepath.Base(os.Args[0]))
		return
	}
	url_one, err := url.Parse(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	c := &http.Client{
		Timeout: 15 * time.Second,
	}
	request, err := http.NewRequest(http.MethodGet, url_one.String(), nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	httpData, err := c.Do(request)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Http status:", httpData.Status)
	header, err := httputil.DumpResponse(httpData, false)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(header))
	contentType := httpData.Header.Get("Content-Type")
	charset := strings.SplitAfter(contentType, "charset=")
	if len(charset) > 1 {
		fmt.Println("Charset:", charset[1])
	}
	if httpData.ContentLength == -1 {
		fmt.Println("ContentLength is unknown :(")
	} else {
		fmt.Println("ContentLength:", httpData.ContentLength)
	}

	length := 0
	var buffer [1024]byte
	r := httpData.Body
	for {
		n, err := r.Read(buffer[0:])
		if err != nil {
			fmt.Println(err)
			break
		}
		length += n
	}
	fmt.Println("Response data length:", length)
}
