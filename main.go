package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("start")

	cfg, err := readConfig("config.json")
	if err != nil {
		log.Fatalf("error read config: %v", err)
	}

	r := gin.Default()
	r.GET("/")
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": cfg["name"],
			"time":    time.Now(),
		})
	})
	port := cfg["port"]
	r.Run(fmt.Sprintf(":%v", port))
}

func readConfig(path string) (map[string]interface{}, error) {
	file, err := os.Open(path)
	if err != nil {
		panic("LoadConfig error: " + err.Error())
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	var result map[string]interface{}
	err = decoder.Decode(&result)
	return result, err
}
