package main

import (
	"ashitomW/goPix/apptype"
	"ashitomW/goPix/swatch"
	"ashitomW/goPix/ui"
	"image/color"

	"fyne.io/fyne/v2/app"
)


func main(){
	goPix := app.New()
	pxlWindow := goPix.NewWindow("GoPix")


	state := apptype.State{
		BrushColor: color.NRGBA{255,255,255,255},
		SwatchSelected: 0,
	}

	appInit := ui.AppInit{
		PxlWindow: pxlWindow,
		State: &state,
		Swatches: make([] *swatch.Swatch,0,64),
	}


	ui.Setup(&appInit)


	appInit.PxlWindow.ShowAndRun()
}
