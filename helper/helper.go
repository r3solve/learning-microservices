package helper

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Quote struct {
	Category string `json:"category"`
	Quote    string `json:"quote"`
	Author   string `json:"author"`
}

func GetQuote(category string) ([]Quote, error) {
	apiURL := fmt.Sprintf("https://api.api-ninjas.com/v1/quotes?category=%s", category)
	request, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	request.Header.Set("X-Api-Key", "EoJtN3S7jZgFB4kZmeJPbxhyc6DdbyCUo96J84vo")
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("error making request: %v", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: received status code %d", response.StatusCode)
	}

	resBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	var quotes []Quote
	err = json.Unmarshal(resBody, &quotes)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling response: %v", err)
	}

	return quotes, nil
}

// func main() {
// 	quotes, err := GetQuote("happiness")
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}

// 	fmt.Println(quotes)
// }
