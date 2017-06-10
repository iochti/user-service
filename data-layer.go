package main

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/namsral/flag"

	"github.com/iochti/user-service/models"
	_ "github.com/lib/pq"
)

// DataLayerInterface is here to abstract DB usage in our code
// so we can use any DB and not break other functions
// used for testing too
type DataLayerInterface interface {
	CreateUser(user *models.User) error
	GetUserByID(id int) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
	GetUserByLogin(login string) (*models.User, error)
	DeleteUser(id int) error
}

// PostgresDL implements the DataLayerInterface
type PostgresDL struct {
	DBName string
	Host   string
	Db     *sql.DB
}

const USER_TABLE = "user"

// Init fetches db flags and init the db conn
func (p *PostgresDL) Init() error {
	// Parse DB flags
	host := flag.String("pq-host", "localhost", "PostgresSQL database host")
	user := flag.String("pq-user", "postgres", "PostgresSQL user")
	dbName := flag.String("pq-db", "iochti", "PostgresSQL DBName")
	password := flag.String("pq-pwd", "", "PostgresSQL user password")
	flag.Parse()

	// Create a db connection
	db, err := sql.Open("postgres", fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", *user, *password, *host, *dbName))
	if err != nil {
		return err
	}
	p.Db = db
	return nil
}

// CreateUser creates a user passed as parameter
func (p *PostgresDL) CreateUser(user *models.User) error {
	var userID int
	timeCreated := time.Now()
	err := p.Db.QueryRow("INSERT INTO "+USER_TABLE+`(name, login, avatar, email, created_at, updated_at)
		VALUES($1, $2, $3, $4, $5, $5) RETURNING id;`,
		user.Name, user.Login, user.AvatarURL, user.Email, timeCreated).Scan(&userID)

	if err != nil {
		return err
	}
	user.ID = userID
	user.Created = timeCreated
	user.Updated = timeCreated
	return nil
}

// GetUserByID fetch a user by its user id
func (p *PostgresDL) GetUserByID(id int) (*models.User, error) {
	// Check ID values
	if id <= 0 {
		return nil, fmt.Errorf("Error, invalid search ID: id must be > 0")
	}
	user := new(models.User)
	err := p.Db.QueryRow("SELECT id, name, login, avatar, email, created_at, updated_at FROM "+USER_TABLE+" WHERE id=$1;", id).Scan(
		&user.ID,
		&user.Name,
		&user.Login,
		&user.AvatarURL,
		&user.Email,
		&user.Created,
		&user.Updated,
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}

// GetUserByLogin fetch a user by its github login
func (p *PostgresDL) GetUserByLogin(login string) (*models.User, error) {
	// Check ID values
	if login == "" {
		return nil, fmt.Errorf("Error, invalid search: login must not be empty")
	}
	user := new(models.User)
	err := p.Db.QueryRow("SELECT id, name, login, avatar, email, created_at, updated_at FROM "+USER_TABLE+" WHERE login=$1;", login).Scan(
		&user.ID,
		&user.Name,
		&user.Login,
		&user.AvatarURL,
		&user.Email,
		&user.Created,
		&user.Updated,
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}

// GetUserByEmail fetch a user by its github email
func (p *PostgresDL) GetUserByEmail(email string) (*models.User, error) {
	// Check ID values
	if email == "" {
		return nil, fmt.Errorf("Error, invalid search: login must not be empty")
	}
	user := new(models.User)
	err := p.Db.QueryRow("SELECT id, name, login, avatar, email, created_at, updated_at FROM "+USER_TABLE+" WHERE email=$1;", email).Scan(
		&user.ID,
		&user.Name,
		&user.Login,
		&user.AvatarURL,
		&user.Email,
		&user.Created,
		&user.Updated,
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}

// DeleteUser delete a user identified by its id
func (p *PostgresDL) DeleteUser(id int) error {
	if id <= 0 {
		return fmt.Errorf("Error, invalid argument: id must be > 0")
	}

	res, err := p.Db.Exec("DELETE FROM "+USER_TABLE+" WHERE id = $1;", id)
	if err != nil {
		return err
	}
	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return fmt.Errorf("Error: no row deleted")
	}
	return nil
}
