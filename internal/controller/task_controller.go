package controller

import (
	"net/http"
	"strconv"
	"task-manager/internal/model"
	"task-manager/internal/repository"

	"github.com/gin-gonic/gin"
)

// Nama struct diubah menjadi TaskController
type TaskController struct {
	taskRepository *repository.TaskRepository
}

// Nama fungsi diubah menjadi NewTaskController
func NewTaskController(repo *repository.TaskRepository) *TaskController {
	return &TaskController{taskRepository: repo}
}

// controller untuk membuat task
func (controller *TaskController) CreateTask(context *gin.Context) {
	var task model.Task

	if err := context.ShouldBindJSON(&task); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	savedTask, err := controller.taskRepository.Save(task)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, savedTask)
}

// controler untuk mencari semua tasks
func (controller *TaskController) FindAll(context *gin.Context) {
	tasks, err := controller.taskRepository.FindAll()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "ID Harus berupa angka"})
		return
	}
	context.JSON(http.StatusCreated, tasks)
}

// controller untuk menghapus task berdasarkan id
func (controller *TaskController) DeleteTask (context *gin.Context) {
	idString := context.Param("id")

	id, err := strconv.Atoi(idString)
	if err != nil {
		context.JSON(http.StatusBadRequest, err)
	}

	err = controller.taskRepository.DeleteByID(id)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	}

	context.JSON(http.StatusOK, gin.H{"message": "Task berhasil dihapus"})
}

// controller untuk mencari task berdasarkan id
func (controller *TaskController) GetTaskByID (context *gin.Context) {
	idString := context.Param("id")

	id, err := strconv.Atoi(idString)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "ID harus berupa angka!"})
		return
	}

	task, err := controller.taskRepository.FindByID(id)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Task tidak ditemukan"})
		return
	}

	context.JSON(http.StatusOK, task)
}