# GODO

GoDo is a simple REST API built in Go language. It provides basic CRUD functionality for a To-Do list.

## Installation

To install GoDo, clone this repository to your local machine and navigate to the project directory.

```
git clone https://github.com/your-username/go-do.git
cd go-do
```

## Usage

To run the server, execute the following command:

```
go run main.go
```

By default, the server will listen on localhost:42069. You can change this by modifying the server.Run() function call in the main function.

## Routes

The following routes are available in Go-Do:

|      **Route**       |                    **Description**                      | **Methods** |
|----------------------|---------------------------------------------------------|-------------|
| /godo                | Returns all objectives.                                 | GET         |
| /godo/:id            | Returns a specific objective by ID.                     | GET         |
| /godo                | Adds a new objective.                                   | POST        |
| /godo/:id            | Removes a specific objective by ID.                     | POST        |
| /godo/:id            | Toggles whether an objective is done or not by ID.      | PATCH       |

## Structure

Go-Do is structured as follows:

main.go - Contains the main function and server setup.

README.md - Contains information about the project.

go.mod - Contains module information and dependencies.

go.sum - Contains checksum information for dependencies.

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## Note

This version of project will store everything while server is running once it is stoped or restarted everything except default values will be removed. This problem will be fixed in next update where I will add database support and where everything will be stored in datbase.
