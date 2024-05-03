package ui

func (m *Model) InitBaseMenu() {
	items := []Item{
		{ID: 1, Name: "Show Events", Description: "to show all existing events"},
		{ID: 2, Name: "Book Ticket", Description: "to book ticket for an event"},
		{ID: 3, Name: "New event", Description: "to create a new event"},
		{ID: 4, Name: "Help", Description: "Show help of the client app"},
		{ID: 5, Name: "Exit", Description: "To exit the app"},
	}
	m.Items = items
}
