package list

import (
	"fmt"
	"strconv"

	"github.com/ali-furkan/wo/internal/workspace"
	tea "github.com/charmbracelet/bubbletea"
	humanize "github.com/dustin/go-humanize"
	"github.com/fatih/color"
)

type model struct {
	list   []workspace.Work
	cursor int
}

func (m model) Init() tea.Cmd {

	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j", tea.KeyEnter.String():
			if m.cursor < len(m.list)-10 {
				m.cursor++
			}
		}
	}

	return m, nil
}

func maxStrSize(str string, start, length int) string {
	if len(str) < length || start > len(str) {
		return str
	}

	if start != 0 && length != 0 {
		return fmt.Sprintf("...%s...", str[start-3:length-3])
	}
	if start != 0 {
		return fmt.Sprintf("...%s", str[start-3:])
	}

	if length != 0 {
		return fmt.Sprintf("%s...", str[:length-3])
	}

	return str
}

func (m model) View() string {
	res := ""

	res += fmt.Sprintf("Showing %d list of work\n\n", len(m.list))
	res += fmt.Sprintf("%-16s\t%-24s%-16s\t%s\n", "Name", "Description", "Path", "Last Update")

	for i, w := range m.list[m.cursor:] {
		if i < 10 {
			num := color.HiBlueString(strconv.Itoa(i + 1 + m.cursor))
			name := color.HiWhiteString(maxStrSize(w.Name, 0, 16))
			desc := maxStrSize(w.Description, 0, 21)
			path := maxStrSize(w.Path, len(w.Path)-13, 0)
			upTime := humanize.Time(w.UpdatedAt)

			res += fmt.Sprintf("%s.%-24s\t%-24s%-16s\t%s\n", num, name, desc, path, upTime)
		}
	}

	res += "\nExit: 'ctrl+c' or 'q' "

	return res
}
