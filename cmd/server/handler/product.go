package handler

import (
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
