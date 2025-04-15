package main

import (
	"GIN_GORM/db"
	"GIN_GORM/models"
	"GIN_GORM/routes"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func main() {
	// Conectar a la base de datos
	db.ConnectDatabase()

	// Ejecutar migraciones solo si es necesario
	shouldMigrate := len(os.Args) > 1 && os.Args[1] == "migrate"

	if shouldMigrate {
		log.Println("Ejecutando migraciones...")
		models.ArtistMigration()
		models.RoleMigration()
		models.MemberMigration()
		models.AlbumMigration()
		models.SongMigration()
		log.Println("Migraciones completadas.")
	}

	// Inicializar el router de Gin
	router := gin.Default()

	// Configurar proxies confiables
	err := router.SetTrustedProxies([]string{"127.0.0.1"})
	if err != nil {
		return
	}

	// Configurar rutas
	routes.SetupRoutes(router)

	// Iniciar el servidor en puerto 8080
	log.Println("Servidor iniciado en http://localhost:8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Error al iniciar el servidor:", err)
	}
}
