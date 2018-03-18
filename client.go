package main

import (
	"github.com/gin-gonic/gin"
	"net/http/httptest"
	"net/http"
	"io"
)

type Client struct {
	router *gin.Engine
}


func (c Client) get(url string) *httptest.ResponseRecorder {
	response := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", url, nil)
	c.router.ServeHTTP(response, req)

	return response
}

func (c Client) post(url string, body io.Reader) *httptest.ResponseRecorder {
	response := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", url, body)
	req.Header.Set("Content-Type", "application/json")
	c.router.ServeHTTP(response, req)

	return response
}

func (c Client) put(url string, body io.Reader) *httptest.ResponseRecorder {
	response := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", url, body)
	req.Header.Set("Content-Type", "application/json")
	c.router.ServeHTTP(response, req)

	return response
}

func (c Client) delete(url string) *httptest.ResponseRecorder {
	response := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", url, nil)
	c.router.ServeHTTP(response, req)

	return response
}


func (c Client) options(url string) *httptest.ResponseRecorder {
	response := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", url, nil)
	c.router.ServeHTTP(response, req)

	return response
}