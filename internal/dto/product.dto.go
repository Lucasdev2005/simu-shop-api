package dto

import "simushop/internal/entity"

type CreateProductDTO struct {
	ProductName             string  `json:"productName" validate:"required"`
	ProductValue            float64 `json:"productValue" validate:"required,min=1"`
	ProductCentimeterWidth  float64 `json:"productCentimeterWidth" validate:"required,min=1"`
	ProductDescription      string  `json:"productDescription" validate:"required,min=1"`
	ProductKgWeitght        float64 `json:"productKgWeitght" validate:"required,min=1"`
	ProductCentimeterHeight float64 `json:"productCentimeterHeight" validate:"required,min=1"`
	ProductCentimeterLength float64 `json:"productCentimeterLength" validate:"required,min=1"`
	ProductDiscountPercent  float64 `json:"productDiscountPercent"`
}

type UpdateProductDTO struct {
	ProductName             string  `json:"productName"`
	ProductValue            float64 `json:"productValue"`
	ProductCentimeterWidth  float64 `json:"productCentimeterWidth" validate:"omitempty,min=1"`
	ProductDescription      string  `json:"productDescription" validate:"omitempty,min=5"`
	ProductKgWeitght        float64 `json:"productKgWeitght" validate:"omitempty,min=1"`
	ProductCentimeterHeight float64 `json:"productCentimeterHeight" validate:"omitempty,min=1"`
	ProductCentimeterLength float64 `json:"productCentimeterLength" validate:"omitempty,min=1"`
	ProductDiscountPercent  float64 `json:"productDiscountPercent"`
}

func (dto CreateProductDTO) ToProduct() entity.Product {
	return entity.Product{
		ProductName:             dto.ProductName,
		ProductValue:            dto.ProductValue,
		ProductCentimeterWidth:  dto.ProductCentimeterWidth,
		ProductDescription:      dto.ProductDescription,
		ProductKgWeight:         dto.ProductKgWeitght,
		ProductCentimeterHeight: dto.ProductCentimeterHeight,
		ProductCentimeterLength: dto.ProductCentimeterLength,
		ProductDiscountPercent:  dto.ProductDiscountPercent,
	}
}

func (dto UpdateProductDTO) ToProduct() entity.Product {
	return entity.Product{
		ProductName:             dto.ProductName,
		ProductValue:            dto.ProductValue,
		ProductCentimeterWidth:  dto.ProductCentimeterWidth,
		ProductDescription:      dto.ProductDescription,
		ProductKgWeight:         dto.ProductKgWeitght,
		ProductCentimeterHeight: dto.ProductCentimeterHeight,
		ProductCentimeterLength: dto.ProductCentimeterLength,
		ProductDiscountPercent:  dto.ProductDiscountPercent,
	}
}
