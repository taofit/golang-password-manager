package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

const sideWidth = 220

type appLayout struct {
	top, content fyne.CanvasObject
	dividers     [2]fyne.CanvasObject
}

// Layout implements fyne.Layout.
func (a *appLayout) Layout(objects []fyne.CanvasObject, size fyne.Size) {
	topHeight := a.top.MinSize().Height
	a.top.Resize(fyne.NewSize(size.Width, topHeight))

	a.content.Move(fyne.NewPos(0, topHeight))
	a.content.Resize(fyne.NewSize(size.Width, size.Height-topHeight))

	dividerThickness := theme.SeparatorThicknessSize()
	a.dividers[0].Move(fyne.NewPos(0, topHeight))
	a.dividers[0].Resize(fyne.NewSize(size.Width, dividerThickness))

	a.dividers[1].Move(fyne.NewPos(0, size.Height))
	a.dividers[1].Resize(fyne.NewSize(size.Width, dividerThickness))
}

// MinSize implements fyne.Layout.
func (a *appLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	borderSize := fyne.NewSize(sideWidth*2, a.top.MinSize().Height)
	return borderSize.AddWidthHeight(100, 100)
}

func newAppLayout(top, content fyne.CanvasObject, dividers [2]fyne.CanvasObject) fyne.Layout {
	return &appLayout{top: top, content: content, dividers: dividers}
}
