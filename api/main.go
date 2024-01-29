package main

import (
	"Visma/api/handlers"
	"Visma/helpers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	//USER/ADMIN
	router.POST("/register-team", handlers.RegistrationHandler)

	//USER/ADMIN
	router.POST("/login", handlers.LoginHandler)

	//USER/ADMIN
	router.POST("/logout", handlers.LogoutHandler)
	//USER neriesil som aby admin nemohol
	router.POST("/submit-task", handlers.SubmitTaskHandler) // submits code to languages by choice
	//ADMIN
	router.POST("/submit-new-task", handlers.SubmitNewTaskHandler) //submits new challenges
	//all tasks aj get language task som chapal ze z databazy co su ale tieto dve chcem preriesit
	router.GET("/get-all-tasks", handlers.GetTasksHandler)             //it gives me all tasks that are currently in database
	router.GET("/get-language-tasks", handlers.GetLanguageTaskHandler) //it gives me all tasks that are currently in database
	//ADMIN
	router.PUT("/update-task/:id", handlers.EditTaskHandler) // it updates (edits task)
	//ADMIN
	router.POST("/add-task", handlers.AddTaskHandler) // adds task
	//ADMIN
	router.DELETE("/remove-task/:id", handlers.RemoveTaskByIDHandler) //remove task
	//ADMIN
	router.POST("/create-event", handlers.AddCompetitionHandler) //creates competition
	//ADMIN
	router.POST("/start-competition/:id", handlers.StartCompetitionHandler) //start competition
	//ADMIN
	router.POST("/end-competition/:id", handlers.EndCompetitionHandler) //ends competition
	//this is try comment

	router.Use(handlers.CorsMiddleware())
	err := router.Run("" + helpers.IpAddress() + ":9090")
	if err != nil {
		return
	}
}
