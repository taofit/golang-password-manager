package gui

import (
	"errors"
	"image/color"
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
	top := makeBanner()
	left := widget.NewLabel("Left")
	right := widget.NewLabel("Right")
	directory := widget.NewLabelWithData(g.title)
	content := container.NewStack(canvas.NewRectangle(color.Gray{Y: 0xEE}), directory)

	// return container.NewBorder(makeBanner(), nil, left, right, content)
	dividers := [3]fyne.CanvasObject{
		widget.NewSeparator(), widget.NewSeparator(), widget.NewSeparator(),
	}
	objs := []fyne.CanvasObject{content, top, left, right, dividers[0], dividers[1], dividers[2]}

	return container.New(newAppLayout(top, left, right, content, dividers), objs...)
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
		log.Println(dir)
		if dir == nil {
			return
		}
		g.OpenFolder(dir)
	}, g.win)
}

func (g *gui) OpenFolder(dir fyne.ListableURI) {
	name := dir.Name()
	// g.win.SetTitle("Password App: " + name)
	g.title.Set(name)
}

func (g *gui) MakeMenu() *fyne.MainMenu {
	items := fyne.NewMenu(
		"File",
		fyne.NewMenuItem("Open project", g.OpenFolderDialog),
	)

	return fyne.NewMainMenu(items)
}

func (g *gui) ShowAndCreate() {
	var wizard *dialogs.Wizard
	intro := widget.NewLabel(`creating a new project here!!!
or open an existing one that is created earlier 
	`)

	open := widget.NewButton("Open Project", func() {
		wizard.Hide()
		g.OpenFolderDialog()
	})
	create := widget.NewButton("Create Project", func() {
		wizard.Push("Project details", g.makeCreateDetail(wizard))
	})
	create.Importance = widget.HighImportance

	buttons := container.NewGridWithColumns(2, open, create)

	homeContent := container.NewVBox(intro, buttons)
	wizard = dialogs.NewWizard("create project", homeContent)
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
