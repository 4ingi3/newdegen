package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "os"
)

const apiURL = "https://api.coingecko.com/api/v3/simple/price?ids=bitcoin&vs_currencies=usd"

type Response struct {
    Bitcoin struct {
        USD float64 `json:"usd"`
    } `json:"bitcoin"`
}

func main() {
    resp, err := http.Get(apiURL)
    if err != nil {
        fmt.Println("Error fetching data:", err)
        os.Exit(1)
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        fmt.Println("Error: received non-200 response code")
        os.Exit(1)
    }

    var response Response
    if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
        fmt.Println("
