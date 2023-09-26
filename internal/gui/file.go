package gui

import (
	"errors"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
	"github.com/taofit/golang-password-manager/internal/dialogs"
	"github.com/taofit/golang-password-manager/internal/project"
)

func (g *gui) ShowFileDialog() {
	var wizard *dialogs.Wizard
	intro := widget.NewLabel(`creating a new file here!!!
or open an existing one that is created earlier 
	`)

	open := widget.NewButton("Open File", func() {
		wizard.Hide()
		g.openFolderDialog()
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

func (g *gui) makeCreateDetail(wizard *dialogs.Wizard) fyne.CanvasObject {
	homeDir, _ := os.UserHomeDir()
	parent := storage.NewFileURI(homeDir)
	chosen, _ := storage.ListerForURI(parent)
	name := widget.NewEntry()
	name.Validator = func(s string) error {
		if s == "" {
			return errors.New("project name is required")
		}
		return nil
	}
	var dir *widget.Button
	dir = widget.NewButton(chosen.Name(), func() {
		d := dialog.NewFolderOpen(func(lu fyne.ListableURI, err error) {
			if err != nil || lu == nil {
				return
			}
			chosen = lu
			dir.SetText(lu.Name())
		}, g.win)
		d.SetLocation(chosen)
		d.Show()
	})
	form := widget.NewForm(
		widget.NewFormItem("Name", name),
		widget.NewFormItem("parent Directory", dir),
	)

	form.OnSubmit = func() {
		project, err := project.CreateProject(name.Text, chosen)
		if err != nil {
			dialog.ShowError(err, g.win)
			return
		}
		wizard.Hide()
		g.OpenFolder(project)
	}

	return form
}
