package main

import (
	"example/hello/conf"
	"example/hello/docs"
	"example/hello/handler"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      go-playground-gllp.onrender.com
//BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
//var db *sql.DB

func main() {
	// programmatically set swagger info
	// docs.SwaggerInfo.Title = "Swagger Example API"
	// docs.SwaggerInfo.Description = "This is a sample server Petstore server."
	// docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = os.Getenv("GO_URL")
	// docs.SwaggerInfo.BasePath = "/api/v1"
	// docs.SwaggerInfo.Schemes = []string{"http", "https"}

	//fmt.Println("Hello, World!")
	//fmt.Println(quote.Go())

	port := os.Getenv("GO_PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.Default()

	//Get a database handle.
	var err error

	db, err := conf.ConnectDB()

	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")

	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "static")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl.html", nil)
	})
	router.GET("/aboutus", func(c *gin.Context) {
		c.HTML(http.StatusOK, "aboutus.tmpl.html", nil)
	})

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	router.GET("/healthcheck", handler.HealthCheckHandler)
	// https://smalldoc124.medium.com/golang-%E0%B9%80%E0%B8%8A%E0%B8%B7%E0%B9%88%E0%B8%AD%E0%B8%A1%E0%B8%81%E0%B8%B1%E0%B8%9A-db-postgres-2c43e9555eeb

	groupApiV1 := router.Group("/api/v1")
	{
		eg := groupApiV1.Group("/example")
		{
			eg.GET("/helloworld", Helloworld)
		}
		// https://medium.com/linedevth/%E0%B8%A3%E0%B8%A7%E0%B8%A1-tips-tricks%E0%B9%83%E0%B8%99%E0%B8%81%E0%B8%B2%E0%B8%A3%E0%B9%83%E0%B8%8A%E0%B9%89-swaggo-%E0%B8%AA%E0%B8%A3%E0%B9%89%E0%B8%B2%E0%B8%87-swagger-ui-%E0%B9%83%E0%B8%AB%E0%B9%89%E0%B8%81%E0%B8%B1%E0%B8%9A-gin-rest-api-76d08985e873

		users := groupApiV1.Group("/users")
		{
			userHandler := handler.UserHandler{
				DB: db,
			}
			users.GET("", userHandler.ListUserHandler)
			users.GET(":id", userHandler.GetUserHandler)
			users.POST("", userHandler.CreateUserHandler)
			users.PATCH(":id", userHandler.UpdateUserHandler)
			users.DELETE(":id", userHandler.DeleteUserHandler)
		}
		albums := groupApiV1.Group("/albums")
		{
			albumHandler := handler.AlbumHandler{
				DB: db,
			}
			albums.GET("", albumHandler.GetAlbums)
			albums.GET(":id", albumHandler.GetAlbumByID)
			albums.POST("", albumHandler.PostAlbums)
		}
	}

	router.Run(":" + port)
}

// @BasePath /api/v1
// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /api/v1/example/helloworld [get]
func Helloworld(g *gin.Context) {
	g.JSON(http.StatusOK, "helloworld")
}

// Simple implementation of an integer minimum
// Adapted from: https://gobyexample.com/testing-and-benchmarking
func IntMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}
