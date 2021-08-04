package main

import (
	"github.com/gin-gonic/gin"
	"HR_module/DB"
	"HR_module/routes"
)

func main(){
	DB.ConnectDB()
    
	r := gin.Default()

	route:= r.Group("/api/v1")
	{
	route.POST("/users/create", routes.Create_users)
    route.PUT("/users/update/:id", routes.Update_user)
    route.POST("/jobs/create", routes.Create_jobs)
    route.GET("/jobs/show", routes.Show_jobs)
    route.PUT("/jobs/update/:id", routes.Update_job)
	route.DELETE("/jobs/delete/:id", routes.Delete_job)
    route.POST("/questions/create", routes.Create_questions)
    route.GET("/questions/show", routes.Show_questions)
    route.PUT("/questions/update/:id", routes.Update_question)
    route.DELETE("/questions/delete/:id", routes.Delete_question)
    route.POST("/tags/create", routes.Create_tags)
    route.GET("/tags/show", routes.Show_tags)
	route.PUT("/tags/update/:id", routes.Update_tag)
    route.DELETE("/tags/delete/:id", routes.Delete_tag)
    route.POST("/applicants/create", routes.Create_applicants)
    route.GET("/applicants/show", routes.Show_applicants)
	route.PUT("/applicant/update/:id", routes.Update_applicant)
    route.DELETE("/applicant/delete/:id", routes.Delete_applicant)
    route.POST("/options/create", routes.Create_options)
    route.GET("/options/show", routes.Show_options)
	route.POST("/interviews/create", routes.Create_interviews)
    route.GET("/ntierviews/show", routes.Show_interviews)
	route.PUT("/interview/update/:id", routes.Update_interview)
    route.DELETE("/interview/delete/:id", routes.Delete_interview)
	route.POST("/notifications/create", routes.Create_notifications)
    route.GET("/notifications/show-user_from/:User_from", routes.Show_notification_from)
    route.GET("/notifications/show-user_to/:User_to", routes.Show_notification_to)
	route.PUT("/notification/update/:id", routes.Update_notification)
    route.DELETE("/notification/delete/:id", routes.Delete_notification)
///////////////////////////////////////////////////////
    route.POST("/employees/create", routes.Create_employees)
    route.GET("/employees/show", routes.Show_employees)
	route.PUT("/employee/update/:id", routes.Update_employee)
    route.DELETE("/employee/delete/:id", routes.Delete_employee)
    route.POST("/departments/create", routes.Create_departments)
    route.GET("/departments/show", routes.Show_departments)
	route.PUT("/department/update/:id", routes.Update_department)
    route.DELETE("/department/delete/:id", routes.Delete_department)
    route.POST("/attendance_and_departure/create", routes.Create_attendance_and_departure)
    route.GET("/attendance_and_departure/show/:HR_employee_id", routes.Show_attendance_and_departure)
    route.POST("/discounts/create", routes.Create_discounts)
    route.GET("/discount/show/:HR_employee_id", routes.Show_discount)
	route.PUT("/discount/update/:id", routes.Update_discount)
    route.DELETE("/discount/delete/:id", routes.Delete_discount)
    route.POST("/loationcs/create", routes.Create_loationcs)
    route.GET("/loationc/show/:id", routes.Show_loationc)
	route.PUT("/loationc/update/:id", routes.Update_loationc)
    route.DELETE("/loationc/delete/:id", routes.Delete_loationc)
    route.POST("/countries/create", routes.Create_countries)
    route.GET("/countrie/show/:id", routes.Show_countrie)
	route.PUT("/countrie/update/:id", routes.Update_countrie)
    route.DELETE("/countrie/delete/:id", routes.Delete_countrie)
    route.POST("/regions/create", routes.Create_regions)
    route.GET("/region/show/:id", routes.Show_region)
	route.PUT("/region/update/:id", routes.Update_region)
    route.DELETE("/region/delete/:id", routes.Delete_region)
    

	}

	r.Run(":8080")
}