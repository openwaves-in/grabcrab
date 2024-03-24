package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/openwaves-in/grabcrab/urlcooker"
)

func main() {
	// Create a new Fyne application
	myApp := app.New()

	// Create a new window
	myWindow := myApp.NewWindow("GRAB CRAB")
	myWindow.Resize(fyne.NewSize(800, 600))

	// Create widgets for location selection, skills input, and search button
	locationOptions := []string{"chennai", "hyderabad", "bangalore", "kochin"}
	locationSelect := widget.NewSelect(locationOptions, nil)
	skillsEntry := widget.NewEntry()
	searchButton := widget.NewButton("Search", func() {
		location := locationSelect.Selected
		skills := skillsEntry.Text

		// Validate inputs
		if location == "" || skills == "" {
			showErrorDialog(myWindow, "Please select a location and enter your skills")
			return
		}

		// Display processing message
		processingLabel := widget.NewLabel("Processing your request...")
		showPage(myWindow, container.NewVBox(processingLabel))

		// Fetch data from URL
		csvFileLocation := urlcooker.Urlcook(skills, location)
		
		// Show completion message and option to open the CSV file location
		completionMessage := fmt.Sprintf("Process completed. CSV file generated at:\n%s", csvFileLocation)
		dialog := widget.NewModalPopUp(container.NewVBox(widget.NewLabel(completionMessage)), myWindow.Canvas())
		dialog.Show()
	})

	// Create a form container
	form := container.NewVBox(
		widget.NewLabel("Location:"),
		locationSelect,
		widget.NewLabel("Skills:"),
		skillsEntry,
		searchButton,
	)

	// Set the window content
	myWindow.SetContent(form)

	// Show and run the application
	myWindow.ShowAndRun()
}

func showErrorDialog(window fyne.Window, message string) {
	errorDialog := widget.NewModalPopUp(container.NewVBox(widget.NewLabel(message)), window.Canvas())
	errorDialog.Show()
}

func showPage(window fyne.Window, content fyne.CanvasObject) {
	window.SetContent(content)
}
