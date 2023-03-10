# ToDo List API
This is a simple API for managing to-do items. The API allows you to view, create, update and delete to-do items.

# Getting started

To use this API, you will need to have Go installed.

1. Clone the repository: ```git clone https://github.com/vertionn/GO-Todo-Rest-API.git```
2. Open the project directory: ```cd GO-Todo-Rest-API-main```
3. Start the server: ```go run main.go```
4. Send requests to the server using a tool like Postman or cURL.

The server should now be running on http://localhost:8080.

# Endpoints
`/todos`

Returns all ToDos in the array.

- Method: GET
- Response:
  - 200: Successful request with an array of ToDos.
  - 200: Empty array with an error message.
  - 500: Internal server error.
 
 
`/todo/:title`

Returns the ToDo with the title specified.

- Method: GET
- Response:
  - 200: Successful request with the ToDo object.
  - 404: ToDo not found.
  - 500: Internal server error.
  
  
`/create/todo`

Creates a new ToDo and adds it to the array.

- Method: POST
- Payload:
  - title: The title of the ToDo.
  - description: The description of the ToDo.
  - date: The due date of the ToDo (in MM/DD/YYYY format).
- Response:
  - 200: Successful request with the new ToDo object and a success message.
  - 400: Invalid request payload.
  - 422: Missing required fields.
  - 500: Internal server error.
  
  
 
`/update/todo/:title`

Updates the specified ToDo with the new information.

- Method: PATCH
- Payload:
  - title: The new title of the ToDo.
  - description: The new description of the ToDo.
  - date: The new due date of the ToDo (in MM/DD/YYYY format).
- Response:
  - 200: Successful request with the updated ToDo object and a success message.
  - 400: Invalid request payload.
  - 404: ToDo not found.
  - 500: Internal server error.


`/remove/todo/:title`

Deletes the specified ToDo from the array.

- Method: DELETE
- Response:
  - 200: Successful request with a success message.
  - 404: ToDo not found.
  - 500: Internal server error.
  
