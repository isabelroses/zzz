package main

import "github.com/charmbracelet/lipgloss"

// SnippetsStyle is the style struct to handle the focusing and blurring of the
// snippets pane in the application.
type SnippetsStyle struct {
	Focused SnippetsBaseStyle
	Blurred SnippetsBaseStyle
}

// FoldersStyle is the style struct to handle the focusing and blurring of the
// folders pane in the application.
type FoldersStyle struct {
	Focused FoldersBaseStyle
	Blurred FoldersBaseStyle
}

// ContentStyle is the style struct to handle the focusing and blurring of the
// content pane in the application.
type ContentStyle struct {
	Focused ContentBaseStyle
	Blurred ContentBaseStyle
}

// SnippetsBaseStyle holds the neccessary styling for the snippets pane of
// the application.
type SnippetsBaseStyle struct {
	Base               lipgloss.Style
	Title              lipgloss.Style
	TitleBar           lipgloss.Style
	SelectedSubtitle   lipgloss.Style
	UnselectedSubtitle lipgloss.Style
	SelectedTitle      lipgloss.Style
	UnselectedTitle    lipgloss.Style
	CopiedTitleBar     lipgloss.Style
	CopiedTitle        lipgloss.Style
	CopiedSubtitle     lipgloss.Style
	DeletedTitleBar    lipgloss.Style
	DeletedTitle       lipgloss.Style
	DeletedSubtitle    lipgloss.Style
}

// FoldersBaseStyle holds the neccessary styling for the folders pane of
// the application.
type FoldersBaseStyle struct {
	Base       lipgloss.Style
	Title      lipgloss.Style
	TitleBar   lipgloss.Style
	Selected   lipgloss.Style
	Unselected lipgloss.Style
}

// ContentBaseStyle holds the neccessary styling for the content pane of the
// application.
type ContentBaseStyle struct {
	Base         lipgloss.Style
	Title        lipgloss.Style
	Separator    lipgloss.Style
	LineNumber   lipgloss.Style
	EmptyHint    lipgloss.Style
	EmptyHintKey lipgloss.Style
}

// Styles is the struct of all styles for the application.
type Styles struct {
	Snippets SnippetsStyle
	Folders  FoldersStyle
	Content  ContentStyle
}

var marginStyle = lipgloss.NewStyle().Margin(1, 0, 0, 1)

// DefaultStyles is the default implementation of the styles struct for all
// styling in the application.
func DefaultStyles(config Config) Styles {
	text := lipgloss.Color(config.TextColor)
	subtext := lipgloss.Color(config.SubTextColor)
	textInvert := lipgloss.Color(config.TextInvertColor)
	gray := lipgloss.Color(config.GrayColor)
	green := lipgloss.Color(config.GreenColor)
	brightGreen := lipgloss.Color(config.BrightGreenColor)
	primary := lipgloss.Color(config.PrimaryColor)
	subprimary := lipgloss.Color(config.PrimaryColorSubdued)
	red := lipgloss.Color(config.RedColor)
	brightRed := lipgloss.Color(config.BrightRedColor)

	return Styles{
		Snippets: SnippetsStyle{
			Focused: SnippetsBaseStyle{
				Base:               lipgloss.NewStyle().Width(35),
				TitleBar:           lipgloss.NewStyle().Background(primary).Width(35-2).Margin(0, 1, 1, 1).Padding(0, 1).Foreground(textInvert),
				SelectedSubtitle:   lipgloss.NewStyle().Foreground(subprimary),
				UnselectedSubtitle: lipgloss.NewStyle().Foreground(subtext),
				SelectedTitle:      lipgloss.NewStyle().Foreground(primary),
				UnselectedTitle:    lipgloss.NewStyle().Foreground(text),
				CopiedTitleBar:     lipgloss.NewStyle().Background(green).Width(35-2).Margin(0, 1, 1, 1).Padding(0, 1).Foreground(textInvert),
				CopiedTitle:        lipgloss.NewStyle().Foreground(brightGreen),
				CopiedSubtitle:     lipgloss.NewStyle().Foreground(green),
				DeletedTitleBar:    lipgloss.NewStyle().Background(red).Width(35-2).Margin(0, 1, 1, 1).Padding(0, 1).Foreground(textInvert),
				DeletedTitle:       lipgloss.NewStyle().Foreground(brightRed),
				DeletedSubtitle:    lipgloss.NewStyle().Foreground(red),
			},
			Blurred: SnippetsBaseStyle{
				Base:               lipgloss.NewStyle().Width(35),
				TitleBar:           lipgloss.NewStyle().Background(gray).Width(35-2).Margin(0, 1, 1, 1).Padding(0, 1).Foreground(primary),
				SelectedSubtitle:   lipgloss.NewStyle().Foreground(primary),
				UnselectedSubtitle: lipgloss.NewStyle().Foreground(subtext),
				SelectedTitle:      lipgloss.NewStyle().Foreground(primary),
				UnselectedTitle:    lipgloss.NewStyle().Foreground(subtext),
				CopiedTitleBar:     lipgloss.NewStyle().Background(green).Width(35-2).Margin(0, 1, 1, 1).Padding(0, 1),
				CopiedTitle:        lipgloss.NewStyle().Foreground(brightGreen),
				CopiedSubtitle:     lipgloss.NewStyle().Foreground(green),
				DeletedTitleBar:    lipgloss.NewStyle().Background(red).Width(35-2).Margin(0, 1, 1, 1).Padding(0, 1),
				DeletedTitle:       lipgloss.NewStyle().Foreground(brightRed),
				DeletedSubtitle:    lipgloss.NewStyle().Foreground(red),
			},
		},
		Folders: FoldersStyle{
			Focused: FoldersBaseStyle{
				Base:       lipgloss.NewStyle().Width(22),
				Title:      lipgloss.NewStyle().Padding(0, 1).Foreground(textInvert),
				TitleBar:   lipgloss.NewStyle().Background(primary).Width(22-2).Margin(0, 1, 1, 1),
				Selected:   lipgloss.NewStyle().Foreground(primary),
				Unselected: lipgloss.NewStyle().Foreground(text),
			},
			Blurred: FoldersBaseStyle{
				Base:       lipgloss.NewStyle().Width(22),
				Title:      lipgloss.NewStyle().Padding(0, 1).Foreground(primary),
				TitleBar:   lipgloss.NewStyle().Background(gray).Width(22-2).Margin(0, 1, 1, 1),
				Selected:   lipgloss.NewStyle().Foreground(primary),
				Unselected: lipgloss.NewStyle().Foreground(subtext),
			},
		},
		Content: ContentStyle{
			Focused: ContentBaseStyle{
				Base:         lipgloss.NewStyle().Margin(0, 1),
				Title:        lipgloss.NewStyle().Background(primary).Foreground(textInvert).Margin(0, 0, 1, 1).Padding(0, 1),
				Separator:    lipgloss.NewStyle().Foreground(text).Margin(0, 0, 1, 1),
				LineNumber:   lipgloss.NewStyle().Foreground(text),
				EmptyHint:    lipgloss.NewStyle().Foreground(gray),
				EmptyHintKey: lipgloss.NewStyle().Foreground(primary),
			},
			Blurred: ContentBaseStyle{
				Base:         lipgloss.NewStyle().Margin(0, 1),
				Title:        lipgloss.NewStyle().Background(gray).Foreground(primary).Margin(0, 0, 1, 1).Padding(0, 1),
				Separator:    lipgloss.NewStyle().Foreground(text).Margin(0, 0, 1, 1),
				LineNumber:   lipgloss.NewStyle().Foreground(subtext),
				EmptyHint:    lipgloss.NewStyle().Foreground(gray),
				EmptyHintKey: lipgloss.NewStyle().Foreground(primary),
			},
		},
	}
}
