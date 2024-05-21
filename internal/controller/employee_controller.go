package controller

import (
	"github.com/Facalder/Planify/models"
	"github.com/Facalder/Planify/pkg"
	"github.com/pterm/pterm"
	"time"
)

func LoginEmployee(employees models.TabEmployee, n int) bool {
	var isValidated bool = false

	userName, _ := pterm.DefaultInteractiveTextInput.Show("Please input your UserName")
	password, _ := pterm.DefaultInteractiveTextInput.Show("Please input your Password")

	for i := 0; i < n; i++ {
		if employees[i].UserName == userName && employees[i].Password == password {
			isValidated = true
			return isValidated
		}
	}

	return isValidated
}

func RegisterEmployee(employees *models.TabEmployee, n *int, menu func()) {
	if *n < pkg.NMAX {
		fullName, _ := pterm.DefaultInteractiveTextInput.Show("Please input your FullName")
		shortName, _ := pterm.DefaultInteractiveTextInput.Show("Please input your ShortName")
		userName, _ := pterm.DefaultInteractiveTextInput.Show("Please input your UserName")
		password, _ := pterm.DefaultInteractiveTextInput.WithMask("*").Show("Please input your Password")
		bio, _ := pterm.DefaultInteractiveTextInput.Show("Please input your Bio")
		department, _ := pterm.DefaultInteractiveTextInput.Show("Please input your Department")

		employee := models.EmployeeModel{
			Id:         *n,
			FullName:   fullName,
			ShortName:  shortName,
			UserName:   userName,
			Password:   password,
			Bio:        bio,
			Department: department,
			CreatedAt:  time.Now(),
		}

		employees[*n] = employee
		*n += 1

		menu()
	} else {
		pterm.Error.Println("You Have Reached The Maksimum Data Allowed!")
	}
}

func CreateEmployee(employees *models.TabEmployee, n *int) {
	if *n < pkg.NMAX {
		fullName, _ := pterm.DefaultInteractiveTextInput.Show("Please input your FullName")
		shortName, _ := pterm.DefaultInteractiveTextInput.Show("Please input your ShortName")
		userName, _ := pterm.DefaultInteractiveTextInput.Show("Please input your UserName")
		password, _ := pterm.DefaultInteractiveTextInput.WithMask("*").Show("Please input your Password")
		bio, _ := pterm.DefaultInteractiveTextInput.Show("Please input your Bio")
		department, _ := pterm.DefaultInteractiveTextInput.Show("Please input your Department")
		manager, _ := pterm.DefaultInteractiveTextInput.Show("Please input your Manager")

		employee := models.EmployeeModel{
			Id:         *n,
			FullName:   fullName,
			ShortName:  shortName,
			UserName:   userName,
			Password:   password,
			Bio:        bio,
			Department: department,
			Manager:    manager,
			CreatedAt:  time.Now(),
		}

		employees[*n] = employee
		*n += 1
	} else {
		pterm.Error.Println("You Have Reached The Maksimum Data Allowed!")
	}
}

func UpdateEmployee() {}

func ShowAllEmployee() {}

func DeleteEmployee() {}
