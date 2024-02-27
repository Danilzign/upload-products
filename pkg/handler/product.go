package handler

import (
	"net/http"

	test "test"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createProduct(c *gin.Context) {

	var input test.Product
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Product.Create(input)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})

}

type getAllProductsResponse struct {
	Data []test.Product `json:"data"`
}

func (h *Handler) getAllProducts(c *gin.Context) {
	limit := c.Request.URL.Query().Get("limit")
	page := c.Request.URL.Query().Get("page")

	products, err := h.services.Product.GetAll(limit, page)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllProductsResponse{
		Data: products,
	})

}

func (h *Handler) getProductById(c *gin.Context) {
	id := c.Param("id")

	product, err := h.services.Product.GetById(id)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, product)
}

func (h *Handler) updateProduct(c *gin.Context) {
	id := c.Param("id")

	var input test.UpdateProductInput
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.Product.UpdateProduct(id, input); err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}

func (h *Handler) deleteProduct(c *gin.Context) {

	id := c.Param("id")

	_, err := h.services.Product.DeleteProduct(id)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})

	// c.JSON(http.StatusOK, map[string]interface{}{
	// 	"imageName": imageName,
	// })
}
