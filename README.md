## REST API with mux

This is a simple REST API with [mux](https://github.com/gorilla/mux).
 ## Project Structure
 
```plaintext
3_gin_Framework/
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
## api endpoints table

|     | method | endpoint       | description               |
| --- | ------ | -------------- | ------------------------- |
| 1   | GET    | /products      | get all products          |
| 2   | GET    | /products/{id} | get product by id         |
| 3   | POST   | /products/{id} | create new product        |
| 4   | PUT    | /products/{id} | full update product by id |
| 5   | PATCH  | /products/{id} | partial product by id     |
| 6   | DELETE | /products/{id} | delete product by id      |

## Api structure

```go
type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

```
