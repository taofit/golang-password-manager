//go:generate fyne bundle -o bundled.go ../../assets
package gui

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

type AppTheme struct {
	fyne.Theme
}

func NewAppTheme() fyne.Theme {
	return &AppTheme{Theme: theme.DefaultTheme()}
}

func (t *AppTheme) Color(name fyne.ThemeColorName, _ fyne.ThemeVariant) color.Color {
	return t.Theme.Color(name, theme.VariantLight)
}

func (t *AppTheme) Size(name fyne.ThemeSizeName) float32 {
	if name == theme.SizeNameText {
		return 12
	}
	return t.Theme.Size(name)
}
