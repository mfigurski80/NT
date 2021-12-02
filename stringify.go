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
	case "heading_1":
		return stringifyHeading1Block(block.(*notionapi.Heading1Block))
	case "heading_2":
		return stringifyHeading2Block(block.(*notionapi.Heading2Block))
	case "heading_3":
		return stringifyHeading3Block(block.(*notionapi.Heading3Block))
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

func stringifyHeading1Block(block *notionapi.Heading1Block) string {
	return "# " + decipherRichText(block.Heading1.Text) + "\n\n"
}

func stringifyHeading2Block(block *notionapi.Heading2Block) string {
	return "## " + decipherRichText(block.Heading2.Text) + "\n\n"
}

func stringifyHeading3Block(block *notionapi.Heading3Block) string {
	return "### " + decipherRichText(block.Heading3.Text) + "\n\n"
}

func stringifyDividerBlock(block *notionapi.DividerBlock) string {
	return "***\n\n"
}
