package controllers

import (
	domain "task_management/Domain"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserUseCase domain.UserUseCase
}

func (uc *UserController) Register(c *gin.Context) {
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request payload"})
		return
	}

	msg, err := uc.UserUseCase.Register(user)
	if err != nil {
		c.JSON(500, gin.H{"error": err})
		return
	}

	c.JSON(200, gin.H{"message": msg})
}

func (uc *UserController) Login(c *gin.Context) {
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request payload"})
		return
	}

	msg, token, err := uc.UserUseCase.Login(user)
	if err != nil {
		c.JSON(400, gin.H{"error": err})
		return
	}

	c.JSON(200, gin.H{"message": msg, "token": token})
}

func (uc *UserController) UpdateUserRole(c *gin.Context) {
	var req struct {
		Email   string `json:"email"`
		NewRole string `json:"new_role"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request payload"})
		return
	}

	err := uc.UserUseCase.UpdateUserRole(req.Email, req.NewRole)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "User role updated successfully"})
}
