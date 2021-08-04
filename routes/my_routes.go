package routes

import (

	"net/http"
	"github.com/gin-gonic/gin"
	"HR_module/DB"
	"HR_module/structs"
)

	func Create_users(c *gin.Context) {
		var users structs.Hr_users
		if err := c.ShouldBindJSON(&users); err != nil {
		  c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		  return
		}
		DB.ConnectDB().Create(&users)
		c.JSON(200, gin.H{
		  "view": users,
		})
	  }
	  
	  func Update_user(c *gin.Context) {
		var userRequest structs.User_update
		if err := c.ShouldBindJSON(&userRequest); err != nil {
		  c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		  return
		}
		id := c.Param("id")
		var Hr_user structs.Hr_users
		DB.ConnectDB().Where("id=?", id).First(&Hr_user)
	  
		Hr_user.Name = userRequest.Name
		Hr_user.Email = userRequest.Email
		Hr_user.Password = userRequest.Password
	  
		DB.ConnectDB().Save(&Hr_user)
	  
		c.JSON(200, gin.H{
		  "view": Hr_user,
		})
	  }
	  
	  func Create_jobs(c *gin.Context) {
		var jobs structs.Hr_jobs
		if err := c.ShouldBindJSON(&jobs); err != nil {
		  c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		  return
		}
		DB.ConnectDB().Create(&jobs)
		c.JSON(200, gin.H{
		  "view": jobs,
		})
	  }
	  
	  func Show_jobs(c *gin.Context) {
		var jobs []structs.Hr_jobs
		DB.ConnectDB().Find(&jobs)
		c.JSON(200, gin.H{
		  "view": jobs,
		})
	  }
	  
	  func Update_job(c *gin.Context) {
		var jobRequest structs.Job_update
		if err := c.ShouldBindJSON(&jobRequest); err != nil {
		  c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		  return
		}
		id := c.Param("id")
		var Hr_job structs.Hr_jobs
		DB.ConnectDB().Where("id=?", id).First(&Hr_job)
	  
		Hr_job.Name = jobRequest.Name
		Hr_job.Salary = jobRequest.Salary
		Hr_job.Description = jobRequest.Description
		Hr_job.Vailable_date = jobRequest.Vailable_date
		Hr_job.Is_show_in_home = jobRequest.Is_show_in_home
		Hr_job.Is_archive = jobRequest.Is_archive
		Hr_job.Is_question = jobRequest.Is_question
		Hr_job.Send_alert = jobRequest.Send_alert
	  
		DB.ConnectDB().Save(&Hr_job)
	  
		c.JSON(200, gin.H{
		  "view": Hr_job,
		})
	  }
	  
	  func Delete_job(c *gin.Context) {
		id := c.Param("id")
		var job structs.Hr_jobs
		DB.ConnectDB().Where("id=?", id).Delete(&job)
		c.JSON(200, gin.H{
		  "view": "Deleted",
		})
	  }
	  
	  func Create_questions(c *gin.Context) {
		var question structs.Hr_job_questions
		if err := c.ShouldBindJSON(&question); err != nil {
		  c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		  return
		}
		DB.ConnectDB().Create(&question)
		c.JSON(200, gin.H{
		  "view": question,
		})
	  }
	  
	  func Show_questions(c *gin.Context) {
		var questions []structs.Hr_job_questions
		DB.ConnectDB().Find(&questions)
		c.JSON(200, gin.H{
		  "view": questions,
		})
	  }
	  
	  func Update_question(c *gin.Context) {
		var questionRequest structs.Question_update
		if err := c.ShouldBindJSON(&questionRequest); err != nil {
		  c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		  return
		}
		id := c.Param("id")
		var Hr_question structs.Hr_job_questions
		DB.ConnectDB().Where("id=?", id).First(&Hr_question)
	  
		Hr_question.Question = questionRequest.Question
		Hr_question.Question_type = questionRequest.Question_type
	  
		DB.ConnectDB().Save(&Hr_question)
	  
		c.JSON(200, gin.H{
		  "view": Hr_question,
		})
	  }
	  func Delete_question(c *gin.Context) {
		id := c.Param("id")
		var question structs.Hr_job_questions
		DB.ConnectDB().Where("id=?", id).Delete(&question)
		c.JSON(200, gin.H{
		  "view": "deleted",
		})
	  }
	  
	  func Create_tags(c *gin.Context) {
		var tags structs.Hr_tags	  
		if err := c.ShouldBindJSON(&tags); err != nil {
		  c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		  return
		}
		DB.ConnectDB().Create(&tags)
		c.JSON(200, gin.H{
		  "view": tags,
		})
	  }
	  func Show_tags(c *gin.Context) {
		var tags []structs.Hr_tags
		DB.ConnectDB().Find(&tags)
		c.JSON(200, gin.H{
		  "view": tags,
		})
	  }
	  func Update_tag(c *gin.Context) {
		var tagRequest structs.Tag_update
		if err := c.ShouldBindJSON(&tagRequest); err != nil {
		  c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		  return
		}
		id := c.Param("id")
		var Hr_tag structs.Hr_tags
		DB.ConnectDB().Where("id=?", id).First(&Hr_tag)
	  
		Hr_tag.Name = tagRequest.Name
	  
		DB.ConnectDB().Save(&Hr_tag)
	  
		c.JSON(200, gin.H{
		  "view": Hr_tag,
		})
	  }
	  func Delete_tag(c *gin.Context) {
		id := c.Param("id")
		var tag structs.Hr_tags
		DB.ConnectDB().Where("id=?", id).Delete(&tag)
		c.JSON(200, gin.H{
		  "view": "deleted",
		})
	  }


	  func Create_applicants(c *gin.Context) {
		var applicants structs.Hr_applicants
		CV:=c.PostForm("CV")
		if err := c.ShouldBindJSON(&applicants); err != nil {
		  c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		  return
		}
		DB.ConnectDB().Create(&applicants)
		c.JSON(200, gin.H{
		  "view": applicants,
		  "CV":CV,
		})
	  }
	  func Show_applicants(c *gin.Context) {
		var applicants []structs.Hr_applicants
		DB.ConnectDB().Find(&applicants)
		c.JSON(200, gin.H{
		  "view": applicants,
		})
	  }
	  func Update_applicant(c *gin.Context) {
		var applicantRequest structs.Applicant_update
		CV:=c.PostForm("CV")
		if err := c.ShouldBindJSON(&applicantRequest); err != nil {
		  c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		  return
		}
		id := c.Param("id")
		var Hr_applicant structs.Hr_applicants
		DB.ConnectDB().Where("id=?", id).First(&Hr_applicant)
	  
		Hr_applicant.Name = applicantRequest.Name
		//Hr_applicant.CV = applicantRequest.CV
		Hr_applicant.Is_archive = applicantRequest.Is_archive
		Hr_applicant.Is_favourit = applicantRequest.Is_favourit

		DB.ConnectDB().Save(&Hr_applicant)
	  
		c.JSON(200, gin.H{
		  "view": Hr_applicant,
		  "CV":CV,
		})
	  }
	  func Delete_applicant(c *gin.Context) {
		id := c.Param("id")
		var applicant structs.Hr_applicants
		DB.ConnectDB().Where("id=?", id).Delete(&applicant)
		c.JSON(200, gin.H{
		  "view": "deleted",
		})
	  }


	  func Create_options(c *gin.Context) {
		var options structs.Hr_question_options
		if err := c.ShouldBindJSON(&options); err != nil {
		  c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		  return
		}
		DB.ConnectDB().Create(&options)
		c.JSON(200, gin.H{
		  "view": options,
		})
	  }
	  
	  func Show_options(c *gin.Context) {
		var options []structs.Hr_question_options
		DB.ConnectDB().Find(&options)
		c.JSON(200, gin.H{
		  "view": options,
		})
	  }

	  func Create_interviews(c *gin.Context) {
		var interview structs.Hr_applicant_interviews
		if err := c.ShouldBindJSON(&interview); err != nil {
		  c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		  return
		}
		DB.ConnectDB().Create(&interview)
		c.JSON(200, gin.H{
		  "view": interview,
		})
	  }
	  func Show_interviews(c *gin.Context) {
		var interview []structs.Hr_applicant_interviews
		DB.ConnectDB().Find(&interview)
		c.JSON(200, gin.H{
		  "view": interview,
		})
	  }
	  func Update_interview(c *gin.Context) {
		var interviewRequest structs.Interview_update
		if err := c.ShouldBindJSON(&interviewRequest); err != nil {
		  c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		  return
		}
		id := c.Param("id")
		var Hr_interview structs.Hr_applicant_interviews
		DB.ConnectDB().Where("id=?", id).First(&Hr_interview)
	  
		Hr_interview.Interview_mothed = interviewRequest.Interview_mothed
		Hr_interview.DateTime = interviewRequest.DateTime
		Hr_interview.Responsible = interviewRequest.Responsible
		Hr_interview.NB = interviewRequest.NB


		DB.ConnectDB().Save(&Hr_interview)
	  
		c.JSON(200, gin.H{
		  "view": Hr_interview,
		})
	  }
	  func Delete_interview(c *gin.Context) {
		id := c.Param("id")
		var interview structs.Hr_applicant_interviews
		DB.ConnectDB().Where("id=?", id).Delete(&interview)
		c.JSON(200, gin.H{
		  "view": "deleted",
		})
	  }

	  
	  func Create_notifications(c *gin.Context) {
		var notification structs.Notifications
		if err := c.ShouldBindJSON(&notification); err != nil {
		  c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		  return
		}
		DB.ConnectDB().Create(&notification)
		c.JSON(200, gin.H{
		  "view": notification,
		})
	  }
	  func Show_notification_to(c *gin.Context) {
		var notification []structs.Notifications
		User_to:=c.Param("User_to")
		DB.ConnectDB().Where("User_to=?", User_to).Find(&notification)
		c.JSON(200, gin.H{
		  "view": notification,
		})
	  }
	  func Show_notification_from(c *gin.Context) {
		var notification []structs.Notifications
		User_from:=c.Param("User_from")
		DB.ConnectDB().Where("User_from=?", User_from).Find(&notification)
		c.JSON(200, gin.H{
		  "view": notification,
		})
	  }
	  func Update_notification(c *gin.Context) {
		var notificationRequest structs.Notification_update
		if err := c.ShouldBindJSON(&notificationRequest); err != nil {
		  c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		  return
		}
		id := c.Param("id")
		var notification structs.Notifications
		DB.ConnectDB().Where("id=?", id).First(&notification)
	  
		notification.Message = notificationRequest.Message
		notification.User_to = notificationRequest.User_to
		notification.User_from = notificationRequest.User_from
		notification.Category = notificationRequest.Category

		DB.ConnectDB().Save(&notification)
	  
		c.JSON(200, gin.H{
		  "view": notification,
		})
	  }
	  func Delete_notification(c *gin.Context) {
		id := c.Param("id")
		var notification structs.Notifications
		DB.ConnectDB().Where("id=?", id).Delete(&notification)
		c.JSON(200, gin.H{
		  "view": "deleted",
		})
	  }
 