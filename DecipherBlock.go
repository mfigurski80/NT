package main

import (
	"github.com/jomei/notionapi"
)

func decipherRichText(text []notionapi.RichText) string {
	buf := ""
	for _, p := range text {
		buf += p.PlainText
	}
	return buf
}

func decipherParagraphBlock(block *notionapi.ParagraphBlock) string {
	return decipherRichText(block.Paragraph.Text) + "\n"
}
