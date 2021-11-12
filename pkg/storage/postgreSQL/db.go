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

	database.Pdb.Exec("CREATE TABLE users (\n    userID VARCHAR(40) PRIMARY KEY NOT NULL,\n    name VARCHAR(50) NOT NULL,\n  age INT);")
	database.Pdb.Exec("CREATE TABLE posts (\n    postID VARCHAR(40) PRIMARY KEY NOT NULL,\n    userID VARCHAR(50) NOT NULL,\n    title VARCHAR(50) NOT NULL\n, message VARCHAR(155));")
	database.Pdb.Exec("CREATE TABLE read (\n    userID VARCHAR(40) NOT NULL,\n    name VARCHAR(50) NOT NULL,\n  age INT,\n postID VARCHAR(40)\n, title VARCHAR(50) NOT NULL\n, message VARCHAR(155));")

	return database, nil
}

func (pdb *PostgresDB) CreateUser(u models.User) (models.User, error) {
	pdb.mu.Lock()
	defer pdb.mu.Unlock()

	id := uuid.New()
	idStr := id.String()

	_, err := pdb.Pdb.Exec(
		`INSERT INTO users (userID, name, age) VALUES ($1, $2, $3)`, idStr, u.Name, u.Age)
	if err != nil {
		return models.User{}, errors.New("couldn't create user in database")
	}

	u.ID = idStr

	return u, nil
}

func (pdb *PostgresDB) CreatePost(p models.Post) (models.Post, error) {
	pdb.mu.Lock()
	defer pdb.mu.Unlock()

	id := uuid.New()
	idStr := id.String()

	_, err := pdb.Pdb.Exec(
		"INSERT INTO posts (postID, userID, title, message) VALUES ($1, $2, $3, $4)", idStr, p.UserID, p.Title, p.Message)
	if err != nil {
		return models.Post{}, errors.New("couldn't create user in database")
	}

	p.ID = idStr

	return p, nil
}

func (pdb *PostgresDB) Get(id string) (models.Read, error) {

	var r models.Read

	err := pdb.Pdb.QueryRow(
		`SELECT userID, name, age, postID, title, message FROM read WHERE userID=$1`, id).Scan(&r.User.ID, &r.User.Name, &r.User.Age, &r.PostRead.ID, &r.PostRead.Title, &r.PostRead.Message)
	if err != nil {
		return models.Read{}, errors.New("user doesn't exist")
	}

	return r, nil
}

func (pdb *PostgresDB) CreateReadInfo(u models.User) error {
	var r models.Read

	r.User.ID = u.ID
	r.User.Name = u.Name
	r.User.Age = u.Age
	r.PostRead.ID = "empty"
	r.PostRead.Title = "empty"
	r.PostRead.Message = "empty"

	_, err := pdb.Pdb.Exec(
		"INSERT INTO read (userID, name, age, postID, title, message) VALUES ($1, $2, $3, $4, $5, $6)", r.User.ID, r.User.Name, r.User.Age, r.PostRead.ID, r.PostRead.Title, r.PostRead.Message)
	if err != nil {
		return errors.New("couldn't create user in database")
	}

	return nil
}

func (pdb *PostgresDB) GetUser(id string) (models.User, error) {

	var u models.User

	err := pdb.Pdb.QueryRow(
		`SELECT userID, name, age FROM users WHERE userID=$1`, id).Scan(&u.ID, &u.Name, &u.Age)
	if err != nil {
		return models.User{}, errors.New("user doesn't exist")
	}

	return u, nil
}

func (pdb *PostgresDB) AddPostToUserRead(r models.Read) error {
	_, err := pdb.Pdb.Exec(`UPDATE read SET userID=$1, name=$2, age=$3, postID=$4, title=$5, message=$6 WHERE userID=$7`, r.User.ID, r.User.Name, r.User.Age, r.PostRead.ID, r.PostRead.Title, r.PostRead.Message, r.User.ID)
	if err != nil {
		return errors.New("couldn't updateuserreadinfo")
	}

	return nil
}

func (pdb *PostgresDB) GetPosts(userID string) ([]models.PostRead, error) {
	pdb.mu.Lock()
	defer pdb.mu.Unlock()

	rows, err := pdb.Pdb.Query(
		`SELECT postID, title, message FROM posts WHERE userID=$1`, userID)
	if err != nil {
		return []models.PostRead{}, errors.New("db problems")
	}
	defer rows.Close()

	postsRead := []models.PostRead{}

	for rows.Next() {
		p := models.PostRead{}
		err = rows.Scan(&p.ID, &p.Title, &p.Message)
		postsRead = append(postsRead, p)
	}

	return postsRead, nil
}
