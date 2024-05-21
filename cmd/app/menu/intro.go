package menu

import (
	"github.com/Facalder/Planify/pkg"
	"github.com/Facalder/Planify/pkg/component"
	"github.com/pterm/pterm"
	"time"
)

func IntroScreen() {
	component.HeaderBigText("planify")

	component.Header("Welcome to Planify!, the Task Management App - Made with love for you!")

	pterm.Info.Println("This project was generated with the latest version of PTerm and Golang!" +
		"\nPlanify works on nearly every terminal and operating system." +
		"\nIt's super easy to use!" +
		"\nIf you want the sourcecode, hmm this project is not open source 😭" +
		"\nYou can see the code of this demo in the " + pterm.LightMagenta("github.com/Facalder/Planify") +
		"\nThis app was updated at: " + pterm.Green(time.Now().Format("02 Jan 2006 - 15:04:05 MST")))

	component.Spinner(pkg.InitApp, "Waiting init for 3 seconds", func() {
		InitialMenu()
	})
}
