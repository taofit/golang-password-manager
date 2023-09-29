package gui

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"github.com/taofit/golang-password-manager/internal/account"
	"github.com/taofit/golang-password-manager/internal/category"
)

func (g *gui) generateLogin() *fyne.Container {
	logo := getLogo()

	loginLabel := widget.NewLabelWithStyle("LOG IN", fyne.TextAlignLeading, fyne.TextStyle{Bold: true})
	loginLabel.Importance = widget.HighImportance
	registerBtn := widget.NewButton("go to register", func() {
		g.makeAppContentView(g.generateRegister)
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
		g.makeAppContentView(g.generateCategory)
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

func (g *gui) generateRegister() *fyne.Container {
	logo := getLogo()

	registerLabel := widget.NewLabelWithStyle("REGISTER", fyne.TextAlignLeading, fyne.TextStyle{Bold: true})
	registerLabel.Importance = widget.HighImportance

	loginBtn := widget.NewButton("go to log in", func() {
		g.makeAppContentView(g.generateLogin)
	})
	loginAndRgt := container.NewBorder(nil, nil, registerLabel, loginBtn, widget.NewSeparator())

	userName := widget.NewEntry()
	userName.SetPlaceHolder("enter user name")
	password := widget.NewPasswordEntry()
	password.SetPlaceHolder("enter user password")
	btn := widget.NewButton("register", func() {
		name := "good name"
		password := "superb password"
		account, err := account.SaveAccount(name, password)
		if err != nil {
			log.Println("error from creating account", account, err)
		}
		g.setAccount(account.Name, account.Password, account.CategoryList)
		category.SaveCategories(name)
		g.makeAppContentView(g.generateCategory)
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
