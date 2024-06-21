package handler

import (
	"database/sql"
	"example/hello/model"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// // albums slice to seed record album data.
var albums = []model.Album{
	{Id: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{Id: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{Id: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

type AlbumHandler struct {
	DB *sql.DB
}

func (h AlbumHandler) GetAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// albumsByArtist queries for albums that have the specified artist name.
func (h AlbumHandler) AlbumsByArtist(name string) ([]model.Album, error) {
	// An albums slice to hold data from returned rows.
	var albums []model.Album

	rows, err := h.DB.Query("SELECT * FROM album WHERE artist = ?", name)
	if err != nil {
		fmt.Printf("Eroor: %v\n", err)
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var alb model.Album
		if err := rows.Scan(&alb.Id, &alb.Title, &alb.Artist, &alb.Price); err != nil {
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
func (h AlbumHandler) AlbumByID(id int64) (model.Album, error) {
	// An model.Album to hold data from the returned row.
	var alb model.Album

	row := h.DB.QueryRow("SELECT * FROM model.Album WHERE id = ?", id)
	if err := row.Scan(&alb.Id, &alb.Title, &alb.Artist, &alb.Price); err != nil {
		if err == sql.ErrNoRows {
			return alb, fmt.Errorf("albumsById %d: no such album", id)
		}
		return alb, fmt.Errorf("albumsById %d: %v", id, err)
	}
	return alb, nil
}

// addAlbum adds the specified album to the database,
// returning the album ID of the new entry
func (h AlbumHandler) AddAlbum(alb model.Album) (int64, error) {
	result, err := h.DB.Exec("INSERT INTO album (title, artist, price) VALUES (?, ?, ?)", alb.Title, alb.Artist, alb.Price)
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
func (h AlbumHandler) PostAlbums(c *gin.Context) {
	var newAlbum model.Album
	var albums []model.Album
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
func (h AlbumHandler) GetAlbumByID(c *gin.Context) {
	var albums []model.Album
	id := c.Param("id")

	// Loop through the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range albums {
		if a.Id == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
