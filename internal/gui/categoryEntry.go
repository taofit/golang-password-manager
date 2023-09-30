package gui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	c "github.com/taofit/golang-password-manager/internal/category"
)

func (g *gui) generateEntryListArea() *fyne.Container {
	g.title.Set(fmt.Sprintf("%s list", g.account.selectedCate))
	accName := g.account.accName
	entryList, err := c.FetchEntryList(accName, g.account.selectedCate)
	if err != nil {
		contentArea := widget.NewLabel(err.Error())

		return container.NewBorder(nil, nil, nil, nil, contentArea)
	}
	g.account.cateEntryList = entryList

	addBtn := widget.NewButtonWithIcon(fmt.Sprintf("add new %s", g.account.selectedCate), theme.ContentAddIcon(), func() {
		g.AddCateEntry()
	})
	backBtn := widget.NewButtonWithIcon("", theme.NavigateBackIcon(), func() {
		g.makeAppContentView(g.generateCateListArea)
	})
	btmBar := container.NewBorder(nil, nil, backBtn, nil, addBtn)

	list := g.generateEntryList()
	objs := []fyne.CanvasObject{list}
	contentArea := container.NewStack(objs...)

	return container.NewBorder(nil, btmBar, nil, nil, contentArea)
}

func (g *gui) generateEntryList() *widget.List {
	g.cateEntryListBindData = binding.BindStringList(&g.account.cateEntryList)
	list := widget.NewListWithData(
		g.cateEntryListBindData,
		func() fyne.CanvasObject {
			return widget.NewLabel("category entry item")
		},
		func(di binding.DataItem, co fyne.CanvasObject) {
			co.(*widget.Label).Bind(di.(binding.String))
		},
	)
	// show each individual entry
	list.OnSelected = func(id widget.ListItemID) {
		g.makeAppContentView(
			func() *fyne.Container {
				return g.generateEntryArea(id)
			})
	}
	return list
}

func (g *gui) generateEntryArea(id int) *fyne.Container {
	form := g.getEntryForm()

	return container.NewBorder(nil, nil, nil, nil, form)
}

func (g *gui) AddCateEntry() {
	name := "category item name"
	if isInSlice(g.account.cateEntryList, name) {
		dialog.NewError(fmt.Errorf("category item '%s' already exists", name), g.win).Show()
		return
	}
	g.cateEntryListBindData.Append(name)
	c.SaveCateEntry(g.account.accName, g.account.selectedCate, name)
}
