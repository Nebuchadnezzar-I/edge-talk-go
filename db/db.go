package db

import (
    "encoding/json"
    "fmt"
    "os"

    "github.com/supabase-community/supabase-go"
    "github.com/joho/godotenv"
)

type Negotiation struct {
    ID        string  `json:"id"`
    SessionID *string `json:"session_id"`
    Name      string  `json:"name"`
    CreatedAt string  `json:"created_at"`
    UpdatedAt string  `json:"updated_at"`
}

var client *supabase.Client

func Initialize() error {
    err := godotenv.Load()
    if err != nil {
        return fmt.Errorf("error loading .env file: %v", err)
    }

    supabaseURL := os.Getenv("SUPABASE_URL")
    supabaseKey := os.Getenv("SUPABASE_KEY")

    client, err = supabase.NewClient(supabaseURL, supabaseKey, &supabase.ClientOptions{})
    if err != nil {
        return fmt.Errorf("cannot initialize client: %v", err)
    }

    return nil
}

func FetchNegotiations() ([]Negotiation, error) {
    data, _, err := client.From("negotiation").Select("*", "exact", false).Execute()
    if err != nil {
        return nil, fmt.Errorf("error fetching data: %v", err)
    }

    var negotiations []Negotiation
    err = json.Unmarshal(data, &negotiations)
    if err != nil {
        return nil, fmt.Errorf("error unmarshalling data: %v", err)
    }

    return negotiations, nil
}
