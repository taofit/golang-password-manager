package gui

import (
	"errors"
	"fmt"
	"log"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/taofit/golang-password-manager/internal/account"
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
				return g.generateCateEntryListArea()
			})
	}

	return list
}

func (g *gui) AddCategory() {
	name := widget.NewEntry()
	formContent := []*widget.FormItem{widget.NewFormItem("name", name)}
	dialog.ShowForm("Add a new category", "Save", "Cancel", formContent, func(b bool) {
		if !b {
			return
		}
		if strings.TrimSpace(name.Text) == "" {
			return
		}
		if isInSlice(g.account.categoryList, name.Text) {
			dialog.ShowError(fmt.Errorf("category '%s' already exists", name.Text), g.win)
			return
		}
		g.categoryListBindData.Append(name.Text)
		c.AddCategory(g.account.accName, name.Text)
		err := account.UpdateAccount(g.account.accName, g.account.categoryList)
		if err != nil {
			errMsg := fmt.Sprintf("error to add category: '%s', message: %s", name.Text, err)
			dialog.ShowError(errors.New(errMsg), g.win)
			log.Println(errMsg)
		}
	}, g.win)
}

func isInSlice(slice []string, needle string) bool {
	for _, item := range slice {
		if item == needle {
			return true
		}
	}
	return false
}
