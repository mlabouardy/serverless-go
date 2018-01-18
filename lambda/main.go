package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/aws/aws-lambda-go/lambda"
)

var (
	API_KEY      = os.Getenv("API_KEY")
	ErrorBackend = errors.New("Something went wrong")
)

type Request struct {
	ID int `json:"id"`
}

type MovieDBResponse struct {
	Movies []Movie `json:"results"`
}

type Movie struct {
	Title       string `json:"title"`
	Description string `json:"overview"`
	Cover       string `json:"poster_path"`
	ReleaseDate string `json:"release_date"`
}

func Handler(request Request) ([]Movie, error) {
	url := fmt.Sprintf("https://api.themoviedb.org/3/discover/movie?api_key=%s", API_KEY)

	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return []Movie{}, ErrorBackend
	}

	if request.ID > 0 {
		q := req.URL.Query()
		q.Add("with_genres", strconv.Itoa(request.ID))
		req.URL.RawQuery = q.Encode()
	}

	resp, err := client.Do(req)
	if err != nil {
		return []Movie{}, ErrorBackend
	}
	defer resp.Body.Close()

	var data MovieDBResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return []Movie{}, ErrorBackend
	}

	return data.Movies, nil
}

func main() {
	lambda.Start(Handler)
}
