package todo

import "errors"

type Product struct {
	Id        string `json:"id" db:"id"`
	Name      string `json:"name" db:"name" binding:"required"`
	Price     string `json:"price" db:"price"`
	Quantity  string `json:"quantity" db:"quantity"`
	ImageName string `json:"image_name" db:"image_name"`
}

type UpdateProductInput struct {
	Name      *string `json:"name"`
	Price     *string `json:"price"`
	Quantity  *string `json:"quantity"`
	ImageName *string `json:"imagename"`
}

func (i UpdateProductInput) Validate() error {
	if i.Name == nil && i.Price == nil && i.Quantity == nil && i.ImageName == nil {
		return errors.New("No values")
	}
	return nil
}
