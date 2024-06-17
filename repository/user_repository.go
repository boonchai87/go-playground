package repository

import (
	"database/sql"
	"example/hello/model"
	"fmt"
)

type UserRepository struct {
	DB *sql.DB
}

func (h UserRepository) GetAllUser() ([]model.User, error) {
	var objects []model.User

	rows, err := h.DB.Query("SELECT id,name,username,email FROM users ")
	if err != nil {
		fmt.Printf("Eroor: %v\n", err)
		return nil, fmt.Errorf("error %q", err)
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var obj model.User
		if err := rows.Scan(&obj.ID, &obj.Name, &obj.Username, &obj.Email); err != nil {
			return nil, fmt.Errorf("error %q", err)
		}
		objects = append(objects, obj)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error %q", err)
	}
	return objects, nil
}
func (h UserRepository) GetUser(id string) (*model.User, error) {
	query := "SELECT * FROM users WHERE id = $1"
	row := h.DB.QueryRow(query, id)

	user := &model.User{}
	err := row.Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (h UserRepository) CreateUser(user model.UserForCreate) (int, error) {

	_, err := h.DB.Exec("INSERT INTO users (name, email,password,username) VALUES ($1, $2,$3,$4)", user.Name, user.Email, user.Password, user.Username)
	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}
	rows, err := h.DB.Query("SELECT * FROM users")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	lastId := 0
	for rows.Next() {
		var alb model.User
		if err := rows.Scan(&alb.ID, &alb.Name, &alb.Email); err != nil {
			panic(err)
		}
		//lastId, err := strconv.Atoi(alb.ID)
		lastId = alb.ID
		if err != nil {
			// ... handle error
			panic(err)
		}
		fmt.Printf("ID: %d", alb.ID)
	}
	//id, err := result.LastInsertId()
	// fmt.Printf("ddddddddddddddd %d", id)
	// if err != nil {
	// 	return 0, err
	// }
	return lastId, nil
}

func (h UserRepository) UpdateUser(user model.UserForUpdate) error {

	query := "UPDATE users SET name = $1, email = $2,password=$3,username=$4 WHERE id = $5"
	_, err := h.DB.Exec(query, user.Name, user.Email, user.Password, user.Username, user.ID)
	if err != nil {
		return err
	}
	return nil
}

func (h UserRepository) DeleteUser(id string) error {

	query := "DELETE FROM users WHERE id = $1"
	_, err := h.DB.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
