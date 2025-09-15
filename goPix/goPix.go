package main

import (
	"ashitomW/goPix/apptype"
	"ashitomW/goPix/pxlcanvas"
	"ashitomW/goPix/swatch"
	"ashitomW/goPix/ui"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)


func main(){
	goPix := app.New()
	pxlWindow := goPix.NewWindow("GoPix")



	pixlCanvasConfig := apptype.PxlCanvasConfig{
		DrawingArea: fyne.NewSize(600,600),
		CanvasOffset: fyne.NewPos(0,0),
		PxlRows: 10,
		PxlCols: 10,
		PxSize: 30,
	}
 


	state := apptype.State{
		BrushColor: color.NRGBA{255,255,255,255},
		SwatchSelected: 0,
	}
	pixlCanvas := pxlcanvas.NewPxlCanvas(&state, pixlCanvasConfig)
	appInit := ui.AppInit{
		PxlWindow: pxlWindow,
		PixlCanvas: pixlCanvas,
		State: &state,
		Swatches: make([] *swatch.Swatch,0,64),
	}


	ui.Setup(&appInit)


	appInit.PxlWindow.ShowAndRun()
}
