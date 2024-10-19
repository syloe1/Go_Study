package ziface

// define server interface
type IServer interface {
	//start
	Start()
	//stop
	Stop()
	//start Server
	Serve()
}
