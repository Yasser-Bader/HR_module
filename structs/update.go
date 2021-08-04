package structs

type User_update struct{
	Name        string      `json:"name" gorm:"type:varchar(100)"`    
	Email       string	    `json:"email" gorm:"type:varchar(50)"`
	Password    string		`json:"password" gorm:"type:varchar(50)"`
}


type Job_update struct {
	Name            string `json:"name" gorm:"type:varchar(100)"`
	Salary          int    `json:"salary" gorm:"type:decimal(10,2)"`
	Description     string `json:"description" gorm:"type:varchar(300)"`
	Vailable_date   int    `json:"vailable_date" gorm:"type:date"`
	Is_show_in_home bool   `json:"is_show_in_home" gorm:"type:bool"`
	Is_archive      bool   `json:"is_archive" gorm:"type:bool"`
	Is_question     bool   `json:"is_question" gorm:"type:bool"`
	Send_alert      string `json:"send_alert" gorm:"type:varchar(100)"`
  }
  
  type Question_update struct {
	Question      string `json:"question" gorm:"type:varchar(100)"`
	Question_type string `json:"question_type" gorm:"type:enum('date','text','choose')"`
  }

  type Tag_update struct{
	Name string 	`json:"name" gorm:"type:varchar(50)"`
  }

  type Applicant_update struct{
	Hr_job_id	int		`json:"hr_job_id" gorm:"type:int"`
	Name		string	`json:"name" gorm:"type:varchar(100)"`
	CV 			string	`json:"cv" gorm:"type:binary"`
	Is_archive  bool	`json:"is_archive" gorm:"type:bool"`
	Is_favourit bool	`json:"is_favourit" gorm:"type:bool"`
  }

  type Notification_update struct{
	Message		string	`json:"message" gorm:"type:varchar(200)"`
	User_to		int		`json:"user_to" gorm:"type:int"`
	User_from	int		`json:"user_from" gorm:"type:int"`
	Category 	string	`json:"gategory" gorm:"type:enum('','')"`
  }

  type Interview_update struct{

	Interview_mothed	string	`json:"interview_mothed" gorm:"type:enum('online','in the place')"`
	DateTime			string	`json:"datetime" gorm:"type:DATETIME"`
	Responsible			string	`json:"Responsible" gorm:"type:varchar(50)"`
	NB					string	`json:"nb" gorm:"type:varchar(200)"`

  }
  ////////////////////////////////////////////////

  type HR_employeeUpdate struct{
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

  type HR_departmentUpdate struct{
	Name       string    `json:"name" gorm:"type:varchar(100)"`
  }

  type HR_discountUpdate struct{
	Discount       string    `json:"discount" gorm:"type:varchar(50)"`
	Reason         string    `json:"reason" gorm:"type:varchar(100)"`
  }

  type HR_loationcUpdate struct{
	Street_address string    `json:"street_address" gorm:"type:"varchar(50)`
	Postal_code    string    `json:"postal_code" gorm:"type:"varchar(50)`
	City           string    `json:"city" gorm:"type:"varchar(50)`
	State_province string    `json:"state_province" gorm:"type:"varchar(50)`
  }

  type HR_countrieUpdate struct{
	Country_name string      `json:"country_name" gorm:"type:varchar(50)"`
  }

  type HR_regionUpdate struct{
	Region_name string      `json:"region_name" gorm:"type:varchar(50)"`
  }