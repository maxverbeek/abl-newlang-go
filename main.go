package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/address/:id", getAddress)

	err := router.Run("0.0.0.0:3000")

	fmt.Printf("errored: %s\n", err.Error())
}

func getAddress(c *gin.Context) {
	id := c.Params.ByName("id")

	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "needs ID"})
		return
	}

	idnum, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "ID needs to be numeric"})
		return
	}

	response, err := http.Get("https://jsonplaceholder.typicode.com/users")

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("failed to make request: %s", err.Error())})
		return
	}

	users, err := parseUsers(response.Body)

	for _, user := range users {
		if user.ID == idnum {
			c.JSON(http.StatusOK, gin.H{"id": user.ID, "address": user.Address.FormatAddr()})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "user ID not found"})
}

func parseUsers(body io.Reader) ([]ForeignUser, error) {
	users := []ForeignUser{}

	err := json.NewDecoder(body).Decode(&users)

	return users, err
}
