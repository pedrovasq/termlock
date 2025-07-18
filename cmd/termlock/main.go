 package main

 import (
	 "fmt"
	 "os"
	 "strings"

	 "termlock/internal/models"
	 "termlock/internal/styles"
	 "termlock/internal/storage"
	 "termlock/internal/tui"

	 "github.com/sahilm/fuzzy"
	 "github.com/atotto/clipboard"
	 tea "github.com/charmbracelet/bubbletea"
	 lip "github.com/charmbracelet/lipgloss"
)

var (
	paneWidth 	= 0
	paneHeight	= 0
)

type model struct {
	cursor 			int
	entries			[]models.PasswordEntry
	filtered		[]models.PasswordEntry
	searchMode		bool
	searchQuery		string
	importMode		bool
	importPath		string
	width, height 	int
	copied 			bool
}

type clearCopiedMsg struct {}

func initialModel() model {
	return model{
		cursor: 0,
		entries: []models.PasswordEntry{
			{
				Title:    "GitHub",
				Username: "raider54",
				Password: "ghp_********",
				Sites:    []string{"https://github.com"},
				Note:     "Personal repo access token",
			},
			{
				Title:    "Gmail",
				Username: "topdog69@gmail.com",
				Password: "********",
				Sites:    []string{"https://mail.google.com"},
				Note:     "Main email",
			},
			{
				Title:    "CanineCloud",
				Username: "doglover22",
				Password: "woofwoof",
				Sites:    []string{"https://caninecloud.com"},
				Note:     "Website for looking up pictures of dog\nThis Website has helped through the pandemic",
			},
			{
				Title:    "Youtube",
				Username: "minecraft_merchant",
				Password: "mine****",
				Sites:    []string{"youtube.com"},
				Note:     "My minecraft youtube channel",
			},
			{
				Title:    "ChatGPT",
				Username: "topdog69@gmail.com",
				Password: "ai-sux420",
				Sites:    []string{"https://chatgpt.com"},
				Note:     "My Personal ChatGPT Account",
			},
			{
				Title:    "X (Twitter)",
				Username: "doglover22",
				Password: "barkbark",
				Sites:    []string{"https://x.com", "https://twitter.com"},
				Note:     "Why did elon do it?",
			},
		},
	}
}
	 
func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if m.searchMode {
			switch msg.Type {
			case tea.KeyEsc:
				m.searchMode = false
				m.searchQuery = ""
				m.filtered = nil
			case tea.KeyEnter:
				selected := m.entries[m.cursor]
				clipboard.WriteAll(selected.Password)
				m.copied = true
			case tea.KeyBackspace:
				if len(m.searchQuery) > 1 {
					m.searchQuery = m.searchQuery[:len(m.searchQuery)-1]
					m.filtered = filterEntries(m.entries, m.searchQuery)
				} else {
					m.searchQuery = ""
					m.filtered = nil
				}
			case tea.KeyUp, tea.KeyCtrlK:
				if m.cursor > 0 {
					m.cursor--
					m.copied = false
				}
			case tea.KeyDown, tea.KeyCtrlJ:
				if m.cursor < len(m.filtered) - 1 {
					m.cursor++
					m.copied = false
				}
			case tea.KeyCtrlC:
				return m, tea.Quit	
			default:
				m.cursor = 0
				if len(msg.String()) == 1 { 
					m.searchQuery += msg.String()
					m.filtered = filterEntries(m.entries, m.searchQuery)
				}
			}
			return m, nil
		}
		if m.importMode {
			switch msg.Type {
			case tea.KeyEnter:
				newEntries, err := storage.ImportCSV(m.importPath)
				if err == nil {
					m.entries = append(m.entries, newEntries...)
				} else {
					
				}
				m.importPath = ""
				m.importMode = false
			case tea.KeyEsc:
				m.importMode = false
				m.importPath = ""
			case tea.KeyBackspace:
				if len(m.importPath) > 0 {
					m.importPath = m.importPath[:len(m.importPath)-1]
				}
			default:
				m.importPath += msg.String()
			}
			return m, nil
		}
		switch msg.String() {
		case "ctrl+c", "q":
			m.copied = false
			return m, tea.Quit
		case "j", "down":
			if m.cursor < len(m.entries) - 1 {
				m.cursor++
				m.copied = false
			}
		case "k", "up":
			if m.cursor > 0 {
				m.cursor--
				m.copied = false
			}
		case "enter":
			selected := m.entries[m.cursor]
			clipboard.WriteAll(selected.Password)
			m.copied = true
		case "/":
			m.searchMode = true
			m.searchQuery = ""
			m.filtered = m.entries
		case "i":

		}
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		paneWidth = min(msg.Width, 80) / 2
		paneHeight = min(msg.Height, 20) - 4
	}
	return m, nil
}

func (m model) View() string {
	leftPaneStyle := styles.Border.Width(paneWidth).Height(paneHeight)
	rightPaneStyle := styles.Border.Width(paneWidth).Height(paneHeight)
	
	var leftContent string;
	entriesToShow := m.entries
	if m.searchMode {
		entriesToShow = m.filtered
		leftContent = styles.Title.Render("Search:") + " " +
		styles.Faded.Render(m.searchQuery) + "\n\n"
	} else if m.importMode {
		entriesToShow = []models.PasswordEntry{}
		leftContent = styles.Title.Render("Import:") + " " +
		styles.Faded.Render(m.importPath) + "\n\n"
	} else {
		leftContent = styles.Title.Render("Termlock") + "\n\n"
	}
	
	if len(entriesToShow) == 0 {
		leftContent += styles.Error.Render("No entries found.") + "\n"
	} else {
		for i, entry := range entriesToShow {
			if i == m.cursor {
				leftContent += styles.HighLight.Render("> " + entry.Title) + "\n"
			} else {
				leftContent += styles.Normal.Render("  " + entry.Title) + "\n"
			}
		}
	}
	leftBox := leftPaneStyle.Render(leftContent)

	rightContent := styles.Title.Render("Details") + "\n\n" 

	if len(entriesToShow) == 0 {
		rightContent += styles.Error.Render("No entry selected.") + "\n"
	} else {

		selected := entriesToShow[m.cursor]
		rightContent += fmt.Sprintf("Username: 	%s\n", selected.Username) +
		fmt.Sprintf("Password: 	%s\n\n", styles.Password.Render(selected.Password)) +
		fmt.Sprintf("Sites:\n%s\n\n", styles.Link.Render(strings.Join(selected.Sites, ", "))) +
		fmt.Sprintf("Note:\n%s\n", selected.Note) 
		if m.copied {
			rightContent += "\n\n" + styles.HighLight.Render("✔  Password copied to clipboard")
		}
	}

	rightBox := rightPaneStyle.Render(rightContent)

	var footer string
	if m.searchMode {
		footer = tui.RenderFooter("search")
	} else {
		footer = tui.RenderFooter("default")
	}

	mainContent := lip.JoinHorizontal(lip.Top, leftBox, rightBox)
	combined := lip.JoinVertical(lip.Center, mainContent, footer)

	centered := lip.Place(m.width, m.height,
		lip.Center, lip.Center, combined,
	)
	
	return centered
}

func filterEntries(entries []models.PasswordEntry, query string) []models.PasswordEntry {
	titles := make([]string, len(entries))
	
	if query == "" {
		return entries
	}

	for i, e := range entries {
		titles[i] = e.Title
	}

	matches := fuzzy.Find(query, titles)

	result := make([]models.PasswordEntry, len(matches))
	for i, match := range matches {
		result[i] = entries[match.Index]
	}
	return result
}

func main() {
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		fmt.Println("error starting termlock:", err)
		os.Exit(1)
	}
}
