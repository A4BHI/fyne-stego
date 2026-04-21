package main

import (
	"fmt"
	"image/color"
	"io"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
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
	pickfilebutton = widget.NewButton("Pick File", func() {

		filepicker := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {

			sourceImage, err := io.ReadAll(reader)
			if err != nil {
				fmt.Println(err)
			}

			fmt.Println(sourceImage)
			os.WriteFile("/home/a4bhi/Desktop/fyne-stego/test.png", sourceImage, 0644)
		}, mainWindow)
		filepicker.SetFilter(filefilter)
		filepicker.SetConfirmText("Choose Image")
		filepicker.Show()

	})
	layout := container.NewBorder(header, footer, nil, nil, pickfilebutton)

	mainWindow.SetContent(layout)

	mainWindow.ShowAndRun()

}
