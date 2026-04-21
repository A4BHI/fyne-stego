package ui

import (
	"fmt"
	"image/color"
	"io"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
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
	ui.pickfilebutton = widget.NewButton("Pick File", func() {

		filepicker := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {

			sourceImage, err := io.ReadAll(reader)
			if err != nil {
				fmt.Println(err)
			}

			fmt.Println(sourceImage)
			os.WriteFile("/home/a4bhi/Desktop/fyne-stego/test.png", sourceImage, 0644)
		}, mainwindow)
		filepicker.SetFilter(ui.filefilter)
		filepicker.SetConfirmText("Choose Image")
		filepicker.Show()

	})
	layout := container.NewBorder(ui.header, ui.footer, nil, nil, ui.pickfilebutton)

	return layout
}

func NewAppUI() *AppUI {
	ui := &AppUI{}
	ui.LoadUI()
	return ui
}
