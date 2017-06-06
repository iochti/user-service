package main

import (
	"database/sql"
	"fmt"

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
	GetUserByGhubID(id int) (*models.User, error)
	GetUserByToken(token string) (*models.User, error)
	GetUserByLogin(login string) (*models.User, error)
	DeleteUser(id int) error
}

// PostgresDL implements the DataLayerInterface
type PostgresDL struct {
	DBName string
	Host   string
	Db     *sql.DB
}

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
	err := p.Db.QueryRow(`INSERT INTO users(name, login, avatar, ghubid, token)
		VALUES($1, $2, $3, $4, $5) RETURNING id;`,
		user.Name, user.Login, user.AvatarURL, user.GhubID, user.AuthToken).Scan(&userID)

	if err != nil {
		return err
	}
	user.ID = userID
	return nil
}

// GetUserByID fetch a user by its user id
func (p *PostgresDL) GetUserByID(id int) (*models.User, error) {
	// Check ID values
	if id <= 0 {
		return nil, fmt.Errorf("Error, invalid search ID: id must be > 0")
	}
	user := new(models.User)
	err := p.Db.QueryRow("SELECT id, name, login, avatar, ghubid, token FROM users WHERE id=$1;", id).Scan(
		&user.ID,
		&user.Name,
		&user.Login,
		&user.AvatarURL,
		&user.GhubID,
		&user.AuthToken,
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}

// GetUserByGhubID fetch a user by its github id
func (p *PostgresDL) GetUserByGhubID(id int) (*models.User, error) {
	// Check ID values
	if id <= 0 {
		return nil, fmt.Errorf("Error, invalid search ID: id must be > 0")
	}
	user := new(models.User)
	err := p.Db.QueryRow("SELECT id, name, login, avatar, ghubid, token FROM users WHERE ghubid=$1;", id).Scan(
		&user.ID,
		&user.Name,
		&user.Login,
		&user.AvatarURL,
		&user.GhubID,
		&user.AuthToken,
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}

// GetUserByToken fetch a user by its last known token
func (p *PostgresDL) GetUserByToken(token string) (*models.User, error) {
	// Check ID values
	if token == "" {
		return nil, fmt.Errorf("Error, invalid search: token should not be empty")
	}
	user := new(models.User)
	err := p.Db.QueryRow("SELECT id, name, login, avatar, ghubid, token FROM users WHERE token=$1;", token).Scan(
		&user.ID,
		&user.Name,
		&user.Login,
		&user.AvatarURL,
		&user.GhubID,
		&user.AuthToken,
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
	err := p.Db.QueryRow("SELECT id, name, login, avatar, ghubid, token FROM users WHERE login=$1;", login).Scan(
		&user.ID,
		&user.Name,
		&user.Login,
		&user.AvatarURL,
		&user.GhubID,
		&user.AuthToken,
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

	res, err := p.Db.Exec("DELETE FROM users WHERE id = $1;", id)
	if err != nil {
		return err
	}
	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return fmt.Errorf("Error: no row deleted !")
	}
	return nil
}
