package redis

import (
	"CQRS-simple/pkg/models"
	"encoding/json"
	"github.com/alicebob/miniredis"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"log"
	"testing"
)

// newTestRedis returns a redis.Cmdable.
func newTestRedis() *RedisDB {
	mr, err := miniredis.Run()
	if err != nil {
		panic(err)
	}

	client := NewRedisDB(
		mr.Addr(),
		"0",
	)

	return client
}

func TestRedisDB_CreateUser(t *testing.T) {
	r := newTestRedis()
	u := models.User{Name: "John", Age: 26}
	read := models.Read{User: u}

	tests := []struct {
		name  string
		r     *RedisDB
		param models.Read
		want  models.User
	}{
		{
			name:  "Everything good",
			r:     r,
			param: read,
			want:  u,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := r.CreateUser(tc.param)
			if err != nil {
				t.Errorf("error = %v", err.Error())
				return
			}
			assert.Equal(t, tc.want.Name, got.Name)
			assert.Equal(t, tc.want.Age, got.Age)
		})
	}
}

func TestRedisDB_CreatePost(t *testing.T) {
	r := newTestRedis()
	u := models.User{ID: "00000000-0000-0000-0000-000000000000", Name: "John", Age: 26}
	read := models.Read{User: u}
	jr, err := json.Marshal(read)
	if err != nil {
		t.Errorf("error = %v", err.Error())
	}
	err = r.Client.Set(read.User.ID, jr, 0).Err()
	if err != nil {
		t.Errorf("error = %v", err.Error())
	}
	post := models.Post{
		UserID:  "00000000-0000-0000-0000-000000000000",
		Title:   "asd222",
		Message: "dsa222",
	}
	post2 := models.Post{
		UserID:  "00000000-0000-0000-0000-000000000001",
		Title:   "asd222",
		Message: "dsa222",
	}

	postRead := models.PostRead{
		Title:   post.Title,
		Message: post.Message,
	}

	tests := []struct {
		name    string
		r       *RedisDB
		param   models.Post
		want    models.PostRead
		wantErr string
	}{
		{
			name:  "Everything good",
			r:     r,
			param: post,
			want:  postRead,
		},
		{
			name:    "Wrong userID",
			r:       r,
			param:   post2,
			wantErr: "redis internal problem",
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := r.CreatePost(tc.param)
			if err != nil {
				if error.Error(err) == tc.wantErr {
					require.Error(t, err)
					return
				} else {
					t.Errorf("error = %v", err.Error())
					return
				}
			}
			assert.Equal(t, tc.want.Title, got.Title)
			assert.Equal(t, tc.want.Message, got.Message)
		})
	}
}

func TestRedisDB_GetAllUsers(t *testing.T) {
	r := newTestRedis()
	u1 := models.User{ID: "00000000-0000-0000-0000-000000000000", Name: "John", Age: 26}
	read := models.Read{User: u1}
	jr, err := json.Marshal(read)
	if err != nil {
		t.Errorf("error = %v", err.Error())
	}
	err = r.Client.Set(read.User.ID, jr, 0).Err()
	if err != nil {
		t.Errorf("error = %v", err.Error())
	}
	u2 := models.User{ID: "00000000-0000-0000-0000-000000000001", Name: "John", Age: 26}
	read2 := models.Read{User: u2}
	jr2, err := json.Marshal(read2)
	if err != nil {
		t.Errorf("error = %v", err.Error())
	}
	err = r.Client.Set(read2.User.ID, jr2, 0).Err()
	if err != nil {
		t.Errorf("error = %v", err.Error())
	}
	users := []models.User{
		u1,
		u2,
	}

	tests := []struct {
		name    string
		r       *RedisDB
		want    *[]models.User
		wantErr string
	}{
		{
			name: "Everything good",
			r:    r,
			want: &users,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := tc.r.GetAllUsers()
			log.Println(err)
			if err != nil {
				if error.Error(err) == tc.wantErr {
					log.Println(err)
					require.Error(t, err)
					return
				} else {
					t.Errorf("error = %v", err.Error())
					return
				}
			}
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestRedisDB_GetUserPosts(t *testing.T) {
	r := newTestRedis()
	u := models.User{ID: "00000000-0000-0000-0000-000000000000", Name: "John", Age: 26}
	post := models.PostRead{
		ID:      "00000000-0000-0000-0000-000000000000",
		Title:   "asd222",
		Message: "dsa222",
	}

	post2 := models.PostRead{
		ID:      "00000000-0000-0000-0000-000000000000",
		Title:   "asd222",
		Message: "dsa222",
	}

	posts := []models.PostRead{
		post,
		post2,
	}
	userPosts := models.UserPosts{
		User:  u,
		Posts: posts,
	}
	jrRes, err := json.Marshal(userPosts)
	if err != nil {
		t.Errorf("error = %v", err.Error())
	}
	err = r.Client.Set(userPosts.User.ID, jrRes, 0).Err()
	if err != nil {
		t.Errorf("error = %v", err.Error())
	}

	tests := []struct {
		name    string
		r       *RedisDB
		param   string
		want    models.UserPosts
		wantErr string
	}{
		{
			name:  "Everything good",
			r:     r,
			param: u.ID,
			want:  userPosts,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := tc.r.GetUserPosts(tc.param)
			log.Println(err)
			if err != nil {
				if error.Error(err) == tc.wantErr {
					log.Println(err)
					require.Error(t, err)
					return
				} else {
					t.Errorf("error = %v", err.Error())
					return
				}
			}
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestRedisDB_UpdateUser(t *testing.T) {
	r := newTestRedis()
	u := models.User{ID: "00000000-0000-0000-0000-000000000000", Name: "Ash", Age: 26}
	p := models.PostRead{ID: "00000000-0000-0000-0000-000000000000", Title: "John", Message: "26"}
	posts := []models.PostRead{
		p,
	}
	userPost := models.UserPosts{
		User:  u,
		Posts: posts,
	}
	jrRes, err := json.Marshal(userPost)
	if err != nil {
		t.Errorf("error = %v", err.Error())
	}
	err = r.Client.Set(userPost.User.ID, jrRes, 0).Err()
	if err != nil {
		t.Errorf("error = %v", err.Error())
	}
	//u := models.User{ID: "00000000-0000-0000-0000-000000000000", Name: "John", Age: 26}

	tests := []struct {
		name    string
		r       *RedisDB
		param   models.User
		wantErr error
	}{
		{
			name:    "Everything good",
			r:       r,
			param:   u,
			wantErr: nil,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.r.UpdateUser(tc.param)
			log.Println(err)
			if err != nil {
				if err == tc.wantErr {
					log.Println(err)
					require.Error(t, err)
					return
				} else {
					t.Errorf("error = %v", err.Error())
					return
				}
			}
			assert.Equal(t, tc.wantErr, err)
		})
	}
}

func TestRedisDB_UpdatePost(t *testing.T) {
	r := newTestRedis()
	u := models.User{ID: "00000000-0000-0000-0000-000000000000", Name: "Ash", Age: 26}
	p := models.PostRead{ID: "00000000-0000-0000-0000-000000000000", Title: "John", Message: "26"}
	posts := []models.PostRead{
		p,
	}
	userPost := models.UserPosts{
		User:  u,
		Posts: posts,
	}
	jrRes, err := json.Marshal(userPost)
	if err != nil {
		t.Errorf("error = %v", err.Error())
	}
	err = r.Client.Set(userPost.User.ID, jrRes, 0).Err()
	if err != nil {
		t.Errorf("error = %v", err.Error())
	}
	pp := models.Post{ID: "00000000-0000-0000-0000-000000000000", UserID: "00000000-0000-0000-0000-000000000000", Title: "John", Message: "26"}

	tests := []struct {
		name    string
		r       *RedisDB
		param   models.Post
		wantErr error
	}{
		{
			name:    "Everything good",
			r:       r,
			param:   pp,
			wantErr: nil,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.r.UpdatePost(tc.param)
			log.Println(err)
			if err != nil {
				if err == tc.wantErr {
					log.Println(err)
					require.Error(t, err)
					return
				} else {
					t.Errorf("error = %v", err.Error())
					return
				}
			}
			assert.Equal(t, tc.wantErr, err)
		})
	}
}

func TestRedisDB_DeleteUser(t *testing.T) {
	r := newTestRedis()
	u := models.User{ID: "00000000-0000-0000-0000-000000000000", Name: "Ash", Age: 26}
	userPost := models.UserPosts{
		User: u,
	}
	jrRes, err := json.Marshal(userPost)
	if err != nil {
		t.Errorf("error = %v", err.Error())
	}
	err = r.Client.Set(userPost.User.ID, jrRes, 0).Err()
	if err != nil {
		t.Errorf("error = %v", err.Error())
	}

	tests := []struct {
		name    string
		r       *RedisDB
		param   string
		wantErr error
	}{
		{
			name:    "Everything good",
			r:       r,
			param:   u.ID,
			wantErr: nil,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.r.DeleteUser(tc.param)
			log.Println(err)
			if err != nil {
				if err == tc.wantErr {
					log.Println(err)
					require.Error(t, err)
					return
				} else {
					t.Errorf("error = %v", err.Error())
					return
				}
			}
			assert.Equal(t, tc.wantErr, err)
		})
	}
}

func TestRedisDB_DeletePost(t *testing.T) {
	r := newTestRedis()
	u := models.User{ID: "00000000-0000-0000-0000-000000000000", Name: "Ash", Age: 26}
	p := models.PostRead{ID: "00000000-0000-0000-0000-000000000000", Title: "John", Message: "26"}
	posts := []models.PostRead{
		p,
	}
	userPost := models.UserPosts{
		User:  u,
		Posts: posts,
	}
	jrRes, err := json.Marshal(userPost)
	if err != nil {
		t.Errorf("error = %v", err.Error())
	}
	err = r.Client.Set(userPost.User.ID, jrRes, 0).Err()
	if err != nil {
		t.Errorf("error = %v", err.Error())
	}

	tests := []struct {
		name    string
		r       *RedisDB
		param1  string
		param2  string
		wantErr error
	}{
		{
			name:    "Everything good",
			r:       r,
			param1:  p.ID,
			param2:  u.ID,
			wantErr: nil,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.r.DeletePost(tc.param1, tc.param2)
			log.Println(err)
			if err != nil {
				if err == tc.wantErr {
					log.Println(err)
					require.Error(t, err)
					return
				} else {
					t.Errorf("error = %v", err.Error())
					return
				}
			}
			assert.Equal(t, tc.wantErr, err)
		})
	}
}

func TestRedisDB_Exist(t *testing.T) {
	r := newTestRedis()
	u := models.User{ID: "00000000-0000-0000-0000-000000000000", Name: "Ash", Age: 26}
	userPost := models.UserPosts{
		User: u,
	}
	jrRes, err := json.Marshal(userPost)
	if err != nil {
		t.Errorf("error = %v", err.Error())
	}
	err = r.Client.Set(userPost.User.ID, jrRes, 0).Err()
	if err != nil {
		t.Errorf("error = %v", err.Error())
	}
	tests := []struct {
		name  string
		r     *RedisDB
		param string
		want  bool
	}{
		{
			name:  "Everything good",
			r:     r,
			param: u.ID,
			want:  true,
		},
		{
			name:  "Everything good",
			r:     r,
			param: "00000000-0000-0000-0000-000000000001",
			want:  false,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			b := tc.r.Exist(tc.param)
			assert.Equal(t, tc.want, b)
		})
	}

}
