package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
    //check and validate api key
    congressApiKey := os.Getenv("CONGRESS_API_KEY")
    if congressApiKey == "" {
        log.Fatalln("CONGRESS_API_KEY env variable is not set")
    }

    client := &http.Client{}
    req, err := http.NewRequest("GET", "https://api.congress.gov/v3/bill", nil)
    if err != nil {
        log.Fatalln(err)
    }

    req.Header.Add("X-Api-Key", congressApiKey)
    resp, err := client.Do(req)
    if err != nil {
        log.Fatalln(err)
    }

    defer resp.Body.Close()
    b, err := io.ReadAll(resp.Body)
    if err != nil {
        log.Fatalln(err)
    }

    log.Println(string(b))
}
