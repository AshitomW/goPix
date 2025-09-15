package ui

import (
	"errors"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)




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



func BuildMenus(app *AppInit) *fyne.Menu{
	return fyne.NewMenu(
		"File",
		BuildNewMenu(app),
		)
}



func SetupMenus(app *AppInit) {
	menus:= BuildMenus(app)
	mainMenu := fyne.NewMainMenu(menus)
	app.PxlWindow.SetMainMenu(mainMenu)
}
