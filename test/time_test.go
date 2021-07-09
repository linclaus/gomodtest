package test

import (
	"testing"
	"time"

	"github.com/sirupsen/logrus"
)

func TestTime(t *testing.T) {
	start := time.Now().Second()
	time.Sleep(5 * time.Second)
	end := time.Now().Second()
	duration := end - start
	logrus.Infof("duration: %d", duration)
}
