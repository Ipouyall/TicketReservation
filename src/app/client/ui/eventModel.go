package ui

import (
	"TicketReservation/src/model"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/fatih/color"
	"strconv"
)

type EventModel struct {
	Events   []model.Event
	Selected int
}

func NewEventModel(events []model.Event) *EventModel {
	return &EventModel{
		Events:   events,
		Selected: 0,
	}
}

func (m EventModel) Init() tea.Cmd {
	return nil
}

func (m EventModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc", "q":
			return m, tea.Quit
		case "up":
			m.Selected--
			if m.Selected < 0 {
				m.Selected = len(m.Events) - 1
			}
		case "down":
			m.Selected++
			if m.Selected >= len(m.Events) {
				m.Selected = 0
			}
		case "enter":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m EventModel) View() string {
	view := ""
	for i, event := range m.Events {
		if i == m.Selected {
			nameStyle := lipgloss.NewStyle().Bold(true)
			view += nameStyle.Render(color.GreenString("[x]")+"\t"+event.Name) + "\n"
			view += "\tDate: " + event.Date.Format("2024-01-01 11:11") + "\n"
			view += "\tAvailable Tickets: " + strconv.Itoa(event.AvailableTickets) + "\n"
			view += "\tTotal Tickets: " + strconv.Itoa(event.TotalTickets) + "\n"
			view += "\t(ID: " + event.ID + ")\n\n"
		} else {
			view += "[ ] " + event.Name + " (" + event.Date.Format("2024-01-01") + ")\n"
			view += "\tAvailable Tickets: " + strconv.Itoa(event.AvailableTickets) + "\n\n"
		}
	}
	return view
}
