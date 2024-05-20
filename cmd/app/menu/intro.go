package menu

import (
	"github.com/Facalder/Planify/pkg/component"
)

func IntroScreen() {
	component.HeaderBigText("planify")

	component.Description("Welcome to Planify the Task Management App - Made with love for you!")

	component.Spinner(3, "Waiting init for 3 seconds", func() {
		InitialMenu()
	})
}
