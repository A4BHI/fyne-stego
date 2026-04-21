package ui

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
)

type AppUI struct {
	header         *canvas.Text
	footer         *canvas.Text
	pickfilebutton *widget.Button
	sourceImage    []byte
	filefilter     storage.FileFilter
}

func (ui *AppUI) LoadUI() {
	ui.header = canvas.NewText("Steganography", color.NRGBA{R: 255, G: 170, B: 0, A: 255})
	ui.header.Alignment = fyne.TextAlignCenter
	ui.header.TextStyle = fyne.TextStyle{Bold: true, Italic: true}
	ui.header.TextSize = 50

	ui.footer = canvas.NewText("Made by a4bhi <3", color.NRGBA{R: 255, G: 170, B: 0, A: 255})
	ui.footer.Alignment = fyne.TextAlignCenter
	ui.footer.TextStyle.Monospace = true
	ui.footer.TextSize = 10

	ui.filefilter = storage.NewExtensionFileFilter([]string{".png"})
}

func (ui *AppUI) BuildUI(mainwindow fyne.Window) fyne.CanvasObject {

}

func NewAppUI() *AppUI {
	ui := &AppUI{}
	ui.LoadUI()
	return ui
}
