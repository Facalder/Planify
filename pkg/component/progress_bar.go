package component

import "github.com/pterm/pterm"

func ProgressBar(total int, title string, content func()) {
	pb, _ := pterm.DefaultProgressbar.WithTotal(total).WithTitle(title).Start()
	pb.Increment()

	content()
}
