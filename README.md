# VideoClub / Go-REST-API
A simple Go-REST-API to manage movie rentals.

This REST API was created with Go using the Gin framework, the GORM ORM and PostgreSQL for relational database management.

## Entity relationship
![Entity relationship](https://raw.githubusercontent.com/jorgemvv01/go-api/master/entity_relationship.jpg)

## Rental price based on type of movies
```
1. New releases - The unit price for each of the rental days.
2. Regular movies - Unit price for the first three days. Each additional day will be an increase of 15% of the unit price per day.
3. Old movies - Unit price for the first five days. Each additional day will be an increase of 10% of the unit price per day.
```

## Installation & Run
**Step 1:**

Download or clone this repo by using the link below:
```
https://github.com/jorgemvv01/go-api
```

**Step 2:**

Create videoclub database:
```sql
CREATE DATABASE videoclub;
```
Configures the database connection with the environment variables on [storage.go](https://github.com/jorgemvv01/go-api/tree/master/storage/storage.go):
```go
func GetInstance() *gorm.DB {
	if db == nil {
		var err error
		dsn := os.Getenv("DATABASE_URL")
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			panic("failed to connect database")
		}
	}
	return db
}
```
```
DATABASE_URL=postgresql://YOUR_USER:YOUR_PASSWORD@HOST:PORT/videoclub
```

**Step 3:**

Run all test:
```bash
go test ./...
```

**Step 4:**

Build and run:
```bash
cd go-api
go build
go run main.go
```

## Documentation
Go to:
```
your_host/api/docs/index.html
```

## Structure
```
├── controllers
├── docs
├── models
├── repositories
├── routes
├── storage
├── tests
│   ├── controllers
├── utils
└── main.go
```

## API

#### Movie Genre
* `/genres` - `GET`: Get all genres
* `/genres/{ID}` - `GET`: Get genre by ID
* `/genres/create` - `POST`: Create genre
* `/genres/update/{ID}` - `PUT`: Update genre
* `/genres/delete/{ID}` - `DELETE`: Delete genre

#### Movie Type
* `/types` - `GET`: Get all types
* `/types/{ID}` - `GET`: Get type by ID

#### Movies
* `/movies` - `GET`: Get all movies
* `/movies/{ID}` - `GET`: Get movie by ID
* `/movies/create` - `POST`: Create movie
* `/movies/update/{ID}` - `PUT`: Update movie
* `/movies/delete/{ID}` - `DELETE`: Delete movie

#### Users
* `/users` - `GET`: Get all users
* `/users/{ID}` - `GET`: Get user by ID
* `/users/create` - `POST`: Create user
* `/users/update/{ID}` - `PUT`: Update user
* `/users/delete/{ID}` - `DELETE`: Delete user

#### Rent
* `/rent/create` - `POST`: Create rent
