package main

import (
	"context"

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
		}
	}
	return blocks
}
