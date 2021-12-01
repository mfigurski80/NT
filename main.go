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

	fmt.Println("Connecting to notion api...")

	INTEGRATION_KEY := notionapi.Token(os.Getenv("NOTION_INTEGRATION_TOKEN"))
	client := notionapi.NewClient(INTEGRATION_KEY)
	// searchForQuery(client, "bio")
	getPage(client, "aaf43ebf6bf6404eaf978588b6c3f0ca")
}

func getPage(client *notionapi.Client, id string) {
	res, err := client.Page.Get(context.Background(), notionapi.PageID(id))
	if err != nil {
		panic(err)
	}
	fmt.Printf("Result: %-v\n", res)
	fmt.Printf("Page: %-v\n", res)
}

func searchForQuery(client *notionapi.Client, q string) {
	got, err := client.Search.Do(context.Background(), &notionapi.SearchRequest{
		Query: q,
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Results: %-v\n", got)
	for i, res := range got.Results {
		fmt.Printf("Got #%d: %-v\n", i, res)
	}
}
