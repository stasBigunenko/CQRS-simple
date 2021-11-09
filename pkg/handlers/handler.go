package handlers

import (
	"CQRS-simple/pkg/models"
	"CQRS-simple/pkg/services"
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

type UserHandler struct {
	handler services.ServiceInterface
}

func NewHandler(service services.ServiceInterface) *UserHandler {
	return &UserHandler{
		handler: service,
	}
}

func (h *UserHandler) Routes(r *mux.Router) *mux.Router {
	r.HandleFunc("/create", h.CreateUser).Methods("POST")
	r.HandleFunc("/{id}", h.GetUser).Methods("GET")
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
	}

	res, err := h.handler.Create(user)
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

	res, _ := h.handler.Get(key)
	//if err != nil {
	//	msg := "Internal problem" //TODO
	//	msgJson, err := json.Marshal(msg)
	//	if err != nil {
	//		log.Fatalf("error")
	//	}
	//	w.WriteHeader(http.StatusInternalServerError)
	//	w.Write(msgJson)
	//}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&res)
}
