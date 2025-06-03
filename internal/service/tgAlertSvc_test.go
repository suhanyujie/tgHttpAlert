package service

import (
	"fmt"
	"testing"
)

func TestSendMsg(t *testing.T) {
	if err := NewAlertMsg("dev", "local", "nil", "test alert msg...").Send(); err != nil {
		fmt.Printf("%+W \n", err)
	}
}
