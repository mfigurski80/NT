package main

import (
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

func stringifyBlock(block notionapi.Block) string {
	switch block.GetType() {
	case "paragraph":
		return stringifyParagraphBlock(block.(*notionapi.ParagraphBlock))
	case "divider":
		return stringifyDividerBlock(block.(*notionapi.DividerBlock))
	default:
		return fmt.Sprintf("//![%s]\n", block.GetType().String())
	}
}

func stringifyParagraphBlock(block *notionapi.ParagraphBlock) string {
	p := decipherRichText(block.Paragraph.Text) + "\n"
	// fmt.Printf("\t%-v\n", block.ID)
	for _, ch := range block.Paragraph.Children {
		p += stringifyBlock(ch)
	}
	return p + "\n"
}

func stringifyDividerBlock(block *notionapi.DividerBlock) string {
	return "***\n\n"
}
