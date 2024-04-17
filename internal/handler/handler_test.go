package handler_test

import (
	"net/http"
	"simushop/internal/core"
	"simushop/internal/entity"
	"simushop/internal/handler"
	"strconv"
	"testing"
)

func productTable() map[string]entity.Product {
	products := map[string]entity.Product{}

	for _, element := range []string{"Lucas Moreira", "produto de teste", "produto de teste 2"} {
		products[element] = entity.Product{
			ProductId:               3,
			ProductName:             element,
			ProductValue:            123,
			ProductDiscountPercent:  1,
			ProductDescription:      "descrição de teste",
			ProductKgWeight:         23,
			ProductCentimeterWidth:  120,
			ProductCentimeterHeight: 11,
			ProductCentimeterLength: 11,
		}
	}

	return products
}

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
