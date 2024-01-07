package main

import (
	"fyne.io/fyne/v2"
	app2 "fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
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

	openMenuItem := fyne.NewMenuItem("Open...", func() {})

	saveMenuItem := fyne.NewMenuItem("Save...", func() {})

	app.SaveMenuItem = saveMenuItem
	app.SaveMenuItem.Disabled = true

	saveMenuAsItem := fyne.NewMenuItem("SaveAs...", app.saveAsFunc(win))

	fileMenu := fyne.NewMenu("Menu", openMenuItem, saveMenuItem, saveMenuAsItem)

	menu := fyne.NewMainMenu(fileMenu)

	win.SetMainMenu(menu)

}

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

			write.Write([]byte(app.EditWidget.Text))
			app.CurrentFile = write.URI()

			defer write.Close()
			win.SetTitle(win.Title() + " - " + write.URI().Name())
			app.SaveMenuItem.Disabled = false
		}, win)
		saveDialog.Show()
	}
}
