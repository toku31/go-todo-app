package main

import (
	"fmt"
	// "log"

	"todo_app/app/controllers"
	"todo_app/app/models"
)

func main() {

	fmt.Println(models.Db)

	controllers.StartMainServer()

	/*
		//セッション作成
		user, _ := models.GetUserByEmail("test@test.com")
		fmt.Println(user)

		session, err := user.CreateSession()
		if err != nil {
			log.Println(err)
		}
		fmt.Println(session)

		valid, _ := session.CheckSession()
		fmt.Println(valid)
	*/

	/*
		fmt.Println(config.Config.Port)
		fmt.Println(config.Config.SQLDriver)
		fmt.Println(config.Config.DbName)
		fmt.Println(config.Config.LogFile)

		log.Println("test")
	*/

	// DB作成
	//fmt.Println(models.Db)

	// user　作成
	// u := &models.User{}
	// u.Name = "test2"
	// u.Email = "test2@example.com"
	// u.PassWord = "test"
	// fmt.Println(u)

	// u.CreateUser()

	//===============
	// u, _ := models.GetUser(1)
	// fmt.Println(u)

	// u.Name = "Test2"
	// u.Email = "test2@example.com"
	// u.UpdateUser()
	// u, _ = models.GetUser(1)
	// fmt.Println(u)

	// u.DeleteUser()
	// u, _ = models.GetUser(1)
	// fmt.Println(u)

	// ==============
	// user, _ := models.GetUser(2)
	// u, _ := models.GetUser(2)
	// u.CreateTodo("First Todo")

	//===============
	// todoの所得
	// t, _ := models.GetTodo(1)
	// fmt.Println(t)

	// todosの取得
	// user, _ := models.GetUser(3)
	// user.CreateTodo("Third Todo")

	// todos, _ := models.GetTodos()
	// for _, v := range todos {
	// 	fmt.Println(v)
	// }

	// user2, _ := models.GetUser(3)
	// todos, _ := user2.GetTodosByUser()
	// for _, v := range todos {
	// 	fmt.Println(v)
	// }

	// Update
	// t, _ := models.GetTodo(8)
	// t.Content = "Update Todo"
	// t.UpdateTodo()

	// Delete
	// t, _ := models.GetTodo(10)
	// t.DeleteTodo()

}
