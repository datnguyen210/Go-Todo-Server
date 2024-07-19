package http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"datnguyen/todo/internal/entity"
	"datnguyen/todo/internal/usecase"
)

const errInvalidTodoID = "Invalid todo ID"

type TodoHandler struct {
	UseCase *usecase.TodoUseCase
}

func NewTodoHandler(uc *usecase.TodoUseCase) *TodoHandler {
	return &TodoHandler{UseCase: uc}
}

func (h *TodoHandler) IndexTodos(c *gin.Context) {
	todos, err := h.UseCase.GetAllTodos()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, todos)
}

func (h *TodoHandler) ReadTodoByID(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errInvalidTodoID})
		return
	}

	todo, err := h.UseCase.GetTodoByID(id)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, todo)
}

func (h *TodoHandler) CreateTodo(c *gin.Context) {
	var newTodo entity.Todo
	if err := c.BindJSON(&newTodo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	createdTodo, err := h.UseCase.CreateTodo(newTodo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusCreated, createdTodo)
}

func (h *TodoHandler) UpdateTodo(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errInvalidTodoID})
		return
	}

	var newTodo entity.Todo
	if err := c.BindJSON(&newTodo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedTodo, err := h.UseCase.UpdateTodo(id, newTodo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedTodo)
}

func (h *TodoHandler) DeleteTodoByID(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errInvalidTodoID})
		return
	}

	err = h.UseCase.DeleteTodo(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Todo deleted successfully"})
}
