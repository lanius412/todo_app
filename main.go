package main

import (
	"fmt"
	//"log"

	//"todo_app/config"
	"todo_app/app/models"
	"todo_app/app/controllers"
)

func main() {
	/*
	fmt.Println(config.Config.Port)
	fmt.Println(config.Config.SQLDriver)
	fmt.Println(config.Config.DbName)
	fmt.Println(config.Config.LogFile)

	log.Println("test")
	*/
	fmt.Println(models.Db)
	fmt.Println()

	controllers.StartMainServer()

	/*
	user, _ := models.GetUserByEmail("test@example.com")
	fmt.Println(user)

	session, err := user.CreateSession()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(session)

	valid, _ := session.CheckSession()
	fmt.Println(valid)
	*/

	/*		//Create User
	user := &models.User{}
	user.Name = "test"
	user.Email = "test@example.com"
	user.Password = "testtest"
	fmt.Println(user)
	user.CreateUser()
	*/

	/*		//Read(GET) User
	user, _ := models.GetUser(2)
	fmt.Println(user)
	*/

	/*		//Update User
	user, _ := models.GetUser(2)
	user.Name = "test2"
	user.Email = "test2@example.com"
	user.UpdateUser()
	*/

	/*		//Delete User
	user, _ := models.GetUser(2)
	user.DeleteUser()
	*/


	/*		//Create Todo  Read(GET) One ToDo
	user, _ := models.GetUser(2)
	user.CreateTodo("First Todo")
	todo, _ := models.GetTodo(1)
	fmt.Println(todo)
	*/

	/*		//Read(GET) All Todos
	user, _ := models.GetUser(2)
	user.CreateTodo("Second Todo")
	todos, _ := models.GetTodos()
	for _, t := range todos {
		fmt.Println(t)
	}
	*/

	/*		//Read(GET) TodosByUser
	user, _ := models.GetUser(1)
	todos, _ := user.GetTodosByUser()
	for _, v := range todos {
		fmt.Println(v)
	}
	*/

	/*		//Update Todo
	todo, _ := models.GetTodo(4)
	todo.Content = "UpdateTodo"
	todo.UpdateTodo()
	*/

	/*		//Delete Todo
	user, _ := models.GetUser(2)
	user.CreateTodo("2nd Todo")
	todo, _ := models.GetTodo(3)
	todo.DeleteTodo()
	*/




}