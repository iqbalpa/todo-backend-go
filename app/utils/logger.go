package utils

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Log the entire request details
func LogRequest(c *gin.Context, filename string) {
	request := c.Request
	log.Printf("\n===== %s =====", filename)
	log.Printf("Method: %s\n URL: %s\n Headers: %v\n", request.Method, request.URL, request.Header)
	// Log the request body, if any
    body, err := ioutil.ReadAll(c.Request.Body)
    if err != nil {
        log.Println("Error reading request body:", err)
        c.AbortWithStatus(http.StatusInternalServerError)
        return
    }
    log.Println("Request Body:", string(body))
	log.Println("\n===== end log =====")
}