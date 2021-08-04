package routes

import (

	"net/http"
	"github.com/gin-gonic/gin"
	"HR_module/DB"
	"HR_module/structs"
)
func Create_employees(c *gin.Context) {
		var employees structs.HR_employees
		if err := c.ShouldBindJSON(&employees); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		DB.ConnectDB().Create(&employees)
		c.JSON(200, gin.H{
			"view": employees,
		})
	}
	
func Show_employees(c *gin.Context) {
		var employees []structs.HR_employees
		DB.ConnectDB().Find(&employees)
		c.JSON(200, gin.H{
			"view": employees,
		})
	}

func Update_employee(c *gin.Context) {
		var employeeRequest structs.HR_employeeUpdate
		if err := c.ShouldBindJSON(&employeeRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		id := c.Param("id")
		var employee structs.HR_employees
		DB.ConnectDB().Where("id=?", id).First(&employee)

		employee.Name = employeeRequest.Name
		employee.HR_employee_department_id = employeeRequest.HR_employee_department_id
		employee.HR_employee_applicant_id = employeeRequest.HR_employee_applicant_id
		employee.Data_hiring = employeeRequest.Data_hiring
		employee.Salary = employeeRequest.Salary
		employee.HR_employee_location_id = employeeRequest.HR_employee_location_id
		employee.Job_title = employeeRequest.Job_title
		employee.Work_hour = employeeRequest.Work_hour
		employee.Is_block = employeeRequest.Is_block

		DB.ConnectDB().Save(&employee)

		c.JSON(200, gin.H{
			"view": employee,
		})
	}

	func Delete_employee(c *gin.Context) {
		id := c.Param("id")
		var employee structs.HR_employees
		DB.ConnectDB().Where("id=?", id).Delete(&employee)
		c.JSON(200, gin.H{
			"view": "Deleted",
		})
	}

	func Create_departments(c *gin.Context) {
		var departments structs.HR_employee_departments
		if err := c.ShouldBindJSON(&departments); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		DB.ConnectDB().Create(&departments)
		c.JSON(200, gin.H{
			"view": departments,
		})
	}
	
	func Show_departments(c *gin.Context) {
		var departments []structs.HR_employee_departments
		DB.ConnectDB().Find(&departments)
		c.JSON(200, gin.H{
			"view": departments,
		})
	}

	func Update_department(c *gin.Context) {
		var departmentRequest structs.HR_departmentUpdate
		if err := c.ShouldBindJSON(&departmentRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		id := c.Param("id")
		var department structs.HR_employee_departments
		DB.ConnectDB().Where("id=?", id).First(&department)

		department.Name = departmentRequest.Name

		DB.ConnectDB().Save(&department)

		c.JSON(200, gin.H{
			"view": department,
		})
	}

	func Delete_department(c *gin.Context) {
		id := c.Param("id")
		var department structs.HR_employee_departments
		DB.ConnectDB().Where("id=?", id).Delete(&department)
		c.JSON(200, gin.H{
			"view": "Deleted",
		})
	}

	func Create_attendance_and_departure(c *gin.Context) {
		var attendance_and_departure structs.HR_employee_attendance_and_departure
		if err := c.ShouldBindJSON(&attendance_and_departure); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		DB.ConnectDB().Create(&attendance_and_departure)
		c.JSON(200, gin.H{
			"view": attendance_and_departure,
		})
	}
	
	func Show_attendance_and_departure(c *gin.Context) {
		var attendance_and_departure []structs.HR_employee_attendance_and_departure
		HR_employee_id:=c.Param("HR_employee_id")
		DB.ConnectDB().Where("HR_employee_id=?", HR_employee_id).Find(&attendance_and_departure)
		c.JSON(200, gin.H{
			"view": attendance_and_departure,
		})
	}
	
	func Create_discounts(c *gin.Context) {
		var discounts structs.HR_employee_discounts
		if err := c.ShouldBindJSON(&discounts); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		DB.ConnectDB().Create(&discounts)
		c.JSON(200, gin.H{
			"view": discounts,
		})
	}
	
	func Show_discount(c *gin.Context) {
		var discount []structs.HR_employee_discounts
		HR_employee_id:=c.Param("HR_employee_id")
		DB.ConnectDB().Where("HR_employee_id=?", HR_employee_id).Find(&discount)
		c.JSON(200, gin.H{
			"view": discount,
		})
	}

	func Update_discount(c *gin.Context) {
		var discountRequest structs.HR_discountUpdate
		if err := c.ShouldBindJSON(&discountRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		id := c.Param("id")
		var discount structs.HR_employee_discounts
		DB.ConnectDB().Where("id=?", id).First(&discount)

		discount.Discount = discountRequest.Discount
		discount.Reason = discountRequest.Reason


		DB.ConnectDB().Save(&discount)

		c.JSON(200, gin.H{
			"view": discount,
		})
	}

	func Delete_discount(c *gin.Context) {
		id := c.Param("id")
		var discount structs.HR_employee_discounts
		DB.ConnectDB().Where("id=?", id).Delete(&discount)
		c.JSON(200, gin.H{
			"view": "Deleted",
		})
	}

	func Create_loationcs(c *gin.Context) {
		var loationcs structs.HR_employee_locations
		if err := c.ShouldBindJSON(&loationcs); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		DB.ConnectDB().Create(&loationcs)
		c.JSON(200, gin.H{
			"view": loationcs,
		})
	}
	
	func Show_loationc(c *gin.Context) {
		var loationc []structs.HR_employee_locations
		id:=c.Param("id")
		DB.ConnectDB().Where("id=?", id).Find(&loationc)
		c.JSON(200, gin.H{
			"view": loationc,
		})
	}

	func Update_loationc(c *gin.Context) {
		var loationcRequest structs.HR_loationcUpdate
		if err := c.ShouldBindJSON(&loationcRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		id := c.Param("id")
		var loationc structs.HR_employee_locations
		DB.ConnectDB().Where("id=?", id).First(&loationc)

		loationc.Street_address = loationcRequest.Street_address
		loationc.Postal_code = loationcRequest.Postal_code
		loationc.City = loationcRequest.City
		loationc.State_province = loationcRequest.State_province


		DB.ConnectDB().Save(&loationc)

		c.JSON(200, gin.H{
			"view": loationc,
		})
	}

	func Delete_loationc(c *gin.Context) {
		id := c.Param("id")
		var loationc structs.HR_employee_locations
		DB.ConnectDB().Where("id=?", id).Delete(&loationc)
		c.JSON(200, gin.H{
			"view": "Deleted",
		})
	}

	func Create_countries(c *gin.Context) {
		var countries structs.Countries
		if err := c.ShouldBindJSON(&countries); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		DB.ConnectDB().Create(&countries)
		c.JSON(200, gin.H{
			"view": countries,
		})
	}
	
	func Show_countrie(c *gin.Context) {
		var countrie []structs.Countries
		id:=c.Param("id")
		DB.ConnectDB().Where("id=?", id).Find(&countrie)
		c.JSON(200, gin.H{
			"view": countrie,
		})
	}

	func Update_countrie(c *gin.Context) {
		var countrieRequest structs.HR_countrieUpdate
		if err := c.ShouldBindJSON(&countrieRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		id := c.Param("id")
		var countrie structs.Countries
		DB.ConnectDB().Where("id=?", id).First(&countrie)

		countrie.Country_name = countrieRequest.Country_name

		DB.ConnectDB().Save(&countrie)

		c.JSON(200, gin.H{
			"view": countrie,
		})
	}

	func Delete_countrie(c *gin.Context) {
		id := c.Param("id")
		var countrie structs.Countries
		DB.ConnectDB().Where("id=?", id).Delete(&countrie)
		c.JSON(200, gin.H{
			"view": "Deleted",
		})
	}

	func Create_regions(c *gin.Context) {
		var regions structs.Regions
		if err := c.ShouldBindJSON(&regions); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		DB.ConnectDB().Create(&regions)
		c.JSON(200, gin.H{
			"view": regions,
		})
	}
	
	func Show_region(c *gin.Context) {
		var region []structs.Regions
		id:=c.Param("id")
		DB.ConnectDB().Where("id=?", id).Find(&region)
		c.JSON(200, gin.H{
			"view": region,
		})
	}

	func Update_region(c *gin.Context) {
		var regionRequest structs.HR_regionUpdate
		if err := c.ShouldBindJSON(&regionRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		id := c.Param("id")
		var region structs.Regions
		DB.ConnectDB().Where("id=?", id).First(&region)

		region.Region_name = regionRequest.Region_name
		
		DB.ConnectDB().Save(&region)

		c.JSON(200, gin.H{
			"view": region,
		})
	}

	func Delete_region(c *gin.Context) {
		id := c.Param("id")
		var region structs.Regions
		DB.ConnectDB().Where("id=?", id).Delete(&region)
		c.JSON(200, gin.H{
			"view": "Deleted",
		})
	}
