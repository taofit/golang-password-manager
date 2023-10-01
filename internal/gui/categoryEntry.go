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

func (g *gui) generateCateEntryListArea() *fyne.Container {
	g.title.Set(fmt.Sprintf("%s list", g.account.selectedCate))
	accName := g.account.accName
	entryList, err := c.FetchEntryList(accName, g.account.selectedCate)
	if err != nil {
		contentArea := widget.NewLabel(err.Error())

		return container.NewBorder(nil, nil, nil, nil, contentArea)
	}
	g.account.cateEntryList = entryList

	addBtn := widget.NewButtonWithIcon(fmt.Sprintf("add new %s", g.account.selectedCate), theme.ContentAddIcon(), func() {
		if fyne.CurrentDevice().IsMobile() {
			g.makeAppContentView(
				func() *fyne.Container {
					return g.generateCateEntryArea(-1)
				},
			)
			return
		}
		g.AddCateEntryDlg()
	})
	backBtn := widget.NewButtonWithIcon("", theme.NavigateBackIcon(), func() {
		g.makeAppContentView(g.generateCateListArea)
	})
	btmBar := container.NewBorder(nil, nil, backBtn, nil, addBtn)

	list := g.generateCateEntryList()

	var contentArea *fyne.Container
	if list.Length() == 0 {
		contentArea = container.NewCenter(widget.NewLabel(fmt.Sprintf("There is no %s", g.account.selectedCate)))
	} else {
		objs := []fyne.CanvasObject{list}
		contentArea = container.NewStack(objs...)
	}

	return container.NewBorder(nil, btmBar, nil, nil, contentArea)
}

func (g *gui) generateCateEntryList() *widget.List {
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
		if fyne.CurrentDevice().IsMobile() {
			g.makeAppContentView(
				func() *fyne.Container {
					return g.generateCateEntryArea(id)
				},
			)
			return
		}
		g.AddCateEntryDlg()
	}
	return list
}

func (g *gui) generateCateEntryArea(id int) *fyne.Container {
	form := g.getEntryForm(nil)

	return container.NewBorder(nil, nil, nil, nil, form)
}

func (g *gui) AddCateEntryDlg() {
	formTitle := fmt.Sprintf("Add a new %s", g.account.selectedCate)
	// name := widget.NewEntry()
	// formContent := []*widget.FormItem{widget.NewFormItem("name", name)}

	// dialog.ShowForm(formTitle, "Save", "Cancel", formContent, func(b bool) {
	// 	if !b {
	// 		return
	// 	}
	// 	if strings.TrimSpace(name.Text) == "" {
	// 		return
	// 	}
	// 	if isInSlice(g.account.cateEntryList, name.Text) {
	// 		dialog.ShowError(fmt.Errorf("%s '%s' already exists", g.account.selectedCate, name.Text), g.win)
	// 		return
	// 	}
	// 	g.cateEntryListBindData.Append(name.Text)
	// 	c.SaveCateEntry(g.account.accName, g.account.selectedCate, name.Text)
	// }, g.win)

	form := g.getEntryForm(nil)
	dialog.ShowCustom(formTitle, "Close", form, g.win)
}
