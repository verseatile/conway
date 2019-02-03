package fsm

// Machine - state machine
type Machine struct {
	current *State
	events  *Events
}

// add name property to describe state?
type State struct {
	name  string
	State map[string]interface{}
}

type Events struct {
	bus       map[string]chan string
	callbacks map[string][]func(string)
}

/*
 *
 *	MACHINE METHODS
 *
 */

// NewMachine - returns new state machine
func NewMachine() *Machine {
	// return Machine{ state: make(map[string]interface{}, 0) }
	return &Machine{
		current: &State{State: make(map[string]interface{}, 0)},
		events:  &Events{bus: make(map[string]chan string, 0), callbacks: make(map[string][]func(string))}}
}

func (m *Machine) SetCurrent(s *State) {
	m.current = s
}

func (m *Machine) SetState(prop string, value interface{}) {
	m.current.State[prop] = value
}

func (m *Machine) GetState(prop string) interface{} {
	return m.current.State[prop]
}

// GetCurrent - returns current machine state
func (m *Machine) GetCurrent() *State {
	return m.current
}

/*
 *
 *	EVENT METHODS
 *
 */
type EventCallback func(string)

func (m *Machine) GetCallbacks(evtName string) []func(string) {
	return m.events.callbacks[evtName]
}

// On - also knwon as add callback
func (m *Machine) On(evtName string, cb EventCallback) {
	m.events.callbacks[evtName] = append(m.events.callbacks[evtName], cb)
}

// Emit - fires an event.
func (m *Machine) EmitEvent(evtName string, buff string) chan string {
	if m.events.bus[evtName] == nil {
		m.events.bus[evtName] = make(chan string)
	}

	// lifecycle methods can go within here
	go func() {
		for {
			select {
			case data := <-m.events.bus[evtName]:
				// if strings.HasPrefix(data, "alt") {
				// 	fmt.Println("alternate path selected. flourish family")
				// 	return
				// }
				// conditional if needed/can pass in
				// remove println with real behavior
				// fmt.Println(evtName, "has been fired.", data)
				go func() {
					for _, cb := range m.events.callbacks[evtName] {
						// go cb()
						cb(data)
					}
				}()
			}
		}
	}()

	// send data - something has to get send over the channel or it does nothing
	m.events.bus[evtName] <- buff

	// callback - maxed at 10 for now
	if m.events.callbacks[evtName] == nil {
		m.events.callbacks[evtName] = make([]func(string), 10)
	}

	// should probably go inside select
	// go func() {
	// 	for _, cb := range m.events.callbacks[evtName] {
	// 		// go cb()
	// 		cb(buff)
	// 	}
	// }()

	// return the channel to allow for closing it in the future
	return m.events.bus[evtName]
}
