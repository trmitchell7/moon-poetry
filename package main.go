package main

import (
	"fmt"
	"math/rand"
	"math/mod"
	"time"
	"bytes"
	"io/ioutil"
	"net/http"
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


func currentMoonPhase() {
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
    if daysSinceNewMoon < 1 {
        phase = "New"
    } else if daysSinceNewMoon < 7 {
        phase = "Waxing Crescent"
    } else if daysSinceNewMoon < 8 {
        phase = "First Quarter"
    } else if daysSinceNewMoon < 16 {
        phase = "Waxing Gibbous"
    } else if daysSinceNewMoon < 17 {
        phase = "Full"
    } else if daysSinceNewMoon < 25 {
        phase = "Waning Gibbous"
    } else if daysSinceNewMoon < 26 {
        phase = "Third Quarter"
    } else {
        phase = "Waning Crescent"
    }

    // Return the current phase of the moon.
    return phase
}


func callChatGptApi(prompt string) {
	// Set up the HTTP request to the ChatGPT API
	url := "https://api.openai.com/v1/chatbot/generate"
	payload := bytes.NewBuffer([]byte(prompt))
	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
	  return "", err
	}
  
	// Add the necessary headers to the request
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer <YOUR_API_KEY>")
  
	// Send the request and get the response
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
	  return "", err
	}
	defer resp.Body.Close()
  
	// Read the response body and return it
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
	  return "", err
	}
	return string(body), nil
}


func CreateDiscordAPICall(message string) {
    // Set the base URL for the Discord API
    baseURL := "https://discordapp.com/api/v6"

    // Set the endpoint for the API call
    endpoint := "/channels/{channel.id}/messages"

    // Set the HTTP method for the API call
    method := "POST"

    // Create the HTTP client
    client := &http.Client{}

    // Create the request
    req, err := http.NewRequest(method, baseURL+endpoint, nil)
    if err != nil {
        fmt.Println(err)
        return
    }

    // Set the necessary headers for the API call
    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("Authorization", "Bot <YOUR_BOT_TOKEN>")

    // Add the message to the request body
    req.Body = ioutil.NopCloser(strings.NewReader(`{"content": "` + message + `"}`))

    // Execute the API call
    resp, err := client.Do(req)
    if err != nil {
        fmt.Println(err)
        return
    }

    // Print the response
    fmt.Println(resp)
}
  

func main() {
	// Get the current phase of the moon.
	currentPhase := currentMoonPhase()

	// Choose a random poet from the list of poets
	poetChoice := poets[rand.Intn(len(poets))]

	// Choose a random artist from the list of artists
	artistChoice := artists[rand.Intn(len(artists))]

	// Form the string to write a poem that includes the current moon phase and randomly chosen poet
	poemPrompt := fmt.Printf("Write a lovely poem about a beautiful woman named Linda as if she was the moon, taking into account last night's %s Moon and what the phase of the moon represents to Linda, done in the style of %s in mixed English and Icelandic.\n\n", currentPhase, poetChoice)

	// Form the string to generate an image that includes the current moon phase in the style of the randomly chosen artist
	imagePrompt := fmt.Printf("/imagine prompt: A beautiful landscape scene of the starry night sky with a silhouette of a woman with curly dark copper hair looking up at the perfect %s Moon in the stunning nightsky, painting by %s, astrophotography, accurate %s phase, spirituality, elegant, ethereal, 8k --v 4 --q 2 --ar 2:3 \n", currentPhase, artistChoice, currentPhase)

	fmt.Println(poemPrompt)
	fmt.Println(imagePrompt)

	response, err := callChatGptApi(poemPrompt)
	if err != nil {
	  fmt.Println(err)
	} else {
	  fmt.Println(response)
	}

	// Call the CreateDiscordAPICall() function and pass the message
	CreateDiscordAPICall(imagePrompt)

}

