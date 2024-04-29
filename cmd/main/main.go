package main

import (
	"fmt"

	"fyne.io/fyne"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Temperature Converter")

	tempEntry := widget.NewEntry()
	tempEntry.SetPlaceHolder("Enter temperature")

	var selectedUnit string
	//raio button for unit selection
	celsiusRadio := widget.NewRadioGroup([]string{"Celsius"}, func(string) {
		selectedUnit = "c"
	})
	fahrenheitRadio := widget.NewRadioGroup([]string{"Fahrenheit"}, func(string) {
		selectedUnit = "f"
	})
	kelvinRadio := widget.NewRadioGroup([]string{"Kelvin"}, func(string) {
		selectedUnit = "k"
	})

	resultLabel := widget.NewLabel("Result")
	convertButton := widget.NewButton("Convert", func() {
		temp := tempEntry.Text

		var result string
		switch selectedUnit {
		case "c":
			result = fmt.Sprintf("Fahrenheit: %.2f\nKelvin: %.2f", celsiusToFahrenheit(parseTemp(temp)), celsiusToKelvin(parseTemp(temp)))
		case "f":
			result = fmt.Sprintf("Celsius: %.2f\nKelvin: %.2f", fahrenheitToCelsius(parseTemp(temp)), fahrenheitToKelvin(parseTemp(temp)))
		case "k":
			result = fmt.Sprintf("Celsius: %.2f\nFahrenheit: %.2f", kelvinToCelsius(parseTemp(temp)), kelvinToFahrenheit(parseTemp(temp)))
		default:
			result = "Unit is Invalid"
		}
		resultLabel.SetText(result)
	})

	content := widget.NewVBox(
		tempEntry,
		celsiusRadio,
		fahrenheitRadio,
		kelvinRadio,
		convertButton,
		layout.NewSpacer(),
		resultLabel,
	)
	myWindow.SetContent(content)
	myWindow.Resize(fyne.NewSize(300, 200))
	myWindow.ShowAndRun()
}

func parseTemp(tempStr string) float64 {
	var temp float64
	fmt.Sscanf(tempStr, "%f", &temp)
	return temp
}

// Celsius to Fahrenheit
func celsiusToFahrenheit(celsius float64) float64 {
	return celsius*1.8 + 32
}

// Celsius to Kelvin
func celsiusToKelvin(celsius float64) float64 {
	return celsius + 273.15
}

// Fahrenheit to Celsius
func fahrenheitToCelsius(fahrenheit float64) float64 {
	return (fahrenheit - 32) / 1.8
}

// Fahrenheit to Kelvin
func fahrenheitToKelvin(fahrenheit float64) float64 {
	return (fahrenheit-32)/1.8 + 273.15
}

// Kelvin to Celsius
func kelvinToCelsius(kelvin float64) float64 {
	return kelvin - 273.15
}

// Kelvin to Fahrenheit
func kelvinToFahrenheit(kelvin float64) float64 {
	return (kelvin-273.15)*1.8 + 32
}
