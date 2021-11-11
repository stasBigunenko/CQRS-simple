package postgreSQL

import (
	"CQRS-simple/pkg/models"
	"database/sql"
	"errors"
	"fmt"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"sync"
)

type PostgresDB struct {
	Pdb *sql.DB
	mu  sync.RWMutex
}

func NewPDB(host string, port string, user string, psw string, dbname string, ssl string) (*PostgresDB, error) {
	connStr := "host=" + host + " port=" + port + " user=" + user + " password=" + psw + " dbname=" + dbname + " sslmode=" + ssl

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to connect database %w\n", err)
	}

	database := &PostgresDB{Pdb: db}

	database.Pdb.Exec("CREATE TABLE users (\n    id VARCHAR(40) PRIMARY KEY NOT NULL,\n    name VARCHAR(50) NOT NULL,\n    surname VARCHAR(150) NOT NULL\n, age INT,\n sex VARCHAR(10));")

	return database, nil
}

func (pdb *PostgresDB) Create(u models.User) (models.User, error) {
	pdb.mu.Lock()
	defer pdb.mu.Unlock()

	id := uuid.New()
	idStr := id.String()

	_, err := pdb.Pdb.Exec(
		"INSERT INTO users (id, name, surname, age, sex) VALUES ($1, $2, $3, $4, $5)", idStr, u.Name, u.Surname, u.Age, u.Sex)
	if err != nil {
		return models.User{}, errors.New("couldn't create user in database")
	}

	u.ID = idStr

	return u, nil
}

func (pdb *PostgresDB) Get(id string) (models.User, error) {

	var u models.User

	u.ID = id

	err := pdb.Pdb.QueryRow(
		`SELECT name, surname, age, sex FROM users WHERE id=$1`, u.ID).Scan(&u.Name, &u.Surname, &u.Age, &u.Sex)
	if err != nil {
		return models.User{}, errors.New("user doesn't exist")
	}

	return u, nil
}
