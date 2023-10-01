package gui

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type gui struct {
	win   fyne.Window
	title binding.String
}

func NewGui(win fyne.Window) *gui {
	return &gui{win: win, title: binding.NewString()}
}

func (g *gui) MakeGUI() {
	g.makeAppContentView(g.createLogin)
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

func (g *gui) openFolderDialog() {
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
		fyne.NewMenuItem("Open file", g.openFolderDialog),
	)

	return fyne.NewMainMenu(items)
}

func (g *gui) ShowAndRun() {

}
