package handlers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserHandler struct {
	DB *gorm.DB
}

func (h *UserHandler) HandleGetUsers(c *gin.Context) {
	var users []User
	if result := h.DB.Find(&users); result.Error != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Failed to retrieve users"})
		return
	}
	c.IndentedJSON(http.StatusOK, users)
}

func (h *UserHandler) HandleGetUsersByID(c *gin.Context) {
	id := c.Param("id")
	var user User

	if result := h.DB.First(&user, id); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "User not found"})
		} else {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Failed to retrieve user"})
		}
		return
	}
	c.IndentedJSON(http.StatusOK, user)
}

func (h *UserHandler) HandleCreateUser(c *gin.Context) {
	var user User

	if err := c.BindJSON(&user); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid user data."})
		return
	}

	if result := h.DB.Create(&user); result.Error != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Failed to create user"})
		return
	}
	c.IndentedJSON(http.StatusOK, user)
}

func (h *UserHandler) HandleDeleteUser(c *gin.Context) {
	var user User
	id := c.Param("id")

	if result := h.DB.First(&user, id); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "User not found"})
		} else {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Failed to retrieve user"})
		}
		return
	}

	if result := h.DB.Delete(&user); result.Error != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Failed to delete user"})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "User deleted"})
}

func (h *UserHandler) HandleUpdateUsers(c *gin.Context) {
	id := c.Param("id")
	var user User

	if result := h.DB.First(&user, id); result.Error != nil {
		c.IndentedJSON(http.StatusBadGateway, gin.H{"message": "Failed to retrieve user"})
		return
	}

	if err := c.BindJSON(&user); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid user data"})
		return
	}

	if result := h.DB.Save(&user); result.Error != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "User updated"})
}
