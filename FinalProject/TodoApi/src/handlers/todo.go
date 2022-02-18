package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nullseed/logruseq"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"host.local/go/golang-todo-api/src/models"
)

func init() {
	log.AddHook(logruseq.NewSeqHook("http://localhost:5341"))
	log.Info("Starting the application")
}

type Todos struct{}

var todoModel models.Todo = models.Todo{}

// NewTodo creates an instance of todos
func NewTodo() *Todos {
	return &Todos{}
}

func (p *Todos) CreateTodo(c *gin.Context) {

	c.BindJSON(&todoModel)
	log.Info("Creating a new todo")
	if len(todoModel.Task) < 2 {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error": "todo must be at least 3 chars length"},
		)
		return
	}

	result, err := todoModel.InsertOne()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	oid, _ := result.InsertedID.(primitive.ObjectID)
	c.JSON(http.StatusCreated, map[string]interface{}{
		"id": oid.Hex(),
	})
}

func (todo *Todos) GetTodos(c *gin.Context) {
	log.Info("Getting all todos")
	todoModel, err := todoModel.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	// c.JSON(http.StatusOK, gin.H{"tasks": todoModel})
	c.JSON(http.StatusOK, todoModel)
}

func (todo *Todos) UpdateTodo(c *gin.Context) {
	id := c.Param("id")
	log.Info("Updating a todo with id %s", id)
	c.BindJSON(&todoModel)

	err := todoModel.Update(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.Status(http.StatusNoContent)
}

func (todo *Todos) DeleteTodo(c *gin.Context) {
	id := c.Param("id")
	log.Info("Deleting a todo with id %s", id)
	err := todoModel.Delete(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return

	}

	c.Status(http.StatusNoContent)
}
