package readServ

import (
	mock2 "CQRS-simple/pkg/mock"
	"CQRS-simple/pkg/models"
	"CQRS-simple/pkg/storage/postgreSQL"
	"CQRS-simple/pkg/storage/redis"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestReadServ_GetAllUsers(t *testing.T) {
	postgeSQLDB := new(mock2.DBInterface)
	redisDB := new(mock2.RedisDBInterface)

	posts := []models.User{
		{"00000000-0000-0000-0000-000000000000", "Stas", 12},
		{"00000000-0000-0000-0000-000000000001", "Stas2", 23},
	}

	p11 := models.PostRead{"00000000-0000-0000-0000-000000000000", "raz", "dva"}
	p22 := models.PostRead{"00000000-0000-0000-0000-000000000001", "raz1", "dv1a"}
	postsRead := []models.PostRead{
		p11,
		p22,
	}

	postgeSQLDB.On("GetAllUsers").Return(&posts, nil)
	redisDB.On("CreateUser", mock.Anything).Return(models.User{}, nil)
	postgeSQLDB.On("GetPosts", mock.Anything).Return(postsRead, nil)
	redisDB.On("CreatePost", mock.Anything).Return(models.PostRead{}, nil)

	postgeSQLDB2 := new(mock2.DBInterface)
	redisDB2 := new(mock2.RedisDBInterface)
	postgeSQLDB2.On("GetAllUsers").Return(&[]models.User{}, errors.New("db problems"))

	tests := []struct {
		name    string
		pdb     postgreSQL.DBInterface
		rdb     redis.RedisDBInterface
		want    *[]models.User
		wantErr string
	}{
		{
			name:    "Everything good",
			pdb:     postgeSQLDB,
			rdb:     redisDB,
			want:    &posts,
			wantErr: "",
		},
		{
			name:    "no users in postgreSQL",
			pdb:     postgeSQLDB2,
			rdb:     redisDB2,
			want:    &[]models.User{},
			wantErr: "db problems",
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			rs := NewReadServ(tc.pdb, tc.rdb)
			got, err := rs.GetAllUsers()
			if (err != nil) && (err.Error() != tc.wantErr) {
				t.Errorf("error = %v, wantErr = %v", err.Error(), tc.wantErr)
				return
			}
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestReadServ_UserPosts(t *testing.T) {
	postgeSQLDB := new(mock2.DBInterface)
	redisDB := new(mock2.RedisDBInterface)

	m := models.User{"00000000-0000-0000-0000-000000000000", "Stas", 12}
	p1 := models.PostRead{"00000000-0000-0000-0000-000000000000", "raz", "dva"}
	p2 := models.PostRead{"00000000-0000-0000-0000-000000000001", "raz1", "dv1a"}
	posts := []models.PostRead{
		p1,
		p2,
	}
	up := models.UserPosts{
		User:  m,
		Posts: posts,
	}
	redisDB.On("GetUserPosts", m.ID).Return(up, nil)

	postgeSQLDB2 := new(mock2.DBInterface)
	redisDB2 := new(mock2.RedisDBInterface)

	mm := models.Read{User: m}
	redisDB2.On("GetUserPosts", m.ID).Return(models.UserPosts{}, errors.New("no such user"))
	postgeSQLDB2.On("GetUser", m.ID).Return(mm, nil)
	redisDB2.On("CreateUser", mm).Return(m, nil)
	postgeSQLDB2.On("GetPosts", m.ID).Return(posts, nil)
	redisDB2.On("CreatePost", mock.Anything).Return(models.PostRead{}, nil)
	posts2 := []models.PostRead{models.PostRead{ID: "", Title: "", Message: ""}, models.PostRead{ID: "", Title: "", Message: ""}}
	up2 := models.UserPosts{
		User:  m,
		Posts: posts2,
	}
	tests := []struct {
		name    string
		pdb     postgreSQL.DBInterface
		rdb     redis.RedisDBInterface
		param   string
		want    models.UserPosts
		wantErr error
	}{
		{
			name:    "Everything good",
			pdb:     postgeSQLDB,
			rdb:     redisDB,
			param:   m.ID,
			want:    up,
			wantErr: nil,
		},
		{
			name:    "User not in redis DB",
			pdb:     postgeSQLDB2,
			rdb:     redisDB2,
			param:   m.ID,
			want:    up2,
			wantErr: nil,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			rs := NewReadServ(tc.pdb, tc.rdb)
			got, err := rs.UserPosts(tc.param)
			if (err != nil) && err != tc.wantErr {
				t.Errorf("error = %v, wantErr %v", err.Error(), tc.wantErr)
				return
			}
			assert.Equal(t, tc.want, got)
		})
	}
}
