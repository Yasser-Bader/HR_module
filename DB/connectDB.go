package DB

import (
        "log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"HR_module/structs"
)

func ConnectDB()*gorm.DB{
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbURL := os.Getenv("DB_URL")

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
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
