package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/marcos-dev88/first-go-redis/cache"
)

type (
	Test struct {
		Name    string `json:"name"`
		Age     int    `json:"age"`
		Address string `json:"address"`
	}
	Result struct {
		Data interface{} `json:"data"`
	}
)

var redisURI = "redis-cache-poc-go:6079"

func main() {
	router := http.NewServeMux()

	router.HandleFunc("/address/", GetByAddr)
	router.HandleFunc("/create/", Create)
	log.Fatal(http.ListenAndServe(":8075", router))
}

func GetByAddr(rw http.ResponseWriter, req *http.Request) {
	var dataResult Result

	addr := strings.TrimPrefix(req.URL.Path, "/address/")

	redisCall := cache.NewRedis(0, redisURI, "somePass")

	c := cache.NewCache(redisCall)

	data, err := c.GetByAddr(addr)

	if err != nil {
		if err.Error() == "redis: nil" {
			dataResult.Data = ""
		} else {
			log.Fatalf("error: %v", err)
		}
	}

	dataResult.Data = data

	jsonVal, _ := json.Marshal(dataResult)

	rw.WriteHeader(http.StatusOK)
	rw.Header().Add("content-type", "application-json")
	rw.Write(jsonVal)
}

func Create(rw http.ResponseWriter, req *http.Request) {

	var testData Test

	body, err := io.ReadAll(req.Body)

	json.Unmarshal(body, &testData)

	redisCall := cache.NewRedis(0, redisURI, "somePass")

	c := cache.NewCache(redisCall)

	err = c.Create(testData.Address, body, 0)

	if err != nil {
		log.Fatalf("error: %v", err)
	}

	rw.WriteHeader(http.StatusCreated)
	rw.Header().Add("content-type", "application-json")
	rw.Write([]byte("created"))
}
