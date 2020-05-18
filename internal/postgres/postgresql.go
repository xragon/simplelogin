package postgresql

import (
	"errors"

	"github.com/gofrs/uuid"

	// Postgresql Driver
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
)

// User represents the users of the app
type User struct {
	ID       uuid.UUID `db:"id"`
	Username string    `db:"username"`
	Password string    `db:"password"`
}

type store struct {
	DB *sqlx.DB
}

// GoodreadsStore exposes prostgres functions for books db
type GoodreadsStore interface {
	WriteRecord(User) error
	GetUser(string) (User, error)
}

// NewStore returns an instance of the GoodreadsStore
func NewStore() (GoodreadsStore, error) {
	s := &store{}
	var err error
	s.DB, err = sqlx.Connect("pgx", "postgresql://localhost:5432/books?user=books&password=books")
	if err != nil {
		return nil, err
	}

	return s, nil
}

// ReadSqlx is a practice function to be replaced
// func ReadSqlx() {
// 	db, err := sqlx.Connect("pgx", "postgresql://localhost:5432/books?user=books&password=books")
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	bookid, _ := uuid.NewV4()

// 	book := Book{
// 		ID:     bookid,
// 		Title:  "Example Book",
// 		Author: "blah",
// 		Rating: 5,
// 		ISBN:   "22222",
// 		ISBN13: "33333",
// 		Status: "read",
// 	}

// 	err = db.Get(&book, "SELECT * FROM books LIMIT 1")

// 	fmt.Println(book)
// }

// WriteRecord will store a Book object in the database
func (s *store) WriteRecord(u User) error {
	query := `INSERT INTO users (id, username, password) VALUES ($1, $2, $3)`

	_, err := s.DB.Exec(query, u.ID, u.Username, u.Password)
	if err != nil {
		return err
	}

	return nil
}

func (s *store) GetUser(un string) (User, error) {
	users := []User{}

	err := s.DB.Select(&users, `SELECT * FROM users WHERE username = $1`, un)
	if err != nil {
		return User{}, err
	}

	if len(users) == 0 {
		return User{}, errors.New("no records found")
	}

	return users[0], nil
}
