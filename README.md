## GO GORM PostgreSQL CRUD REST API
Build a `CRUD` `REST API` in `Go (Golang)` using **GORM** with `PostgreSQL` and the `Gorilla Mux framework`.

>For PostgreSQL
## Installation
1. **Clone the Repository**
 ```bash
https://github.com/alamgir-ahosain/go-gorm-postgres-crud-rest-api.git
```
2. **Install Dependencies**<br>
 ```bash
go get -u gorm.io/gorm
go get -u gorm.io/driver/postgres
go get -u github.com/gorilla/mux
go get github.com/joho/godotenv
```

##  Database Setup
Login to PostgreSQL and create the database and table:
```sql
CREATE DATABASE goDB;
USE goDB;

CREATE TABLE users (
    id SERIAL PRIMARY KEY,         
    sid VARCHAR(128) NOT NULL,
    name VARCHAR(128) NOT NULL,
    cgpa NUMERIC(5,2)             
);
```
## Run the Application:
1. Start PostgreSQL server.  
2. Update database connection in `db/postgreSql.go`
3. Configure `.env` in `config/.env`
4. Run the Project:
```go
go run main.go
```
   
##  Project Structure
```plaintext
GO-GORM-POSTGRES-CRUD-REST-API

│── cmd/
│   │─── MyApp
│   │    ├── server.go                # main serve file
│   │
│── internal/
│   ├── api/
│   │   ├── handlers/
│   │   │   ├── get_users/            
│   │   │   │   ├── get_user_by_id.go # get user by ID function
│   │   │   │   ├── get_users.go      # get all users function
│   │   │   │── create_user.go        # create user function
│   │   │   │── delete_user.go        # delete user function
│   │   │   │── update_user.go        # update user function 
│   │   ├── routes/
│   │       └── routes.go             # all routes declared here
│   │
│   │── config/
│   │   │── .env                      # configuration environment 
│   │
│   │── db/
│   │   │── postgreSql.go             # postgreSQL database connection setup
│   │
│   │── models/
│   │   │── model.go                  # User model definition
│   │
│   │── services/
│   │   │── service.go                # Helper function 
│
│── main.go                           # Main entry file 
│── go.mod                            # Go module definition
│── go.sum                            # Go dependencies checksums
│── README.md                         # Project documentation 

```
--- 
##  API endpoints table

| number | method | endpoint       |     description               |
| ------ | ------ | -------------- | ------------------------------|
|    1   | GET    | /users         | get all users                 |
|    2   | GET    | /users/{id}    | get user by id                |
|    3   | POST   | /users         | create new user               |
|    4   | PUT    | /users/{id}    | full update user by id        |
|    5   | DELETE | /users/{id}    | delete user by id             |

## Testing the CRUD API with Postman or Thunder Client( VS code extension)
#### 1. Get All Users
  Method: GET<br> URL: http://localhost:8000/users<br>
  
#### 2. Get a single User
  Method: GET<br> URL: http://localhost:8000/users/1<br>
  
#### 3. Create a new  Users
  Method: POST<br> URL: http://localhost:8000/users<br>
  
#### 4. Update a User
  Method: PUT<br> URL: http://localhost:8000/users/1<br>

#### 5. Delete a User
  Method: DELETE<br> URL: http://localhost:8000/users/1<br>

---
# Thank You for Checking Out This Repository!
Your feedback is valuable! Please share your thoughts and suggestions for improvement via [GitHub Issues](https://github.com/alamgir-ahosain/go-gorm-postgres-crud-rest-api/issues).

# Contributing  
Contributions are welcome! Feel free to fork the repo and create a pull request.


