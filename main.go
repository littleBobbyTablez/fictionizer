package main

import (
    "fmt"
    "log"
    "net/http"

    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()

    r.LoadHTMLGlob("statics/*.html")

    r.GET("/", func(c *gin.Context) {
        c.HTML(http.StatusOK, "index.html", gin.H{})
    })

    r.POST("/project", func(c *gin.Context) {
        name := c.PostForm("name")

        c.HTML(http.StatusOK, "newProject.html", gin.H{ "name": name })
    })

    r.GET("/project", func(c *gin.Context) {
        c.HTML(http.StatusOK, "project.html", gin.H{})
    })

    r.StaticFile("/output.css", "./statics/css/output.css")

    fmt.Printf("Starting server at port 8089\n")
    if err := r.Run(":8089"); err != nil {
        log.Fatal(err)
    }
}
