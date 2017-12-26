package flash

// Messagebox map
type Messagebox map[string][]Message

// Flash message
type Flash struct {
	Messages Messagebox
}

// Add message
func (f *Flash) Add(t, msg string) {
	f.Messages[t] = append(f.Messages[t], NewMessage(t, msg))
}

// AddSuccess message
func (f *Flash) AddSuccess(msg interface{}) {
	switch v := msg.(type) {
	case string:
		f.Add("success", v)
	case []string:
		for _, m := range v {
			f.Add("success", m)
		}
	}
}

// AddDanger message
func (f *Flash) AddDanger(msg interface{}) {
	switch v := msg.(type) {
	case string:
		f.Add("danger", v)
	case []string:
		for _, m := range v {
			f.Add("danger", m)
		}
	}
}

// AddInfo message
func (f *Flash) AddInfo(msg interface{}) {
	switch v := msg.(type) {
	case string:
		f.Add("info", v)
	case []string:
		for _, m := range v {
			f.Add("info", m)
		}
	}
}

// AddWarning message
func (f *Flash) AddWarning(msg interface{}) {
	switch v := msg.(type) {
	case string:
		f.Add("warning", v)
	case []string:
		for _, m := range v {
			f.Add("warning", m)
		}
	}
}

// GetSuccess Messages
func (f *Flash) GetSuccess() []Message {
	return f.Messages["success"]
}

// GetDanger Messages
func (f *Flash) GetDanger() []Message {
	return f.Messages["danger"]
}

// GetInfo Messages
func (f *Flash) GetInfo() []Message {
	return f.Messages["info"]
}

// GetWarning Messages
func (f *Flash) GetWarning() []Message {
	return f.Messages["warning"]
}

// GetAll Messages
func (f *Flash) GetAll() Messagebox {
	return f.Messages
}

// New flash message manager
func New() *Flash {
	return &Flash{Messages: make(Messagebox)}
}

// Message type
type Message struct {
	Type string
	Data string
}

// NewMessage with type
func NewMessage(t, d string) Message {
	return Message{Type: t, Data: d}
}
