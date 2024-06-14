package handler

// import (
// 	"database/sql"
// 	"fmt"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// type Env struct {
// 	db *sql.DB
// }

// func (e *Env) GetUserHandler(c *gin.Context) {

// 	userId := c.Param("id")
// 	fmt.Print(userId)
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

// func GetUser(id string) (*model.User, error) {
// 	query := "SELECT * FROM users WHERE id = ?"
// 	row := db.QueryRow(query, id)

// 	user := &model.User{}
// 	err := row.Scan(&user.ID, &user.Name, &user.Email)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return user, nil
// }

// func CreateUserHandler(c *gin.Context) {
// 	var newUser model.User
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

// func UpdateUserHandler(c *gin.Context) {
// 	userID := c.Param("id")
// 	var oldUser model.User
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

// func DeleteUserHandler(c *gin.Context) {

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
