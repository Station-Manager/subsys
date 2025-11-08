package subsys

import "github.com/Station-Manager/types"

const (
	errMsgNilService         = "Subsys service is nil."
	errMsgNotInitialized     = "Subsys service is not initialized."
	errMsgAlreadyStarted     = "Subsys service already started."
	errMsgAlreadyStopped     = "Subsys service already stopped."
	errMsgNilConfigService   = "Config service is nil."
	errMsgNilDatabaseService = "Database service is nil."
)

const (
	ServiceName = types.SubSystemServiceName
)
