package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
	//"bytes"
	//"io/ioutil"
	//"net/http"
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


func currentMoonPhase() string {
    // Calculate the current phase of the moon based on the given New Moon date.
	// We'll arbitrarily use the New Moon on November 23, 2022 as the seed date.
    newMoon := time.Date(2022, time.November, 23, 0, 0, 0, 0, time.UTC)
	// Get the current date so we can calculate the delta from the seed.
    currentTime := time.Now()
    // Calculate the delta since the newMoon seed date.
    daysSinceNewMoon := currentTime.Sub(newMoon).Hours() / 24

    // Calculate the number of days in the current lunar cycle and mod the delta since the seed date.
    daysInLunarCycle := 29.530588853
    daysSinceNewMoon = math.Mod(daysSinceNewMoon, daysInLunarCycle)

    // Determine the current phase of the moon based on the number of days since the last New Moon
    var phase string
    switch {
    case daysSinceNewMoon < 1:
        phase = "New"
    case daysSinceNewMoon < 7:
        phase = "Waxing Crescent"
    case daysSinceNewMoon < 8:
        phase = "First Quarter"
    case daysSinceNewMoon < 16:
        phase = "Waxing Gibbous"
    case daysSinceNewMoon < 17:
        phase = "Full"
    case daysSinceNewMoon < 25:
        phase = "Waning Gibbous"
    case daysSinceNewMoon < 26:
        phase = "Third Quarter"
    default:
        phase = "Waning Crescent"
    }
    // Return the current phase of the moon.
    return phase

}
  

func main() {
	// Get the current phase of the moon.
	currentPhase := currentMoonPhase()

	// Choose a random poet from the list of poets
	poetChoice := poets[rand.Intn(len(poets))]

	// Choose a random artist from the list of artists
	artistChoice := artists[rand.Intn(len(artists))]

	// Form the string to write a poem that includes the current moon phase and randomly chosen poet
	poemPrompt, _ := fmt.Printf("Write a lovely poem about a beautiful woman named Linda as if she was the moon, taking into account last night's %s Moon and what the phase of the moon represents to Linda, done in the style of %s in mixed English and Icelandic.\n\n", currentPhase, poetChoice)

	// fmt.Printf("Write a lovely poem about a beautiful woman named Linda as if she was the moon, taking into account last night's %s Moon and what the phase of the moon represents to Linda, done in the style of %s in mixed English and Icelandic.\n\n", currentPhase, poetChoice)

	// Form the string to generate an image that includes the current moon phase in the style of the randomly chosen artist
	imagePrompt, _ := fmt.Printf("/imagine prompt: A beautiful landscape scene of the starry night sky with a silhouette of a woman with curly dark copper hair looking up at the perfect %s Moon in the stunning nightsky, painting by %s, astrophotography, accurate %s phase, spirituality, elegant, ethereal, 8k --v 4 --q 2 --ar 2:3 \n", currentPhase, artistChoice, currentPhase)

	// fmt.Printf("/imagine prompt: A beautiful landscape scene of the starry night sky with a silhouette of a woman with curly dark copper hair looking up at the perfect %s Moon in the stunning nightsky, painting by %s, astrophotography, accurate %s phase, spirituality, elegant, ethereal, 8k --v 4 --q 2 --ar 2:3 \n", currentPhase, artistChoice, currentPhase)

	fmt.Println(poemPrompt)
	fmt.Println(imagePrompt)

}

