package server

var (
	defaultPanicMessage = "wunderDB panicked on request"
)

// WIP/To-Do: Add custom startup message.
// Ref: https://github.com/TanmoySG/wunderDB/issues/129
func (ws wdbServer) startupMessage(address, port string) {}
