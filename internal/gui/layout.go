package gui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

const sideWidth = 220

type appLayout struct {
	top, left, right, content fyne.CanvasObject
	dividers                  [3]fyne.CanvasObject
}

// Layout implements fyne.Layout.
func (a *appLayout) Layout(objects []fyne.CanvasObject, size fyne.Size) {
	fmt.Println(objects)
	topHeight := a.top.MinSize().Height
	a.top.Resize(fyne.NewSize(size.Width, topHeight))
	a.left.Move(fyne.NewPos(0, topHeight))
	a.left.Resize(fyne.NewSize(sideWidth, size.Height-topHeight))

	a.right.Move(fyne.NewPos(size.Width-sideWidth, topHeight))
	a.right.Resize(fyne.NewSize(sideWidth, size.Height-topHeight))

	a.content.Move(fyne.NewPos(sideWidth, topHeight))
	a.content.Resize(fyne.NewSize(size.Width-sideWidth*2, size.Height-topHeight))

	dividerThickness := theme.SeparatorThicknessSize()
	a.dividers[0].Move(fyne.NewPos(0, topHeight))
	a.dividers[0].Resize(fyne.NewSize(size.Width, dividerThickness))

	a.dividers[1].Move(fyne.NewPos(sideWidth, topHeight))
	a.dividers[1].Resize(fyne.NewSize(dividerThickness, size.Height-topHeight))

	a.dividers[2].Move(fyne.NewPos(size.Width-sideWidth, topHeight))
	a.dividers[2].Resize(fyne.NewSize(dividerThickness, size.Height-topHeight))

}

// MinSize implements fyne.Layout.
func (a *appLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	borderSize := fyne.NewSize(sideWidth*2, a.top.MinSize().Height)
	return borderSize.AddWidthHeight(100, 100)
}

func newAppLayout(top, left, right, content fyne.CanvasObject, dividers [3]fyne.CanvasObject) fyne.Layout {
	return &appLayout{top: top, left: left, right: right, content: content, dividers: dividers}
}
