package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	r := gin.Default()
	MainRouter(r)
	r.Run(fmt.Sprintf("%s:%d", "0.0.0.0", GetPort()))
}

func MainRouter(r *gin.Engine) {
	r.POST("/check_key", checkKey)
}

func checkKey(c *gin.Context) {
	_body, err := c.GetRawData()
	if err != nil {
		log.Println(err)
		c.JSON(500, nil)
		return
	}

	var body map[string]interface{}
	if err := json.Unmarshal(_body, &body); err != nil {
		log.Println(err)
		c.JSON(500, nil)
		return
	}
	bodyKey, ok := body["key"]
	if !ok {
		c.JSON(401, nil)
		return
	}

	keys := GetKeys()
	for _, key := range keys {
		hashedKey := getKeyHash(key)
		if hashedKey == bodyKey.(string) {
			c.JSON(200, nil)
			return
		}
	}
	c.JSON(401, nil)
}

func getKeyHash(key string) string {
	h := sha256.New()
	h.Write([]byte(key))
	bs := h.Sum(nil)
	return hex.EncodeToString(bs[:])
}