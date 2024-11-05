package helper

import (
	"github.com/Takina-Space/backend-go/libraries/logger"
	"github.com/sirupsen/logrus"
	"github.com/ztrue/tracerr"
)

func CatchError() {
	if err := recover(); err != nil {

		logger.SetLogFileAndConsole(logger.LogData{
			Message: "unexpected error",
			CustomFields: logrus.Fields{
				"message": err,
			},
			Level: "ERROR",
		})

		dataErrr := tracerr.Wrap(err.(error))
		tracerr.PrintSourceColor(dataErrr)
	}
}

func ErrorHandler(err error) {
	panic(err)
}
