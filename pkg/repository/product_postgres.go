package repository

import (
	"fmt"
	"os"
	"strings"

	"github.com/danilzign/todo-app"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type ProductPostgres struct {
	db *sqlx.DB
}

func NewProductPostgres(db *sqlx.DB) *ProductPostgres {
	return &ProductPostgres{db: db}
}

func (r *ProductPostgres) Create(product todo.Product) (int, error) {

	var id int
	createProductQuery := fmt.Sprintf("INSERT INTO %s (name, price, amount, imagename) VALUES ($1, $2, $3, $4) RETURNING id", productsTable)
	row := r.db.QueryRow(createProductQuery, product.Name, product.Price, product.Amount, product.ImageName)

	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *ProductPostgres) GetAll() ([]todo.Product, error) {
	var products []todo.Product

	query := fmt.Sprintf("SELECT * FROM %s", productsTable)
	if err := r.db.Select(&products, query); err != nil {
		return nil, err
	}

	return products, nil
}

func (r *ProductPostgres) GetById(productId int) (todo.Product, error) {
	var product todo.Product

	query := fmt.Sprintf(`SELECT * FROM %s WHERE id=$1`, productsTable)
	if err := r.db.Get(&product, query, productId); err != nil {
		return product, err
	}

	return product, nil
}

func (r *ProductPostgres) DeleteProduct(listId int) error {
	stmt, err := r.db.Prepare("SELECT imagename FROM products WHERE id=$1")
	if err != nil {
		return err
	}

	var imageName string

	err = stmt.QueryRow(listId).Scan(&imageName)
	if err != nil {
		return err
	}

	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", productsTable)
	if _, err := r.db.Exec(query, listId); err != nil {
		return err
	}

	directory := fmt.Sprintf("dev/TaskTest/image/product/default/%s", imageName)

	os.Remove(directory)

	return nil
}

func (r *ProductPostgres) UpdateProduct(listId int, input todo.UpdateProductInput) error {
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

	if input.Amount != nil {
		setValues = append(setValues, fmt.Sprintf("amount=$%d", argId))
		args = append(args, *input.Amount)
		argId++
	}

	if input.Amount != nil {
		setValues = append(setValues, fmt.Sprintf("imagename=$%d", argId))
		args = append(args, *input.ImageName)
		argId++
	}
	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s SET %s WHERE id=$%d", productsTable, setQuery, argId)
	args = append(args, listId)

	logrus.Debugf("updateQuery: %s", query)
	logrus.Debugf("args: %s", args)

	_, err := r.db.Exec(query, args...)
	return err

}
