package pxlcanvas

import "fyne.io/fyne/v2"


func (pxlCanvas *PxlCanvas) Pan(previousCoords, currentCoords fyne.PointEvent){
	xDiff := currentCoords.Position.X - previousCoords.Position.X 
	yDiff := currentCoords.Position.Y - previousCoords.Position.Y
	

	pxlCanvas.CanvasOffset.X += xDiff
	pxlCanvas.CanvasOffset.Y += yDiff


	pxlCanvas.Refresh()

}



func (pxlCanvas *PxlCanvas) Scale(direction int){
	switch {
	case direction > 0:
		pxlCanvas.PxSize +=1
	case direction <0:
		if pxlCanvas.PxSize >2 {
			pxlCanvas.PxSize -=1
		}
	default:
		pxlCanvas.PxSize = 10
	}
}




