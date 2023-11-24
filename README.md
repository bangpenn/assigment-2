# Assignment 2 - Golang API

This repository contains the implementation of Assignment 2 for the "Scalable With Golang" course at Hacktiv8. The Golang API is designed to manage orders and items in a simple system.

## Features
- Create a new order with detailed information.
- Retrieve the list of orders with associated items.
- Update order information based on its ID.
- Delete an order based on its ID.

## Technologies Used
- Go (Golang): The primary programming language for developing the API.
- Gin: The web framework used to facilitate the development of the API.
- SQLite: The database used to store order and item information.

## How to Use
To get started with the Golang API, follow these steps:

1. Clone this repository to your local machine.
2. Make sure you have Go installed.
3. Run the command `go run main.go` to start the server.
4. Use tools like Postman to test the API by accessing the appropriate URLs for the desired endpoints.

## Endpoints
- **Create Order**: `POST /orders`
- **Get Orders**: `GET /orders`
- **Update Order**: `PUT /orders/{orderId}`
- **Delete Order**: `DELETE /orders/{orderId}`

## Notes
This project was created as part of completing the "Scalable With Golang" course at Hacktiv8. Please use this project as a reference or learning material for developing your own back-end applications.

Feel free to modify this README.md to better suit your specific project details.
