package handler

import (
	"database/sql"
	"example/hello/model"
	"example/hello/repository"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	DB *sql.DB
}

// @summary Get User
// @description  Get User by id
// @tags users
// @security ApiKeyAuth
// @id GetUser
// @accept json
// @produce json
// @param id path int true "id of customer to be gotten"
// @response 200 {object} model.User "OK"
// @response 400 {object} model.Response "Bad Request"
// @response 401 {object} model.Response "Unauthorized"
// @response 409 {object} model.Response "Conflict"
// @response 500 {object} model.Response "Internal Server Error"
// @Router /api/v1/users/:id [get]
func (h UserHandler) GetUserHandler(c *gin.Context) {

	userId := c.Param("id")

	// Call the GetUser function to fetch the user data from the database
	userRepository := repository.UserRepository{
		DB: h.DB,
	}
	user, err := userRepository.GetUser(userId)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "User not found"})
		return
	}
	// Convert the user object to JSON and send it in the response
	c.IndentedJSON(http.StatusOK, user)
}

// GetCustomer godoc
// @summary List User
// @description  Get all user
// @tags users
// @security ApiKeyAuth
// @id ListUsers
// @accept json
// @produce json
// @response 200 {object} model.Users "OK"
// @response 400 {object} model.Response "Bad Request"
// @response 401 {object} model.Response "Unauthorized"
// @response 409 {object} model.Response "Conflict"
// @response 500 {object} model.Response "Internal Server Error"
// @Router /api/v1/customers/:id [get]
func (h UserHandler) ListUserHandler(c *gin.Context) {

	// Call the GetUser function to fetch the user data from the database
	userRepository := repository.UserRepository{
		DB: h.DB,
	}
	users, err := userRepository.GetAllUser()
	fmt.Print(err)
	//fmt.Print(users)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "User not found"})
		return
	}
	// Convert the user object to JSON and send it in the response
	c.IndentedJSON(http.StatusOK, users)
}

// CreateUser godoc
// @summary Create User
// @description Create new user
// @tags users
// @security ApiKeyAuth
// @id CreateUser
// @accept json
// @produce json
// @param Customer body model.UserForCreate true "User data to be created"
// @response 200 {object} model.Response "OK"
// @response 400 {object} model.Response "Bad Request"
// @response 401 {object} model.Response "Unauthorized"
// @response 500 {object} model.Response "Internal Server Error"
// @Router /api/v1/users [post]
func (h UserHandler) CreateUserHandler(c *gin.Context) {
	var newUser model.User
	var err error
	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newUser); err != nil {
		return
	}

	// insert into database
	userRepository := repository.UserRepository{
		DB: h.DB,
	}
	id, err := userRepository.CreateUser(newUser.Name, newUser.Email)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, "failed to create userxxxx")
		return
	}
	newUser.ID = id
	c.IndentedJSON(http.StatusCreated, newUser)
}

// DeleteCustomer godoc
// @summary Delete User
// @description Delete user by id
// @tags users
// @security ApiKeyAuth
// @id DeleteUser
// @accept json
// @produce json
// @param id path int true "id of user to be deleted"
// @response 200 {object} model.Response "OK"
// @response 400 {object} model.Response "Bad Request"
// @response 401 {object} model.Response "Unauthorized"
// @response 500 {object} model.Response "Internal Server Error"
// @Router /api/v1/users/:id [delete]
func (h UserHandler) DeleteUserHandler(c *gin.Context) {

	userId := c.Param("id")
	userRepository := repository.UserRepository{
		DB: h.DB,
	}
	err := userRepository.DeleteUser(userId)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "User not found"})
		return
	}
	c.IndentedJSON(http.StatusNoContent, "User deleted  successfully")
}

// UpdateCustomer godoc
// @summary Update Customer
// @description Update customer by id
// @tags users
// @security ApiKeyAuth
// @id UpdateCustomer
// @accept json
// @produce json
// @param id path int true "id of user to be updated"
// @param Customer body model.UserForUpdate true "User data to be updated"
// @response 200 {object} model.Response "OK"
// @response 400 {object} model.Response "Bad Request"
// @response 401 {object} model.Response "Unauthorized"
// @response 500 {object} model.Response "Internal Server Error"
// @Router /api/v1/users/:id [post]
func (h UserHandler) UpdateUserHandler(c *gin.Context) {
	userID := c.Param("id")
	var oldUser model.User
	//var err error
	if err := c.BindJSON(&oldUser); err != nil {
		return
	}
	// insert into database
	userRepository := repository.UserRepository{
		DB: h.DB,
	}
	err := userRepository.UpdateUser(userID, oldUser.Name, oldUser.Email)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "User not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, "User updated successfully")
}
