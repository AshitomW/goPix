package ui

import (
	"ashitomW/goPix/apptype"
	"ashitomW/goPix/pxlcanvas"
	"ashitomW/goPix/swatch"

	"fyne.io/fyne/v2"
)



type AppInit struct{
	PixlCanvas *pxlcanvas.PxlCanvas
	PxlWindow fyne.Window
	State * apptype.State
	Swatches []*swatch.Swatch
}





