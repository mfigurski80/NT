package main

import (
	"context"
	"fmt"

	"github.com/jomei/notionapi"
)

func decipherRichText(text []notionapi.RichText) string {
	buf := ""
	for _, p := range text {
		buf += p.PlainText
	}
	return buf
}

func decipherParagraphBlock(block *notionapi.ParagraphBlock, client *notionapi.Client) string {
	p := decipherRichText(block.Paragraph.Text)
	fmt.Printf("\t%-v\n", block.ID)
	client.Block.GetChildren(context.Background(), block.ID, nil)
	return p + "\n\n"
}

func decipherDivider(block *notionapi.DividerBlock) string {
	return "***\n\n"
}
