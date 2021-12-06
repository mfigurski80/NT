package main

import (
	"context"
	"fmt"

	"github.com/jomei/notionapi"
)

func fillBlockChildren(ctx context.Context, client *notionapi.Client, id string) []notionapi.Block {
	// get blocks
	resp, err := client.Block.GetChildren(ctx, notionapi.BlockID(id), nil)
	if err != nil {
		panic(err)
	}
	blocks := resp.Results
	// hydrate blocks (check for children)
	for _, block := range blocks {
		// fmt.Printf("Block #%-2d: %s\n", i, block.GetType())
		switch block.GetType() {
		case "paragraph":
			p := block.(*notionapi.ParagraphBlock)
			if p.HasChildren {
				p.Paragraph.Children = fillBlockChildren(ctx, client, p.ID.String())
			}
		case "toggle":
			t := block.(*notionapi.ToggleBlock)
			if t.HasChildren {
				// fmt.Printf("\t[%-v] toggle has children\n", decipherRichText(t.Toggle.Text))
				t.Toggle.Children = fillBlockChildren(ctx, client, t.ID.String())
			}
		case "synced_block":
			b := block.(*notionapi.SyncedBlock)
			if b.HasChildren {
				b.SyncedBlock.Children = fillBlockChildren(ctx, client, b.ID.String())
			}
		case "bulleted_list_item":
			b := block.(*notionapi.BulletedListItemBlock)
			if b.HasChildren {
				b.BulletedListItem.Children = fillBlockChildren(ctx, client, b.ID.String())
			}
		case "numbered_list_item":
			b := block.(*notionapi.NumberedListItemBlock)
			if b.HasChildren {
				b.NumberedListItem.Children = fillBlockChildren(ctx, client, b.ID.String())
			}
		}

	}
	return blocks
}

func getPageMeta(ctx context.Context, client *notionapi.Client, id string) *notionapi.Page {
	res, err := client.Page.Get(ctx, notionapi.PageID(id))
	if err != nil {
		panic(err)
	}
	// fmt.Printf("Result: %-v\n", res)
	return res
}


func readPage(client *notionapi.Client, id string, readPageTitle bool) {
	ctx := context.Background()
	if readPageTitle {
		fmt.Println("Getting page meta...")
		met := getPageMeta(ctx, client, id)
		fmt.Printf(stringifyPageMeta(met))
	}
	// fmt.Printf("# %s %s\n",
	// *met.Icon.Emoji,
	// decipherRichText(met.Properties["title"].(*notionapi.TitleProperty).Title),
	// )
	blocks := fillBlockChildren(context.Background(), client, id)
	for _, b := range blocks {
		fmt.Printf(stringifyBlock(b))
	}
}
