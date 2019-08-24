![alt text](https://github.com/verseatile/conway/blob/master/conway.png?raw=true "Conway")

# conway
a small finite state machine written in Go

>I mean the plan is simple...store some state, swap them out and do whatever with them. Can only go up from here

### Goals
As of now, primary goals are the following:
* a state machine flexible enough for several scenarios (parsers, games, etc.)
* flexibility, ready to be dropped into any Go project

### Install
Use go get.

```bash
go get github.com/verseatile/conway
```

```go
// create a new state machine
machine := fsm.NewMachine()
// create a state instance, a default one is already initialized if not
s := &fsm.State{
    State: make(map[string]interface{}, 0)}

// Set the state machine's state to the one you just created
machine.SetCurrent(s)

// set state property
machine.SetState("hello", "world")

fmt.Println(machine.GetState("hello"))

```
