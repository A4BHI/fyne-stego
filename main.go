package main

import (
	"image/color"
	"stego/ui"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
)

var (
	header         *canvas.Text
	footer         *canvas.Text
	pickfilebutton *widget.Button
	sourceImage    []byte
	filefilter     storage.FileFilter
)

func init() {
	header = canvas.NewText("Steganography", color.NRGBA{R: 255, G: 170, B: 0, A: 255})
	header.Alignment = fyne.TextAlignCenter
	header.TextStyle = fyne.TextStyle{Bold: true, Italic: true}
	header.TextSize = 50

	footer = canvas.NewText("Made by a4bhi <3", color.NRGBA{R: 255, G: 170, B: 0, A: 255})
	footer.Alignment = fyne.TextAlignCenter
	footer.TextStyle.Monospace = true
	footer.TextSize = 10

	filefilter = storage.NewExtensionFileFilter([]string{".png"})

}

func main() {
	stego := app.New()
	mainWindow := stego.NewWindow("Image Steganography")
	mainWindow.SetMaster()

	ui := ui.NewAppUI()

	mainWindow.SetContent(ui.BuildUI(mainWindow))

	mainWindow.ShowAndRun()

}
