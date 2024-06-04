package entities

type Product struct {
	ProductTitle string  `json:"product_title" validate:"required"`
	ProductPrice float64 `json:"product_price" validate:"required"`
	PathToImage  string  `json:"path_to_image" validate:"required,url"`
}
