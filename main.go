package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	_ "api/docs" // 👈 Replace "api" with your module name from go.mod

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// =========================================================
// ============== Data Models ==============================
// =========================================================

type Blog struct {
	ID      string `json:"id" example:"1"`
	Title   string `json:"title" example:"My First Blog"`
	Content string `json:"content" example:"This is a sample blog post."`
}

type ErrorResponse struct {
	Error string `json:"error" example:"Blog not found"`
}

type SuccessResponse struct {
	Message string `json:"message" example:"Blog deleted successfully"`
}

// =========================================================
// ============== Sample Data ==============================
// =========================================================

var blogs = []Blog{
	{ID: "1", Title: "Go Basics", Content: "Learn Go programming..."},
	{ID: "2", Title: "Gin Framework", Content: "Building APIs in Go..."},
}

// =========================================================
// ============== Handlers =================================
// =========================================================

// @Summary Get all blogs
// @Description Retrieve a list of all blogs
// @Tags blogs
// @Accept  json
// @Produce  json
// @Success 200 {array} Blog
// @Router /blogs [get]
func GetBlogs(c *gin.Context) {
	c.JSON(http.StatusOK, blogs)
}

// @Summary Get a blog by ID
// @Description Retrieve blog details by its ID
// @Tags blogs
// @Accept  json
// @Produce  json
// @Param id path string true "Blog ID"
// @Success 200 {object} Blog
// @Failure 404 {object} ErrorResponse
// @Router /blog/{id} [get]
func GetBlogByID(c *gin.Context) {
	id := c.Param("id")
	for _, b := range blogs {
		if b.ID == id {
			c.JSON(http.StatusOK, b)
			return
		}
	}
	c.JSON(http.StatusNotFound, ErrorResponse{Error: "Blog not found"})
}

// @Summary Create a new blog
// @Description Add a new blog post
// @Tags blogs
// @Accept  json
// @Produce  json
// @Param blog body Blog true "Blog Data"
// @Success 201 {object} Blog
// @Failure 400 {object} ErrorResponse
// @Router /blog [post]
func CreateBlog(c *gin.Context) {
	var newBlog Blog
	if err := c.ShouldBindJSON(&newBlog); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid JSON"})
		return
	}

	blogs = append(blogs, newBlog)
	c.JSON(http.StatusCreated, newBlog)
}

// @Summary Delete a blog by ID
// @Description Delete a specific blog post by its ID
// @Tags blogs
// @Param id path string true "Blog ID"
// @Success 200 {object} SuccessResponse
// @Failure 404 {object} ErrorResponse
// @Router /blog/{id} [delete]
func DeleteBlogByID(c *gin.Context) {
	id := c.Param("id")
	for i, b := range blogs {
		if b.ID == id {
			blogs = append(blogs[:i], blogs[i+1:]...)
			c.JSON(http.StatusOK, SuccessResponse{Message: "Blog deleted successfully"})
			return
		}
	}
	c.JSON(http.StatusNotFound, ErrorResponse{Error: "Blog not found"})
}

// =========================================================
// ============== Main =====================================
// =========================================================

// @title Blog API Example
// @version 1.0
// @description Simple Go Blog API with Swagger documentation.
// @host localhost:8080
// @BasePath /
func main() {
	r := gin.Default()

	// Auth
	r.POST("/token")
	r.POST("/rtoken")
	// Routes
	r.GET("/blogs", GetBlogs)
	r.GET("/blog/:id", GetBlogByID)
	r.POST("/blog", CreateBlog)
	r.DELETE("/blog/:id", DeleteBlogByID)

	// Swagger UI
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Start server
	r.Run(":8080")
}
