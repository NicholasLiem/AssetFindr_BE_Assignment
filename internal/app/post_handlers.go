package app

import (
	"github.com/NicholasLiem/AssetFindr_BackendAssignment/internal/datastruct"
	"github.com/NicholasLiem/AssetFindr_BackendAssignment/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (m *MicroserviceServer) CreatePost(c *gin.Context) {
	var post datastruct.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	result, httpErr := m.postService.CreatePost(post)
	if httpErr != nil {
		c.JSON(httpErr.StatusCode, gin.H{"error": httpErr.Message})
		return
	}

	c.JSON(http.StatusCreated, result)
}

func (m *MicroserviceServer) GetPost(c *gin.Context) {
	postId := c.Param("id")

	parsedPostId, err := utils.ParseStrToUint(postId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post id"})
		return
	}

	result, httpErr := m.postService.GetPostWithTag(*parsedPostId)
	if httpErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not retrieve post"})
		return
	}

	c.JSON(http.StatusOK, result)
}

func (m *MicroserviceServer) UpdatePost(c *gin.Context) {
	postId := c.Param("id")

	parsedPostId, err := utils.ParseStrToUint(postId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post id"})
		return
	}

	var updatedPost datastruct.Post
	if err := c.ShouldBindJSON(&updatedPost); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	result, httpErr := m.postService.UpdatePost(*parsedPostId, updatedPost)
	if httpErr != nil {
		c.JSON(httpErr.StatusCode, gin.H{"error": httpErr.Message})
		return
	}

	c.JSON(http.StatusOK, result)
}

func (m *MicroserviceServer) DeletePost(c *gin.Context) {
	postId := c.Param("id")

	parsedPostId, err := utils.ParseStrToUint(postId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post id"})
		return
	}

	result, httpErr := m.postService.DeletePost(*parsedPostId)
	if httpErr != nil {
		c.JSON(httpErr.StatusCode, gin.H{"error": httpErr.Message})
		return
	}

	c.JSON(http.StatusOK, result)
}

func (m *MicroserviceServer) GetAllPost(c *gin.Context) {
	result, httpErr := m.postService.GetAllPostWithTags()
	if httpErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not retrieve posts"})
		return
	}

	c.JSON(http.StatusOK, result)
}
