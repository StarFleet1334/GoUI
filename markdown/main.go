package main

import (
	"fyne.io/fyne/v2"
	app2 "fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
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

	saveMenuAsItem := fyne.NewMenuItem("SaveAs...", func() {})

	fileMenu := fyne.NewMenu("Menu", openMenuItem, saveMenuItem, saveMenuAsItem)

	menu := fyne.NewMainMenu(fileMenu)

	win.SetMainMenu(menu)

}
