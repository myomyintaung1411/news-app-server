package controller

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

//  var slice = []string{}
//single file upload
func Upload(c *gin.Context) {
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("file err : %s", err.Error()))
		return
	}
	filename := header.Filename
	out, err := os.Create("public/" + filename)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		log.Fatal(err)
	}
	//filepath := "http://localhost:3000/public/" + filename
	c.JSON(http.StatusOK, gin.H{"filepath": filename})
}

//multiple files upload
func MultipleUpload(c *gin.Context) {
	// Multiple Form
	form, err := c.MultipartForm()
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("err: %s", err.Error()))
		return
	}
	// Files
	files := form.File["files"]
	fmt.Println(form)
	// For range
	var link string
	for _, file := range files {
		path := "images/" + file.Filename

		if err := c.SaveUploadedFile(file, path); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("err: %s", err.Error()))
			return
		}

		// slice = append(slice,path)
		// fmt.Println(len(slice))
		link = link + path + ","
		//filepath := "http://localhost:3000/" + path
		//c.JSON(http.StatusOK, gin.H{"filepath": filepath})
	}
	c.JSON(http.StatusOK, gin.H{"link": link})

}
