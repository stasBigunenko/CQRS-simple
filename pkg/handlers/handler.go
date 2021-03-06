package handlers

import (
	"CQRS-simple/pkg/models"
	"CQRS-simple/pkg/rabbitMQ/createQueue"
	"CQRS-simple/pkg/services/readServ"
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

type UserHandler struct {
	readServ    readServ.ReadServInterface
	createQueue createQueue.QueueCreateInterface
}

func NewHandler(rs readServ.ReadServInterface, cq createQueue.QueueCreateInterface) *UserHandler {
	return &UserHandler{
		readServ:    rs,
		createQueue: cq,
	}
}

func (h *UserHandler) Routes(r *mux.Router) *mux.Router {
	r.HandleFunc("/create", h.CreateUser).Methods("POST")
	r.HandleFunc("/post/{id}/create", h.CreatePost).Methods("POST")

	r.HandleFunc("/update-user/{id}", h.UpdateUser).Methods("PUT")
	r.HandleFunc("/update-post/{id}", h.UpdatePost).Methods("PUT")

	r.HandleFunc("/delete-user/{id}", h.DeleteUser).Methods("DELETE")
	r.HandleFunc("/delete-post/{id}", h.DeletePost).Methods("DELETE")

	r.HandleFunc("/user-posts/{id}", h.GetUserPosts).Methods("GET")
	r.HandleFunc("/", h.GetAllUsers).Methods("GET")
	return r
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodPost {
		msg := "Method Not Allowed" //TODO
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(&msg)
		return
	}

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		msg := "Bad request" //TODO
		msgJson, err := json.Marshal(msg)
		if err != nil {
			log.Fatalf("error")
		}
		w.WriteHeader(http.StatusBadRequest)
		w.Write(msgJson)
		return
	}

	var user models.User

	if err = json.Unmarshal(data, &user); err != nil {
		msg := "Bad request" //TODO
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&msg)
		return
	}

	if user.Name == "" || user.Age == 0 {
		msg := "Bad request"
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&msg)
		return
	}

	var cud models.Cud

	cud.Model = "user"
	cud.Command = "create"
	cud.User = user

	h.createQueue.QueueCreateWrite(cud)

	res := "user created"
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&res)
}

func (h *UserHandler) CreatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodPost {
		msg := "Method Not Allowed" //TODO
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(&msg)
		return
	}

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		msg := "Bad request" //TODO
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&msg)
		return
	}

	var post models.Post

	if err = json.Unmarshal(data, &post); err != nil {
		msg := "Bad request" //TODO
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&msg)
		return
	}

	vars := mux.Vars(r)
	key := vars["id"]

	post.UserID = key

	if post.Title == "" || post.Message == "" {
		msg := "Bad request" //TODO
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&msg)
		return
	}

	var cud models.Cud

	cud.Model = "post"
	cud.Command = "create"
	cud.Post = post

	h.createQueue.QueueCreateWrite(cud)

	res := "post created"
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&res)
}

func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodPut {
		msg := "Method Not Allowed" //TODO
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(&msg)
		return
	}

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		msg := "Bad request" //TODO
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&msg)
		return
	}

	var user models.User

	if err = json.Unmarshal(data, &user); err != nil {
		msg := "Bad request" //TODO
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&msg)
		return
	}

	vars := mux.Vars(r)
	key := vars["id"]

	user.ID = key

	if user.Name == "" && user.Age == 0 {
		msg := "Bad request"
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&msg)
		return
	}

	var cud models.Cud

	cud.Model = "user"
	cud.Command = "update"
	cud.User = user

	h.createQueue.QueueCreateWrite(cud)

	res := "user updated"
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(&res)
}

func (h *UserHandler) UpdatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodPut {
		msg := "Method Not Allowed" //TODO
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(&msg)
		return
	}

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		msg := "Bad request" //TODO
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&msg)
		return
	}

	var post models.Post

	if err = json.Unmarshal(data, &post); err != nil {
		msg := "Bad request" //TODO
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&msg)
		return
	}

	vars := mux.Vars(r)
	key := vars["id"]

	post.ID = key

	if post.Title == "" && post.Message == "" {
		msg := "Bad request"
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&msg)
		return
	}

	var cud models.Cud

	cud.Model = "post"
	cud.Command = "update"
	cud.Post = post

	h.createQueue.QueueCreateWrite(cud)

	res := "post updated"
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(&res)
}

func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodDelete {
		msg := "Method Not Allowed" //TODO
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(&msg)
		return
	}

	vars := mux.Vars(r)
	key := vars["id"]

	var user models.User
	user.ID = key

	var cud models.Cud

	cud.Model = "user"
	cud.Command = "delete"
	cud.User = user

	h.createQueue.QueueCreateWrite(cud)

	w.WriteHeader(http.StatusAccepted)
	msg := "User deleted"
	json.NewEncoder(w).Encode(&msg)
}

func (h *UserHandler) DeletePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodDelete {
		msg := "Method Not Allowed" //TODO
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(&msg)
		return
	}

	vars := mux.Vars(r)
	key := vars["id"]

	var post models.Post
	post.ID = key

	var cud models.Cud

	cud.Model = "post"
	cud.Command = "delete"
	cud.Post = post

	h.createQueue.QueueCreateWrite(cud)

	w.WriteHeader(http.StatusAccepted)
	msg := "Post deleted"
	json.NewEncoder(w).Encode(&msg)
}

func (h *UserHandler) GetUserPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodGet {
		msg := "Method Not Allowed" //TODO
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(&msg)
		return
	}

	vars := mux.Vars(r)
	key := vars["id"]

	res, err := h.readServ.UserPosts(key)
	if err != nil {
		msg := "Internal problem" //TODO
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&msg)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&res)
}

func (h *UserHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodGet {
		msg := "Method Not Allowed" //TODO
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(&msg)
		return
	}

	res, err := h.readServ.GetAllUsers()
	if err != nil {
		msg := "Internal problem" //TODO
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&msg)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&res)
}
