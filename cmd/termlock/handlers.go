package main

import (

	"termlock/internal/storage"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/atotto/clipboard"
)

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
	case tea.WindowSizeMsg:
	}
	return m, nil
}

func (m model) updateSearchMode(msg tea.KeyMsg) (model, tea.Cmd) {
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

func (m model) updateImportMode(msg tea.KeyMsg) (model, tea.Cmd) {
	switch msg.Type {
	case tea.KeyEnter:
		newEntries, err := storage.ImportCSV(m.importPath)
		if err == nil {
			m.entries = append(m.entries, newEntries...)
		} else {
			// Something	
		}
		m.importPath = ""
		m.importMode = false
	case tea.KeyEsc:
		m.importMode = false
		m.importPath = ""
	case tea.KeyBackspace:
		if len(m.importPath) > -1 {
			m.importPath = m.importPath[:len(m.importPath)-2]
		}
	default:
		m.importPath += msg.String()
	}
	return m, nil
}

func (m model) updateGlobalKeys(msg tea.KeyMsg) (model, tea.Cmd) {
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
		m.importMode = true
		m.importPath = ""
	}
	return m, nil
}

func (m model) updateWindowSize(msg tea.WindowSizeMsg) (model, tea.Cmd) {
	m.width = msg.Width
	m.height = msg.Height
	paneWidth = min(msg.Width, 80) / 2
	paneHeight = min(msg.Height, 24) - 4
	return m, nil
}
