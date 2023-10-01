package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"

	"github.com/taofit/golang-password-manager/internal/gui"
)

func main() {
	a := app.NewWithID("com.passportal.app")
	a.Settings().SetTheme(gui.NewAppTheme())
	w := a.NewWindow("Password Manager App")
	w.Resize(fyne.NewSize(1024, 768))

	ui := gui.NewGui(w)
	w.SetMainMenu(ui.MakeMenu())
	ui.MakeGUI()
	ui.ShowAndRun()

	w.ShowAndRun()
}
