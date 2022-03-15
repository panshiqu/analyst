package cache

import (
	"fmt"
	"testing"
)

func TestHandleAlerts(t *testing.T) {
	fmt.Println(HandleAlert(1, "", "", 0, ""))
	PrintAlerts()
	fmt.Println(HandleAlert(2, "", "", 0, ""))
	PrintAlerts()
	fmt.Println(HandleAlert(2, "", "", 0, ""))
	PrintAlerts()
}

func TestNotifyAlerts(t *testing.T) {
	fmt.Println(HandleAlert(1, "", "btc", -2, ""))
	fmt.Println(HandleAlert(1, "", "btc", -4, ""))
	fmt.Println(HandleAlert(1, "", "eth", 6, ""))
	fmt.Println(HandleAlert(1, "", "eth", 8, ""))
	NotifyCurrentAlerts("", "eth")
	//NotifyAlerts("eth", 2)
	PrintAlerts()
}

func TestLoadAlerts(t *testing.T) {
	fmt.Println(LoadAlerts())
	PrintAlerts()
	// HandleAlert(1, "", "", 0, "")
	// PrintAlerts()
	// fmt.Println(SaveAlerts())
}
