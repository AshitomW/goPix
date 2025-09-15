package pxlcanvas

import (
	"ashitomW/goPix/pxlcanvas/brush"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/desktop"
)

// mouse interface implementation


func (pxlCanvas *PxlCanvas) Scrolled(ev *fyne.ScrollEvent){
	pxlCanvas.Scale((int(ev.Scrolled.DY)))
	pxlCanvas.Refresh()
}

func (pxlCanvas *PxlCanvas) MouseMoved(ev *desktop.MouseEvent){

	if x, y:=pxlCanvas.MouseToCanvasXY(ev); x!= nil && y!=nil {
		brush.TryBrush(pxlCanvas.appState,pxlCanvas,ev)
		cursor := brush.Cursor(pxlCanvas.PxlCanvasConfig, pxlCanvas.appState.BrushType,ev, *x,*y)
		pxlCanvas.renderer.SetCursor(cursor)
	}else {
		pxlCanvas.renderer.SetCursor(make([]fyne.CanvasObject,0))
	}


	pxlCanvas.TryPan(pxlCanvas.mouseState.previousCoords,ev)
	pxlCanvas.Refresh()
	pxlCanvas.mouseState.previousCoords = &ev.PointEvent
}
func(pxlCanvas *PxlCanvas) MouseIn(ev *desktop.MouseEvent){}
func(pxlCanvas *PxlCanvas) MouseOut(){}



func (pxlCanvas *PxlCanvas) MouseUp(ev *desktop.MouseEvent){}

func (pxlCanvas *PxlCanvas) MouseDown(ev *desktop.MouseEvent){
	brush.TryBrush(pxlCanvas.appState,pxlCanvas,ev)
}

