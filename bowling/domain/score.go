package domain

import "fmt"

type Score interface {
	PrintInfoClients()
	PrintInfoRows()
}

type consoleScore struct {
	manager InfoManager
}

func NewConsoleScore(manager InfoManager) *consoleScore {
	return &consoleScore{
		manager: manager,
	}
}

func (c *consoleScore) PrintInfoClients() {
	clients := c.manager.InfoClient()
	fmt.Printf("Информация о клиентах:\n")
	for i, client := range clients {
		fmt.Printf("%d. %s\n", i+1, client.Info())
	}
}

func (c *consoleScore) PrintInfoRows() {
	rows := c.manager.InfoRow()
	fmt.Printf("Информация о дорожках:\n")
	for i, row := range rows {
		fmt.Printf("%d. %s\n", i+1, row.Info())
	}
}