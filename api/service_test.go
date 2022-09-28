package main

import (
	"net/http"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestPostItem(t *testing.T) {
	// given
	json := "{\"name\":\"test\"}"
	reader := strings.NewReader(json)
	request, err := http.NewRequest("POST", "/items", reader)
	if err != nil {
		t.Fail()
	}

	gc := gin.Context{
		Request: request,
	}

	// when
	PostItem(&gc)
}
