package gui

import (
	"errors"
	"log"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/taofit/golang-password-manager/internal/dialogs"
	"github.com/taofit/golang-password-manager/internal/project"
)

type gui struct {
	win   fyne.Window
	title binding.String
}

func NewGui(win fyne.Window) *gui {
	return &gui{win: win, title: binding.NewString()}
}

func (g *gui) MakeGUI() fyne.CanvasObject {
	return appEntryContent(g.createLogin)
}

func (g *gui) BindWindowTitle() {
	g.title.AddListener(binding.NewDataListener(func() {
		name, _ := g.title.Get()
		g.win.SetTitle("Password Manager App: " + name)
	}))
}

func makeBanner() fyne.CanvasObject {
	toolbar := widget.NewToolbar(
		widget.NewToolbarAction(theme.MenuIcon(), func() {
			log.Println("toolbar icon is clicked")
		}),
	)
	logo := canvas.NewImageFromResource(resourceLogoPng)
	logo.FillMode = canvas.ImageFillContain

	return container.NewStack(toolbar, container.NewPadded(logo))
}

func (g *gui) OpenFolderDialog() {
	dialog.ShowFolderOpen(func(dir fyne.ListableURI, err error) {
		if err != nil {
			dialog.ShowError(err, g.win)
			return
		}
		if dir == nil {
			return
		}
		g.OpenFolder(dir)
	}, g.win)
}

func (g *gui) OpenFolder(dir fyne.ListableURI) {
	name := dir.Name()
	g.title.Set(name)
}

func (g *gui) MakeMenu() *fyne.MainMenu {
	items := fyne.NewMenu(
		"File",
		fyne.NewMenuItem("Open file", g.OpenFolderDialog),
	)

	return fyne.NewMainMenu(items)
}

func (g *gui) ShowAndRun() {

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
