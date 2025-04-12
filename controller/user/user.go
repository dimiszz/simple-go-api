package user

import (
	"encoding/json"
	"net/http"
	"strconv"

	"dimi/server/controller/route"
	"dimi/server/mistakes"
	"dimi/server/repository/userRepository"
)

type UserController struct {
	routes []route.Route
	path   string
}

func NewUserController(prefix string) *UserController {
	return &UserController{path: prefix}
}

func (controller *UserController) AddRoutes() {
	controller.routes = []route.Route{
		{Pattern: "/hello", Handler: http.HandlerFunc(helloWorldHandler)},
		{Pattern: "/{id}", Handler: http.HandlerFunc(getUserByIdHandler)},
		{Pattern: "/create/{name}/{age}", Handler: http.HandlerFunc(createUserHandler)},
		{Pattern: "/users", Handler: http.HandlerFunc(getUsersHandler)},
	}
}

func (controller *UserController) GetPrefix() string {
	return controller.path
}

func (controller UserController) RegisterRoutes() *http.ServeMux {
	userRouter := http.NewServeMux()
	for _, route := range controller.routes {
		userRouter.HandleFunc(route.Pattern, route.Handler.ServeHTTP)
	}

	return userRouter
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func getUserByIdHandler(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		w.Write([]byte("Número inválido!"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := userRepository.GetUserById(id)

	if err != nil {
		if err == mistakes.ErrNotFound {
			w.Write([]byte("Usuário não encontrado!"))
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.Write([]byte("Erro ao buscar usuário! Erro: " + err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Write([]byte("Usuário encontrado! Nome: " + user.Name + " Idade: " + strconv.Itoa(user.Age)))
	w.WriteHeader(http.StatusOK)
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	age, err := strconv.Atoi(r.PathValue("age"))

	if err != nil {
		w.Write([]byte("Número inválido!"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	name := r.PathValue("name")

	user := userRepository.User{
		Id:   0,
		Name: name,
		Age:  age,
	}

	err = userRepository.CreateUser(&user)
	if err != nil {
		w.Write([]byte("Erro ao criar usuário! Erro: " + err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Usuário criado com sucesso! ID: " + strconv.Itoa(user.Id)))
	w.WriteHeader(http.StatusCreated)
}

func getUsersHandler(w http.ResponseWriter, r *http.Request) {
	users := userRepository.GetUsers()

	encoder := json.NewEncoder(w)

	_ = encoder.Encode(users)

	w.WriteHeader(http.StatusOK)
}
