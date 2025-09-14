package apptype

import (
	"image/color"

	"fyne.io/fyne/v2"
)



type BrushType = int 

type PxlCanvasConfig struct{
	DrawingArea fyne.Size
	CanvasOffset fyne.Position
	PxlRows, PxlCols int 
	PxSize int // scale factor for pixels
}



type State struct {
	BrushColor color.Color 
	BrushType int 
	SwatchSelected int 
	FilePath string 
}


func (state *State) SetFilePath(path string){
	state.FilePath = path
}

