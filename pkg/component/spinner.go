package component

import (
	"github.com/Facalder/Planify/pkg"
	"strconv"
	"time"

	"github.com/pterm/pterm"
)

var second = time.Second

func Spinner(duration int, description string, content func()) {
	spin, _ := pterm.DefaultSpinner.WithShowTimer(false).WithRemoveWhenDone(true).Start(
		description)

	time.Sleep(second)

	for i := duration - 1; i > 0; i-- {
		if i > 1 {
			spin.UpdateText("Remaining time: " + strconv.Itoa(i) + " seconds...")
		} else {
			spin.UpdateText("Remaining time: " + strconv.Itoa(i) + " seconds...")
		}

		time.Sleep(second)
	}

	spin.Stop()

	pkg.Clear()
	content()
}
