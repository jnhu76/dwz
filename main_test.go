package main

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"url/models"

	"github.com/stretchr/testify/assert"
)

func TestHelloWorld(t *testing.T) {

	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "hello world!", w.Body.String())
}

func TestCreateUrl(t *testing.T) {
	setup()

	// originUrl := "https://www.baidu.com"
	// originUrl2 := "https://ghproxy.com"

	body := []byte(`{
		"url": "https://www.baidu.com"
	}`)

	router := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/new", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, 6, len(w.Body.String()))

	shorter := w.Body.String()

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/new", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, 6, len(w.Body.String()))
	assert.Equal(t, shorter, w.Body.String())

	teardown()
}

func TestGetUrl(t *testing.T) {
	setup()

	// url := "https://www.baidu.com"

	router := setupRouter()

	body := []byte(`{
		"url": "https://www.baidu.com"
	}`)

	// add data
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/new", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)
	shorter := w.Body.String()

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/"+shorter, nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 302, w.Code)
	// assert.Equal(t, url, w.Body.String())

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/234dsd", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 404, w.Code)
	assert.Equal(t, "{\"error\":\"Url not found\"}", w.Body.String())

	teardown()
}

func setup() {
	models.ConnectDatabase()
}

func teardown() {
	file := "test.db"
	err := os.Remove(file)

	if err != nil {
		fmt.Println("删除失败")
	} else {
		fmt.Println("删除成功")
	}

}
