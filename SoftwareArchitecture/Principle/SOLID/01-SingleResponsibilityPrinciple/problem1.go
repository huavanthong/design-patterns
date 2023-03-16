package main

type Command struct {
	commandType string
	args        []string
}

func (c Command) Decode(data []byte) {
	// decodes and initialise
	panic("implement me")
}

func (c Command) ValidateCommandType() bool {
	// validates command type
	panic("implement me")
}

func (c Command) ValidateArgs() bool {
	// validate provided args as if input
	panic("implement me")
}

func (c Command) Execute() {
	// Executes seperate types of commands
	panic("implement me")
}
