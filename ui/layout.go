package ui


func Setup(app *AppInit){
	swatchesContainer := BuildSwatches(app)
	app.PxlWindow.SetContent(swatchesContainer)
}
