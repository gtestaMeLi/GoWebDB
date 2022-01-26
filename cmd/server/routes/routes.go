package routes

import (
	"database/sql"

	"github.com/gtestaMeLi/GoWebDB/cmd/server/handler"
	"github.com/gtestaMeLi/GoWebDB/internal/product"

	"github.com/gin-gonic/gin"
)

type Router interface {
	MapRoutes()
}

type router struct {
	r  *gin.Engine
	rg *gin.RouterGroup
	db *sql.DB
}

func NewRouter(r *gin.Engine, db *sql.DB) Router {
	return &router{r: r, db: db}
}

func (r *router) MapRoutes() {
	r.setGroup()

	r.buildProductRoutes()
}

func (r *router) setGroup() {
	r.rg = r.r.Group("/api/v1")
}

func (r *router) buildProductRoutes() {
	repo := product.NewRepository(r.db)
	service := product.NewService(repo)
	handler := handler.NewProduct(service)
	groupsAPI := r.rg
	{
		groupsAPI.GET("/products/:name", handler.GetByName())
		groupsAPI.POST("/products/", handler.Store())
	}

}
