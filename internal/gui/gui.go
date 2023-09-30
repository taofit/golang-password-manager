package gui

import (
	"image/color"
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
	win                   fyne.Window
	title                 binding.String
	account               Account
	categoryListBindData  binding.ExternalStringList
	cateEntryListBindData binding.ExternalStringList
}

type Account struct {
	accName       string
	password      string //TODO will be removed
	categoryList  []string
	selectedCate  string
	cateEntryList []string
}

func NewGui(win fyne.Window) *gui {
	return &gui{win: win, title: binding.NewString()}
}

func (g *gui) MakeGUI() {
	g.bindWindowTitle()
	g.makeAppContentView(g.generateLoginArea)
}

func (g *gui) bindWindowTitle() {
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

func (g *gui) makeAppContentView(entryFunc func() *fyne.Container) {
	top := makeBanner()
	content := container.NewStack()
	content.Objects = []fyne.CanvasObject{canvas.NewRectangle(color.Gray{Y: 0xEE}), entryFunc()}

	g.win.SetContent(container.NewBorder(top, nil, nil, nil, content))
}

func (g *gui) setAccount(name string, password string, cateList []string) {
	g.account = Account{accName: name, password: password, categoryList: cateList}
}

func (g *gui) updateSelectedCateById(cateId int) {
	cateName := g.account.categoryList[cateId]
	g.account.selectedCate = cateName
}

func (g *gui) ShowAndRun() {

}
