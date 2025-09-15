package pxlcanvas

import (
	"ashitomW/goPix/apptype"
	"image"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
)

type PxlCanvasMouseState struct {
	previousCoords *fyne.PointEvent

}


type PxlCanvas struct {
	widget.BaseWidget
	apptype.PxlCanvasConfig
	renderer *PxlCanvasRenderer
	PixelData image.Image
	mouseState PxlCanvasMouseState
	appState *apptype.State
	reloadImage bool
}


func (pxlCanvas *PxlCanvas) Bounds() image.Rectangle{
	
	// top left 
	x0 := int(pxlCanvas.CanvasOffset.X)
	y0 := int(pxlCanvas.CanvasOffset.Y)
	// bottom right
	x1:= int(pxlCanvas.PxlCols * pxlCanvas.PxSize + int(pxlCanvas.CanvasOffset.X))
	y1 := int(pxlCanvas.PxlRows * pxlCanvas.PxSize + int(pxlCanvas.CanvasOffset.Y))

	return image.Rect(x0,y0,x1,y1)
		
}



func InBoundds(position fyne.Position , bounds image.Rectangle ) bool {
	if position.X >= float32(bounds.Min.X) && position.X < float32(bounds.Max.X) && position.Y >= float32(bounds.Min.Y) && position.Y < float32(bounds.Max.Y){
		return true
	}
	return false
}


func NewBlankImage(cols,rows int, c color.Color) image.Image{
	img := image.NewNRGBA(image.Rect(0,0,cols,rows))
	for y:=0;y<rows;y++{
		for x:=0;x<cols;x++{
			img.Set(x,y,c)
		}
	}

	return img
}



func NewPxlCanvas(state *apptype.State,config apptype.PxlCanvasConfig) *PxlCanvas{
	pxlCanvas := &PxlCanvas{
		PxlCanvasConfig: config,
		appState: state,
	}

	pxlCanvas.PixelData = NewBlankImage(pxlCanvas.PxlCols,pxlCanvas.PxlRows,color.NRGBA{128,128,128,255})
	pxlCanvas.ExtendBaseWidget(pxlCanvas)
	return pxlCanvas	
}


func (pxlCanvas *PxlCanvas) CreateRenderer() fyne.WidgetRenderer{
	canvasImage:= canvas.NewImageFromImage(pxlCanvas.PixelData)
	canvasImage.ScaleMode = canvas.ImageScalePixels
	canvasImage.FillMode = canvas.ImageFillContain
	
	canvasBorder := make([]canvas.Line,4)
	for i:=0;i <len(canvasBorder);i++{
		canvasBorder[i].StrokeColor = color.NRGBA{100,100,100,255}
		canvasBorder[i].StrokeWidth = 2 
	}


	renderer := &PxlCanvasRenderer{
		pxlCanvas: pxlCanvas,
		canvasImage: canvasImage,
		canvasBorder: canvasBorder,

	}

	pxlCanvas.renderer = renderer
	return renderer
}





func (pxlCanvas *PxlCanvas) TryPan(previousCoords *fyne.PointEvent, ev *desktop.MouseEvent){	
	if previousCoords != nil && ev.Button == desktop.MouseButtonTertiary{
		pxlCanvas.Pan(*previousCoords, ev.PointEvent)
	}
}

