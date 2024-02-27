package repository

import (
	"fmt"
	"strconv"
	"strings"

	test "test"
	fileName "test/pkg/image"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type ProductPostgres struct {
	db *sqlx.DB
}

func NewProductPostgres(db *sqlx.DB) *ProductPostgres {
	return &ProductPostgres{db: db}
}

func (r *ProductPostgres) Create(product test.Product) (string, error) {

	var id string
	imageName := fileName.RandomFilename()
	createProductQuery := fmt.Sprintf("INSERT INTO %s (name, price, quantity, image_name) VALUES ($1, $2, $3, $4) RETURNING id", productsTable)
	row := r.db.QueryRow(createProductQuery, product.Name, product.Price, product.Quantity, imageName)

	if err := row.Scan(&id); err != nil {
		return "", err
	}

	return id, nil
}

func (r *ProductPostgres) GetAll(limit string, page string) ([]test.Product, error) {
	var products []test.Product

	limitInt, _ := strconv.Atoi(limit)
	pageInt, _ := strconv.Atoi(page)
	offset := limitInt * (pageInt - 1)

	pagination := fmt.Sprintf("SELECT name, price, quantity FROM %s ORDER BY id LIMIT $1 OFFSET $2", productsTable)

	err := r.db.Select(&products, pagination, limit, offset)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (r *ProductPostgres) GetById(productId string) (test.Product, error) {
	var product test.Product

	query := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", productsTable)
	if err := r.db.Get(&product, query, productId); err != nil {
		return product, err
	}

	return product, nil
}

func (r *ProductPostgres) DeleteProduct(productId string) (string, error) {

	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1 RETURNING image_name", productsTable)

	var imageName string

	row := r.db.QueryRow(query, productId)
	_ = row

	err := row.Scan(&imageName)
	_ = err

	return imageName, nil
}

func (r *ProductPostgres) UpdateProduct(productId string, input test.UpdateProductInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Name != nil {
		setValues = append(setValues, fmt.Sprintf("name=$%d", argId))
		args = append(args, *input.Name)
		argId++
	}

	if input.Price != nil {
		setValues = append(setValues, fmt.Sprintf("price=$%d", argId))
		args = append(args, *input.Price)
		argId++
	}

	if input.Quantity != nil {
		setValues = append(setValues, fmt.Sprintf("quantity=$%d", argId))
		args = append(args, *input.Quantity)
		argId++
	}

	if input.ImageName != nil {
		setValues = append(setValues, fmt.Sprintf("image_name=$%d", argId))
		args = append(args, *input.ImageName)
		argId++
	}
	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s SET %s WHERE id=$%d", productsTable, setQuery, argId)
	args = append(args, productId)

	logrus.Debugf("updateQuery: %s", query)
	logrus.Debugf("args: %s", args)

	_, err := r.db.Exec(query, args...)
	return err

}
