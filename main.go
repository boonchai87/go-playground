package main

import (
	"database/sql"
	"example/hello/conf"
	"example/hello/docs"
	"example/hello/handler"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"rsc.io/quote"

	_ "github.com/lib/pq"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// albums slice to seed record album data.
var albums = []Album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

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
var db *sql.DB

type User struct {
	ID    int64
	Name  string
	Email string
}

func main() {
	// programmatically set swagger info
	// docs.SwaggerInfo.Title = "Swagger Example API"
	// docs.SwaggerInfo.Description = "This is a sample server Petstore server."
	// docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = os.Getenv("GO_URL")
	// docs.SwaggerInfo.BasePath = "/api/v1"
	// docs.SwaggerInfo.Schemes = []string{"http", "https"}

	fmt.Println("Hello, World!")
	fmt.Println(quote.Go())

	port := os.Getenv("GO_PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.Default()

	// connect db
	//fmt.Print(os.Getenv("MYSQL_HOST") + "," + os.Getenv("MYSQL_PASS"))
	// var mysql_user = os.Getenv("MYSQL_USER")
	// var mysql_password = os.Getenv("MYSQL_PASSWORD")
	// https://go.dev/doc/tutorial/database-access
	//fmt.Print("xxxx" + mysql_user + "," + mysql_password)
	// cfg := mysql.Config{
	// 	User:   mysql_user,
	// 	Passwd: mysql_password,
	// 	Net:    "tcp",
	// 	Addr:   "127.0.0.1:3306",
	// 	DBName: "mydb",
	// }

	//Get a database handle.
	var err error
	//db, err = sql.Open("mysql", cfg.FormatDSN())

	//postgresql
	//var dataSoruce = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", "localhost", 5432, "postgres", "root", "postgres", "disable")
	//connstring := "user='postgres' dbname='postgres' password='root' host=localhost port=5432 sslmode=disable"
	//fmt.Print(dataSoruce)
	//db, err := sql.Open("postgres", dataSoruce)

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

	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	router.GET("/healthcheck", handler.HealthCheckHandler)
	// https://smalldoc124.medium.com/golang-%E0%B9%80%E0%B8%8A%E0%B8%B7%E0%B9%88%E0%B8%AD%E0%B8%A1%E0%B8%81%E0%B8%B1%E0%B8%9A-db-postgres-2c43e9555eeb

	// userHandler := handler.UserHandler{
	// 	DB: db,
	// }
	// router.POST("/user", userHandler.CreateUserHandler)
	// router.GET("/user/:id", userHandler.GetUserHandler)
	// router.PUT("/user/:id", userHandler.UpdateUserHandler)
	// router.DELETE("/user/:id", userHandler.DeleteUserHandler)

	groupApiV1 := router.Group("/api/v1")
	{
		eg := groupApiV1.Group("/example")
		{
			eg.GET("/helloworld", Helloworld)
		}
		// https://medium.com/linedevth/%E0%B8%A3%E0%B8%A7%E0%B8%A1-tips-tricks%E0%B9%83%E0%B8%99%E0%B8%81%E0%B8%B2%E0%B8%A3%E0%B9%83%E0%B8%8A%E0%B9%89-swaggo-%E0%B8%AA%E0%B8%A3%E0%B9%89%E0%B8%B2%E0%B8%87-swagger-ui-%E0%B9%83%E0%B8%AB%E0%B9%89%E0%B8%81%E0%B8%B1%E0%B8%9A-gin-rest-api-76d08985e873
		// customers := groupApiV1.Group("/customers")
		// {
		// 	customersHandler := handler.CustomerHandler{}
		// 	customers.GET(":id", customersHandler.GetCustomer)
		// 	customers.GET("", customersHandler.ListCustomers)
		// 	customers.POST("", customersHandler.CreateCustomer)
		// 	customers.DELETE(":id", customersHandler.DeleteCustomer)
		// 	customers.PATCH(":id", customersHandler.UpdateCustomer)
		// }
		users := groupApiV1.Group("/users")
		{
			userHandler := handler.UserHandler{
				DB: db,
			}
			users.GET("", userHandler.ListUserHandler)
			users.GET(":id", userHandler.GetUserHandler)
			users.POST("", userHandler.CreateUserHandler)
			users.PUT(":id", userHandler.UpdateUserHandler)
			users.DELETE(":id", userHandler.DeleteUserHandler)
		}
	}
	// mysql
	// const (
	// 	dbDriver = "mysql"
	// 	dbUser   = "root"
	// 	dbPass   = "root"
	// 	dbName   = "mydb"
	// )

	// _, err = db.Exec("INSERT INTO  album(title,artist,price) VALUES($1,$2,$3)", "test", "neng", 35.5)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println("value inserted")
	// }

	// // read
	// rows, err := db.Query("SELECT * FROM album")
	// if err != nil {
	// 	panic(err)
	// }
	// defer rows.Close()
	// lastId := ""
	// for rows.Next() {
	// 	var alb Album
	// 	if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
	// 		panic(err)
	// 	}
	// 	lastId = alb.ID
	// 	fmt.Printf("ID: %d, Name: %s\n", alb.ID, alb.Title)
	// }

	// // update
	// title := "Eddie"
	// _, err = db.Exec("UPDATE album SET title=$1,artist=$2 WHERE id=$3", title, lastId)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println("Data updated")
	// }
	// // delete
	// _, err = db.Exec("DELETE FROM album WHERE id=$1", lastId)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println("Data deleted")
	// }

	// albID, err := addAlbum(Album{
	// 	Title:  "The Modern Sound of Betty Carter",
	// 	Artist: "Betty Carter",
	// 	Price:  49.99,
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("ID of added album: %v\n", albID)

	// albums, err := albumsByArtist("Betty Carter")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("Albums found: %v\n", albums)

	// // Hard-code ID 2 here to test the query.
	// alb, err := albumByID(2)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("Album found: %v\n", alb)

	router.Run(":" + port)
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
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

// albumsByArtist queries for albums that have the specified artist name.
func albumsByArtist(name string) ([]Album, error) {
	// An albums slice to hold data from returned rows.
	var albums []Album

	rows, err := db.Query("SELECT * FROM album WHERE artist = ?", name)
	if err != nil {
		fmt.Printf("Eroor: %v\n", err)
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var alb Album
		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
			return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
		}
		albums = append(albums, alb)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}
	return albums, nil
}

// albumByID queries for the album with the specified ID.
func albumByID(id int64) (Album, error) {
	// An album to hold data from the returned row.
	var alb Album

	row := db.QueryRow("SELECT * FROM album WHERE id = ?", id)
	if err := row.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
		if err == sql.ErrNoRows {
			return alb, fmt.Errorf("albumsById %d: no such album", id)
		}
		return alb, fmt.Errorf("albumsById %d: %v", id, err)
	}
	return alb, nil
}

// addAlbum adds the specified album to the database,
// returning the album ID of the new entry
func addAlbum(alb Album) (int64, error) {
	result, err := db.Exec("INSERT INTO album (title, artist, price) VALUES (?, ?, ?)", alb.Title, alb.Artist, alb.Price)
	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}
	return id, nil
}

// postAlbums adds an album from JSON received in the request body.
func postAlbums(c *gin.Context) {
	var newAlbum Album

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice.
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// Loop through the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

//
// https://www.honeybadger.io/blog/how-to-create-crud-application-with-golang-and-mysql/
// func getUserHandler(c *gin.Context) {

// 	userId := c.Param("id")
// 	fmt.Print(c.Param("id"))
// 	// Convert 'id' to an integer
// 	//userID, err := strconv.Atoi(idStr)

// 	// Call the GetUser function to fetch the user data from the database
// 	user, err := GetUser(userId)
// 	if err != nil {
// 		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "User not found"})
// 		return
// 	}
// 	// Convert the user object to JSON and send it in the response
// 	c.IndentedJSON(http.StatusOK, user)
// }
// func GetUser(id string) (*User, error) {
// 	query := "SELECT * FROM users WHERE id = ?"
// 	row := db.QueryRow(query, id)

// 	user := &User{}
// 	err := row.Scan(&user.ID, &user.Name, &user.Email)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return user, nil
// }

// func createUserHandler(c *gin.Context) {
// 	var newUser User
// 	var err error
// 	// Call BindJSON to bind the received JSON to
// 	// newAlbum.
// 	if err := c.BindJSON(&newUser); err != nil {
// 		return
// 	}

// 	// insert into database
// 	id, err := CreateUser(newUser.Name, newUser.Email)
// 	if err != nil {
// 		c.IndentedJSON(http.StatusInternalServerError, "failed to create user")
// 		return
// 	}
// 	newUser.ID = id
// 	c.IndentedJSON(http.StatusCreated, newUser)
// }
// func CreateUser(name, email string) (int64, error) {
// 	query := "INSERT INTO users (name, email) VALUES (?, ?)"
// 	result, err := db.Exec(query, name, email)
// 	if err != nil {
// 		return 0, fmt.Errorf("addAlbum: %v", err)
// 	}
// 	id, err := result.LastInsertId()
// 	if err != nil {
// 		return 0, err
// 	}
// 	return id, nil
// }

// func updateUserHandler(c *gin.Context) {
// 	userID := c.Param("id")
// 	var oldUser User
// 	//var err error
// 	if err := c.BindJSON(&oldUser); err != nil {
// 		return
// 	}
// 	err := UpdateUser(userID, oldUser.Name, oldUser.Email)
// 	if err != nil {
// 		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "User not found"})
// 		return
// 	}
// 	c.IndentedJSON(http.StatusOK, "User updated successfully")
// }

// func UpdateUser(id string, name, email string) error {
// 	query := "UPDATE users SET name = ?, email = ? WHERE id = ?"
// 	_, err := db.Exec(query, name, email, id)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// func deleteUserHandler(c *gin.Context) {

// 	userId := c.Param("id")

// 	err := DeleteUser(userId)
// 	if err != nil {
// 		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "User not found"})
// 		return
// 	}
// 	c.IndentedJSON(http.StatusNoContent, "User deleted  successfully")
// }

// func DeleteUser(id string) error {
// 	query := "DELETE FROM users WHERE id = ?"
// 	_, err := db.Exec(query, id)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
