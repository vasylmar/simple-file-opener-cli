package main

import (
	"log"
	"opencli/utils"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const instruction = "q: quit, j/down: next, k/up: previous, enter: open file"

var (
	selectedOptions  = lipgloss.NewStyle().Foreground(lipgloss.Color("56"))
	options          = lipgloss.NewStyle().Foreground(lipgloss.Color("15"))
	cursorStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("15"))
	box              = lipgloss.NewStyle().PaddingLeft(2).PaddingTop(1).PaddingBottom(2)
	instructionStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("8")).PaddingTop(1)
	errorStyle       = lipgloss.NewStyle().Foreground(lipgloss.Color("9")).PaddingTop(1)
	emptyStyle       = lipgloss.NewStyle().Foreground(lipgloss.Color("56"))
)

type model struct {
	files   []string
	current int
	error   error
}

func (m model) Init() tea.Cmd {
	if len(m.files) == 0 {
		return tea.Quit
	}
	return nil
}

func NewModel() model {
	currentDirItems := &utils.CurrendDirItems{}

	curDir, err := utils.GetCurDir()
	if err != nil {
		log.Fatal(err)
	}
	readErr := currentDirItems.ReadCurDir(curDir)
	if readErr != nil {
		log.Fatal(err)
	}

	files := currentDirItems.GetFilesFromCurDir()

	return model{files, 0, nil}
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc":
			return m, tea.Quit
		case "j", "down":
			if m.current < len(m.files)-1 {
				m.current++
			}
		case "k", "up":
			if m.current > 0 {
				m.current--
			}
		case "enter":
			err := utils.OpenFile(m.files[m.current])
			if err != nil {
				m.error = err
			} else {
				return m, tea.Quit
			}
		}
	}

	return m, nil
}

func (m model) View() string {
	var output string
	var view string
	cursor := cursorStyle.Render("> ")

	for i, file := range m.files {
		if i == m.current {
			output += cursor + selectedOptions.Render(file) + "\n"
		} else {
			output += "  " + options.Render(file) + "\n"
		}
	}

	if m.error != nil {
		view = output + instructionStyle.Render(instruction) + "\n" +
			errorStyle.Render("Error opening file: "+m.error.Error()+"\n")
	} else {
		view = output + instructionStyle.Render(instruction)
	}

	if len(m.files) == 0 {
		view = emptyStyle.Render("Nothing here...")
	}

	return box.Render(view)
}

func main() {
	p := tea.NewProgram(NewModel())

	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
