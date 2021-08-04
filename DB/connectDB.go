package DB

import (
	"fmt"
	"log"
	"os"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"HR_module/structs"
)

func ConnectDB()*gorm.DB{
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbUser := os.Getenv("DB_User")
	dbPass := os.Getenv("DB_Pass")
	dbPort := os.Getenv("DB_Port")
	dbHost := os.Getenv("DB_Host")
	dbName := os.Getenv("DB_Name")
	// Connect DB
	dbn:=fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",dbUser,dbPass,dbHost,dbPort,dbName)
	db, err := gorm.Open(mysql.Open(dbn), &gorm.Config{})
    if err != nil {
		log.Fatal(err)
	}
	// create tables
	db.AutoMigrate(&structs.Hr_users{},&structs.Hr_jobs{},&structs.Hr_jobs_tags{},&structs.Hr_tags{},
		&structs.Hr_job_questions{},&structs.Hr_question_options{},&structs.Hr_applicants{},&structs.Notifications{},
		&structs.Hr_applicant_interviews{},&structs.HR_employees{},&structs.HR_employee_departments{},
		&structs.HR_employee_attendance_and_departure{},&structs.HR_employee_discounts{},&structs.HR_employee_locations{},
		&structs.Countries{},&structs.Regions{})

		return db
}