package main

import (
	"bowling/domain"
	"fmt"
)

const RowsAmount int = 4

func main() {
	manager := domain.NewManager()
	for range RowsAmount {
		manager.NewRow(domain.NewRow())
	}
	score := domain.NewConsoleScore(manager)
	console := domain.NewConsole(manager, score)
	console.Dialog()
	fmt.Printf("День завершён")
}
