package main

import (
	"fmt"
	"strings"

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
	case "toggle":
		return stringifyToggleBlock(block.(*notionapi.ToggleBlock))
	case "synced_block":
		return stringifySyncedBlock(block.(*notionapi.SyncedBlock))
	case "bulleted_list_item":
		return stringifyBulletedListItemBlock(block.(*notionapi.BulletedListItemBlock))
	case "numbered_list_item":
		return stringifyNumberedListItemBlock(block.(*notionapi.NumberedListItemBlock))
	case "child_page":
		return stringifyChildPageBlock(block.(*notionapi.ChildPageBlock))
	case "callout":
		return stringifyCalloutBlock(block.(*notionapi.CalloutBlock))
	default:
		return fmt.Sprintf("//![%s]\n", block.GetType().String())
	}
}

func stringifyParagraphBlock(block *notionapi.ParagraphBlock) string {
	p := decipherRichText(block.Paragraph.Text) + "\n"
	if block.HasChildren {
		p += "\t"
	}
	// fmt.Printf("\t%-v\n", block.ID)
	for _, ch := range block.Paragraph.Children {
		p += strings.ReplaceAll(stringifyBlock(ch), "\n", "\n\t")
	}
	return strings.TrimSuffix(p, "\t")
}

func stringifyHeading1Block(block *notionapi.Heading1Block) string {
	return "\n# " + decipherRichText(block.Heading1.Text) + " #\n\n"
}

func stringifyHeading2Block(block *notionapi.Heading2Block) string {
	return "\n## " + decipherRichText(block.Heading2.Text) + " ##\n\n"
}

func stringifyHeading3Block(block *notionapi.Heading3Block) string {
	return "\n### " + decipherRichText(block.Heading3.Text) + " ###\n\n"
}

func stringifyDividerBlock(block *notionapi.DividerBlock) string {
	return "\n-----------\n\n"
}

func stringifyToggleBlock(block *notionapi.ToggleBlock) string {
	p := "[ " + decipherRichText(block.Toggle.Text) + " Toggle ]\n"
	// p += fmt.Sprintf("\tchildren: %-v\n", block.Toggle.Children)
	if block.HasChildren {
		p += "\t"
	}
	for _, ch := range block.Toggle.Children {
		p += strings.ReplaceAll(stringifyBlock(ch), "\n", "\n\t")

	}
	return strings.TrimSuffix(p, "\t")
}

func stringifySyncedBlock(block *notionapi.SyncedBlock) string {
	p := ""
	for _, ch := range block.SyncedBlock.Children {
		p += stringifyBlock(ch)
	}
	return p
}

func stringifyBulletedListItemBlock(block *notionapi.BulletedListItemBlock) string {
	p := "- " + decipherRichText(block.BulletedListItem.Text) + "\n"
	for _, ch := range block.BulletedListItem.Children {
		p += stringifyBlock(ch)
	}
	return p
}

func stringifyNumberedListItemBlock(block *notionapi.NumberedListItemBlock) string {
	p := "- " + decipherRichText(block.NumberedListItem.Text) + "\n"
	for _, ch := range block.NumberedListItem.Children {
		p += stringifyBlock(ch)
	}
	return p
}

func stringifyChildPageBlock(block *notionapi.ChildPageBlock) string {
	return fmt.Sprintf("( %s Page )[ %s ]\n", block.ChildPage.Title, block.ID.String())
}

func stringifyCalloutBlock(block *notionapi.CalloutBlock) string {
	return fmt.Sprintf("\n[[ %s %s ]]\n\n", string(*block.Callout.Icon.Emoji), decipherRichText(block.Callout.Text))
}
