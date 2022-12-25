package broadcast

type Broadcaster interface {
	Register(chan<- interface{})
	Unregister(chan<- interface{})

	Submit(interface{})
	Close() error
}

type broadcaster struct {
	register chan chan<- interface{}
	unregister chan chan<- interface{}
	input chan interface{}

	outputs map[chan<- interface{}]bool
}

func NewBroadcaster(buflen int) Broadcaster {
	brc := &broadcaster{
		register:   make(chan chan<- interface{}),
		unregister: make(chan chan<- interface{}),
		input:      make(chan interface{}, buflen),
		outputs:    make(map[chan<- interface{}]bool),
	}

	go brc.run()

	return brc
}

func (brc *broadcaster) broadcast(msg interface{}) {
	for ch := range brc.outputs {
		ch <- msg
	}
}

func (brc *broadcaster) run() {
	for {
		select {
		case msg := <-brc.input:
			brc.broadcast(msg)
		case ch, ok := <-brc.register:
			if ok {
				brc.outputs[ch] = true
			} else {
				return
			}
		case ch := <- brc.unregister:
			delete(brc.outputs, ch)
		}
	}
}

func (brc *broadcaster) Register(ch chan<- interface{}) {
	brc.register <- ch
}

func (brc *broadcaster) Unregister(ch chan<- interface{}) {
	brc.unregister <- ch
}

func (brc *broadcaster) Submit(msg interface{}) {
	brc.input <- msg
}

func (brc *broadcaster) Close() error {
	close(brc.register)
	close(brc.unregister)
	return nil
}
