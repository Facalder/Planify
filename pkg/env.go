package pkg

const (
	NMAX int = 10

	SaveData int = 3

	InitApp int = 3

	ValidateData int = 3

	ExitProgram int = 5

	ReloadProject int = 5

	Loading int = 1
)

var (
	ChooseMenu string
)

func Clear() {
	print("\033[H\033[2J")
}
