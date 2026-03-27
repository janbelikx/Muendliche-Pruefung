package main

import "fmt"

func main() {

	var fortschritt int = 0
	fmt.Printf("Aktueller Fortschritt: %d\n", fortschritt)

	// fmt.Printf("Prüfung gestartet für: %s\n", name)

	// if __________ {
	// 	fmt.Println("Status: Ziel erreicht.")
	// 	fortschritt = 100
	// } else {
	// 	fmt.Println("Status: In Bearbeitung...")
	// }

	pruefungsPhase := "Vorbereitung"

	switch pruefungsPhase {
	case "Vorbereitung":
		fmt.Println("Schritt 1: Repository wurde korrekt geklont.")
		fortschritt = 50
	case "Codierung":
		fmt.Println("Schritt 2: Logik wird implementiert.")
		fortschritt = 80
	default:
		fmt.Println("Unbekannte Phase.")
	}

	fmt.Println("\n************************************")
	fmt.Println("Sie sind bereit für die Prüfung!")
	fmt.Println("************************************")

}
