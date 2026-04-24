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

type appUI struct {
	header *canvas.Text
	footer *canvas.Text

	sourceImage []byte
	filefilter  storage.FileFilter

	decodeTab *decodeTab
	encodeTab *encodeTab
}

type encodeTab struct {
	sourceImageLabel *canvas.Text //choose source image
	selectImageBtn   *widget.Button

	payloadFileLabel *canvas.Text //choose the file that you want to hide inside the image
	selectFilebtn    *widget.Button

	encodeBtn *widget.Button
}
type decodeTab struct {
	pickFileButtpm *widget.Button
}

func (ui *appUI) LoadUI() {
	ui.header = canvas.NewText("Steganography", color.NRGBA{R: 255, G: 170, B: 0, A: 255})
	ui.header.Alignment = fyne.TextAlignCenter
	ui.header.TextStyle = fyne.TextStyle{Bold: true, Italic: true}
	ui.header.TextSize = 50

	ui.footer = canvas.NewText("Made by a4bhi <3", color.NRGBA{R: 255, G: 170, B: 0, A: 255})
	ui.footer.Alignment = fyne.TextAlignCenter
	ui.footer.TextStyle.Monospace = true
	ui.footer.TextSize = 10

	ui.filefilter = storage.NewExtensionFileFilter([]string{".png"})

	ui.encodeTab.sourceImageLabel = canvas.NewText("Choose Image", color.RGBA{R: 52, G: 160, B: 142, A: 255})
	ui.encodeTab.sourceImageLabel.TextSize = 20

	ui.encodeTab.payloadFileLabel = canvas.NewText("Choose Payload", color.RGBA{R: 52, G: 160, B: 142, A: 255})
	ui.encodeTab.payloadFileLabel.TextSize = 20
}

func (ui *appUI) BuildUI(mainwindow fyne.Window) fyne.CanvasObject {

	ui.encodeTab.selectImageBtn = widget.NewButton("select image", func() {

		filepicker := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
			if reader == nil {

				return
			} //if user clicks cancel

			defer reader.Close()

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

	ui.encodeTab.selectFilebtn = widget.NewButton("Pick File", func() {

		filepicker := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
			if reader == nil {

				return
			} //if user clicks cancel

			defer reader.Close()

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

	ui.decodeTab.pickFileButtpm = widget.NewButton("Pick File", func() {

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

	imageLayout := container.NewGridWithColumns(2, ui.encodeTab.sourceImageLabel, ui.encodeTab.selectImageBtn)
	payloadLayout := container.NewGridWithColumns(2, ui.encodeTab.payloadFileLabel, ui.encodeTab.selectFilebtn)

	vbox := container.NewVBox(imageLayout, payloadLayout)
	tabs := container.NewAppTabs(container.NewTabItem("Encode", vbox), container.NewTabItem("Decode", ui.decodeTab.pickFileButtpm))
	layout := container.NewBorder(ui.header, ui.footer, nil, nil, tabs)

	return layout
}

func NewAppUI() *appUI {
	ui := &appUI{
		encodeTab: &encodeTab{},
		decodeTab: &decodeTab{},
	}
	ui.LoadUI()
	return ui
}
