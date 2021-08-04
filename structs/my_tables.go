package structs

import (
	"gorm.io/gorm"
)
type Hr_users struct{
	gorm.Model
	Name        string      `json:"name" gorm:"type:varchar(100)"`    
	Email       string	    `json:"email" gorm:"type:varchar(50)"`
	Password    string		`json:"password" gorm:"type:varchar(50)"`
}

type Hr_jobs struct{
	gorm.Model
	Name            string	`json:"name" gorm:"type:varchar(100)"`
	Salary          int		`json:"salary" gorm:"type:decimal(10,2)"`
	Description     string	`json:"description" gorm:"type:varchar(300)"`
	Vailable_date   int		`json:"vailable_date" gorm:"type:date"`
	Is_show_in_home bool	`json:"is_show_in_home" gorm:"type:bool"`
	Is_archive      bool 	`json:"is_archive" gorm:"type:bool"`
	Is_question     bool 	`json:"is_question" gorm:"type:bool"`
	Send_alert      string	`json:"send_alert" gorm:"type:varchar(100)"`
	User_id         int 	`json:"user_id" gorm:"type:int"`
}

type Hr_jobs_tags struct{
	gorm.Model
	Hr_job_id  int	`json:"hr_job_id" gorm:"type:int"`
	Hr_tag_id  int	`json:"hr_tag_id" gorm:"type:int"`
}

type Hr_tags struct{
	gorm.Model
	Name string 	`json:"name" gorm:"type:varchar(50)"`
}

type Hr_job_questions struct{
	gorm.Model
	Question 		string	`json:"question" gorm:"type:varchar(100)"`
	Question_type 	string	`json:"question_type" gorm:"type:enum('date','text','choose')"`
	Hr_job_id 		int 	`json:"hr_job_id" gorm:"type:int"`
}

type Hr_question_options struct{
	gorm.Model
	Hr_applicant_id int		`json:"hr_applicant_id" gorm:"type:int"`
	Option 			string	`json:"option" gorm:"type:json"`
}

type Hr_applicants struct{
	gorm.Model
	Hr_job_id	int		`json:"hr_job_id" gorm:"type:int"`
	Name		string	`json:"name" gorm:"type:varchar(100)"`
	CV 			string	`json:"cv" gorm:"type:binary"`
	Is_archive  bool	`json:"is_archive" gorm:"type:bool"`
	Is_favourit bool	`json:"is_favourit" gorm:"type:bool"`
}

type Notifications struct{
	gorm.Model
	Message		string	`json:"message" gorm:"type:varchar(200)"`
	User_to		int		`json:"user_to" gorm:"type:int"`
	User_from	int		`json:"user_from" gorm:"type:int"`
	Category 	string	`json:"gategory" gorm:"type:enum('','')"`
}

type Hr_applicant_interviews struct{
	gorm.Model
	Interview_mothed	string	`json:"interview_mothed" gorm:"type:enum('online','in the place')"`
	DateTime			string	`json:"datetime" gorm:"type:DATETIME"`
	Responsible			string	`json:"Responsible" gorm:"type:varchar(50)"`
	NB					string	`json:"nb" gorm:"type:varchar(200)"`

}
////////////////////////////////////////////////////////////////////////////////////////////////////////////////
type HR_employees struct {
	gorm.Model
	Name                     	string    `json:"name" gorm:"type:varchar(100)"`
	HR_employee_department_id   int       `json:"hr_employee_department_id"`
	HR_employee_applicant_id 	int       `json:"hr_employee_applicant_id"`
	Data_hiring              	string    `json:"data_hiring" gorm:"type:date"`
	Salary                   	int       `json:"salary" gorm:"type:decimal(10,2)"`
	HR_employee_location_id  	int       `json:"hr_employee_location_id"`
	Job_title                	string    `json:"job_title " gorm:"type:varchar(100)"`
	Work_hour                	int       `json:"work_hour" gorm:"type:int(10)"`
	Is_block                 	bool      `json:"is_block" gorm:"type:bool"`	
}

type HR_employee_departments struct {
	gorm.Model
	Name       string    `json:"name" gorm:"type:varchar(100)"`

}

type HR_employee_attendance_and_departure struct {
	gorm.Model
	Is_presence    bool    `json:"is_presence" gorm:"type:bool"`
	Is_departure   bool    `json:"is_departure" gorm:"type:bool"`
	Is_absence     bool    `json:"is_absence" gorm:"type:bool"`
	HR_employee_id int     `json:"hr_employee_id" gorm:"type:int"`

}

type HR_employee_discounts struct {
	gorm.Model
	Discount       string    `json:"discount" gorm:"type:varchar(50)"`
	Reason         string    `json:"reason" gorm:"type:varchar(100)"`
	HR_employee_id int       `json:"hr_employee_id" gorm:"int"`

}

type HR_employee_locations struct {
	gorm.Model
	Street_address string    `json:"street_address" gorm:"type:varchar(50)"`
	Postal_code    string    `json:"postal_code" gorm:"type:varchar(50)"`
	City           string    `json:"city" gorm:"type:varchar(50)"`
	State_province string    `json:"state_province" gorm:"type:varchar(50)"`
	Country_id     int       `json:"country_id"`
}

type Countries struct {
	gorm.Model
	Country_name string      `json:"country_name" gorm:"type:varchar(50)"`
	Region_id    int         `json:"region_id" gorm:"type:int"`
}

type Regions struct {
	gorm.Model
	Region_name string      `json:"region_name" gorm:"type:varchar(50)"`
}
