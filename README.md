# moon-poetry
Just a silly side project to make Linda smile, feel free to fork it and make someone else smile :)


Maybe add this stuff in the future:



```

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


main:
    response, err := callChatGptApi(poemPrompt)
	if err != nil {
	  fmt.Println(err)
	} else {
	  fmt.Println(response)
	}

	// Call the CreateDiscordAPICall() function and pass the message
	CreateDiscordAPICall(imagePrompt)


```