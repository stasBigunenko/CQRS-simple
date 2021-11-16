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
	//database.Pdb.Exec("CREATE TABLE read (\n    userID VARCHAR(40) NOT NULL,\n    name VARCHAR(50) NOT NULL,\n  age INT,\n postID VARCHAR(40)\n, title VARCHAR(50) NOT NULL\n, message VARCHAR(155));")

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

//
//func (pdb *PostgresDB) Get(id string) (models.Read, error) {
//
//	var r models.Read
//
//	err := pdb.Pdb.QueryRow(
//		`SELECT userID, name, age, postID, title, message FROM read WHERE userID=$1`, id).Scan(&r.User.ID, &r.User.Name, &r.User.Age, &r.PostRead.ID, &r.PostRead.Title, &r.PostRead.Message)
//	if err != nil {
//		return models.Read{}, errors.New("user doesn't exist")
//	}
//
//	return r, nil
//}

func (pdb *PostgresDB) CreateReadInfo(res models.Read) error {

	//res.ID = uuid.New().String()

	_, err := pdb.Pdb.Exec(
		"INSERT INTO read (userID, name, age, postID, title, message) VALUES ($1, $2, $3, $4, $5, $6)", res.User.ID, res.User.Name, res.User.Age, res.PostRead.ID, res.PostRead.Title, res.PostRead.Message)
	if err != nil {
		return errors.New("couldn't create user in database")
	}

	return nil
}

func (pdb *PostgresDB) GetUserRead(id string) (models.Read, error) {

	var u models.Read

	err := pdb.Pdb.QueryRow(
		`SELECT userID, name, age FROM read WHERE userID=$1`, id).Scan(&u.User.ID, &u.User.Name, &u.User.Age)
	if err != nil {
		return models.Read{}, errors.New("user doesn't exist")
	}

	return u, nil
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

func (pdb *PostgresDB) GetAllUsers() (*[]models.User, error) {
	pdb.mu.Lock()
	defer pdb.mu.Unlock()

	rows, err := pdb.Pdb.Query(
		`SELECT userID, name, age FROM users`)
	if err != nil {
		return &[]models.User{}, errors.New("db problems")
	}
	defer rows.Close()

	users := []models.User{}

	for rows.Next() {
		u := models.User{}
		err = rows.Scan(&u.ID, &u.Name, &u.Age)
		users = append(users, u)
	}

	return &users, nil
}

func (pdb *PostgresDB) GetPostRead(id string) (models.Read, error) {

	var p models.Read

	err := pdb.Pdb.QueryRow(
		`SELECT postID, userID, title, message FROM read WHERE postID=$1`, id).Scan(&p.PostRead.ID, &p.User.ID, &p.PostRead.Title, &p.PostRead.Message)
	if err != nil {
		return models.Read{}, errors.New("user doesn't exist")
	}

	return p, nil
}

func (pdb *PostgresDB) UpdateUser(u models.User) (models.User, error) {
	pdb.mu.Lock()
	defer pdb.mu.Unlock()

	var oldUser models.User

	err := pdb.Pdb.QueryRow(
		`SELECT name, age FROM users WHERE userID=$1`, u.ID).Scan(&oldUser.Name, &oldUser.Age)
	if err != nil {
		return models.User{}, errors.New("couldn't find post")
	}

	if u.Name == "" {
		u.Name = oldUser.Name
	}
	if u.Age == 0 {
		u.Age = oldUser.Age
	}

	_, err = pdb.Pdb.Exec(
		`UPDATE users SET name=$1, age=$2 WHERE userID=$3`, u.Name, u.Age, u.ID)
	if err != nil {
		return models.User{}, errors.New("couldn't update post")
	}

	return u, nil
}

func (pdb *PostgresDB) UpdatePost(p models.Post) (models.Post, error) {
	pdb.mu.Lock()
	defer pdb.mu.Unlock()

	var oldPost models.Post

	err := pdb.Pdb.QueryRow(
		`SELECT userID, title, message FROM posts WHERE postID=$1`, p.ID).Scan(&oldPost.UserID, &oldPost.Title, &oldPost.Message)
	if err != nil {
		return models.Post{}, errors.New("couldn't find post")
	}

	if p.UserID == "" {
		p.UserID = oldPost.UserID
	}
	if p.Title == "" {
		p.Title = oldPost.Title
	}
	if p.Message == "" {
		p.Message = oldPost.Message
	}

	_, err = pdb.Pdb.Exec(
		`UPDATE posts SET title=$1, message=$2 WHERE postID=$3`, p.Title, p.Message, p.ID)
	if err != nil {
		return models.Post{}, errors.New("couldn't update post")
	}

	return p, nil
}

func (pdb *PostgresDB) DeleteUser(id string) error {
	pdb.mu.Lock()
	defer pdb.mu.Unlock()

	_, err := pdb.Pdb.Exec(
		`DELETE FROM users where userID = $1`, id)
	if err != nil {
		return errors.New("couldn't delete post")
	}

	return nil
}

func (pdb *PostgresDB) DeletePost(id string) error {
	pdb.mu.Lock()
	defer pdb.mu.Unlock()

	_, err := pdb.Pdb.Exec(
		`DELETE FROM posts where postID = $1`, id)
	if err != nil {
		return errors.New("couldn't delete post")
	}

	return nil
}

func (pdb *PostgresDB) DeleteReadUser(id string) error {
	pdb.mu.Lock()
	defer pdb.mu.Unlock()

	_, err := pdb.Pdb.Exec(
		`DELETE FROM read where userID = $1`, id)
	if err != nil {
		return errors.New("couldn't delete post")
	}

	return nil
}

func (pdb *PostgresDB) DeleteReadPost(id string) error {
	pdb.mu.Lock()
	defer pdb.mu.Unlock()

	_, err := pdb.Pdb.Exec(
		`DELETE FROM read where postID = $1`, id)
	if err != nil {
		return errors.New("couldn't delete post")
	}

	return nil
}

func (pdb *PostgresDB) UpdateReadUser(u models.User) error {

	pdb.mu.Lock()
	defer pdb.mu.Unlock()

	rows, err := pdb.Pdb.Query(
		`SELECT userID, name, age FROM read WHERE userID=$1`, u.ID)
	if err != nil {
		return errors.New("db problems")
	}
	defer rows.Close()

	for rows.Next() {
		_, err = pdb.Pdb.Exec(
			`UPDATE read SET name=$1, age=$2 WHERE userID=$3`, u.Name, u.Age, u.ID)
		if err != nil {
			return errors.New("couldn't update post")
		}
	}

	return nil
}

func (pdb *PostgresDB) UpdateReadPost(p models.Post) error {

	pdb.mu.Lock()
	defer pdb.mu.Unlock()

	_, err := pdb.Pdb.Exec(
		`UPDATE read SET title=$1, message=$2 WHERE postID=$3`, p.Title, p.Message, p.ID)
	if err != nil {
		return errors.New("couldn't update post")
	}

	return nil
}

func (pdb *PostgresDB) GetUser(id string) (models.Read, error) {

	var u models.Read

	err := pdb.Pdb.QueryRow(
		`SELECT userID, name, age FROM users WHERE userID=$1`, id).Scan(&u.User.ID, &u.User.Name, &u.User.Age)
	if err != nil {
		return models.Read{}, errors.New("user doesn't exist")
	}

	return u, nil
}
