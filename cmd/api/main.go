package main

import (
	"log"

	"github.com/juancuadra2/crud-api-go/config"
	"github.com/juancuadra2/crud-api-go/internal/adapters/primary/rest/handlers"
	"github.com/juancuadra2/crud-api-go/internal/adapters/primary/rest/router"
	"github.com/juancuadra2/crud-api-go/internal/adapters/secondary/mysql_adapter"
	"github.com/juancuadra2/crud-api-go/internal/adapters/secondary/mysql_adapter/entity"
	"github.com/juancuadra2/crud-api-go/internal/core/services"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	// Cargar la configuración
	configuration := config.Load()

	// Conectar a la base de datos
	db, err := gorm.Open(mysql.Open(configuration.GetDSN()), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Automigrar el esquema de la base de datos
	if err := db.AutoMigrate(&entity.UserEntity{}); err != nil {
		log.Fatal("Failed to migrate database schema:", err)
	}

	// Inyección de dependencias y arranque del servidor se omiten por brevedad
	userRepository := mysql_adapter.NewUserAdapter(db)
	userService := services.NewUserService(userRepository)
	userHandler := handlers.NewUserHandler(userService)

	// Configurar rutas
	routes := router.SetupRoutes(userHandler)

	log.Printf("Servidor corriendo en el puerto %s", configuration.ServerPort)
	if err := routes.Run(":" + configuration.ServerPort); err != nil {
		log.Fatal("Failed to run server:", err)
	}
}
