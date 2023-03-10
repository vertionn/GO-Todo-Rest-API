package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type ToDoFormat struct {
	Complete    bool   `json:"complete"`
	Title       string `json:"title"`
	Description string `json:"description,omitempty"`
	DueDate     string `json:"date,omitempty"`
}

var ToDos []ToDoFormat

func main() {

	server := gin.Default()

	server.GET("/todos", func(ctx *gin.Context) {

		if len(ToDos) < 1 {
			ctx.IndentedJSON(http.StatusOK, gin.H{
				"error_type":        "empty todo Array",
				"error_description": "the array is empty so we could not return any data",
			})
		} else {
			ctx.IndentedJSON(http.StatusOK, ToDos)
		}
	})

	server.GET("/todo/:title", func(ctx *gin.Context) {

		title := ctx.Param("title")

		var found bool

		for _, item := range ToDos {
			if item.Title == title {

				found = true

				ctx.IndentedJSON(http.StatusOK, item)
				return
			}
		}

		if !found {
			ctx.IndentedJSON(http.StatusNotFound, gin.H{
				"error": "todo could not be found",
			})
		}

	})

	server.POST("/todo", func(ctx *gin.Context) {
		var (
			ToDo    ToDoFormat
			Missing []string
		)

		err := ctx.BindJSON(&ToDo)
		if err != nil {
			ctx.IndentedJSON(http.StatusBadRequest, gin.H{
				"error": "Invalid request payload",
			})
			return
		}

		if ToDo.Title == "" {
			Missing = append(Missing, "Title")
		}
		if ToDo.Description == "" {
			Missing = append(Missing, "Description")
		}
		if ToDo.DueDate == "" {
			Missing = append(Missing, "Date")
		} else {

			dateFormat := "01/02/2006"

			date, err := time.Parse(dateFormat, ToDo.DueDate)
			if err != nil {
				ctx.IndentedJSON(http.StatusBadRequest, gin.H{
					"error": "Invalid date format. Please use " + dateFormat,
				})
				return
			}

			ToDo.DueDate = date.Format(time.RFC3339)

			now := time.Now().Local()
			if date.Before(now) {
				ctx.IndentedJSON(http.StatusBadRequest, gin.H{
					"error": "Date is in the past",
				})
				return
			}

		}

		if len(Missing) > 0 {
			ctx.IndentedJSON(http.StatusUnprocessableEntity, gin.H{
				"error":          "missing fields required",
				"missing_fields": Missing,
			})
			return
		}

		ToDos = append(ToDos, ToDo)

		ctx.IndentedJSON(http.StatusOK, gin.H{
			"content": ToDo,
			"message": "successfully added",
		})
	})

	server.PATCH("/todo/:title", func(ctx *gin.Context) {
		var (
			title      string
			updateToDo ToDoFormat
		)

		title = ctx.Param("title")

		err := ctx.BindJSON(&updateToDo)
		if err != nil {
			ctx.IndentedJSON(http.StatusBadRequest, gin.H{
				"error": "invalid request payload",
			})
			return
		}

		var found bool
		for i, t := range ToDos {
			if t.Title == title {
				found = true

				if updateToDo.Title != "" {
					ToDos[i].Title = updateToDo.Title
				}
				if updateToDo.Description != "" {
					ToDos[i].Description = updateToDo.Description
				}
				if updateToDo.DueDate != "" {
					dateFormat := "01/02/2006"

					date, err := time.Parse(dateFormat, updateToDo.DueDate)
					if err != nil {
						ctx.IndentedJSON(http.StatusBadRequest, gin.H{
							"error": "Invalid date format. Please use " + dateFormat,
						})
						return
					}

					ToDos[i].DueDate = date.Format(time.RFC3339)
				}

				ctx.IndentedJSON(http.StatusOK, gin.H{
					"message": "todo updated",
					"content": ToDos[i],
				})
				return
			}
		}

		if !found {
			ctx.IndentedJSON(http.StatusNotFound, gin.H{
				"error": "todo could not be found",
			})
			return
		}
	})

	server.DELETE("/todo/:title", func(ctx *gin.Context) {

		title := ctx.Param("title")

		var found bool
		for indx, item := range ToDos {
			if item.Title == title {

				found = true

				ToDos = append(ToDos[:indx], ToDos[indx+1:]...)

				ctx.IndentedJSON(http.StatusOK, gin.H{
					"message": "successfully deleted",
				})
				return
			}
		}

		if !found {
			ctx.IndentedJSON(http.StatusNotFound, gin.H{
				"error": "todo could not be found",
			})
		}
	})

	server.Run()
}
