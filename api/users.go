package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"rest-api-server/models"
	"strconv"

	"github.com/gorilla/mux"
)

var users = models.NewUsers()

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	users, _ := users.SeleteAll()

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Incorrect id", http.StatusBadRequest)
		return
	}

	// TODO: 해당 id가 없는 경우 에러 처리
	user, err := users.Selete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func AddUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	name := keyVal["name"]
	if len(name) == 0 {
		http.Error(w, "Incorrect name", http.StatusBadRequest)
		return
	}

	user, err := users.Create(name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func DelUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Incorrect id", http.StatusBadRequest)
		return
	}

	// TODO: 해당 id가 없는 경우 에러 처리
	num, err := users.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if num == 0 {
		http.Error(w, "No User", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
