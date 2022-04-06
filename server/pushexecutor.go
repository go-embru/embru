package server

type PushExecutor struct {
	broker map[string]*router
}

func NewPushExecutor() *PushExecutor {
	v := &PushExecutor{
		broker: make(map[string]*router),
	}
	return v
}

func (p *PushExecutor) put(name string, pushable *router) {
	p.broker[name] = pushable
}

func (p *PushExecutor) push(name string, msg string) {
	p.broker[name].push2self(msg)
}
