package todo

import "errors"

type Product struct {
	Id        int    `json:"id" db:"id"`
	Name      string `json:"name" db:"name" binding:"required"`
	Price     string `json:"price" db:"price"`
	Amount    string `json:"amount" db:"amount"`
	ImageName string `json:"imagename" db:"imagename"`
}

type UpdateProductInput struct {
	Name      *string `json:"name"`
	Price     *string `json:"price"`
	Amount    *string `json:"amount"`
	ImageName *string `json:"imagename"`
}

func (i UpdateProductInput) Validate() error {
	if i.Name == nil && i.Price == nil && i.Amount == nil && i.ImageName == nil {
		return errors.New("No values")
	}
	return nil
}
