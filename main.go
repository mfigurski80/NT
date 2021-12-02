package main

import (
	"context"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/jomei/notionapi"
)

func main() {
	godotenv.Load()

	INTEGRATION_KEY := notionapi.Token(os.Getenv("NOTION_INTEGRATION_TOKEN"))
	client := notionapi.NewClient(INTEGRATION_KEY)
	// searchForQuery(client, "bio")
	readPage(client, "aaf43ebf6bf6404eaf978588b6c3f0ca")
}
func searchForQuery(client *notionapi.Client, q string) {
	got, err := client.Search.Do(context.Background(), &notionapi.SearchRequest{
		Query: q,
	})
	if err != nil {
		panic(err)
	}
	// fmt.Printf("Results: %-v\n", got)
	for i, res := range got.Results {
		fmt.Printf("Got #%d: %-v\n", i, res)
	}
}
