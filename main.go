package main

import (
	"flag"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/storage"
	"github.com/taofit/golang-password-manager/internal/gui"
)

func main() {
	a := app.New()
	a.Settings().SetTheme(gui.NewAppTheme())
	w := a.NewWindow("Password Manager App")
	w.Resize(fyne.NewSize(1024, 768))
	ui := gui.NewGui(w)
	ui.BindWindowTitle()
	w.SetMainMenu(ui.MakeMenu())
	w.SetContent(ui.MakeGUI())

	flag.Usage = func() {
		fmt.Println("Usage: project[]")
	}
	flag.Parse()
	if len(flag.Args()) > 0 {
		dirPath := flag.Args()[0]
		dirURI := storage.NewFileURI(dirPath)
		dir, err := storage.ListerForURI(dirURI)

		if err != nil {
			fmt.Println("Error opening project: ", err)
			return
		}
		ui.OpenFolder(dir)
	} else {
		ui.ShowAndCreate()
	}

	w.ShowAndRun()
}
