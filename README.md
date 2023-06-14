


# JSONStore

JSONStore is a Go application that provides a RESTful API for managing users and orders using a PostgreSQL database. It uses the GORM library for database operations and the Gorilla Mux library for routing HTTP requests.

## Prerequisites

Before running the application, make sure you have the following installed:

- Go (version 1.16 or higher)
- PostgreSQL
- `github.com/jinzhu/gorm` package
- `github.com/lib/pq` package
- `github.com/gorilla/mux` package

## Installation

1. Clone the repository:

   ```shell
   git clone https://github.com/rassulmagauin/jsonstore.git
   ```

2. Change to the project directory:

   ```shell
   cd jsonstore
   ```

3. Install the dependencies:

   ```shell
   go mod download
   ```

4. Configure the PostgreSQL database connection in the `models.InitDB` function in the `models.go` file. Update the following line with your database credentials:

   ```go
   db, err := gorm.Open("postgres", "postgres://username:password@localhost/mydb?sslmode=disable")
   ```

   Replace `username` and `password` with your PostgreSQL credentials, and `mydb` with the name of the database.

5. Initialize the database by running the following command:

   ```shell
   go run models.go
   ```

   This will create the necessary tables (`user` and `order`) in the database and perform any pending migrations.

6. Start the server:

   ```shell
   go run main.go
   ```

   The server will start listening on `http://127.0.0.1:8000`.

## API Endpoints

The following API endpoints are available:

- **GET /v1/user/{id}**

  Retrieves a user by ID.

- **POST /v1/user**

  Creates a new user. The user data should be provided in the request body as JSON.

- **GET /v1/user?first_name={name}**

  Retrieves a list of users matching the specified first name.

- **POST /v1/order**

  Creates a new order. The order data should be provided in the request body as JSON.

- **GET /v1/order/{id}**

  Retrieves an order by ID.

## Example Usage

### Create a User

```shell
curl -X POST -H "Content-Type: application/json" -d '{"first_name":"John","last_name":"Doe","email":"john@example.com"}' http://127.0.0.1:8000/v1/user
```

### Get a User

```shell
curl -X GET http://127.0.0.1:8000/v1/user/{user_id}
```

Replace `{user_id}` with the actual ID of the user.

### Get Users by First Name

```shell
curl -X GET "http://127.0.0.1:8000/v1/user?first_name=John"
```

Replace `John` with the desired first name.

### Create an Order

```shell
curl -X POST -H "Content-Type: application/json" -d '{"product":"iPhone","quantity":2}' http://127.0.0.1:8000/v1/order
```

### Get an Order

```shell
curl -X GET http://127.0.0.1:8000/v1/order/{order_id}
```

Replace `{order_id}` with the actual ID of the order.

## Conclusion

JSONStore provides
