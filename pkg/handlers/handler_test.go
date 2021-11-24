package handlers

import (
	mock2 "CQRS-simple/pkg/mock"
	"CQRS-simple/pkg/models"
	"CQRS-simple/pkg/rabbitMQ/createQueue"
	"CQRS-simple/pkg/services/readServ"
	"errors"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestUserHandler_CreateUser(t *testing.T) {
	r := new(mock2.ReadServInterface)
	cq := new(mock2.QueueCreateInterface)

	cq.On("QueueCreateWrite", mock.Anything).Return(nil)

	tests := []struct {
		name       string
		rsi        readServ.ReadServInterface
		crq        createQueue.QueueCreateInterface
		method     string
		url        string
		body       string
		wantStatus int
	}{
		{
			name:       "POST user Everything ok",
			rsi:        r,
			crq:        cq,
			method:     "POST",
			url:        "localhost:8080/create",
			body:       `{"name":"asd","Age":12}`,
			wantStatus: http.StatusCreated,
		},
		{
			name:       "Get wrong method",
			rsi:        r,
			crq:        cq,
			method:     "Get",
			url:        "localhost:8080/create",
			body:       `{"name":"asd","Age":12}`,
			wantStatus: http.StatusMethodNotAllowed,
		},
		{
			name:       "wrong body",
			rsi:        r,
			crq:        cq,
			method:     "POST",
			url:        "localhost:8080/create",
			body:       "1111",
			wantStatus: http.StatusBadRequest,
		},
		{
			name:       "not included name",
			rsi:        r,
			crq:        cq,
			method:     "POST",
			url:        "localhost:8080/create",
			body:       `{"name":"asd"}`,
			wantStatus: http.StatusBadRequest,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			h := NewHandler(tc.rsi, tc.crq)
			req, err := http.NewRequest(tc.method, tc.url, strings.NewReader(tc.body))
			if err != nil {
				t.Fatalf("something goes wrong %v", err)
			}
			rr := httptest.NewRecorder()

			h.CreateUser(rr, req)

			res := rr.Result()

			assert.Equal(t, tc.wantStatus, res.StatusCode)
		})
	}
}

func TestUserHandler_CreatePost(t *testing.T) {
	r := new(mock2.ReadServInterface)
	cq := new(mock2.QueueCreateInterface)

	cq.On("QueueCreateWrite", mock.Anything).Return(nil)

	tests := []struct {
		name       string
		rsi        readServ.ReadServInterface
		crq        createQueue.QueueCreateInterface
		method     string
		url        string
		body       string
		vars       map[string]string
		wantStatus int
	}{
		{
			name:   "POST post Everything ok",
			rsi:    r,
			crq:    cq,
			method: "POST",
			url:    "http://localhost:8080/post/{id}/create",
			body:   `{"title":"asd","message":"das"}`,
			vars: map[string]string{
				"id": "00000000-0000-0000-0000-000000000000",
			},
			wantStatus: http.StatusCreated,
		},
		{
			name:   "Get wrong method",
			rsi:    r,
			crq:    cq,
			method: "Get",
			url:    "http://localhost:8080/post/{id}/create",
			body:   `{"title":"asd","message":"das"}`,
			vars: map[string]string{
				"id": "00000000-0000-0000-0000-000000000000",
			},
			wantStatus: http.StatusMethodNotAllowed,
		},
		{
			name:   "wrong body model",
			rsi:    r,
			crq:    cq,
			method: "POST",
			url:    "http://localhost:8080/post/{id}/create",
			body:   `1111`,
			vars: map[string]string{
				"id": "00000000-0000-0000-0000-000000000000",
			},
			wantStatus: http.StatusBadRequest,
		},
		{
			name:   "wrong body model (empty)",
			rsi:    r,
			crq:    cq,
			method: "POST",
			url:    "http://localhost:8080/post/{id}/create",
			body:   `{"message":"das"}`,
			vars: map[string]string{
				"id": "00000000-0000-0000-0000-000000000000",
			},
			wantStatus: http.StatusBadRequest,
		},
		{
			name:   "wrong id",
			rsi:    r,
			crq:    cq,
			method: "POST",
			url:    "http://localhost:8080/post/{id}/create",
			body:   `{"message":"das"}`,
			vars: map[string]string{
				"id": "00000000-0000-0000-0000-00000000000",
			},
			wantStatus: http.StatusBadRequest,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			h := NewHandler(tc.rsi, tc.crq)
			req, err := http.NewRequest(tc.method, tc.url, strings.NewReader(tc.body))
			if err != nil {
				t.Fatalf("something goes wrong %v", err)
			}
			rr := httptest.NewRecorder()

			vars := tc.vars
			req = mux.SetURLVars(req, vars)

			h.CreatePost(rr, req)

			res := rr.Result()

			assert.Equal(t, tc.wantStatus, res.StatusCode)
		})
	}
}

func TestUserHandler_UpdateUser(t *testing.T) {
	r := new(mock2.ReadServInterface)
	cq := new(mock2.QueueCreateInterface)

	cq.On("QueueCreateWrite", mock.Anything).Return(nil)

	tests := []struct {
		name       string
		rsi        readServ.ReadServInterface
		crq        createQueue.QueueCreateInterface
		method     string
		url        string
		body       string
		vars       map[string]string
		wantStatus int
	}{
		{
			name:   "PUT user Everything ok",
			rsi:    r,
			crq:    cq,
			method: "PUT",
			url:    "http://localhost:8080/update-user/00000000-0000-0000-0000-000000000000",
			body:   `{"name":"asd","Age":12}`,
			vars: map[string]string{
				"id": "00000000-0000-0000-0000-000000000000",
			},
			wantStatus: http.StatusAccepted,
		},
		{
			name:   "GET wrong method",
			rsi:    r,
			crq:    cq,
			method: "GET",
			url:    "http://localhost:8080/update-user/00000000-0000-0000-0000-000000000000",
			body:   `{"name":"asd","Age":12}`,
			vars: map[string]string{
				"id": "00000000-0000-0000-0000-000000000000",
			},
			wantStatus: http.StatusMethodNotAllowed,
		},
		{
			name:   "PUT wrong body model",
			rsi:    r,
			crq:    cq,
			method: "PUT",
			url:    "http://localhost:8080/update-user/00000000-0000-0000-0000-000000000000",
			body:   `{"1221","Age":12}`,
			vars: map[string]string{
				"id": "00000000-0000-0000-0000-000000000000",
			},
			wantStatus: http.StatusBadRequest,
		},
		{
			name:   "PUT wrong body model (something missing)",
			rsi:    r,
			crq:    cq,
			method: "PUT",
			url:    "http://localhost:8080/update-user/00000000-0000-0000-0000-000000000000",
			body:   `{"Age":0}`,
			vars: map[string]string{
				"id": "00000000-0000-0000-0000-000000000000",
			},
			wantStatus: http.StatusBadRequest,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			h := NewHandler(tc.rsi, tc.crq)
			req, err := http.NewRequest(tc.method, tc.url, strings.NewReader(tc.body))
			if err != nil {
				t.Fatalf("something goes wrong %v", err)
			}
			rr := httptest.NewRecorder()

			vars := tc.vars
			req = mux.SetURLVars(req, vars)

			h.UpdateUser(rr, req)

			res := rr.Result()

			assert.Equal(t, tc.wantStatus, res.StatusCode)
		})
	}
}

func TestUserHandler_UpdatePost(t *testing.T) {
	r := new(mock2.ReadServInterface)
	cq := new(mock2.QueueCreateInterface)

	cq.On("QueueCreateWrite", mock.Anything).Return(nil)

	tests := []struct {
		name       string
		rsi        readServ.ReadServInterface
		crq        createQueue.QueueCreateInterface
		method     string
		url        string
		body       string
		vars       map[string]string
		wantStatus int
	}{
		{
			name:   "PUT post Everything ok",
			rsi:    r,
			crq:    cq,
			method: "PUT",
			url:    "http://localhost:8080/update-user/00000000-0000-0000-0000-000000000000",
			body:   `{"title":"asd","message":"12"}`,
			vars: map[string]string{
				"id": "00000000-0000-0000-0000-000000000000",
			},
			wantStatus: http.StatusAccepted,
		},
		{
			name:   "GET wrong method",
			rsi:    r,
			crq:    cq,
			method: "GET",
			url:    "http://localhost:8080/update-user/00000000-0000-0000-0000-000000000000",
			body:   `{"title":"asd","message":"12"}`,
			vars: map[string]string{
				"id": "00000000-0000-0000-0000-000000000000",
			},
			wantStatus: http.StatusMethodNotAllowed,
		},
		{
			name:   "PUT wrong body model",
			rsi:    r,
			crq:    cq,
			method: "PUT",
			url:    "http://localhost:8080/update-user/00000000-0000-0000-0000-000000000000",
			body:   `{""Age":12}`,
			vars: map[string]string{
				"id": "00000000-0000-0000-0000-000000000000",
			},
			wantStatus: http.StatusBadRequest,
		},
		{
			name:   "PUT wrong body model",
			rsi:    r,
			crq:    cq,
			method: "PUT",
			url:    "http://localhost:8080/update-user/00000000-0000-0000-0000-000000000000",
			body:   `{"Message":""}`,
			vars: map[string]string{
				"id": "00000000-0000-0000-0000-000000000000",
			},
			wantStatus: http.StatusBadRequest,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			h := NewHandler(tc.rsi, tc.crq)
			req, err := http.NewRequest(tc.method, tc.url, strings.NewReader(tc.body))
			if err != nil {
				t.Fatalf("something goes wrong %v", err)
			}
			rr := httptest.NewRecorder()

			vars := tc.vars
			req = mux.SetURLVars(req, vars)

			h.UpdatePost(rr, req)

			res := rr.Result()

			assert.Equal(t, tc.wantStatus, res.StatusCode)
		})
	}
}

func TestUserHandler_DeleteUser(t *testing.T) {
	r := new(mock2.ReadServInterface)
	cq := new(mock2.QueueCreateInterface)
	cq.On("QueueCreateWrite", mock.Anything).Return(nil)

	tests := []struct {
		name       string
		rsi        readServ.ReadServInterface
		crq        createQueue.QueueCreateInterface
		method     string
		url        string
		vars       map[string]string
		wantStatus int
	}{
		{
			name:   "DELETE user Everything ok",
			rsi:    r,
			crq:    cq,
			method: "DELETE",
			url:    "http://localhost:8080/delete-user/00000000-0000-0000-0000-000000000000",
			vars: map[string]string{
				"id": "00000000-0000-0000-0000-000000000000",
			},
			wantStatus: http.StatusAccepted,
		},
		{
			name:   "GET wrong method",
			rsi:    r,
			crq:    cq,
			method: "GET",
			url:    "http://localhost:8080/delete-user/00000000-0000-0000-0000-000000000000",
			vars: map[string]string{
				"id": "00000000-0000-0000-0000-000000000000",
			},
			wantStatus: http.StatusMethodNotAllowed,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			h := NewHandler(tc.rsi, tc.crq)
			req, err := http.NewRequest(tc.method, tc.url, strings.NewReader(""))
			if err != nil {
				t.Fatalf("something goes wrong %v", err)
			}
			rr := httptest.NewRecorder()

			vars := tc.vars
			req = mux.SetURLVars(req, vars)

			h.DeleteUser(rr, req)

			res := rr.Result()

			assert.Equal(t, tc.wantStatus, res.StatusCode)
		})
	}
}

func TestUserHandler_DeletePost(t *testing.T) {
	r := new(mock2.ReadServInterface)
	cq := new(mock2.QueueCreateInterface)
	cq.On("QueueCreateWrite", mock.Anything).Return(nil)

	tests := []struct {
		name       string
		rsi        readServ.ReadServInterface
		crq        createQueue.QueueCreateInterface
		method     string
		url        string
		vars       map[string]string
		wantStatus int
	}{
		{
			name:   "DELETE post Everything ok",
			rsi:    r,
			crq:    cq,
			method: "DELETE",
			url:    "http://localhost:8080/delete-post/00000000-0000-0000-0000-000000000000",
			vars: map[string]string{
				"id": "00000000-0000-0000-0000-000000000000",
			},
			wantStatus: http.StatusAccepted,
		},
		{
			name:   "GET wrong method",
			rsi:    r,
			crq:    cq,
			method: "GET",
			url:    "http://localhost:8080/delete-post/00000000-0000-0000-0000-000000000000",
			vars: map[string]string{
				"id": "00000000-0000-0000-0000-000000000000",
			},
			wantStatus: http.StatusMethodNotAllowed,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			h := NewHandler(tc.rsi, tc.crq)
			req, err := http.NewRequest(tc.method, tc.url, strings.NewReader(""))
			if err != nil {
				t.Fatalf("something goes wrong %v", err)
			}
			rr := httptest.NewRecorder()

			vars := tc.vars
			req = mux.SetURLVars(req, vars)

			h.DeletePost(rr, req)

			res := rr.Result()

			assert.Equal(t, tc.wantStatus, res.StatusCode)
		})
	}
}

func TestUserHandler_GetUserPosts(t *testing.T) {
	r := new(mock2.ReadServInterface)
	cq := new(mock2.QueueCreateInterface)
	cq.On("QueueCreateWrite", mock.Anything).Return(nil)

	m := models.User{"00000000-0000-0000-0000-000000000000", "raz", 2}
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
	r.On("UserPosts", mock.Anything).Return(up, nil)

	r2 := new(mock2.ReadServInterface)
	r2.On("UserPosts", mock.Anything).Return(models.UserPosts{}, errors.New("Internal problem"))

	tests := []struct {
		name       string
		rsi        readServ.ReadServInterface
		crq        createQueue.QueueCreateInterface
		method     string
		url        string
		vars       map[string]string
		wantStatus int
	}{
		{
			name:   "GET usersPost Everything ok",
			rsi:    r,
			crq:    cq,
			method: "GET",
			url:    "http://localhost:8080/user-posts/00000000-0000-0000-0000-000000000000",
			vars: map[string]string{
				"id": "00000000-0000-0000-0000-000000000000",
			},
			wantStatus: http.StatusOK,
		},
		{
			name:   "wrong method",
			rsi:    r,
			crq:    cq,
			method: "PUT",
			url:    "http://localhost:8080/user-posts/00000000-0000-0000-0000-000000000000",
			vars: map[string]string{
				"id": "00000000-0000-0000-0000-000000000000",
			},
			wantStatus: http.StatusMethodNotAllowed,
		},
		{
			name:   "GET usersPost Everything ok",
			rsi:    r2,
			crq:    cq,
			method: "GET",
			url:    "http://localhost:8080/user-posts/00000000-0000-0000-0000-000000000000",
			vars: map[string]string{
				"id": "00000000-0000-0000-0000-000000000000",
			},
			wantStatus: http.StatusInternalServerError,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			h := NewHandler(tc.rsi, tc.crq)
			req, err := http.NewRequest(tc.method, tc.url, strings.NewReader(""))
			if err != nil {
				t.Fatalf("something goes wrong %v", err)
			}
			rr := httptest.NewRecorder()

			vars := tc.vars
			req = mux.SetURLVars(req, vars)

			h.GetUserPosts(rr, req)

			res := rr.Result()

			assert.Equal(t, tc.wantStatus, res.StatusCode)
		})
	}
}

func TestUserHandler_GetAllUsers(t *testing.T) {
	r := new(mock2.ReadServInterface)
	cq := new(mock2.QueueCreateInterface)
	cq.On("QueueCreateWrite", mock.Anything).Return(nil)

	m := models.User{"00000000-0000-0000-0000-000000000000", "raz", 2}
	m1 := models.User{"00000000-0000-0000-0000-000000000001", "raz", 2}

	users := []models.User{
		m,
		m1,
	}
	r.On("GetAllUsers").Return(&users, nil)

	r2 := new(mock2.ReadServInterface)
	r2.On("GetAllUsers").Return(&[]models.User{}, errors.New("Internal problem"))

	tests := []struct {
		name       string
		rsi        readServ.ReadServInterface
		crq        createQueue.QueueCreateInterface
		method     string
		url        string
		vars       map[string]string
		wantStatus int
	}{
		{
			name:   "GET usersPost Everything ok",
			rsi:    r,
			crq:    cq,
			method: "GET",
			url:    "http://localhost:8080/user-posts/00000000-0000-0000-0000-000000000000",
			vars: map[string]string{
				"id": "00000000-0000-0000-0000-000000000000",
			},
			wantStatus: http.StatusOK,
		},
		{
			name:   "Wrong method",
			rsi:    r,
			crq:    cq,
			method: "POST",
			url:    "http://localhost:8080/user-posts/00000000-0000-0000-0000-000000000000",
			vars: map[string]string{
				"id": "00000000-0000-0000-0000-000000000000",
			},
			wantStatus: http.StatusMethodNotAllowed,
		},
		{
			name:   "Wrong function",
			rsi:    r2,
			crq:    cq,
			method: "GET",
			url:    "http://localhost:8080/user-posts/00000000-0000-0000-0000-000000000000",
			vars: map[string]string{
				"id": "00000000-0000-0000-0000-000000000000",
			},
			wantStatus: http.StatusInternalServerError,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			h := NewHandler(tc.rsi, tc.crq)
			req, err := http.NewRequest(tc.method, tc.url, strings.NewReader(""))
			if err != nil {
				t.Fatalf("something goes wrong %v", err)
			}
			rr := httptest.NewRecorder()

			vars := tc.vars
			req = mux.SetURLVars(req, vars)

			h.GetAllUsers(rr, req)

			res := rr.Result()

			assert.Equal(t, tc.wantStatus, res.StatusCode)
		})
	}
}
