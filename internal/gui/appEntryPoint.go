package gui

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"

	"github.com/taofit/golang-password-manager/internal/account"
)

func (g *gui) generateLoginArea() *fyne.Container {
	g.title.Set("Log in")
	logo := getLogo()

	loginLabel := widget.NewLabelWithStyle("LOG IN", fyne.TextAlignLeading, fyne.TextStyle{Bold: true})
	loginLabel.Importance = widget.HighImportance
	registerBtn := widget.NewButton("go to register", func() {
		g.makeAppContentView(g.generateRegisterArea)
	})
	loginAndRgt := container.NewBorder(nil, nil, loginLabel, registerBtn, widget.NewSeparator())

	userName := widget.NewEntry()
	userName.SetPlaceHolder("enter user name")
	password := widget.NewPasswordEntry()
	password.SetPlaceHolder("enter password")
	btn := widget.NewButton("log in", func() {
		// TODO user name and password check, if it is valid, go to the category view
		// userName.Text
		// password.Text
		userName := "good name"
		password := "superb password"
		acc, err := account.GetAccount(userName, password)
		if err != nil {
			dialog.ShowError(err, g.win)
			log.Printf("error from fetching account: %s, %s", userName, err)
			return
		}
		g.setAccount(acc.Name, acc.Password, acc.CategoryList)
		g.makeAppContentView(g.generateCateListArea)
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

func (g *gui) generateRegisterArea() *fyne.Container {
	g.title.Set("Register")
	logo := getLogo()

	registerLabel := widget.NewLabelWithStyle("REGISTER", fyne.TextAlignLeading, fyne.TextStyle{Bold: true})
	registerLabel.Importance = widget.HighImportance

	loginBtn := widget.NewButton("go to log in", func() {
		g.makeAppContentView(g.generateLoginArea)
	})
	loginAndRgt := container.NewBorder(nil, nil, registerLabel, loginBtn, widget.NewSeparator())

	userName := widget.NewEntry()
	userName.SetPlaceHolder("enter user name")
	password := widget.NewPasswordEntry()
	password.SetPlaceHolder("enter user password")
	btn := widget.NewButton("register", func() {
		name := "good name"
		password := "superb password"
		acc, err := account.CreateAccount(name, password)
		if err != nil {
			log.Println("error from creating account", acc, err)
		}
		g.setAccount(acc.Name, acc.Password, acc.CategoryList)
		g.makeAppContentView(g.generateCateListArea)
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
