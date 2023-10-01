package gui

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	c "github.com/taofit/golang-password-manager/internal/category"
)

func (g *gui) getEntryForm(formDialog dialog.Dialog) *widget.Form {
	cateName := g.account.selectedCate
	form := widget.NewForm()
	switch cateName {
	case c.Website:
		form = g.getWebsiteEntryForm(formDialog)
	case c.CreditCard:
		form = g.getCreditCardEntryForm()
	}
	return form
}

func (g *gui) getWebsiteEntryForm(formDialog dialog.Dialog) *widget.Form {
	name := widget.NewFormItem("name", widget.NewEntry())
	password := widget.NewFormItem("password", widget.NewPasswordEntry())
	form := widget.NewForm(name, password)

	form.OnCancel = func() {
		g.makeAppContentView(g.generateCateEntryListArea)
	}
	form.OnSubmit = func() {
		// c.NewWebsiteEntry(name.Text, password.Text)
		c.SaveCateEntry(g.account.accName, g.account.selectedCate, "category item name")
		g.makeAppContentView(
			func() *fyne.Container {
				return g.generateCateEntryListArea()
			})
		// formDialog.Hide()
	}
	return form
}

func (g *gui) getCreditCardEntryForm() *widget.Form {
	name := widget.NewFormItem("name", widget.NewEntry())
	password := widget.NewFormItem("password", widget.NewPasswordEntry())
	form := widget.NewForm(name, password)

	form.OnCancel = func() {
		g.makeAppContentView(g.generateCateEntryListArea)
	}
	form.OnSubmit = func() {
		log.Println("submitting...")

		// c.NewWebsiteEntry(name.Text, password.Text)
		c.SaveCateEntry(g.account.accName, g.account.selectedCate, "category item name")
		g.makeAppContentView(
			func() *fyne.Container {
				return g.generateCateEntryListArea()
			})
	}

	return form
}
