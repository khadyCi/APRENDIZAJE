package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/smtp"
	"os"

	_ "github.com/joho/godotenv/autoload"

	//"os"

	//"net/smtp"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/khadyCi/bloober/pkg/models"
	"github.com/khadyCi/bloober/pkg/utils"
	//"github.com/khadyCi/bloober/pkg/email"
)

var NewUser models.User

/********** TABLA ALLOW QUE NOS PERMITE CONTROLAR LOS PERMISOS DE USUARIOS*********/

// Controle de Permisos
func CreateAllow(w http.ResponseWriter, r *http.Request) {
	CreateAllow := &models.Allow{}
	// Metemos los datos del JSON en el struct allow
	utils.ParseBody(r, CreateAllow)
	k := CreateAllow.CreateAllow()
	res, _ := json.Marshal(k)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// Obtener Todos los permisos registrados
func GetAllow(w http.ResponseWriter, r *http.Request) {
	newAllows := models.GetAllAllows()
	res, _ := json.Marshal(newAllows)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

// Busqueda de permisos por su Id
func GetAllowById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	allowId := vars["allowId"]
	ID, err := strconv.ParseInt(allowId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	allowDatails, _ := models.GetAllowById(ID)
	res, _ := json.Marshal(allowDatails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//Eliminar Permisos
func DeleteAllow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	allowId := vars["allowId"]
	id, err := strconv.ParseInt(allowId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	allow := models.DeleteAllow(id)
	res, _ := json.Marshal(allow)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//Modificar Permisos

func UpdateAllow(w http.ResponseWriter, r *http.Request) {
	var updateAllow = &models.Allow{}
	utils.ParseBody(r, updateAllow)
	vars := mux.Vars(r)
	allowId := vars["allowId"]
	ID, err := strconv.ParseInt(allowId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	allowDatails, db := models.GetAllowById(ID)
	if updateAllow.Name != "" {
		allowDatails.Name = updateAllow.Name
	}
	if allowDatails.Description != "" {
		allowDatails.Description = allowDatails.Description
	}

	db.Save(&allowDatails)
	res, _ := json.Marshal(allowDatails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

/***************** TABLA USER_ALLOW QUE NOS PERMITE DAR PERMISOS A USUARIO*****************/

// Obtener todas las tarea registrados
func GetUserAllow(w http.ResponseWriter, r *http.Request) {
	newUserAllows := models.GetAllUserAllows()
	res, _ := json.Marshal(newUserAllows)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// Busqueda de permiso asignada por su Id
func GetUserAllowById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userAllowId := vars["userAllowId"]
	ID, err := strconv.ParseInt(userAllowId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	userAllowDatails, _ := models.GetUserAllowById(ID)
	res, _ := json.Marshal(userAllowDatails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//Creacion de un permiso para un usuario especifico
func CreateUserAllow(w http.ResponseWriter, r *http.Request) {
	CreateUserAllow := &models.UserAllow{}
	utils.ParseBody(r, CreateUserAllow)
	c := CreateUserAllow.CreateUserAllow()
	res, _ := json.Marshal(c)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// Eliminacion de Permiso de usuario
func DeleteUserAllow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userAllowId := vars["userAllowId"]
	id, err := strconv.ParseInt(userAllowId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	userAllow := models.DeleteUserAllow(id)
	res, _ := json.Marshal(userAllow)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateUserAllow(w http.ResponseWriter, r *http.Request) {
	var updateUserAllow = &models.UserAllow{}
	utils.ParseBody(r, updateUserAllow)
	vars := mux.Vars(r)
	userAllowId := vars["userAllowId"]
	ID, err := strconv.ParseInt(userAllowId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	userAllowDatails, db := models.GetUserAllowById(ID)
	if updateUserAllow.UserId != "" {
		userAllowDatails.UserId = updateUserAllow.UserId
	}
	if updateUserAllow.AllowId != "" {
		userAllowDatails.AllowId = updateUserAllow.AllowId
	}

	db.Save(&userAllowDatails)
	res, _ := json.Marshal(userAllowDatails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// Obtener Todos los usuarios registrados
func GetUser(w http.ResponseWriter, r *http.Request) {
	newUsers := models.GetAllUsers()
	res, _ := json.Marshal(newUsers)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

// Busqueda de usuario por su Id
func GetUserById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]
	ID, err := strconv.ParseInt(userId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	userDatails, _ := models.GetUserById(ID)
	res, _ := json.Marshal(userDatails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func SendMail(w http.ResponseWriter, r *http.Request) {
	mailPayload := &models.MailPayload{}
	utils.ParseBody(r, mailPayload)
	msg := "From: " + os.Getenv("MAIL_USERNAME") + "\n" +
		"To: " + mailPayload.To + "\n" +
		"Subject: " + mailPayload.Subject + "\n\n" +
		mailPayload.Message
	host := os.Getenv("MAIL_HOST") + ":" + os.Getenv("MAIL_PORT")
	auth := smtp.PlainAuth("", os.Getenv("MAIL_USERNAME"), os.Getenv("MAIL_PASSWORD"), os.Getenv("MAIL_HOST"))
	err := smtp.SendMail(host, auth, os.Getenv("MAIL_USERNAME"), []string{mailPayload.To}, []byte(msg))

	if err != nil {
		log.Printf("smtp error: %s", err)
		return
	}

	log.Print("sent, visit http://foobarbazz.mailinator.com")

}

//Creacion de un nuevo usuario
func CreateUser(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	// Metemos los datos del JSON en el struct user
	utils.ParseBody(r, user)
	user.HashPassword()
	b := user.CreateUser()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func Login(w http.ResponseWriter, r *http.Request) {
	var u models.User
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		fmt.Println("error while parsing")
		return
	}

	defer r.Body.Close()

	storedUser := u.GetByUserName(u.Username)
	if err != nil {
		err := errors.New("No se han encontrado usuarios")
		http.Error(w, err.Error(), 401)
		return
	}

	if !storedUser.PasswordMatch(u.Password) {
		err := errors.New("La contraseña no es correcta")
		http.Error(w, err.Error(), 401)
		return
	}
	res, _ := json.Marshal(storedUser)
	w.Write(res)

}

// Eliminacion de datos de usuarios
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]
	id, err := strconv.ParseInt(userId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	user := models.DeleteUser(id)
	res, _ := json.Marshal(user)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// Modifcacion de los datos de usuarios
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var updateUser = &models.User{}
	utils.ParseBody(r, updateUser)
	vars := mux.Vars(r)
	userId := vars["userId"]
	ID, err := strconv.ParseInt(userId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	userDatails, db := models.GetUserById(ID)
	if updateUser.Dni != "" {
		userDatails.Dni = updateUser.Dni
	}
	if updateUser.Name != "" {
		userDatails.Name = updateUser.Name
	}
	if updateUser.Last_name != "" {
		userDatails.Last_name = updateUser.Last_name
	}
	if updateUser.Phone_number != "" {
		userDatails.Phone_number = updateUser.Phone_number
	}
	if updateUser.Direction != "" {
		userDatails.Direction = updateUser.Direction
	}
	if updateUser.Section != "" {
		userDatails.Section = updateUser.Section
	}
	if updateUser.Postal_code != "" {
		userDatails.Postal_code = updateUser.Postal_code
	}
	if updateUser.Email != "" {
		userDatails.Email = updateUser.Email
	}
	if updateUser.Imagen != "" {
		userDatails.Imagen = updateUser.Imagen
	}
	if updateUser.Password != "" {
		userDatails.HashPassword()
	}
	db.Save(&userDatails)
	res, _ := json.Marshal(userDatails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

/***   Tabla TASK  ***/

var NewTask models.Task

// Obtener todas las tarea registrados
func GetTask(w http.ResponseWriter, r *http.Request) {
	newTasks := models.GetAllTasks()
	res, _ := json.Marshal(newTasks)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// Busqueda de Tarea por su Id
func GetTaskById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskId := vars["taskId"]
	ID, err := strconv.ParseInt(taskId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	taskDatails, _ := models.GetTaskById(ID)
	res, _ := json.Marshal(taskDatails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//Creacion de una nueva Tarea
func CreateTask(w http.ResponseWriter, r *http.Request) {
	CreateTask := &models.Task{}
	utils.ParseBody(r, CreateTask)
	c := CreateTask.CreateTask()
	res, _ := json.Marshal(c)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// Eliminacion de Tarea
func DeleteTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskId := vars["taskId"]
	id, err := strconv.ParseInt(taskId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	task := models.DeleteTask(id)
	res, _ := json.Marshal(task)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// Modifcacion de los datos de usuarios
func UpdateTask(w http.ResponseWriter, r *http.Request) {
	var updateTask = &models.Task{}
	utils.ParseBody(r, updateTask)
	vars := mux.Vars(r)
	taskId := vars["taskId"]
	ID, err := strconv.ParseInt(taskId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	taskDatails, db := models.GetTaskById(ID)
	if updateTask.Title != "" {
		taskDatails.Title = updateTask.Title
	}
	if updateTask.TaskTypeName != "" {
		taskDatails.TaskTypeName = updateTask.TaskTypeName
	}
	if updateTask.Importance != "" {
		taskDatails.Importance = updateTask.Importance
	}
	if updateTask.Description != "" {
		taskDatails.Description = updateTask.Description
	}
	if updateTask.PubDate != "" {
		taskDatails.PubDate = updateTask.PubDate
	}
	if updateTask.FinalDate != "" {
		taskDatails.FinalDate = updateTask.FinalDate
	}
	/*

			if updateTask.PubDate.IsZero() {
			taskDatails.PubDate = updateTask.PubDate
		}
		if updateTask.FinalDate.IsZero() {
			taskDatails.FinalDate = updateTask.FinalDate
		}
	*/
	db.Save(&taskDatails)
	res, _ := json.Marshal(taskDatails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// Obtener todas las tareas registradas para el usuario
func GetUserTasks(w http.ResponseWriter, r *http.Request) {
	newTasks := models.GetAllTasks()
	res, _ := json.Marshal(newTasks)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

/***    TABLA TIPO DE TAREA  ***/
var NewTypeTask models.TypeTask

// Obtener Todos los tipos de tareas registrados
func GetAllTypeTask(w http.ResponseWriter, r *http.Request) {
	newTypeTasks := models.GetAllTypeTasks()
	res, _ := json.Marshal(newTypeTasks)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// Busqueda de tipo de tarea por su Id
func GetTypeTaskById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	typeTaskId := vars["typeTaskId"]
	ID, err := strconv.ParseInt(typeTaskId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	typeTaskDatails, _ := models.GetTypeTaskById(ID)
	res, _ := json.Marshal(typeTaskDatails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//Creacion de un nuevo usuario
func CreateTypeTask(w http.ResponseWriter, r *http.Request) {
	CreateTypeTask := &models.TypeTask{}
	utils.ParseBody(r, CreateTypeTask)
	b := CreateTypeTask.CreateTypeTask()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// Eliminacion de datos de usuarios
func DeleteTypeTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	typeTaskId := vars["typeTaskId"]
	id, err := strconv.ParseInt(typeTaskId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	typeTask := models.DeleteTypeTask(id)
	res, _ := json.Marshal(typeTask)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// Modifcacion de los datos de usuarios
func UpdateTypeTask(w http.ResponseWriter, r *http.Request) {
	var updateTypeTask = &models.TypeTask{}
	utils.ParseBody(r, updateTypeTask)
	vars := mux.Vars(r)
	typeTaskId := vars["typeTaskId"]
	ID, err := strconv.ParseInt(typeTaskId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	typeTaskDatails, db := models.GetTypeTaskById(ID)
	if updateTypeTask.TaskTypeName != "" {
		typeTaskDatails.TaskTypeName = updateTypeTask.TaskTypeName
	}

	db.Save(&typeTaskDatails)
	res, _ := json.Marshal(typeTaskDatails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// Busqueda de  tarea por owner_id
func GetTaskByOwnerId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	owner_id := vars["owner_id"]
	OwnerID, err := strconv.ParseInt(owner_id, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	ownerTaskDatails, _ := models.GetTaskByOwnerId(OwnerID)
	res, _ := json.Marshal(ownerTaskDatails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
