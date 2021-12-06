package main

import (
	"flag"
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

	readCmd := flag.NewFlagSet("read", flag.ExitOnError)
	readTitleFlag := readCmd.Bool("t", false, "Read title of page")

	if len(os.Args) < 2 {
		fmt.Println("No command given")
		os.Exit(1)
	}

	switch(os.Args[1]) {
	case "read":
		readCmd.Parse(os.Args[2:])
		// fmt.Println("Subcommand Read... with title? ", *readTitleFlag)
		// fmt.Println("   Tail:", readCmd.Args())
		for _, arg := range readCmd.Args() {
			readPage(client, arg, *readTitleFlag)
		}
	default:
		fmt.Println("Unknown command:", os.Args[1])
		os.Exit(1)
	}

	// searchForQuery(client, "bio")
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
