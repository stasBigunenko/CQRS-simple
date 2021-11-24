package postgreSQL

//
//import (
//	"CQRS-simple/cmd/http/myConfig"
//	"CQRS-simple/pkg/models"
//	"github.com/stretchr/testify/assert"
//	"log"
//	"testing"
//)
//
//var testUser = models.User{
//	Name: "user",
//	Age:  12,
//}
//
//var testPost = models.Post{
//	Title:   "asd",
//	Message: "dsa",
//}
//
//func TestPostgresDB_CreateUser(t *testing.T) {
//	config := myConfig.SetConfig()
//	db, err := NewPDB(config.PostgresHost, config.PostgresPort, config.PostgresUser, config.PostgresPsw, config.PostgresDB, config.PostgresSSL)
//	if err != nil {
//		if err != nil {
//			log.Fatalf("failed to connect postgreSQL: %s", err)
//		}
//	}
//
//	err = db.Pdb.Ping()
//	log.Println(err)
//
//	got, err := db.CreateUser(testUser)
//	if err != nil {
//		t.Errorf("error = %v", err.Error())
//		return
//	}
//
//	testUser.ID = got.ID
//	assert.Equal(t, testUser, got)
//}
//
//func TestPostgresDB_CreatePost(t *testing.T) {
//	config := myConfig.SetConfig()
//
//	db, err := NewPDB(config.PostgresHost, config.PostgresPort, config.PostgresUser, config.PostgresPsw, config.PostgresDB, config.PostgresSSL)
//	if err != nil {
//		if err != nil {
//			log.Fatalf("failed to connect postgreSQL: %s", err)
//		}
//	}
//	tests := []struct {
//		name  string
//		param *models.Post
//		want  *models.Post
//	}{
//		{
//			name:  "everything ok",
//			param: &testPost,
//			want:  &testPost,
//		},
//	}
//	for _, tc := range tests {
//		tc.param.UserID = testUser.ID
//		t.Run(tc.name, func(t *testing.T) {
//			got, err := db.CreatePost(*tc.param)
//			if err != nil {
//				t.Errorf("error = %v", err.Error())
//				return
//			}
//			testPost.ID = got.ID
//			assert.Equal(t, tc.want, &got)
//		})
//	}
//}
//
//func TestPostgresDB_GetPosts(t *testing.T) {
//	config := myConfig.SetConfig()
//
//	db, err := NewPDB(config.PostgresHost, config.PostgresPort, config.PostgresUser, config.PostgresPsw, config.PostgresDB, config.PostgresSSL)
//	if err != nil {
//		if err != nil {
//			log.Fatalf("failed to connect postgreSQL: %s", err)
//		}
//	}
//
//	got, err := db.GetPosts(testUser.ID)
//	if err != nil {
//		t.Errorf("error = %v", err.Error())
//		return
//	}
//	assert.NotNil(t, got)
//}
//
//func TestPostgresDB_GetAllUsers(t *testing.T) {
//	config := myConfig.SetConfig()
//
//	db, err := NewPDB(config.PostgresHost, config.PostgresPort, config.PostgresUser, config.PostgresPsw, config.PostgresDB, config.PostgresSSL)
//	if err != nil {
//		if err != nil {
//			log.Fatalf("failed to connect postgreSQL: %s", err)
//		}
//	}
//
//	got, err := db.GetAllUsers()
//	if err != nil {
//		t.Errorf("error = %v", err.Error())
//		return
//	}
//	assert.NotNil(t, got)
//}
//
//func TestPostgresDB_GetPostRead(t *testing.T) {
//	config := myConfig.SetConfig()
//
//	db, err := NewPDB(config.PostgresHost, config.PostgresPort, config.PostgresUser, config.PostgresPsw, config.PostgresDB, config.PostgresSSL)
//	if err != nil {
//		if err != nil {
//			log.Fatalf("failed to connect postgreSQL: %s", err)
//		}
//	}
//
//	p := models.PostRead{
//		ID:      testPost.ID,
//		Title:   testPost.Title,
//		Message: testPost.Message,
//	}
//	u := models.User{
//		ID: testUser.ID,
//	}
//	want := models.Read{
//		User:     u,
//		PostRead: p,
//	}
//
//	got, err := db.GetPostRead(testPost.ID)
//	if err != nil {
//		t.Errorf("error = %v", err.Error())
//		return
//	}
//
//	assert.Equal(t, want, got)
//}
//
//func TestPostgresDB_UpdateUser(t *testing.T) {
//	config := myConfig.SetConfig()
//
//	db, err := NewPDB(config.PostgresHost, config.PostgresPort, config.PostgresUser, config.PostgresPsw, config.PostgresDB, config.PostgresSSL)
//	if err != nil {
//		if err != nil {
//			log.Fatalf("failed to connect postgreSQL: %s", err)
//		}
//	}
//
//	testUser.Name = "abdula"
//
//	got, err := db.UpdateUser(testUser)
//	if err != nil {
//		t.Errorf("error = %v", err.Error())
//		return
//	}
//
//	assert.Equal(t, testUser, got)
//}
//
//func TestPostgresDB_UpdatePost(t *testing.T) {
//	config := myConfig.SetConfig()
//
//	db, err := NewPDB(config.PostgresHost, config.PostgresPort, config.PostgresUser, config.PostgresPsw, config.PostgresDB, config.PostgresSSL)
//	if err != nil {
//		if err != nil {
//			log.Fatalf("failed to connect postgreSQL: %s", err)
//		}
//	}
//
//	testPost.Title = "title"
//
//	got, err := db.UpdatePost(testPost)
//	if err != nil {
//		t.Errorf("error = %v", err.Error())
//		return
//	}
//
//	assert.Equal(t, testPost, got)
//}
//
//func TestPostgresDB_GetUser(t *testing.T) {
//	config := myConfig.SetConfig()
//
//	db, err := NewPDB(config.PostgresHost, config.PostgresPort, config.PostgresUser, config.PostgresPsw, config.PostgresDB, config.PostgresSSL)
//	if err != nil {
//		if err != nil {
//			log.Fatalf("failed to connect postgreSQL: %s", err)
//		}
//	}
//
//	want := models.Read{
//		User: testUser,
//	}
//
//	got, err := db.GetUser(testUser.ID)
//	if err != nil {
//		t.Errorf("error = %v", err.Error())
//		return
//	}
//
//	assert.Equal(t, want, got)
//}
//
//func TestPostgresDB_DeletePost(t *testing.T) {
//	config := myConfig.SetConfig()
//
//	db, err := NewPDB(config.PostgresHost, config.PostgresPort, config.PostgresUser, config.PostgresPsw, config.PostgresDB, config.PostgresSSL)
//	if err != nil {
//		if err != nil {
//			log.Fatalf("failed to connect postgreSQL: %s", err)
//		}
//	}
//
//	err = db.DeletePost(testPost.ID)
//	if err != nil {
//		t.Errorf("error = %v", err.Error())
//		return
//	}
//
//	assert.Nil(t, err, "delete post err")
//}
//
//func TestPostgresDB_DeleteUser(t *testing.T) {
//	config := myConfig.SetConfig()
//
//	db, err := NewPDB(config.PostgresHost, config.PostgresPort, config.PostgresUser, config.PostgresPsw, config.PostgresDB, config.PostgresSSL)
//	if err != nil {
//		if err != nil {
//			log.Fatalf("failed to connect postgreSQL: %s", err)
//		}
//	}
//
//	err = db.DeleteUser(testUser.ID)
//	if err != nil {
//		t.Errorf("error = %v", err.Error())
//		return
//	}
//
//	assert.Nil(t, err, "delete post err")
//}
