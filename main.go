package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Markdown Viewer")

	// Markdown content
	markdownContent := `
# Markdown Viewer

This is a simple Markdown viewer using Fyne and Go.

## Features

- Displays Markdown content.
- Easy to use.

### Try it out!

1. Enter Markdown text in the input area.
2. Click the "Render" button to see the rendered Markdown. Click the "Render" button to see the rendered Markdown. Click the "Render" button to see the rendered Markdown.

*Enjoy your Markdown experience!*
# Markdown Viewer

This is a simple Markdown viewer using Fyne and Go.

## Features

- Displays Markdown content.
- Easy to use.

### Try it out!

1. Enter Markdown text in the input area.
2. Click the "Render" button to see the rendered Markdown. Click the "Render" button to see the rendered Markdown. Click the "Render" button to see the rendered Markdown.

*Enjoy your Markdown experience!*
# Markdown Viewer

This is a simple Markdown viewer using Fyne and Go.

## Features

- Displays Markdown content.
- Easy to use.

### Try it out!

1. Enter Markdown text in the input area.
2. Click the "Render" button to see the rendered Markdown. Click the "Render" button to see the rendered Markdown. Click the "Render" button to see the rendered Markdown.

*Enjoy your Markdown experience!*
# Markdown Viewer

This is a simple Markdown viewer using Fyne and Go.

## Features

- Displays Markdown content.
- Easy to use.

### Try it out!

1. Enter Markdown text in the input area.
2. Click the "Render" button to see the rendered Markdown. Click the "Render" button to see the rendered Markdown. Click the "Render" button to see the rendered Markdown.

*Enjoy your Markdown experience!*
# Markdown Viewer

This is a simple Markdown viewer using Fyne and Go.

## Features

- Displays Markdown content.
- Easy to use.

### Try it out!

1. Enter Markdown text in the input area.
2. Click the "Render" button to see the rendered Markdown. Click the "Render" button to see the rendered Markdown. Click the "Render" button to see the rendered Markdown.

*Enjoy your Markdown experience!*
`

	// Create a text area for input
	inputArea := widget.NewMultiLineEntry()
	inputArea.SetPlaceHolder("Enter Markdown text here")
	inputArea.SetText(markdownContent)

	// Create a rich text widget for displaying rendered Markdown
	outputView := widget.NewRichTextFromMarkdown("")
	outputView.ParseMarkdown(markdownContent)

	// Create a button to render Markdown
	renderButton := widget.NewButton("Render", func() {
		// Get Markdown text from the input area
		markdownText := inputArea.Text

		// Set the Markdown content to the rich text widget
		outputView.ParseMarkdown(markdownText)
	})

	// Create a flex container for arranging widgets
	flexContainer := container.New(
		layout.NewAdaptiveGridLayout(2),
		container.New(layout.NewVBoxLayout(),
			widget.NewLabel("Input:"),
			inputArea,
			renderButton,
		),
		container.New(layout.NewGridLayout(1),
			widget.NewLabel("Output:"),
			container.NewVScroll(outputView),
		),
	)

	// Set the flex container as the main content of the window
	myWindow.SetContent(flexContainer)

	// Show the window
	myWindow.ShowAndRun()
	myWindow.Resize(fyne.Size{
		Width:  800,
		Height: 600,
	})
}
