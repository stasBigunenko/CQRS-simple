package handlers

import (
	"CQRS-simple/pkg/models"
	"CQRS-simple/pkg/services/command"
	"CQRS-simple/pkg/services/queue"
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

type UserHandler struct {
	command command.CommandInterface
	queue   queue.QueueInterface
}

func NewHandler(c command.CommandInterface, q queue.QueueInterface) *UserHandler {
	return &UserHandler{
		command: c,
		queue:   q,
	}
}

func (h *UserHandler) Routes(r *mux.Router) *mux.Router {
	r.HandleFunc("/post/{id}/create", h.CreatePost).Methods("POST")
	r.HandleFunc("/create", h.CreateUser).Methods("POST")
	r.HandleFunc("/user-posts/{id}", h.GetUserPosts).Methods("GET")
	r.HandleFunc("/user/{id}", h.GetUser).Methods("GET")
	r.HandleFunc("/post/{id}", h.GetPost).Methods("GET")
	r.HandleFunc("/", h.GetAllUsers).Methods("GET")
	return r
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodPost {
		msg := "Method Not Allowed" //TODO
		msgJson, err := json.Marshal(msg)
		if err != nil {
			log.Fatalf("error")
		}
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write(msgJson)
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
		msgJson, err := json.Marshal(msg)
		if err != nil {
			log.Fatalf("error")
		}
		w.WriteHeader(http.StatusBadRequest)
		w.Write(msgJson)
		return
	}

	res, err := h.command.CreateUser(user)
	if err != nil {
		msg := "Internal problem" //TODO
		msgJson, err := json.Marshal(msg)
		if err != nil {
			log.Fatalf("error")
		}
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(msgJson)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&res)
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodGet {
		msg := "Method Not Allowed" //TODO
		msgJson, err := json.Marshal(msg)
		if err != nil {
			log.Fatalf("error")
		}
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write(msgJson)
		return
	}

	vars := mux.Vars(r)
	key := vars["id"]

	res, err := h.queue.GetUser(key)
	if err != nil {
		msg := "Internal problem" //TODO
		msgJson, err := json.Marshal(msg)
		if err != nil {
			log.Fatalf("error")
		}
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(msgJson)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&res)
}

func (h *UserHandler) CreatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodPost {
		msg := "Method Not Allowed" //TODO
		msgJson, err := json.Marshal(msg)
		if err != nil {
			log.Fatalf("error")
		}
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write(msgJson)
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
	}

	var post models.Post

	if err = json.Unmarshal(data, &post); err != nil {
		msg := "Bad request" //TODO
		msgJson, err := json.Marshal(msg)
		if err != nil {
			log.Fatalf("error")
		}
		w.WriteHeader(http.StatusBadRequest)
		w.Write(msgJson)
	}

	vars := mux.Vars(r)
	key := vars["id"]

	post.UserID = key

	res, err := h.command.CreatePost(post)
	if err != nil {
		msg := "Internal problem" //TODO
		msgJson, err := json.Marshal(msg)
		if err != nil {
			log.Fatalf("error")
		}
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(msgJson)
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&res)
}

func (h *UserHandler) GetUserPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodGet {
		msg := "Method Not Allowed" //TODO
		msgJson, err := json.Marshal(msg)
		if err != nil {
			log.Fatalf("error")
		}
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write(msgJson)
		return
	}

	vars := mux.Vars(r)
	key := vars["id"]

	res, err := h.queue.UserPosts(key)
	if err != nil {
		msg := "Internal problem" //TODO
		msgJson, err := json.Marshal(msg)
		if err != nil {
			log.Fatalf("error")
		}
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(msgJson)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&res)
}

func (h *UserHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodGet {
		msg := "Method Not Allowed" //TODO
		msgJson, err := json.Marshal(msg)
		if err != nil {
			log.Fatalf("error")
		}
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write(msgJson)
		return
	}

	res, err := h.queue.GetAllUsers()
	if err != nil {
		msg := "Internal problem" //TODO
		msgJson, err := json.Marshal(msg)
		if err != nil {
			log.Fatalf("error")
		}
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(msgJson)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&res)
}

func (h *UserHandler) GetPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodGet {
		msg := "Method Not Allowed" //TODO
		msgJson, err := json.Marshal(msg)
		if err != nil {
			log.Fatalf("error")
		}
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write(msgJson)
		return
	}

	vars := mux.Vars(r)
	key := vars["id"]

	res, err := h.queue.GetPost(key)
	if err != nil {
		msg := "Internal problem" //TODO
		msgJson, err := json.Marshal(msg)
		if err != nil {
			log.Fatalf("error")
		}
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(msgJson)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&res)
}
