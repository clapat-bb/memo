package controller

import (
	"net/http"

	"github.com/clapat-bb/memo/logger"
	"github.com/clapat-bb/memo/model"
	"github.com/clapat-bb/memo/util"
	"github.com/gin-gonic/gin"
)

type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Register(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "parameter error"})
		return
	}

	var user model.User
	if err := model.DB.Where("username = ?", req.Username).First(&user).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user already exist"})
		return
	}

	hashed, _ := util.HashPassword(req.Password)

	user = model.User{
		Username: req.Username,
		Password: hashed,
	}
	if err := model.DB.Create(&user).Error; err != nil {
		logger.Log.Errorf("sign up failed: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Sign up failed!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Sign up success!"})
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "parameter error"})
		return
	}

	// 查找用户
	var user model.User
	if err := model.DB.Where("username = ?", req.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "usernaem or password error"})
		return
	}

	// 校验密码
	if !util.CheckPasswordHash(req.Password, user.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "usernaem or password error"})
		return
	}

	// 生成 token
	token, err := util.GenerateToken(user.ID)
	if err != nil {
		logger.Log.Errorf("Generate JWT error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "system error"})
		return
	}

	// 登录成功，返回 token
	c.JSON(http.StatusOK, gin.H{
		"message": "login success",
		"token":   token,
	})
}

func Profile(c *gin.Context) {
	uidVal, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Not login"})
		return

	}
	userID := uidVal.(uint)
	c.JSON(http.StatusOK, gin.H{
		"message": "Login success, you can access now",
		"user_id": userID,
	})
}
