package domain

import (
	"fmt"
	"time"
)

type Console interface {
	start()
}

type consoleImpl struct {
	manager Manager
	score   Score
}

func NewConsole(manager Manager, score Score) *consoleImpl {
	return &consoleImpl{
		manager: manager,
		score:   score,
	}
}

func (c *consoleImpl) addClient() {
	var id int
	fmt.Printf("Введите id пользователя:\n")
	_, err := fmt.Scanf("%d\n", &id)
	if err != nil {
		fmt.Printf("Ошибка: %v", err)
	}
	fmt.Printf("Введите количество секунд игры:\n")
	var playSeconds int
	_, err = fmt.Scanf("%d\n", &playSeconds)
	if err != nil {
		fmt.Printf("Ошибка: %v", err)
	}
	fmt.Printf("Введите количество секунд ожидания:\n")
	var waitSeconds int
	_, err = fmt.Scanf("%d\n", &waitSeconds)
	if err != nil {
		fmt.Printf("Ошибка: %v", err)
	}

	playTime := time.Second * time.Duration(playSeconds)
	waitTime := time.Second * time.Duration(waitSeconds)

	client := NewClient(ClientId(id), playTime, waitTime )
	c.manager.NewClient(client)

}

func (c *consoleImpl) start(sleepTime time.Duration) {
	doneChan := c.manager.Start()
	for {
		c.score.PrintInfoClients()
		c.score.PrintInfoRows()
		time.Sleep(sleepTime)
		select{
		case <- doneChan:
			return
		default:
			
		}
		
	}

}

func (c *consoleImpl) Dialog() {
	for {
		var i int
		fmt.Printf("1 - запуск программы, 2 - создать клиента\n")
		_, err := fmt.Scanf("%d\n", &i)
		if err != nil {
			fmt.Printf("Ошибка ввода %v", err)
		}
		switch i {
		case 1:
			c.start(time.Second)
			return
		case 2:
			c.addClient()
		}

	}
}
