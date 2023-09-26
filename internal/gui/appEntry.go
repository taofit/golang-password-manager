package gui

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func (g *gui) createLogin() *fyne.Container {
	logo := getLogo()

	loginLabel := widget.NewLabelWithStyle("LOG IN", fyne.TextAlignLeading, fyne.TextStyle{Bold: true})
	loginLabel.Importance = widget.HighImportance
	registerBtn := widget.NewButton("go to register", func() {
		g.makeAppContentView(g.createRegister)
	})
	loginAndRgt := container.NewBorder(nil, nil, loginLabel, registerBtn, widget.NewSeparator())

	userName := widget.NewEntry()
	userName.SetPlaceHolder("enter user name")
	password := widget.NewPasswordEntry()
	password.SetPlaceHolder("enter password")
	btn := widget.NewButton("log in", func() {
		// userName.Text
		// password.Text
		g.makeAppContentView(g.createCategoryContent)
	})

	objects := []fyne.CanvasObject{
		logo,
		loginAndRgt,
		userName,
		password,
		btn,
	}
	loginArea := container.NewVBox(objects...)

	return container.NewCenter(loginArea)
}

func (g *gui) createRegister() *fyne.Container {
	logo := getLogo()

	registerLabel := widget.NewLabelWithStyle("REGISTER", fyne.TextAlignLeading, fyne.TextStyle{Bold: true})
	registerLabel.Importance = widget.HighImportance

	loginBtn := widget.NewButton("go to log in", func() {
		g.makeAppContentView(g.createLogin)
	})
	loginAndRgt := container.NewBorder(nil, nil, registerLabel, loginBtn, widget.NewSeparator())

	userName := widget.NewEntry()
	userName.SetPlaceHolder("enter user name")
	password := widget.NewPasswordEntry()
	password.SetPlaceHolder("enter user password")
	btn := widget.NewButton("register", func() {

	})

	objects := []fyne.CanvasObject{
		logo,
		loginAndRgt,
		userName,
		password,
		btn,
	}
	registerArea := container.NewVBox(objects...)

	return container.NewCenter(registerArea)
}

func (g *gui) makeAppContentView(entryFunc func() *fyne.Container) {
	top := makeBanner()
	content := container.NewStack()
	content.Objects = []fyne.CanvasObject{canvas.NewRectangle(color.Gray{Y: 0xEE}), entryFunc()}

	g.win.SetContent(container.NewBorder(top, nil, nil, nil, content))
}
