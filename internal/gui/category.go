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

func (g *gui) generateCateListArea() *fyne.Container {
	g.title.Set("Category list")
	list := g.generateCategoryList() //get widget list
	addBtn := widget.NewButtonWithIcon("add new category", theme.ContentAddIcon(), func() {
		g.AddCategory()
	})
	objs := []fyne.CanvasObject{list}
	contentArea := container.NewStack(objs...)

	return container.NewBorder(nil, addBtn, nil, nil, contentArea)
}

func (g *gui) generateCategoryList() *widget.List {
	g.categoryListBindData = binding.BindStringList(&g.account.categoryList)
	list := widget.NewListWithData(
		g.categoryListBindData,
		func() fyne.CanvasObject {
			return widget.NewLabel("category item")
		},
		func(di binding.DataItem, co fyne.CanvasObject) {
			co.(*widget.Label).Bind(di.(binding.String))
		})
	list.OnSelected = func(id widget.ListItemID) {
		g.updateSelectedCateById(id)
		g.makeAppContentView(
			func() *fyne.Container {
				return g.generateEntryListArea()
			})
	}

	return list
}

func (g *gui) AddCategory() {
	name := "category name"
	// category := category{name: "category name", key: "category key"}
	// g.categories.addCate(category)
	// popup := widget.NewModalPopUp(widget.NewEntry(), g.win.Canvas())
	// popup.Show()
	if isInSlice(g.account.categoryList, name) {
		dialog.NewError(fmt.Errorf("category '%s' already exists", name), g.win).Show()
		return
	}
	g.categoryListBindData.Append(name)
	c.AddCategory(g.account.accName, name)
}

func isInSlice(slice []string, needle string) bool {
	for _, item := range slice {
		if item == needle {
			return true
		}
	}
	return false
}
