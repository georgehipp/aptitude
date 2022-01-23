package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main() {
	app := app.New()
	win := app.NewWindow("Aptitude Test Library")
	win.Resize(fyne.NewSize(200, 200))

	title := widget.NewLabel("Select Test")
	header := container.New(layout.NewHBoxLayout(), title)

	wintitle := widget.NewLabel("Window Title")
	var winName *fyne.Container
	winName = container.New(layout.NewHBoxLayout(), wintitle)

	apptitle := widget.NewLabel("App Title")
	var appwin *fyne.Container
	appwin = container.New(layout.NewHBoxLayout(), apptitle)

	var appcont *fyne.Container
	appcont = container.New(layout.NewHBoxLayout(), apptitle)
	//appwin.Refresh()
	w1 := widget.NewButton("Numbers", func() {
		title.SetText("Numbers Selected - Complete Test Below")
		wintitle.SetText("Numbers Test")
		apptitle.SetText("Instructions will be here")
		b1 := widget.NewButton("b1", func() {})
		b2 := widget.NewButton("b2", func() {})
		b3 := widget.NewButton("b3", func() {})
		appcont = container.New(layout.NewGridLayout(3), b1, b2, b3)
	})
	w2 := widget.NewButton("Perception", func() {
		title.SetText("Perception Selected - Complete Test Below")
		wintitle.SetText("Perception Test")
	})
	w3 := widget.NewButton("Reason", func() {
		title.SetText("Reason Selected - Complete Test Below")
		wintitle.SetText("Reason Test")
	})
	w4 := widget.NewButton("Spatial", func() {
		title.SetText("Spatial Selected - Complete Test Below")
		wintitle.SetText("Spatial Test")
	})
	w5 := widget.NewButton("Words", func() {
		title.SetText("Words Selected - Complete Test Below")
		wintitle.SetText("Words Test")
	})
	w6 := widget.NewButton("Results", func() {
		title.SetText("Results Selected - Results Below")
		wintitle.SetText("Results")
	})
	grid := container.New(layout.NewGridLayout(6), w1, w2, w3, w4, w5, w6)
	appgrid := container.New(layout.NewGridLayout(3), layout.NewSpacer(), appwin, layout.NewSpacer())
	win.SetContent(container.New(layout.NewVBoxLayout(), header, grid, winName, appgrid, appcont))

	win.ShowAndRun()
}
