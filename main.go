package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("statics/*.html")

	r.GET("/", func(c *gin.Context) {
		files, err := os.ReadDir("projects/")
		if err != nil {
			log.Fatal(err)
		}

		c.HTML(http.StatusOK, "index.html", gin.H{"files": files})
	})

	r.POST("/project", func(c *gin.Context) {
		name := c.PostForm("name")
		fileName := fmt.Sprintf("projects/%s.txt", name)
		myfile, e := os.Create(fileName)
		if e != nil {
			log.Fatal(e)
		}
		log.Println(myfile)
		myfile.Close()

		c.HTML(http.StatusOK, "project.html", gin.H{"name": name})
	})

	r.GET("/project", func(c *gin.Context) {
		c.HTML(http.StatusOK, "newProject.html", gin.H{})
	})

	r.GET("/openproject/:Id", func(c *gin.Context) {
		param := c.Param("Id")
		filenName := fmt.Sprintf("projects/%s.txt", param[2:len(param)-4])
		file, _ := os.Open(filenName)
		defer file.Close()

		scanner := bufio.NewScanner(file)
		var text = ""
		for scanner.Scan() {
			text = text + fmt.Sprintln(scanner.Text())
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}

		c.HTML(http.StatusOK, "openProject.html", gin.H{"text": text})
	})

	r.StaticFile("/output.css", "./statics/css/output.css")

	fmt.Printf("Starting server at port 8089\n")
	if err := r.Run(":8089"); err != nil {
		log.Fatal(err)
	}
}
