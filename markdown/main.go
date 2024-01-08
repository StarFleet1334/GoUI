package main

import (
	"fyne.io/fyne/v2"
	app2 "fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
	"io"
	"strings"
)

type config struct {
	EditWidget    *widget.Entry
	PreviewWidget *widget.RichText
	CurrentFile   fyne.URI
	SaveMenuItem  *fyne.MenuItem
}

var cfg config

func main() {

	app := app2.New()

	w := app.NewWindow("MarkDown")

	edit, preview := cfg.makeUI()

	cfg.createMenuItem(w)

	w.SetContent(container.NewHSplit(edit, preview))

	w.Resize(fyne.Size{Width: 800, Height: 500})

	w.CenterOnScreen()

	w.ShowAndRun()
}

func (app *config) makeUI() (*widget.Entry, *widget.RichText) {
	edit := widget.NewMultiLineEntry()
	preview := widget.NewRichTextFromMarkdown("")
	app.EditWidget = edit
	app.PreviewWidget = preview

	edit.OnChanged = preview.ParseMarkdown

	return edit, preview
}

func (app *config) createMenuItem(win fyne.Window) {

	openMenuItem := fyne.NewMenuItem("Open...", app.openFunc(win))

	saveMenuItem := fyne.NewMenuItem("Save...", app.saveAsFunc(win))

	app.SaveMenuItem = saveMenuItem
	app.SaveMenuItem.Disabled = true

	saveMenuAsItem := fyne.NewMenuItem("SaveAs...", app.saveAsFunc(win))

	fileMenu := fyne.NewMenu("Menu", openMenuItem, saveMenuItem, saveMenuAsItem)

	menu := fyne.NewMainMenu(fileMenu)

	win.SetMainMenu(menu)

}

var filter = storage.NewExtensionFileFilter([]string{".md", ".MD"})

func (app *config) saveAsFunc(win fyne.Window) func() {
	return func() {
		saveDialog := dialog.NewFileSave(func(write fyne.URIWriteCloser, err error) {
			if err != nil {
				dialog.ShowError(err, win)
				return
			}

			if write == nil {
				return
			}

			if !strings.HasSuffix(strings.ToLower(write.URI().String()), ".md") {
				dialog.ShowInformation("Error!", "Please name your file with .md extension", win)
				return
			}

			write.Write([]byte(app.EditWidget.Text))
			app.CurrentFile = write.URI()

			defer func(write fyne.URIWriteCloser) {
				err := write.Close()
				if err != nil {
					return
				}
			}(write)
			win.SetTitle(win.Title() + " - " + write.URI().Name())
			app.SaveMenuItem.Disabled = false
		}, win)
		saveDialog.SetFileName("untitled.md")
		saveDialog.SetFilter(filter)
		saveDialog.Show()
	}
}

func (app *config) openFunc(win fyne.Window) func() {
	return func() {
		openDialog := dialog.NewFileOpen(func(read fyne.URIReadCloser, err error) {
			if err != nil {
				dialog.ShowError(err, win)
				return
			}

			if read == nil {
				return
			}

			defer func(read fyne.URIReadCloser) {
				err := read.Close()
				if err != nil {
					return
				}
			}(read)

			data, error := io.ReadAll(read)
			if error != nil {
				dialog.ShowError(error, win)
				return
			}

			app.EditWidget.SetText(string(data))
			app.CurrentFile = read.URI()
			win.SetTitle(win.Title() + " - " + read.URI().Name())
			app.SaveMenuItem.Disabled = false
		}, win)
		openDialog.SetFilter(filter)
		openDialog.Show()
	}
}

func (app *config) saveFunc(win fyne.Window) func() {
	return func() {
		if app.CurrentFile != nil {
			write, err := storage.Writer(app.CurrentFile)
			if err != nil {
				dialog.ShowError(err, win)
				return
			}

			write.Write([]byte(app.EditWidget.Text))
			defer func(write fyne.URIWriteCloser) {
				err := write.Close()
				if err != nil {
					return
				}
			}(write)
		}
	}
}
