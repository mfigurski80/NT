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

	fmt.Printf("Connecting to notion api...\n\n")

	INTEGRATION_KEY := notionapi.Token(os.Getenv("NOTION_INTEGRATION_TOKEN"))
	client := notionapi.NewClient(INTEGRATION_KEY)
	// searchForQuery(client, "bio")
	getPage(client, "5d2e668096784c78bec872d537b90d9d")
}

func getPage(client *notionapi.Client, id string) {
	blocks := fillBlockChildren(context.Background(), client, id)
	// fmt.Printf("FINAL RESULT:\n%-v\n", blocks)
	// fmt.Printf("Text: %-v\n", blocks[2].(*notionapi.ParagraphBlock).Paragraph.Children)

	for _, b := range blocks {
		fmt.Printf(stringifyBlock(b))
	}

	return
	// fmt.Printf("Result: %-v\n", res.Results)
	// for i, block := range res.Results {
	// // fmt.Printf("Block #%-2d: %s\n", i, block.GetType())
	// switch block.GetType() {
	// case "paragraph":
	//
	// fmt.Print(decipherParagraphBlock(
	// block.(*notionapi.ParagraphBlock),
	// client,
	// ))
	// case "divider":
	// fmt.Print(decipherDivider(
	// block.(*notionapi.DividerBlock),
	// ))
	// default:
	// fmt.Printf("#![%d] %s\n\n", i, block.GetType())
	// }
	// }
}

func getPageMeta(client *notionapi.Client, id string) {
	res, err := client.Page.Get(context.Background(), notionapi.PageID(id))
	if err != nil {
		panic(err)
	}
	fmt.Printf("Result: %-v\n", res)
	fmt.Printf(
		"Page: %-v\n",
		res.Properties["title"].(*notionapi.TitleProperty).Title[0].PlainText,
	)
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
