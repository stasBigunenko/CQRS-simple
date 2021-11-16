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
	r.HandleFunc("/create", h.CreateUser).Methods("POST")
	r.HandleFunc("/post/{id}/create", h.CreatePost).Methods("POST")
	r.HandleFunc("/update-user/{id}", h.UpdateUser).Methods("PUT")
	r.HandleFunc("/update-post/{id}", h.UpdatePost).Methods("PUT")
	r.HandleFunc("/user-posts/{id}", h.GetUserPosts).Methods("GET")
	r.HandleFunc("/delete-user/{id}", h.DeleteUser).Methods("DELETE")
	r.HandleFunc("/delete-post/{id}", h.DeletePost).Methods("DELETE")
	//r.HandleFunc("/user/{id}", h.GetUser).Methods("GET")
	//r.HandleFunc("/post/{id}", h.GetPost).Methods("GET")
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

//func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", "application/json")
//	if r.Method != http.MethodGet {
//		msg := "Method Not Allowed" //TODO
//		msgJson, err := json.Marshal(msg)
//		if err != nil {
//			log.Fatalf("error")
//		}
//		w.WriteHeader(http.StatusMethodNotAllowed)
//		w.Write(msgJson)
//		return
//	}
//
//	vars := mux.Vars(r)
//	key := vars["id"]
//
//	res, err := h.queue.GetUser(key)
//	if err != nil {
//		msg := "Internal problem" //TODO
//		msgJson, err := json.Marshal(msg)
//		if err != nil {
//			log.Fatalf("error")
//		}
//		w.WriteHeader(http.StatusInternalServerError)
//		w.Write(msgJson)
//		return
//	}
//
//	w.WriteHeader(http.StatusOK)
//	json.NewEncoder(w).Encode(&res)
//}

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

//func (h *UserHandler) GetPost(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", "application/json")
//	if r.Method != http.MethodGet {
//		msg := "Method Not Allowed" //TODO
//		msgJson, err := json.Marshal(msg)
//		if err != nil {
//			log.Fatalf("error")
//		}
//		w.WriteHeader(http.StatusMethodNotAllowed)
//		w.Write(msgJson)
//		return
//	}
//
//	vars := mux.Vars(r)
//	key := vars["id"]
//
//	res, err := h.queue.GetPost(key)
//	if err != nil {
//		msg := "Internal problem" //TODO
//		msgJson, err := json.Marshal(msg)
//		if err != nil {
//			log.Fatalf("error")
//		}
//		w.WriteHeader(http.StatusInternalServerError)
//		w.Write(msgJson)
//		return
//	}
//
//	w.WriteHeader(http.StatusOK)
//	json.NewEncoder(w).Encode(&res)
//}

func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodPut {
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

	vars := mux.Vars(r)
	key := vars["id"]

	user.ID = key

	res, err := h.command.UpdateUser(user)
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

func (h *UserHandler) UpdatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodPut {
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

	post.ID = key

	res, err := h.command.UpdatePost(post)
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

func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodDelete {
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

	err := h.command.DeleteUser(key)
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
	msg := "User deleted"
	json.NewEncoder(w).Encode(&msg)
}

func (h *UserHandler) DeletePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodDelete {
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

	err := h.command.DeletePost(key)
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
	msg := "Post deleted"
	json.NewEncoder(w).Encode(&msg)
}
