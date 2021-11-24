package writeServ

import (
	mock2 "CQRS-simple/pkg/mock"
	"CQRS-simple/pkg/models"
	"CQRS-simple/pkg/rabbitMQ/createQueue"
	"CQRS-simple/pkg/storage/postgreSQL"
	"CQRS-simple/pkg/storage/redis"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestWriteServ_CreateUser(t *testing.T) {
	postgeSQLDB := new(mock2.DBInterface)
	redisDB := new(mock2.RedisDBInterface)
	cq := new(mock2.QueueCreateInterface)
	cq.On("QueueCreateWrite", mock.Anything).Return(nil)

	u := models.User{Name: "Stas", Age: 12}
	u1 := models.User{ID: "00000000-0000-0000-0000-000000000000", Name: "Stas", Age: 12}

	postgeSQLDB.On("CreateUser", u).Return(u1, nil)

	postgeSQLDB2 := new(mock2.DBInterface)
	postgeSQLDB2.On("CreateUser", mock.Anything).Return(models.User{}, errors.New("couldn't create user in database"))

	tests := []struct {
		name    string
		pdb     postgreSQL.DBInterface
		rdb     redis.RedisDBInterface
		crq     createQueue.QueueCreateInterface
		param   models.User
		want    *models.User
		wantErr string
	}{
		{
			name:    "Everything good",
			pdb:     postgeSQLDB,
			rdb:     redisDB,
			crq:     cq,
			param:   u,
			want:    &u1,
			wantErr: "",
		},
		{
			name:    "no users in postgreSQL",
			pdb:     postgeSQLDB2,
			rdb:     redisDB,
			crq:     cq,
			want:    nil,
			wantErr: "couldn't create user in database",
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ws := NewWriteServ(tc.pdb, tc.rdb, tc.crq)
			got, err := ws.CreateUser(tc.param)
			if (err != nil) && (err.Error() != tc.wantErr) {
				t.Errorf("error = %v, wantErr = %v", err.Error(), tc.wantErr)
				return
			}
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestWriteServ_CreatePost(t *testing.T) {
	postgeSQLDB := new(mock2.DBInterface)
	redisDB := new(mock2.RedisDBInterface)
	cq := new(mock2.QueueCreateInterface)
	cq.On("QueueCreateWrite", mock.Anything).Return(nil)

	p := models.Post{UserID: "00000000-0000-0000-0000-000000000000", Title: "asd", Message: "dsa"}
	p1 := models.Post{ID: "00000000-0000-0000-0000-000000000000", UserID: "00000000-0000-0000-0000-000000000000", Title: "asd", Message: "dsa"}
	postgeSQLDB.On("CreatePost", p).Return(p1, nil)
	redisDB.On("Exist", mock.Anything).Return(false)

	postgeSQLDB2 := new(mock2.DBInterface)
	postgeSQLDB2.On("CreatePost", models.Post{}).Return(models.Post{}, errors.New("couldn't create user in database"))

	tests := []struct {
		name    string
		pdb     postgreSQL.DBInterface
		rdb     redis.RedisDBInterface
		crq     createQueue.QueueCreateInterface
		param   models.Post
		want    *models.Post
		wantErr string
	}{
		{
			name:    "Everything good",
			pdb:     postgeSQLDB,
			rdb:     redisDB,
			crq:     cq,
			param:   p,
			want:    &p1,
			wantErr: "",
		},
		{
			name:    "no users in postgreSQL",
			pdb:     postgeSQLDB2,
			rdb:     redisDB,
			crq:     cq,
			param:   models.Post{},
			want:    nil,
			wantErr: "couldn't create user in database",
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ws := NewWriteServ(tc.pdb, tc.rdb, tc.crq)
			got, err := ws.CreatePost(tc.param)
			if (err != nil) && (err.Error() != tc.wantErr) {
				t.Errorf("error = %v, wantErr = %v", err.Error(), tc.wantErr)
				return
			}
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestWriteServ_UpdateUser(t *testing.T) {
	postgeSQLDB := new(mock2.DBInterface)
	redisDB := new(mock2.RedisDBInterface)
	cq := new(mock2.QueueCreateInterface)
	cq.On("QueueCreateWrite", mock.Anything).Return(nil)

	u1 := models.User{ID: "00000000-0000-0000-0000-000000000000", Name: "asd", Age: 12}
	postgeSQLDB.On("UpdateUser", u1).Return(u1, nil)
	redisDB.On("Exist", mock.Anything).Return(false)

	postgeSQLDB2 := new(mock2.DBInterface)
	postgeSQLDB2.On("UpdateUser", mock.Anything).Return(models.User{}, errors.New("couldn't find user"))

	tests := []struct {
		name    string
		pdb     postgreSQL.DBInterface
		rdb     redis.RedisDBInterface
		crq     createQueue.QueueCreateInterface
		param   models.User
		want    *models.User
		wantErr string
	}{
		{
			name:    "Everything good",
			pdb:     postgeSQLDB,
			rdb:     redisDB,
			crq:     cq,
			param:   u1,
			want:    &u1,
			wantErr: "",
		},
		{
			name:    "no users in postgreSQL",
			pdb:     postgeSQLDB2,
			rdb:     redisDB,
			crq:     cq,
			param:   models.User{},
			want:    &models.User{},
			wantErr: "couldn't find user",
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ws := NewWriteServ(tc.pdb, tc.rdb, tc.crq)
			got, err := ws.UpdateUser(tc.param)
			if (err != nil) && (err.Error() != tc.wantErr) {
				t.Errorf("error = %v, wantErr = %v", err.Error(), tc.wantErr)
				return
			}
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestWriteServ_UpdatePost(t *testing.T) {
	postgeSQLDB := new(mock2.DBInterface)
	redisDB := new(mock2.RedisDBInterface)
	cq := new(mock2.QueueCreateInterface)
	cq.On("QueueCreateWrite", mock.Anything).Return(nil)

	p1 := models.Post{ID: "00000000-0000-0000-0000-000000000000", UserID: "00000000-0000-0000-0000-000000000000", Title: "asd", Message: "dsa"}
	postgeSQLDB.On("UpdatePost", p1).Return(p1, nil)
	redisDB.On("Exist", mock.Anything).Return(false)

	postgeSQLDB2 := new(mock2.DBInterface)
	postgeSQLDB2.On("UpdatePost", mock.Anything).Return(models.Post{}, errors.New("couldn't find post"))

	tests := []struct {
		name    string
		pdb     postgreSQL.DBInterface
		rdb     redis.RedisDBInterface
		crq     createQueue.QueueCreateInterface
		param   models.Post
		want    *models.Post
		wantErr string
	}{
		{
			name:    "Everything good",
			pdb:     postgeSQLDB,
			rdb:     redisDB,
			crq:     cq,
			param:   p1,
			want:    &p1,
			wantErr: "",
		},
		{
			name:    "no users in postgreSQL",
			pdb:     postgeSQLDB2,
			rdb:     redisDB,
			crq:     cq,
			param:   models.Post{},
			want:    nil,
			wantErr: "couldn't find post",
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ws := NewWriteServ(tc.pdb, tc.rdb, tc.crq)
			got, err := ws.UpdatePost(tc.param)
			if (err != nil) && (err.Error() != tc.wantErr) {
				t.Errorf("error = %v, wantErr = %v", err.Error(), tc.wantErr)
				return
			}
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestWriteServ_DeleteUser(t *testing.T) {
	postgeSQLDB := new(mock2.DBInterface)
	redisDB := new(mock2.RedisDBInterface)
	cq := new(mock2.QueueCreateInterface)
	cq.On("QueueCreateWrite", mock.Anything).Return(nil)

	u1 := "00000000-0000-0000-0000-000000000000"
	postgeSQLDB.On("DeleteUser", u1).Return(nil)
	redisDB.On("Exist", mock.Anything).Return(false)

	postgeSQLDB2 := new(mock2.DBInterface)
	postgeSQLDB2.On("DeleteUser", mock.Anything).Return(errors.New("couldn't delete user"))

	tests := []struct {
		name    string
		pdb     postgreSQL.DBInterface
		rdb     redis.RedisDBInterface
		crq     createQueue.QueueCreateInterface
		param   string
		wantErr string
	}{
		{
			name:    "Everything good",
			pdb:     postgeSQLDB,
			rdb:     redisDB,
			crq:     cq,
			param:   u1,
			wantErr: "",
		},
		{
			name:    "no users in postgreSQL",
			pdb:     postgeSQLDB2,
			rdb:     redisDB,
			crq:     cq,
			param:   u1,
			wantErr: "couldn't delete user",
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ws := NewWriteServ(tc.pdb, tc.rdb, tc.crq)
			err := ws.DeleteUser(tc.param)
			if (err != nil) && (err.Error() != tc.wantErr) {
				t.Errorf("error = %v, wantErr = %v", err.Error(), tc.wantErr)
				return
			}
		})
	}
}

func TestWriteServ_DeletePost(t *testing.T) {
	postgeSQLDB := new(mock2.DBInterface)
	redisDB := new(mock2.RedisDBInterface)
	cq := new(mock2.QueueCreateInterface)
	cq.On("QueueCreateWrite", mock.Anything).Return(nil)

	p1 := "00000000-0000-0000-0000-000000000000"
	//p := models.Post{ID: "00000000-0000-0000-	0000-000000000000", UserID: "00000000-0000-0000-0000-000000000000", Title: "asd", Message: "dsa"}

	redisDB.On("Exist", mock.Anything).Return(false)
	pp := models.PostRead{ID: "00000000-0000-0000-0000-000000000000", Title: "asd", Message: "dsa"}
	ppp := models.Read{PostRead: pp}
	postgeSQLDB.On("GetPostRead", mock.Anything).Return(ppp, nil)
	postgeSQLDB.On("DeletePost", p1).Return(nil)

	postgeSQLDB2 := new(mock2.DBInterface)
	postgeSQLDB2.On("GetPostRead", mock.Anything).Return(models.Read{}, errors.New("user doesn't exist"))

	tests := []struct {
		name    string
		pdb     postgreSQL.DBInterface
		rdb     redis.RedisDBInterface
		crq     createQueue.QueueCreateInterface
		param   string
		want    *models.Post
		wantErr string
	}{
		{
			name:    "Everything good",
			pdb:     postgeSQLDB,
			rdb:     redisDB,
			crq:     cq,
			param:   p1,
			wantErr: "",
		},
		{
			name:    "no users in postgreSQL",
			pdb:     postgeSQLDB2,
			rdb:     redisDB,
			crq:     cq,
			param:   p1,
			want:    nil,
			wantErr: "user doesn't exist",
		},
		{
			name:    "incorrect id",
			pdb:     postgeSQLDB2,
			rdb:     redisDB,
			crq:     cq,
			param:   "00000000-0000-0000-0000-00000000000",
			want:    nil,
			wantErr: "service: couldn't parse id",
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ws := NewWriteServ(tc.pdb, tc.rdb, tc.crq)
			err := ws.DeletePost(tc.param)
			if (err != nil) && (err.Error() != tc.wantErr) {
				t.Errorf("error = %v, wantErr = %v", err.Error(), tc.wantErr)
				return
			}
		})
	}
}

func TestWriteServ_GetPost(t *testing.T) {
	postgeSQLDB := new(mock2.DBInterface)
	redisDB := new(mock2.RedisDBInterface)
	cq := new(mock2.QueueCreateInterface)
	cq.On("QueueCreateWrite", mock.Anything).Return(nil)

	p1 := "00000000-0000-0000-0000-000000000000"
	pp := models.PostRead{ID: "00000000-0000-0000-0000-000000000000", Title: "asd", Message: "dsa"}
	p := models.Read{PostRead: pp}
	post := models.Post{ID: "00000000-0000-0000-0000-000000000000", UserID: "", Title: "asd", Message: "dsa"}
	postgeSQLDB.On("GetPostRead", p1).Return(p, nil)

	p2 := "00000000-0000-0000-0000-00000000000"

	postgeSQLDB2 := new(mock2.DBInterface)
	postgeSQLDB2.On("GetPostRead", mock.Anything).Return(models.Read{}, errors.New("user doesn't exist"))

	tests := []struct {
		name    string
		pdb     postgreSQL.DBInterface
		rdb     redis.RedisDBInterface
		crq     createQueue.QueueCreateInterface
		param   string
		want    *models.Post
		wantErr string
	}{
		{
			name:    "Everything good",
			pdb:     postgeSQLDB,
			rdb:     redisDB,
			crq:     cq,
			param:   p1,
			want:    &post,
			wantErr: "",
		},
		{
			name:    "no users in postgreSQL",
			pdb:     postgeSQLDB,
			rdb:     redisDB,
			crq:     cq,
			param:   p2,
			want:    nil,
			wantErr: "service: couldn't parse id",
		},
		{
			name:    "no users in postgreSQL",
			pdb:     postgeSQLDB2,
			rdb:     redisDB,
			crq:     cq,
			param:   p1,
			want:    nil,
			wantErr: "user doesn't exist",
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ws := NewWriteServ(tc.pdb, tc.rdb, tc.crq)
			got, err := ws.GetPost(tc.param)
			if (err != nil) && (err.Error() != tc.wantErr) {
				t.Errorf("error = %v, wantErr = %v", err.Error(), tc.wantErr)
				return
			}
			assert.Equal(t, tc.want, got)
		})
	}
}
