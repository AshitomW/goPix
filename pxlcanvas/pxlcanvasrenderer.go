package pxlcanvas

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)


type PxlCanvasRenderer struct {
	pxlCanvas *PxlCanvas
	canvasImage *canvas.Image
	canvasBorder []canvas.Line
	canvasCursor []fyne.CanvasObject 
}


func (renderer *PxlCanvasRenderer) SetCursor(objects []fyne.CanvasObject){
	renderer.canvasCursor = objects
}


// interface implementation
func (renderer *PxlCanvasRenderer) MinSize() fyne.Size{
	return renderer.pxlCanvas.DrawingArea
}

func (renderer *PxlCanvasRenderer) Objects() []fyne.CanvasObject{
	objects := make([]fyne.CanvasObject,0,5)
	for i:=0 ; i< len(renderer.canvasBorder);i++{
		objects = append(objects,&renderer.canvasBorder[i])
	}
	objects = append(objects, renderer.canvasImage)
	objects = append(objects, renderer.canvasCursor...)
	return objects
}



func(renderer *PxlCanvasRenderer) Layout(size fyne.Size){
	renderer.LayoutCanvas(size)
	renderer.LayoutBorder(size)

}

func(renderer *PxlCanvasRenderer) Refresh(){
	if(renderer.pxlCanvas.reloadImage){
		renderer.canvasImage = canvas.NewImageFromImage(renderer.pxlCanvas.PixelData)
		renderer.canvasImage.ScaleMode = canvas.ImageScalePixels
		renderer.canvasImage.FillMode = canvas.ImageFillContain
		renderer.pxlCanvas.reloadImage = false
	}
	renderer.Layout(renderer.pxlCanvas.Size())
	canvas.Refresh(renderer.canvasImage)
}


func(renderer *PxlCanvasRenderer) LayoutCanvas(size fyne.Size){
	imgPxlWidth := renderer.pxlCanvas.PxlCols
	imgPxHeight := renderer.pxlCanvas.PxlRows
	pxSize := renderer.pxlCanvas.PxSize



	renderer.canvasImage.Move(fyne.NewPos(renderer.pxlCanvas.CanvasOffset.X, renderer.pxlCanvas.CanvasOffset.Y))
	renderer.canvasImage.Resize(fyne.NewSize(float32(imgPxlWidth *pxSize), float32(imgPxHeight*pxSize)))

}

func(renderer *PxlCanvasRenderer) LayoutBorder(size fyne.Size){
	offset := renderer.pxlCanvas.CanvasOffset
	imgWidth := renderer.canvasImage.Size().Width
	imgHeight := renderer.canvasImage.Size().Height

	
	leftB := &renderer.canvasBorder[0]
	leftB.Position1 = fyne.NewPos(offset.X, offset.Y)
	leftB.Position2 = fyne.NewPos(offset.X, offset.Y + imgHeight)


	topB := &renderer.canvasBorder[1]
	topB.Position1 = fyne.NewPos(offset.X, offset.Y)
	topB.Position2 = fyne.NewPos(offset.X + imgWidth, offset.Y)


	rightB := &renderer.canvasBorder[2]
	rightB.Position1 = fyne.NewPos(offset.X+imgWidth,offset.Y)
	rightB.Position2 = fyne.NewPos(offset.X+imgWidth, offset.Y + imgHeight)

	bottomB := &renderer.canvasBorder[3]
	bottomB.Position1 = fyne.NewPos(offset.X,offset.Y+imgHeight)
	bottomB.Position2 = fyne.NewPos(offset.X +imgWidth, offset.Y+imgHeight)



	
}








func (renderer *PxlCanvasRenderer) Destroy(){}








