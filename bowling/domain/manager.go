package domain

import (
	"context"
	"fmt"
	"slices"
	"sync"
)

type managerRowState int

const (
	free managerRowState = iota
	reserved
)

type Manager interface {
	NewClient(client Client)
	NewRow(row Row)
	Start() chan struct{}
}

type InfoManager interface {
	InfoClient() []Client
	InfoRow() []Row
}

type managerImpl struct {
	clients      []Client
	rows         []Row
	rowStates    []managerRowState
	clientsCount int
	clientMutex  sync.Mutex
	rowMutex     sync.Mutex
	clientCond   *sync.Cond
	rowCond      *sync.Cond
	startChan    chan struct{}
}

func NewManager() *managerImpl {
	manager := managerImpl{
		startChan: make(chan struct{}),
	}
	manager.clientCond = sync.NewCond(&manager.clientMutex)
	manager.rowCond = sync.NewCond(&manager.rowMutex)
	return &manager
}

func (m *managerImpl) NewClient(client Client) {
	m.clientMutex.Lock()
	defer m.clientMutex.Unlock()
	m.clients = append(m.clients, client)
	go m.clientRoutine(client)
}

func (m *managerImpl) InfoClient() []Client {
	m.clientMutex.Lock()
	defer m.clientMutex.Unlock()
	result := slices.Clone(m.clients)
	return result
}

func (m *managerImpl) clientRoutine(client Client) {
	<-m.startChan
	m.rowMutex.Lock()
	defer m.rowMutex.Unlock()
	freeRow := rowCheck(m.rows, m.rowStates)
	clientCtx, cancelFunc := client.Wait(context.Background())
	context.AfterFunc(clientCtx, func() {
		m.rowCond.Broadcast()
	})
	for freeRow == len(m.rows) {
		m.rowCond.Wait()
		select {
		case <-clientCtx.Done():
			m.clientsCount++
			return
		default:

		}
		freeRow = rowCheck(m.rows, m.rowStates)
		fmt.Printf("%d\n", freeRow) 
	}
	m.rowStates[freeRow] = reserved

	go func() {
		err := client.Play(m.rows[freeRow])
		if err != nil {
			fmt.Printf("Ошибка %v\n", err)
		}
		m.rowMutex.Lock()
		defer m.rowMutex.Unlock()
		m.rowStates[freeRow] = free
		m.clientsCount++
		m.rowCond.Signal()
		cancelFunc()
	}()
}
func (m *managerImpl) NewRow(row Row) {
	m.rowMutex.Lock()
	defer m.rowMutex.Unlock()
	m.rows = append(m.rows, row)
	m.rowStates = append(m.rowStates, free)
	go m.rowRoutine(row)
}

func (m *managerImpl) InfoRow() []Row {
	m.rowMutex.Lock()
	defer m.rowMutex.Unlock()
	result := slices.Clone(m.rows)
	return result
}

func (m *managerImpl) Start() chan struct{} {
	close(m.startChan)
	doneChan := make(chan struct{})
	go func() {
		m.rowMutex.Lock()
		defer m.rowMutex.Unlock()
		for m.clientsCount < len(m.clients) {
			m.rowCond.Wait()
		}
		close(doneChan)
	}()
	return doneChan
}

func (m *managerImpl) rowRoutine(row Row) {

}
func rowCheck(rows []Row, rowStates []managerRowState) int {
	for i, row := range rows {
		if row.State() == Free && rowStates[i] == free {
			return i
		}
	}
	return len(rows)
}
