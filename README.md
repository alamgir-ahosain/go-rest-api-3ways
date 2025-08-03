## Go REST API 
>A RESTful API built using the Go standard library, Gorilla Mux, and Gin to create, read, update, and delete products.<br>**No database implementation yet**.

## Installation
1. **Clone the Repository**
 ```bash
https://github.com/alamgir-ahosain/go-rest-api-3ways.git
```
2. **Install Dependencies**<br>

**For Gin**
 ```bash
go get -u github.com/gin-gonic/gin
```
**For Gorilla Mux**
```bash
go get -u github.com/gorilla/mux
```
3. **Run the Application:** 
```bash
go run main.go
```
 ##  Project Structure
 
```plaintext
1_standardWay/
├── cmd/
│   └── myApp/
│       └── main.go        # main file
├── internal/
│   ├── api/
│   │   └── routes.go      # API routes 
│   ├── handlers/
│   │   └── handler.go     # HTTP handlers
│   ├── models/
│   │   └── model.go       # Data models
│   └── services/
│       └── service.go     # Business logic 
├── .gitignore             # Git ignore file
├── go.mod                 # Go module definition
├── go.sum                 # Go dependencies checksums
└── README.md              # Project readme file
```
##  api endpoints table

| number | method | endpoint       | description                      |
| ------ | ------ | -------------- | -------------------------------- |
|    1   | GET    | /products      | get all products                 |
|    2   | GET    | /products/{id} | get product by id                |
|    3   | POST   | /products/{id} | create new product               |
|    4   | PUT    | /products/{id} | full update product by id        |
|    5   | PATCH  | /products/{id} | partial update product by id     |
|    6   | DELETE | /products/{id} | delete product by id             |

##  Product Data

```go
type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

```

# Thank You for Checking Out This Repository!
Your feedback is valuable! Please share your thoughts and suggestions for improvement by creating a pull request.
