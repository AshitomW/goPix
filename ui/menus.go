package ui

import (
	util "ashitomW/goPix/utils"
	"errors"
	"image"
	"image/png"
	"os"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)



func saveFileDialog(app *AppInit) {
	 dialog.ShowFileSave(func (uri fyne.URIWriteCloser, e error){
		if uri == nil {
			return 
		} else {
			err := png.Encode(uri,app.PixlCanvas.PixelData)
			if err != nil {
				dialog.ShowError(err,app.PxlWindow)
				return 
			}
			app.State.SetFilePath(uri.URI().Path())
		} 
	},app.PxlWindow)
}

func BuildSaveAsMenu(app *AppInit) *fyne.MenuItem{
	return fyne.NewMenuItem("Save As...",func(){
		saveFileDialog(app)
	})
}

func BuildSaveMenu(app *AppInit) *fyne.MenuItem{
	return fyne.NewMenuItem("Save",func(){
		if (app.State.FilePath == ""){
			saveFileDialog(app)
		}else {
				tryClose := func(fh *os.File) {
				err:=fh.Close()
				if err !=nil {
					dialog.ShowError(err, app.PxlWindow)
				}
			}

			fh,err:= os.Create(app.State.FilePath)
			defer tryClose(fh)

			if err != nil {
				dialog.ShowError(err, app.PxlWindow)
				return
			}
			err = png.Encode(fh,app.PixlCanvas.PixelData)
			if err != nil{
				dialog.ShowError(err, app.PxlWindow)
				return 
			}
		}
	})
}



func BuildNewMenu(app *AppInit) *fyne.MenuItem{
	return fyne.NewMenuItem("New Drawing",func(){
		sizeValidator := func(s string) error {
			width, err := strconv.Atoi(s)
			if err != nil {
				return errors.New("must be a positive integer")  
			}
			if width <= 0 {
				return errors.New("must be greater than 0")
			}
			return nil
		}
		widthEntry := widget.NewEntry()
		widthEntry.Validator = sizeValidator
		heightEntry:= widget.NewEntry()
		heightEntry.Validator = sizeValidator
		

		widthEntryForm := widget.NewFormItem("Width",widthEntry)
		heightEntryForm := widget.NewFormItem("Width",heightEntry)



		formItems := []*widget.FormItem{widthEntryForm,heightEntryForm}

		dialog.ShowForm("new Image","Create","Cancer",formItems,func (ok bool){
			if ok {
				pixelWidth := 0
				pixelHeight := 0
				if widthEntry.Validate() != nil {
					dialog.ShowError(errors.New("Invalid Width"),app.PxlWindow)
				}else {
					pixelWidth,_ =strconv.Atoi(widthEntry.Text)
				}
					if heightEntry.Validate() != nil {
					dialog.ShowError(errors.New("Invalid Height"),app.PxlWindow)
				}else {
					pixelHeight,_ =strconv.Atoi(widthEntry.Text)
				}

				app.PixlCanvas.NewDrawing(pixelWidth,pixelHeight)
			}
		},app.PxlWindow)

	})
}


func BuildOpenMenu(app *AppInit) *fyne.MenuItem{
	return fyne.NewMenuItem("Open...",func(){
		dialog.ShowFileOpen(func(uri fyne.URIReadCloser,e error){
			if uri == nil {
				return 
			}else {
				image, _, err := image.Decode(uri)
				if err != nil {
					dialog.ShowError(err, app.PxlWindow)
					return;
				}
				app.PixlCanvas.LoadImage(image)
				app.State.SetFilePath(uri.URI().Path())
				imgColors := util.GetImageColor(image)
				i := 0 
				for c:=range imgColors {
					if i==len(app.Swatches){
						break
					}
					app.Swatches[i].SetColor(c)
					i++
				}
			}
		},app.PxlWindow)
	})
}




func BuildMenus(app *AppInit) *fyne.Menu{
	return fyne.NewMenu(
		"File",
		BuildNewMenu(app),
		BuildOpenMenu(app),
		BuildSaveMenu(app),
		BuildSaveAsMenu(app),
		)
}



func SetupMenus(app *AppInit) {
	menus:= BuildMenus(app)
	mainMenu := fyne.NewMainMenu(menus)
	app.PxlWindow.SetMainMenu(mainMenu)
}
