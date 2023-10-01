package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

func getLogo() *canvas.Image {
	logo := canvas.NewImageFromResource(resourceLogoPng)
	logo.FillMode = canvas.ImageFillContain
	logo.SetMinSize(fyne.NewSize(120, 120))

	return logo
}
