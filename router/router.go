package router

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"task-manager/internal/controller"
	"task-manager/internal/repository"
)

func SetupRouter(db *sql.DB) *gin.Engine {
	// `request` atau `r` menjadi `routerEngine` agar lebih jelas
	routerEngine := gin.Default()

	// Inisialisasi Dependencies
	taskRepository := repository.NewTaskRepository(db)
	// Nama variabel diubah menjadi taskController
	taskController := controller.NewTaskController(taskRepository)

	// Rute Ping
	routerEngine.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Rute untuk tambah Tasks
	routerEngine.POST("/tasks", taskController.CreateTask)
	// Rute untuk ambil Tasks
	routerEngine.GET("/tasks", taskController.FindAll)
	// Rute untuk hapus Tasks berdasarkan ID
	routerEngine.DELETE("tasks/:id", taskController.DeleteTask)
	// Rute untuk ambil task berdasarkan ID
	routerEngine.GET("tasks/:id", taskController.GetTaskByID)
	// Rute untuk update task berdasarkan ID
	routerEngine.PUT("/tasks/:id", taskController.UpdateTask)

	return routerEngine
}
