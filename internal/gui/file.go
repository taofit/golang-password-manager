package gui

import (
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/taofit/golang-password-manager/internal/dialogs"
)

func (g *gui) ShowFileDialog() {
	var wizard *dialogs.Wizard
	intro := widget.NewLabel(`creating a new file here!!!
or open an existing one that is created earlier 
	`)

	open := widget.NewButton("Open File", func() {
		wizard.Hide()
		g.OpenFolderDialog()
	})
	create := widget.NewButton("Create File", func() {
		wizard.Push("File details", g.makeCreateDetail(wizard))
	})
	create.Importance = widget.HighImportance

	buttons := container.NewGridWithColumns(2, open, create)

	homeContent := container.NewVBox(intro, buttons)
	wizard = dialogs.NewWizard("create file", homeContent)
	wizard.Show(g.win)
	wizard.Resize(homeContent.MinSize().AddWidthHeight(40, 80))
}
