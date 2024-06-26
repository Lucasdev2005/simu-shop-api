package handler_test

import (
	"net/http"
	"simushop/internal/core"
	"simushop/internal/handler"
	"strconv"
	"testing"
)

var (
	mockRepository  = &mockRepositoryImpl{}
	handlerInstance = handler.NewHandler(mockRepository)
)

func BadRequest(result core.Response, t *testing.T) {
	if result.Code != http.StatusBadRequest {
		t.Error(result.Data)
		t.Errorf("response Code Should be a 400 bad Request but got " + strconv.Itoa(result.Code))
	}
}

func Created(result core.Response, t *testing.T) {
	if result.Code != http.StatusCreated {
		t.Error(result.Data)
		t.Errorf("response Code Should be a 201 Created but got " + strconv.Itoa(result.Code))
	}
}

func Ok(result core.Response, t *testing.T) {
	if result.Code != http.StatusOK {
		t.Error(result.Data)
		t.Errorf("response Code Should be a 200 Created but got " + strconv.Itoa(result.Code))
	}
}
