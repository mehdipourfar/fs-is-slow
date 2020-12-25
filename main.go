package main

import (
	"github.com/valyala/fasthttp"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

var image_path, _ = filepath.Abs("./image1.jpg")

func handleRequests(ctx *fasthttp.RequestCtx) {
	path := string(ctx.Path())

	if path == "/1" {
		serveWithReadingFile(ctx)
	} else if path == "/2" {
		fasthttp.ServeFile(ctx, image_path)
	} else {
		ctx.Error("Path not found", 404)
	}

}

func serveWithReadingFile(ctx *fasthttp.RequestCtx) {
	_, err := os.Stat(image_path)
	buf, err := ioutil.ReadFile(image_path)
	if err != nil {
		ctx.Error("File not found", 404)
	}
	ctx.SetBody(buf)
	ctx.SetContentType(http.DetectContentType(buf))
}

func main() {
	log.Println(fasthttp.ListenAndServe("127.0.0.1:8080", handleRequests))
}
