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

    r.StaticFile("/output.css", "./statics/css/output.css")

    fmt.Printf("Starting server at port 8080\n")
    if err := r.Run(":8080"); err != nil {
        log.Fatal(err)
    }
}
