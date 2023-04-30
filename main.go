package main    // Including the main package

import (    // Including the packages
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type godo struct {    // Defining a structure named godo
	ID          string `json:"ID"`
	Title       string `json:"Title"`
	Description string `json:"Description"`
	Done        bool   `json:"Done"`
}

var objectives = []godo{    // Initialize objectives with array of godo structure
	{ID: "1", Title: "Clean the room", Description: "Dummy just go clear your room", Done: false},
	{ID: "2", Title: "Do laundry", Description: "Wash all clothes including bedsheet and towels", Done: false},
	{ID: "3", Title: "Buy groceries", Description: "Buy milk, bread, cheese and eggs from the grocery store", Done: true},
	{ID: "4", Title: "Take out the trash", Description: "Dispose off all the trash from the dustbins", Done: false},
}

func main() {    // Defining the main function
	//Starting the server and listening for requests
	server := gin.Default()

	//Get all objectives
	server.GET("/godo", func(context *gin.Context) {    // When user send a GET request to /godo route
		context.IndentedJSON(http.StatusOK, objectives)    // Send the objectives array back to user
	})

	//Get specific objective by id
	server.GET("/godo:id", func(context *gin.Context) {    // When user send a GET request to /godo:id route
		id := context.Param("id")    // Store the id parameter sent by user in a variable
		godo, err := objectiveById(id)    // Get the goal by id

		if err != nil {    // Check if goal is found or not
			context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Couldn't find objective by that ID"})    // If goal is not found, send 404 response to user
		}

		context.IndentedJSON(http.StatusOK, godo)    // If goal is found then send 200 response to user with goal
	})

	//Set new objective
	server.POST("/godo", func(context *gin.Context) {    // When user send a POST request to /godo route
		var newObjective godo

		if err := context.BindJSON(&newObjective); err != nil {    // Bind the user submitted data to the godo structure
			return
		}

		objectives = append(objectives, newObjective)    // Append the new goal to the objectives array
		context.IndentedJSON(http.StatusOK, newObjective)    // Send the new goal back to user
	})

	//Remove objective
	server.POST("/godo/:id", func(context *gin.Context) {    // When user send a POST request to /godo/:id
		id := context.Param("id")    // Store the id parameter in id variable
		objectives = removeObjectiveById(objectives, id)    // Remove the goal by id
		context.IndentedJSON(http.StatusOK, gin.H{"message": "Successfully removed objective"})    // Send the success message to user
	})

	//Toggle weather objective is done or not
	server.PATCH("/godo/:id", func(context *gin.Context) {    // When user send a PATCH request to /godo/:id
		id := context.Param("id")    // Store the id parameter in a variable
		godo, err := objectiveById(id)    // Get the goal by id

		if err != nil {    // Check if goal is found or not
			context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Couldn't find objective"})    // If goal is not found, send 404 response to user
		}

		godo.Done = !godo.Done    // Toggle the boolean value of done property
		context.IndentedJSON(http.StatusOK, godo)    // Send the goal back to user
	})

	server.Run("localhost:42069")    // Run the server on localhost:42069
}

// Helpers
func objectiveById(id string) (*godo, error) {    // Defining the objectiveById function to get the goal by id
	for _, value := range objectives {    // Iterate over the objectives array
		if value.ID == id {    // Compare id with the id property of goal
			return &value, nil    // If ids are equal then return the goal
		}
	}
	return nil, errors.New("[X] Couldn't find objective you are looking for")    // If goal is not found, return nil and an error message
}

func removeObjectiveById(objectives []godo, id string) []godo {    // Defining the removeObjectiveById function to remove the goal by id
	for i, obj := range objectives {    // Iterate over the objectives array
		if obj.ID == id {    // Compare id with the id property of goal
			return append(objectives[:i], objectives[i+1:]...)    // If ids are equal then remove the goal
		}
	}
	return objectives    // If goal is not found then return the objectives array
}
