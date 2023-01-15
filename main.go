package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type User struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}
type Profile struct {
	Department  string `json:"department"`
	Designation string `json:"designation"`
	Employee    User   `json:"employee"`
}

var profiles []Profile = []Profile{}

func addItem(res http.ResponseWriter, req *http.Request) {
	var newProfile Profile
	json.NewDecoder(req.Body).Decode(&newProfile)
	res.Header().Set("Content-Type", "application/json")
	profiles = append(profiles, newProfile)
	json.NewEncoder(res).Encode(profiles)
}

func getAllProfiles(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(profiles)

}

func getProfile(res http.ResponseWriter, req *http.Request) {
	var idParam string = mux.Vars(req)["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		res.WriteHeader(400)
		res.Write([]byte("Id couldn't be converted to Integer."))
		return
	}

	if id >= len(profiles) {
		res.WriteHeader(404)
		res.Write([]byte("No profile found with specified ID."))
		return
	}

	profile := profiles[id]
	res.Header().Set("Content-type", "application/json")
	json.NewEncoder(res).Encode(profile)
}

func updateProfile(res http.ResponseWriter, req *http.Request) {
	var idParam string = mux.Vars(req)["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		res.WriteHeader(400)
		res.Write([]byte("Id couldn't be converted to Integer."))
		return
	}

	if id >= len(profiles) {
		res.WriteHeader(404)
		res.Write([]byte("No profile found with specified ID."))
		return
	}

	var updateProfile Profile
	json.NewDecoder(req.Body).Decode(&updateProfile)

	profiles[id] = updateProfile

	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(updateProfile)
}

func deleteProfile(res http.ResponseWriter, req *http.Request) {
	var idParam string = mux.Vars(req)["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		res.WriteHeader(400)
		res.Write([]byte("Id couldn't be converted to Integer."))
		return
	}

	if id >= len(profiles) {
		res.WriteHeader(404)
		res.Write([]byte("No profile found with specified ID."))
		return
	}

	profiles = append(profiles[:id], profiles[id+1:]...)
	res.WriteHeader(200)
	res.Write([]byte("Profile Deleted."))

}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/v1/profiles", addItem).Methods("POST")

	router.HandleFunc("/api/v1/profiles", getAllProfiles).Methods("GET")

	router.HandleFunc("/api/v1/profiles/{id}", getProfile).Methods("GET")

	router.HandleFunc("/api/v1/profiles/{id}", updateProfile).Methods("PUT")

	router.HandleFunc("/api/v1/profiles/{id}", deleteProfile).Methods("DELETE")

	http.ListenAndServe(":5000", router)

}
