package watchdog

import "github.com/fatih/color"

var (
	InformationText = color.New(color.FgBlue)
	SuccessText     = color.New(color.FgGreen)
	DangerText      = color.New(color.FgRed).Add(color.Bold)
)
