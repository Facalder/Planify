package controller

import (
	"fmt"
	"github.com/Facalder/Planify/internal/repositories"
	"github.com/Facalder/Planify/models"
	"github.com/pterm/pterm"
	"golang.org/x/exp/rand"
	"time"
)

func LoginAdmin() bool {
	var isValidated bool = false

	userName, _ := pterm.DefaultInteractiveTextInput.Show("Please input your UserName")
	password, _ := pterm.DefaultInteractiveTextInput.Show("Please input your Password")

	admins := repositories.GetAllAdmin()

	for i := 0; i < len(admins); i++ {
		if admins[i].UserName == userName && admins[i].Password == password {
			isValidated = true
			return isValidated
		}
	}

	return isValidated
}

func RegisterAdmin(menu func()) {

	fullName, _ := pterm.DefaultInteractiveTextInput.Show("Please input your FullName")
	shortName, _ := pterm.DefaultInteractiveTextInput.Show("Please input your ShortName")
	userName, _ := pterm.DefaultInteractiveTextInput.Show("Please input your UserName")
	password, _ := pterm.DefaultInteractiveTextInput.Show("Please input your Password")

	rand.Seed(uint64(time.Now().UnixNano()))
	randomId := int64(rand.Intn(10))

	admin := models.AdminModel{
		Id:        randomId,
		FullName:  fullName,
		ShortName: shortName,
		UserName:  userName,
		Password:  password,
	}

	createdAdmin := repositories.CreateAdmin(admin)

	menu()

	fmt.Printf("Admin registered successfully: %+v\n", createdAdmin)
}

func GetAllAdmin() {}

func UpdateAdmin() {}

func DeleteAdmin() {}
