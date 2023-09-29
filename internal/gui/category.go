package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	c "github.com/taofit/golang-password-manager/internal/category"
)

// type category struct {
// 	key  string
// 	name string
// }

// type categoryList struct {
// 	list []category
// }

// func (cl *categoryList) addCate(c category) {
// 	cl.list = append(cl.list, c)
// }

func (g *gui) generateCategory() *fyne.Container {
	g.fetchDefaultCategories()
	list := g.generateCategoryList()
	addBtn := widget.NewButtonWithIcon("add new category", theme.ContentAddIcon(), func() {
		g.AddCategory()
	})
	objs := []fyne.CanvasObject{list}
	contentArea := container.NewStack(objs...)

	return container.NewBorder(nil, addBtn, nil, nil, contentArea)
}

func (g *gui) generateCategoryList() *widget.List {
	g.categoryListBindData = binding.BindStringList(&g.categoryList)
	list := widget.NewListWithData(g.categoryListBindData,
		func() fyne.CanvasObject {
			return widget.NewLabel("category item")
		},
		func(di binding.DataItem, co fyne.CanvasObject) {
			co.(*widget.Label).Bind(di.(binding.String))
		})

	return list
}

func (g *gui) fetchDefaultCategories() {
	g.categoryList = c.CategoryList
}

func (g *gui) AddCategory() {
	name := "category name"
	// category := category{name: "category name", key: "category key"}
	// g.categories.addCate(category)
	// popup := widget.NewModalPopUp(widget.NewEntry(), g.win.Canvas())
	// popup.Show()
	g.categoryListBindData.Append(name)
	c.AddCategory(g.accName, name)
}
