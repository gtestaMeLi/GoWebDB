package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gtestaMeLi/GoWebDB/internal/product"
	"github.com/gtestaMeLi/GoWebDB/pkg/web"
)

type Product struct {
	productService product.Service
}

func NewProduct(p product.Service) *Product {
	return &Product{
		productService: p,
	}
}

func (p *Product) GetByName() gin.HandlerFunc {
	return func(c *gin.Context) {
		name := c.Param("name")
		if name == "" {
			web.Error(c, 400, "Error: el campo nombre debe contener al menos 1 caracter")
			return
		}
		pr, err := p.productService.GetByName(c, name)
		if err != nil {
			web.Error(c, 404, err.Error())
			return
		}
		web.Success(c, 200, pr)
	}
}

func (p *Product) Store() gin.HandlerFunc {
	return func(c *gin.Context) {
		type ProductCreate struct {
			ID    *int     `json:"id"`
			Name  *string  `json:"name"`
			Type  *string  `json:"type"`
			Count *int     `json:"count"`
			Price *float64 `json:"price"`
		}

		var req ProductCreate

		if err := c.ShouldBindJSON(&req); err != nil {
			web.Error(c, 422, "%s", err.Error())
			return
		}

		//validamos los valores default de cada paramentro
		if req.Name == nil {
			web.Error(c, 422, "%s", fmt.Errorf("Error: el campo Name es obligatorio"))
			return
		}
		if req.Type == nil {
			web.Error(c, 422, "%s", fmt.Errorf("Error: el campo Type es obligatorio"))
			return
		}
		if req.Count == nil {
			web.Error(c, 422, "%s", fmt.Errorf("Error: el campo Count es obligatorio"))
			return
		}
		if req.Price == nil {
			web.Error(c, 422, "%s", fmt.Errorf("Error: el campo Price es obligatorio"))
			return
		}

		prod, err := p.productService.Store(c, *req.Name, *req.Type, *req.Count, *req.Price)
		if err != nil {
			web.Error(c, 422, "%s", fmt.Errorf("Error: no se pudo hacer el insert"))
			return
		}

		web.Success(c, 200, prod)
	}
}
