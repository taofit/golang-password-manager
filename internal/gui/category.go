package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

func (g *gui) createCategoryContent() *fyne.Container {
	objs := []fyne.CanvasObject{}
	return container.NewStack(objs...)
}
