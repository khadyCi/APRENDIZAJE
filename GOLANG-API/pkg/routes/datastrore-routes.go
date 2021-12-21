package routes

import (
	"github.com/gorilla/mux"
	"github.com/khadyCi/bloober/pkg/controllers"
)

// Routa para interactuar sobre la tabla Usuarios
func RegisterUserStoreRoutes(router *mux.Router) {
	router.HandleFunc("/user", controllers.CreateUser).Methods("POST", "OPTIONS")
	router.HandleFunc("/user", controllers.GetUser).Methods("GET", "OPTIONS")
	router.HandleFunc("/user/{userId}/tasks", controllers.GetUserTasks).Methods("GET", "OPTIONS")
	router.HandleFunc("/user/{userId}", controllers.GetUserById).Methods("GET")
	router.HandleFunc("/user/{userId}", controllers.UpdateUser).Methods("PUT")
	router.HandleFunc("/user/{userId}", controllers.DeleteUser).Methods("DELETE", "OPTIONS")
	router.HandleFunc("/user/login", controllers.Login).Methods("POST", "OPTIONS")
}

//Routa para interactuar sobre la tabla de permisos
func RegisterAllow(router *mux.Router) {
	router.HandleFunc("/allow", controllers.CreateAllow).Methods("POST", "OPTIONS")
	router.HandleFunc("/allow", controllers.GetAllow).Methods("GET", "OPTIONS")
	router.HandleFunc("/allow/{allowId}", controllers.GetAllowById).Methods("GET")
	router.HandleFunc("/allow/{allowId}", controllers.UpdateAllow).Methods("PUT")
	router.HandleFunc("/allow/{allowId}", controllers.DeleteAllow).Methods("DELETE")
}

//Routa para interactuar sobre la tabla de permisos
func RegisterUserAllow(router *mux.Router) {
	router.HandleFunc("/user_allow", controllers.CreateUserAllow).Methods("POST", "OPTIONS")
	router.HandleFunc("/user_allow", controllers.GetUserAllow).Methods("GET", "OPTIONS")
	router.HandleFunc("/user_allow/{user_allowId}", controllers.GetUserAllowById).Methods("GET")
	router.HandleFunc("/user_allow/{user_allowId}", controllers.UpdateUserAllow).Methods("PUT")
}

// Routa para interactuar sobre la tabla Tarea
var RegisterTaskStoreRoutes = func(router *mux.Router) {
	//router.HandleFunc("/email", controllers.SendMail).Methods("POST")
	router.HandleFunc("/task", controllers.CreateTask).Methods("POST")
	router.HandleFunc("/task", controllers.GetTask).Methods("GET")
	router.HandleFunc("/task/{taskId:[0-9]+}", controllers.GetTaskById).Methods("GET")
	router.HandleFunc("/task/{taskId}", controllers.UpdateTask).Methods("PUT")
	router.HandleFunc("/task/{taskId}", controllers.DeleteTask).Methods("DELETE")
	router.HandleFunc("/task/mail", controllers.SendMail).Methods("POST", "OPTIONS")
	router.HandleFunc("/task/owner/{owner_id}", controllers.GetTaskByOwnerId).Methods("GET")
}

// Routa para interactuar sobre la tabla Tarea
var RegisterTypeTaskStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/typeTask", controllers.CreateTypeTask).Methods("POST")
	router.HandleFunc("/typeTask", controllers.GetAllTypeTask).Methods("GET")
	router.HandleFunc("/typeTask/{typeTaskId}", controllers.GetTypeTaskById).Methods("GET")
	router.HandleFunc("/typeTask/{typeTaskId}", controllers.UpdateTypeTask).Methods("PUT")
	router.HandleFunc("/typeTask/{typeTaskId}", controllers.DeleteTypeTask).Methods("DELETE")
}

// Send Mail
var SendMailsRoutes = func(router *mux.Router) {
	router.HandleFunc("/sendMail", controllers.SendMail).Methods("POST", "OPTIONS")
}
