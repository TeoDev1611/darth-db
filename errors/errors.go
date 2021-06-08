package errors

import (
	"errors"
	"os"

	log "github.com/sirupsen/logrus"
)

/* Set the log options :D */

func init() {
	log.SetFormatter(&log.JSONFormatter{PrettyPrint: true})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.ErrorLevel)
}

/* Generate a Error util */

func NewError(msg string, hight bool) {
	/* Check if the error is hight priority */
	if hight {
		err := errors.New(msg)
		/* Write a log message for hight priority */
		if err != nil {
			log.WithFields(log.Fields{
				"error": err,
				"msg":   "Error Hight level detected",
			}).Fatal("FIX THIS ERROR NOW!")
		}
		/* If not is priority make a simple log :D */
	} else {
		err := errors.New(msg)
		if err != nil {
			log.WithFields(log.Fields{
				"error": err,
				"msg":   "Error not hight priority",
			}).Error("Fix the error but not more fast or yes ?)")
		}
	}
}

/* Check errors util function */
func CheckErrors(err error) {
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
			"msg":   "Error detected read the error field for more information",
		}).Error("New error detected :c if this is a bug report in the repo")
	}
}
