package main

import "github.com/fatih/color"

var (
	informationText = color.New(color.FgBlue)
	successText     = color.New(color.FgGreen)
	dangerText      = color.New(color.FgRed).Add(color.Bold)
)
