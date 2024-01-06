package main

import (
	"strconv"

	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

func main() {
	a := app.New()
	c := 0
	l := widget.NewLabel("Hello Fyne!")
	w := a.NewWindow("Hello")
	w.SetContent(
		widget.NewVBox(
			l,
			widget.NewButton("Click me!", func() {
				c++
				l.SetText("count: " + strconv.Itoa(c))
			}),
		),
	)
	w.ShowAndRun()
}
