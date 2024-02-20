package main

import (
	"fmt"
	"time"
)

func main() {

	currentTime := time.Now()
	hour := currentTime.Hour()
	var textLine string

	if hour > 7 && hour < 12 {
		textLine = "Goedemorgen!"
	} else if hour > 12 && hour < 18 {
		textLine = "Goedemiddag!"
	} else if hour > 18 && hour < 23 {
		textLine = "Goedeavond!"
	} else {
		textLine = "Sorry, de parkeerplaats is ’s nachts gesloten"
	}

	fmt.Println(textLine)

	if hour > 7 && hour < 23 {
		fmt.Println("Welkom bij Fonteyn Vakantieparken")
	}

}
