package domain

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type ClientId int

type Client interface {
	Play(row Row) error
	Id() ClientId
	PlayTime() time.Duration
	Info() string
	Wait(ctx context.Context) (context.Context, context.CancelFunc)
}

type clientImpl struct {
	id       ClientId
	playTime time.Duration
	score    int
	mutex    sync.Mutex
	waitTime time.Duration
}

func NewClient(
	id ClientId,
	playTime time.Duration,
	waitTime time.Duration,
) *clientImpl {
	client := clientImpl{
		id:       id,
		playTime: playTime,
		waitTime: waitTime,
	}

	return &client
}

func (c *clientImpl) PlayTime() time.Duration {
	return c.playTime
}

func (c *clientImpl) Id() ClientId {
	return c.id
}

func (c *clientImpl) Info() string {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	return fmt.Sprintf("Информация о клиенте:\nid: %d\nОчки: %d", c.id, c.score)
}

func (c *clientImpl) Play(row Row) error {
	ok := row.Acquire()
	if !ok {
		return fmt.Errorf("Дорожка уже занята")
	}
	for range c.playTime / time.Second {
		c.mutex.Lock()
		c.score += rand.Int() % 4
		c.mutex.Unlock()
		time.Sleep(time.Second)
	}
	row.Release()

	return nil
}

func (c *clientImpl) Wait(ctx context.Context) (context.Context, context.CancelFunc){
	return context.WithTimeout(ctx, c.waitTime)
}