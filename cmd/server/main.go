package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gtestaMeLi/GoWebDB/cmd/server/routes"
	_ "github.com/gtestaMeLi/GoWebDB/pkg/web"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error al intentar levantar las variables de entorno")
	}

	db, err := sql.Open("mysql", os.Getenv("DB"))

	r := gin.Default()

	router := routes.NewRouter(r, db)
	router.MapRoutes()

	if err := r.Run(); err != nil {
		panic(err)
	}

}
