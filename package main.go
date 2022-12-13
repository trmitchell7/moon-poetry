package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Create a list of poets
var poets = []string{
	"William Shakespeare",
	"Elizabeth Barrett Browning",
	"Edgar Allan Poe",
	"Christina Rossetti",
	"Emily Dickinson",
	"Robert Frost",
	"Langston Hughes",
	"Sylvia Plath",
	"Pablo Neruda",
	"Rumi",
	"Skaldic Verse",
	"Einar Már Guðmundsson",
	"Halldór Laxness",
	"Sigurjón Birgir Sigurðsson",
	"Sjón",
	"Ragna Sigurðardóttir",
	"Einar Kárason",
}

// Create a list of artists
var artists = []string{
	"Vincent van Gogh",
	"John Constable",
	"Caspar David Friedrich",
	"William Blake",
	"Henri de Toulouse-Lautrec",
	"Gustav Klimt",
	"Yves Tanguy",
	"Edward Hopper",
	"Studio Ghibli",
	"Giovanni Battista Tiepolo",
}


// getCurrentMoonPhase returns the current phase of the moon as a string.
func getCurrentMoonPhase() string {
	now := time.Now()
	// Use the date to calculate the moon's age in days.
	moonAge := now.Sub(time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)).Hours() / 24
	// Use the moon's age to determine its phase.
	fmt.Printf("current day %s \n current phase int: %s \n\n", now, moonAge)
	switch {
	case moonAge < 1.84566:
		return "New"
	case moonAge < 5.53699:
		return "Waxing Crescent"
	case moonAge < 9.22831:
		return "First Quarter"
	case moonAge < 12.91963:
		return "Waxing Gibbous"
	case moonAge < 16.61096:
		return "Full"
	case moonAge < 20.30228:
		return "Waning Gibbous"
	case moonAge < 23.99361:
		return "Last Quarter"
	case moonAge < 27.68493:
		return "Waning Crescent"
	default:
		return "New"
	}
}


func main() {
	// Get the current phase of the moon.
	currentPhase := getCurrentMoonPhase()

	// Choose a random poet from the list of poets
	poetChoice := poets[rand.Intn(len(poets))]

	// Choose a random artist from the list of artists
	artistChoice := artists[rand.Intn(len(artists))]

	// Output a string to write a poem that includes the current moon phase and randomly chosen poet
	fmt.Printf("Write a lovely poem about a beautiful woman named Linda as if she was the moon, taking into account last night's %s Moon and what the phase of the moon represents to Linda, done in the style of %s in mixed English and Icelandic.\n\n", currentPhase, poetChoice)

	// Output a string to generate an image that includes the current moon phase in the style of the randomly chosen artist
	fmt.Printf("/imagine prompt: A beautiful night sky landscape scene with a silhouette of a woman with short curly copper hair looking up at the perfect %s Moon in the starry nightsky, painting by %s, astrophotography, accurate %s phase, spirituality, elegant, ethereal, 8k --v 4 --q 2 --ar 2:3 \n", currentPhase, artistChoice, currentPhase)

}
