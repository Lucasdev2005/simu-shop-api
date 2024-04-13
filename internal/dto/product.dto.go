package dto

import "simushop/internal/entity"

type CreateProductDTO struct {
	ProductName             string  `json:"productName" validate:"required"`
	ProductValue            float64 `json:"productValue" validate:"required"`
	ProductCentimeterWidth  float64 `json:"productCentimeterWidth" validate:"required"`
	ProductDescription      string  `json:"productDescription" validate:"required"`
	ProductKgWeitght        float64 `json:"productKgWeitght" validate:"required"`
	ProductCentimeterHeight float64 `json:"productCentimeterHeight" validate:"required"`
	ProductCentimeterLength float64 `json:"productCentimeterLength" validate:"required"`
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
