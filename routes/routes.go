package routes

import (
	"GIN_GORM/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {

	member := router.Group("/members")
	{
		member.GET("", func(c *gin.Context) {
			handlers.GetMembers(c.Writer, c.Request)
		})

	}

	role := router.Group("/roles")
	{
		role.POST("/create", func(c *gin.Context) {
			handlers.SetRoles(c.Writer, c.Request)
		})
	}

	artist := router.Group("/artists")
	{
		artist.GET("", func(c *gin.Context) {
			handlers.GetArtists(c.Writer)
		})

		artist.POST("/create", func(c *gin.Context) {
			handlers.SetArtists(c.Writer, c.Request)
		})

		artist.DELETE("/delete", func(c *gin.Context) {
			handlers.DeleteArtists(c.Writer, c.Request)
		})
	}
}
