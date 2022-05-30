package controller

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func (h *Handler) UploadImage(ctx *gin.Context) {
	ctx.Request.ParseMultipartForm(10 << 20)

	file, handler, err := ctx.Request.FormFile("file")

	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}

	dst, err := os.Create("images/" + handler.Filename)
	defer dst.Close()
	if err != nil {
		http.Error(ctx.Writer, err.Error(), http.StatusInternalServerError)
		return
	}
	if _, err := io.Copy(dst, file); err != nil {
		http.Error(ctx.Writer, err.Error(), http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, "images/"+handler.Filename)
}
