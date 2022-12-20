package service

import "project_api/broadcast"

type Manager interface {
	OpenListener() chan interface{}
	CloseListener(channel chan interface{})
	Submit(msg string)
	DeleteBroadcast()
}

type Listener struct {
	Chan chan interface{}
}

type manager struct {
	open  chan *Listener
	close chan *Listener

	messages chan string

	delete chan string

	roomChannels map[string]broadcast.Broadcaster
}

var managerSingleton *manager

func NewRoomManager() Manager {
	if managerSingleton == nil {
		managerSingleton = &manager{
			roomChannels: make(map[string]broadcast.Broadcaster),
			open:         make(chan *Listener, 100),
			close:        make(chan *Listener, 100),
			delete:       make(chan string, 100),
			messages:     make(chan string, 100),
		}

		go managerSingleton.run()
	}

	return managerSingleton
}

func (m *manager) run() {
	currentRoom := m.room()
	for {
		select {
		case listener := <-m.open:
			currentRoom.Register(listener.Chan)
		case listener := <-m.close:
			currentRoom.Unregister(listener.Chan)
		case msg := <-m.messages:
			currentRoom.Submit(msg)
		case room := <-m.delete:
			b, ok := m.roomChannels[room]
			if ok {
				b.Close()
				delete(m.roomChannels, room)
			}
		}
	}
}

func (m *manager) room() broadcast.Broadcaster {
	b, ok := m.roomChannels["payments"]

	if !ok {
		b = broadcast.NewBroadcaster(10)
		m.roomChannels["payments"] = b
	}
	return b
}

func (m *manager) OpenListener() chan interface{} {
	listener := make(chan interface{})
	m.open <- &Listener{
		Chan: listener,
	}

	return listener
}

func (m *manager) CloseListener(channel chan interface{}) {
	m.close <- &Listener{
		Chan: channel,
	}
}

func (m *manager) Submit(text string) {
	m.messages <- text
}

func (m *manager) DeleteBroadcast() {
	m.delete <- "payments"
}
