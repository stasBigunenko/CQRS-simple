package postgreSQL

import (
	"CQRS-simple/pkg/models"
	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"
	"log"
	"testing"
)

//func TestPostgresDB_CreateUser2(t *testing.T) {
//	// build DB Client mock
//	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
//	if err != nil {
//		t.Fatalf("an error '%s' was not expected when opening a stub database connection: %s", err, mock)
//	}
//	defer db.Close()
//
//	//mockedRow := sqlmock.NewRows([]string{"userID", "name", "age"}).AddRow("00000000-0000-0000-0000-000000000000", "name1", 10)
//
//	query := "INSERT INTO users \\(userID, name, age\\) VALUES \\(\\?, \\?, \\?\\)"
//	prep := mock.ExpectPrepare(query)
//	prep.ExpectExec().WithArgs("00000000-0000-0000-0000-000000000000", "name1", 10).
//		WillReturnResult(sqlmock.NewResult(0, 1))
//
//	// describe expected behaviour
//	//mock.ExpectBegin()
//	//mock.ExpectExec(`INSERT INTO users (userID, name, age) VALUES ($1, $2, $3)`).
//	//	WithArgs("00000000-0000-0000-0000-000000000000", "name1", 10).
//	//	WillReturnResult(sqlmock.NewResult(0, 1))
//	//mock.ExpectCommit()
//
//	// make storage for test
//	postgreSQL := &PostgresDB{Pdb: db}
//
//	err = postgreSQL.Pdb.Ping()
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	u := models.User{Name: "name1", Age: 10}
//
//	res ,err := postgreSQL.CreateUser(u)
//
//	require.NoError(t, err)
//	require.NotNil(t, res)
//	//assert.Equal(t, u.Name, res.Name)
//	//assert.Equal(t, u.Age, res.Age)
//}

func TestPostgresDB_GetPosts2(t *testing.T) {
	// build DB Client mock
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection: %s", err, mock)
	}
	defer db.Close()

	mock.ExpectQuery(`SELECT postID, title, message FROM posts WHERE userID=$1`).
		WithArgs("00000000-0000-0000-0000-000000000000").
		WillReturnRows(
			mock.
				NewRows([]string{"postID", "title", "message"}).
				AddRow(
					"00000000-0000-0000-0000-000000000000", "title1", "message1"),
		)

	postgreSQL := &PostgresDB{Pdb: db}

	res, err := postgreSQL.GetPosts("00000000-0000-0000-0000-000000000000")

	exp := []models.PostRead{
		{"00000000-0000-0000-0000-000000000000", "title1", "message1"},
	}

	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, exp, res)
}

func TestPostgresDB_GetAllUsers2(t *testing.T) {
	// build DB Client mock
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection: %s", err, mock)
	}
	defer db.Close()

	mock.ExpectQuery(`SELECT userID, name, age FROM users`).
		WillReturnRows(mock.
			NewRows([]string{"userID", "name", "age"}).
			AddRow("00000000-0000-0000-0000-000000000000", "name1", 10),
		)

	postgreSQL := &PostgresDB{Pdb: db}

	res, err := postgreSQL.GetAllUsers()

	exp := []models.User{
		{"00000000-0000-0000-0000-000000000000", "name1", 10},
	}

	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, &exp, res)
}

func TestPostgresDB_GetPostRead2(t *testing.T) {
	// build DB Client mock
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection: %s", err, mock)
	}
	defer db.Close()

	mock.ExpectQuery(`SELECT postID, userID, title, message FROM posts WHERE postID=$1`).
		WithArgs("00000000-0000-0000-0000-000000000000").
		WillReturnRows(mock.
			NewRows([]string{"postID", "userID", "title", "message"}).
			AddRow("00000000-0000-0000-0000-000000000000", "00000000-0000-0000-0000-000000000000", "title1", "message1"),
		)

	postgreSQL := &PostgresDB{Pdb: db}

	res, err := postgreSQL.GetPostRead("00000000-0000-0000-0000-000000000000")

	exp := models.Read{
		User:     models.User{ID: "00000000-0000-0000-0000-000000000000"},
		PostRead: models.PostRead{ID: "00000000-0000-0000-0000-000000000000", Title: "title1", Message: "message1"},
	}

	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, exp, res)
}

func TestPostgresDB_GetUser2(t *testing.T) {
	// build DB Client mock
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection: %s", err, mock)
	}
	defer db.Close()

	mock.ExpectQuery(`SELECT userID, name, age FROM users WHERE userID=$1`).
		WithArgs("00000000-0000-0000-0000-000000000000").
		WillReturnRows(mock.
			NewRows([]string{"userID", "name", "age"}).
			AddRow("00000000-0000-0000-0000-000000000000", "name1", 10),
		)

	postgreSQL := &PostgresDB{Pdb: db}

	res, err := postgreSQL.GetUser("00000000-0000-0000-0000-000000000000")

	exp := models.Read{
		User: models.User{"00000000-0000-0000-0000-000000000000", "name1", 10},
	}

	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, exp, res)
}

func TestPostgresDB_UpdateUser2(t *testing.T) {
	// build DB Client mock
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection: %s", err, mock)
	}
	defer db.Close()

	exp := models.User{"00000000-0000-0000-0000-000000000000", "name1", 10}

	mock.ExpectQuery(`SELECT name, age FROM users WHERE userID=$1`).
		WithArgs("00000000-0000-0000-0000-000000000000").
		WillReturnRows(mock.
			NewRows([]string{"name", "age"}).
			AddRow("name1", 10),
		)
	mock.ExpectExec(`UPDATE users SET name=$1, age=$2 WHERE userID=$3`).
		WithArgs("name1", 10, "00000000-0000-0000-0000-000000000000").
		WillReturnResult(sqlmock.NewResult(0, 1))

	postgreSQL := &PostgresDB{Pdb: db}

	res, err := postgreSQL.UpdateUser(exp)

	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, exp, res)
}

func TestPostgresDB_UpdatePost2(t *testing.T) {
	// build DB Client mock
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection: %s", err, mock)
	}
	defer db.Close()

	exp := models.Post{"00000000-0000-0000-0000-000000000000", "00000000-0000-0000-0000-000000000000", "title1", "message1"}

	mock.ExpectQuery(`SELECT userID, title, message FROM posts WHERE postID=$1`).
		WithArgs("00000000-0000-0000-0000-000000000000").
		WillReturnRows(mock.
			NewRows([]string{"userID", "title", "message"}).
			AddRow("00000000-0000-0000-0000-000000000000", "title1", "message1"),
		)
	mock.ExpectExec(`UPDATE posts SET title=$1, message=$2 WHERE postID=$3`).
		WithArgs("title1", "message1", "00000000-0000-0000-0000-000000000000").
		WillReturnResult(sqlmock.NewResult(0, 1))

	postgreSQL := &PostgresDB{Pdb: db}

	res, err := postgreSQL.UpdatePost(exp)
	log.Println(res)
	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, exp, res)
}

func TestPostgresDB_DeleteUser2(t *testing.T) {
	// build DB Client mock
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection: %s", err, mock)
	}
	defer db.Close()

	mock.ExpectExec(`DELETE FROM users where userID = $1`).
		WithArgs("00000000-0000-0000-0000-000000000000").
		WillReturnResult(sqlmock.NewResult(0, 1))

	postgreSQL := &PostgresDB{Pdb: db}

	err = postgreSQL.DeleteUser("00000000-0000-0000-0000-000000000000")

	require.NoError(t, err)
}

func TestPostgresDB_DeletePost2(t *testing.T) {
	// build DB Client mock
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection: %s", err, mock)
	}
	defer db.Close()

	mock.ExpectExec(`DELETE FROM posts where postID = $1`).
		WithArgs("00000000-0000-0000-0000-000000000000").
		WillReturnResult(sqlmock.NewResult(0, 1))

	postgreSQL := &PostgresDB{Pdb: db}

	err = postgreSQL.DeletePost("00000000-0000-0000-0000-000000000000")

	require.NoError(t, err)
}
