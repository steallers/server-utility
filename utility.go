package server_utility

var BuildType = "dev" //development and production
var LoggerType = "text"
var ServiceName = "not defined"

const LoggerTypeText = "text"
const LoggerTypeJson = "json"

func SetLoggerTypeToJson() {
	LoggerType = LoggerTypeJson
}

func SetLoggerTypeText() {
	LoggerType = LoggerTypeText
}

func SetUtilityServiceName(serviceName string) {
	ServiceName = serviceName
}
