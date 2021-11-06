package server_utility_test

import (
	"fmt"
	serverUtility "github.com/steallers/server-utility"
	"github.com/steallers/server-utility/loggers"
	"testing"
)

func TestLogger(t *testing.T){
	serverUtility.SetLoggerTypeToJson()
	loggers.PrintError(fmt.Errorf("erasd"),"asdas")
	serverUtility.SetLoggerTypeText()
	loggers.PrintError(fmt.Errorf("erasd"),"asdas")
}
