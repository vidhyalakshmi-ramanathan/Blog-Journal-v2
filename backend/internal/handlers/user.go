package handlers

import (
	"backend/internal/entities"
	"backend/internal/services"
	"encoding/json"
	"net/http"
)

type UserHandler struct {
	service services.UserService
}

func NewUserHandler(service services.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var userData entities.SignupData

	if err := json.NewDecoder(r.Body).Decode(&userData); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	user := entities.User{
		UserName:    userData.Username,
		Password:    userData.Password,
		Email:       userData.Email,
		PhoneNumber: userData.Phonenumber,
	}
	if err := h.service.Create(&user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(entities.ApiResponse{
		Message: "User Created Successfully!",
	})
}

func (h *UserHandler) LoginUser(w http.ResponseWriter, r *http.Request) {
	var loginData entities.LoginData
	if err := json.NewDecoder(r.Body).Decode(&loginData); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	token, err := h.service.Login(loginData.Identifier, loginData.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(entities.ApiResponse{
		Message: "Login successfully!",
		Data:    map[string]string{"token": token},
	})
}
