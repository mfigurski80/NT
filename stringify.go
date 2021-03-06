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

func addIndentedChildren(hasChildren bool, children []notionapi.Block) string {
	if !hasChildren {
		return ""
	}
	ret := "\t"
	for _, ch := range children {
		ret += strings.ReplaceAll(stringifyBlock(ch), "\n", "\n\t")
	}
	return strings.TrimSuffix(ret, "\t")
}

func stringifyPageMeta(page *notionapi.Page) string {
	txt := decipherRichText(page.Properties["title"].(*notionapi.TitleProperty).Title)
	underline := strings.Repeat("=", len(txt))
	if page.Icon != nil && page.Icon.Emoji != nil {
		txt = string(*page.Icon.Emoji) + " " + txt
		underline += "==="
	}
	txt = underline + "\n" + txt + "\n" + underline + "\n\n"

	return txt
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
	case "to_do":
		return stringifyTodoBlock(block.(*notionapi.ToDoBlock))
	case "bookmark":
		return stringifyBookmarkBlock(block.(*notionapi.BookmarkBlock))
	case "equation":
		return stringifyEquationBlock(block.(*notionapi.EquationBlock))
	default:
		return fmt.Sprintf("//![%s]\n", block.GetType().String())
	}
}

func stringifyParagraphBlock(block *notionapi.ParagraphBlock) string {
	return decipherRichText(block.Paragraph.Text) + "\n" + addIndentedChildren(block.HasChildren, block.Paragraph.Children)
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
	p += addIndentedChildren(block.HasChildren, block.Toggle.Children)
	return p
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
	p += addIndentedChildren(block.HasChildren, block.BulletedListItem.Children)
	return p
}

func stringifyNumberedListItemBlock(block *notionapi.NumberedListItemBlock) string {
	p := "- " + decipherRichText(block.NumberedListItem.Text) + "\n"
	p += addIndentedChildren(block.HasChildren, block.NumberedListItem.Children)
	return p
}

func stringifyChildPageBlock(block *notionapi.ChildPageBlock) string {
	return fmt.Sprintf("[ %s Page ](%s)\n", block.ChildPage.Title, block.ID.String())
}

func stringifyCalloutBlock(block *notionapi.CalloutBlock) string {
	return fmt.Sprintf("\n[[ %s %s ]]\n\n", string(*block.Callout.Icon.Emoji), decipherRichText(block.Callout.Text))
}

func stringifyTodoBlock(block *notionapi.ToDoBlock) string {
	check := " "
	if block.ToDo.Checked {
		check = "x"
	}
	return fmt.Sprintf("[%s] %s\n", check, decipherRichText(block.ToDo.Text))
}

func stringifyBookmarkBlock(block *notionapi.BookmarkBlock) string {
	txt := decipherRichText(block.Bookmark.Caption)
	if txt != "" {
		txt += " "
	}
	return fmt.Sprintf("[ %sBookmark ](%s)\n", txt, block.Bookmark.URL)
}

func stringifyEquationBlock(block *notionapi.EquationBlock) string {
	return fmt.Sprintf("[ Equation: %s ]\n", block.Equation.Expression)
}
