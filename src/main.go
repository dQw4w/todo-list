package main

import (
	//"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var schema = `
CREATE TABLE IF NOT EXISTS todo(
	id int primary key,
	title VARCHAR(50) not null,
	complete boolean not null
);
`

var db *sqlx.DB = nil

type Todo struct {
	ID       int    `json:"id" db:"id"`
	Title    string `json:"title" db:"title"`
	Complete bool   `json:"complete" db:"complete"`
}

func getTodos(c *gin.Context) {
	todos := []Todo{}
	if err := db.Select(&todos, "SELECT id,title,complete FROM todo ORDER BY id ASC;"); err != nil{
        c.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
        return
    }
	c.JSON(http.StatusOK, todos)
}
func createTodo(c *gin.Context) {
	var newTodo Todo
	if err := c.ShouldBindJSON(&newTodo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errora": err.Error()})
		return
	}
	newID, newTitle, newComplete := newTodo.ID, newTodo.Title, newTodo.Complete
	tx := db.MustBegin()
	if _, err1 := tx.Exec("INSERT INTO todo (id,title,complete) VALUES ($1, $2, $3)", newID, newTitle, newComplete); err1 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errorb": err1.Error()})
		return
	}
	tx.Commit()
	c.JSON(http.StatusOK, gin.H{"message": "Added"})
}
func deleteTodo(c *gin.Context) {
    input := c.Param("id")
    tx := db.MustBegin()
    if input == "all"{
        if _,err := tx.Exec("TRUNCATE TABLE todo;"/*"DELETE FROM todo;"*/); err != nil{
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
        tx.Commit()
        c.JSON(http.StatusOK, gin.H{"message": "Deleted everything"})
        return
    }
    todoID,err:= strconv.Atoi(input)
    if (err != nil){
        c.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
        return
    }
    result, err1 := tx.Exec("DELETE FROM todo WHERE id = $1;", todoID)
	if err1 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errorb": err1.Error()})
		return
	}
	tx.Commit()
    if affected,_ := result.RowsAffected(); affected == 0{
        c.JSON(http.StatusBadRequest, gin.H{"error": "Todo not found"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Todo deleted"})
}
func updateTodoStatus(c *gin.Context) {
    todoID ,err := strconv.Atoi(c.Param("id"))
    if (err != nil){
      c.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
      return
    }
    tx := db.MustBegin()
    result, err1 := tx.Exec("UPDATE todo SET complete = CASE WHEN complete = true THEN false ELSE true END WHERE id = $1;", todoID)
    if err1 != nil {
        c.JSON(http.StatusBadRequest, gin.H{"errorb": err1.Error()})
        return
    }
    tx.Commit()
    if affected,_ := result.RowsAffected(); affected == 0{
        c.JSON(http.StatusBadRequest, gin.H{"error": "Todo not found"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message" : "Status updated"})
}
    
func main() {
    fmt.Println("Start")
	var err error
	db, err = sqlx.Connect("postgres", "user=postgres dbname=postgres password=mysecretpassword sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}
	db.MustExec(schema)
	r := gin.Default()
	r.GET("/todos", getTodos)
	r.POST("/todos", createTodo)
    r.PUT("/todos/:id", updateTodoStatus)
    r.DELETE("/todos/:id", deleteTodo)
	r.Run(":8080")
}
