package controller

import (
	"net/http"

	"github.com/clapat-bb/memo/model"
	"github.com/clapat-bb/memo/util"
	"github.com/gin-gonic/gin"
)

type MemoRequest struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

func CreateMemo(c *gin.Context) {
	uidVal, _ := c.Get("userID")
	userID := uidVal.(uint)

	var req MemoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, "parameter error")
		return
	}

	memo := model.Memo{
		Title:   req.Title,
		Content: req.Content,
		UserID:  userID,
	}

	if err := model.DB.Create(&memo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "create failed",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "create success", "memo": memo})
}

func ListMemos(c *gin.Context) {
	uidVal, _ := c.Get("userID")
	userID := uidVal.(uint)

	page := c.DefaultQuery("page", "1")
	pageSize := c.DefaultQuery("page_size", "10")
	keyword := c.DefaultQuery("keyword", "")

	var (
		pageInt     = util.Atoi(page, 1)
		pageSizeInt = util.Atoi(pageSize, 10)
		offset      = (pageInt - 1) * pageSizeInt
	)
	var memos []model.Memo
	query := model.DB.Where("user_id = ?", userID)

	if keyword != "" {
		like := "%" + keyword + "%"
		query = query.Where("title ILIKE ? OR content ILIKE ?", like, like)
	}

	var total int64

	query.Model(&model.Memo{}).Count(&total)
	query.Order("created_at desc").Limit(pageSizeInt).Offset(offset).Find(&memos)

	c.JSON(http.StatusOK, gin.H{
		"total":     total,
		"page":      pageInt,
		"page_size": pageSizeInt,
		"data":      memos,
	})

	model.DB.Where("user_id = ?", userID).Order("create_at desc").Find(&memos)
	c.JSON(http.StatusOK, gin.H{"memos": memos})
}

func UpdateMemo(c *gin.Context) {
	uidVal, _ := c.Get("userID")
	userID := uidVal.(uint)

	id := c.Param("id")

	var memo model.Memo
	if err := model.DB.First(&memo, "id = ? AND user_id = ?", id, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "memo don't exist"})
		return
	}

	var req MemoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "parameter error"})
		return
	}

	memo.Title = req.Title
	memo.Content = req.Content

	if err := model.DB.Save(&memo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "update failed"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "update success", "memo": memo})
}

func DeleteMemo(c *gin.Context) {
	uidVal, _ := c.Get("userID")
	userID := uidVal.(uint)

	id := c.Param("id")

	var memo model.Memo
	if err := model.DB.First(&memo, "id = ? AND user_id = ?", id, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "memo don't exist"})
		return
	}

	if err := model.DB.Delete(&memo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "delete failed"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "delete success"})
}
