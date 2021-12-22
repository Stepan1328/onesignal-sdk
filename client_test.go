package onesignal

import (
	"encoding/json"
	"testing"

	"github.com/google/uuid"
)

const (
	TestUserAuthKey = "ZDI1NWQzNTgtZTM5Ni00MjFkLWEyOWUtZGU3NGU0MGQwMjlk"

	TestAppID      = "fc07fb28-1c8a-4ec8-b77a-817476e6e158"
	TestRestAPIKey = "OWM2YmIwYWMtZGMwMS00NDI0LTk5MjAtZGQ0ZjVkMDY2MDBi"

	TestUserID = "8729f50f-7cde-429a-a34b-b05e7a973ddd"
)

func newToken() string {
	return uuid.New().String()
}

func getClient() *OneSignalAPI {
	client, err := NewOneSignalClient(TestUserAuthKey, TestAppID, TestRestAPIKey)
	if err != nil {
		panic(err)
	}

	return client
}

var notificationID string

func TestOneSignalAPI_CreateNotification(t *testing.T) {
	client := getClient()

	config := NewCreateNotification().
		AddMessage(
			English("New message", "Something header"),
			Russian("Новое сообщение", "Какой-то заголовок"),
		).
		AddExternalUserIDs(
			"013c1c9e-4b53-4759-90da-6da148b907e0",
		).
		AddData(map[string]interface{}{"abc": 123, "foo": "bar", "event_performed": true, "amount": 12.1}).
		AddAndroidGrouping("en", "Some header")

	result, err := client.CreateNotification(config)
	if err != nil {
		t.Error("view apps error:", err)
		t.Failed()
	}

	notificationID = result.Id

	t.Logf("result: \n%s\n", FormatData(result))
}

func TestOneSignalAPI_CancelNotification(t *testing.T) {
	client := getClient()

	config := NewCancelNotification(notificationID)

	result, err := client.CancelNotification(config)
	if err != nil {
		t.Error("view apps error:", err)
		t.Failed()
	}

	if !result.Success {
		t.Failed()
	}

	t.Logf("result: \n%s\n", FormatData(result))
}

func TestOneSignalAPI_ViewApps(t *testing.T) {
	client := getClient()

	apps, err := client.ViewApps()
	if err != nil {
		t.Error("view apps error:", err)
		t.Failed()
	}

	t.Logf("apps: \n%s\n", FormatData(apps))
}

func TestOneSignalAPI_ViewApp(t *testing.T) {
	client := getClient()

	app, err := client.ViewApp()
	if err != nil {
		t.Error("view app error:", err)
		t.Failed()
	}

	t.Logf("app: \n%s\n", FormatData(app))
}

//func TestOneSignalAPI_CreatApp(t *testing.T) {
//	client := getClient()
//
//	config := &CreateAppConfig{
//		Name: "Test Create App",
//	}
//
//	app, err := client.CreateApp(config)
//	if err != nil {
//		t.Error("create app:", err)
//		t.Failed()
//	}
//
//	t.Logf("app: \n%s\n", FormatData(app))
//}

func TestOneSignalAPI_UpdateApp(t *testing.T) {
	// TODO
}

func TestOneSignalAPI_ViewDevices(t *testing.T) {
	client := getClient()

	config := &ViewDevicesConfig{
		Limit: 10,
	}

	devices, err := client.ViewDevices(config)
	if err != nil {
		t.Error("view devices error:", err)
		t.Failed()
	}

	t.Logf("devices: \n%s\n", FormatData(devices))
}

func TestOneSignalAPI_ViewDevice(t *testing.T) {
	client := getClient()

	config := &ViewDeviceConfig{
		ID: TestUserID,
	}

	devices, err := client.ViewDevice(config)
	if err != nil {
		t.Error("view devices error:", err)
		t.Failed()
	}

	t.Logf("devices: \n%s\n", FormatData(devices))
}

/*
func TestOneSignalAPI_AddDevice(t *testing.T) {
	client := getClient()

	config := &AddDeviceConfig{
		DeviceType:     IOS,
		ExternalUserID: newToken(),
	}

	result, err := client.AddDevice(config)
	if err != nil {
		t.Error("add device error:", err)
		t.Failed()
	}

	if !result.Success {
		t.Failed()
	}

	t.Logf("device: \n%s\n", FormatData(result))
}
*/

func TestOneSignalAPI_EditDevice(t *testing.T) {
	client := getClient()

	config := &EditDeviceConfig{
		ID:             TestUserID,
		ExternalUserID: newToken(),
	}

	result, err := client.EditDevice(config)
	if err != nil {
		t.Error("edit device error:", err)
		t.Failed()
	}

	if !result.Success {
		t.Failed()
	}

	t.Logf("result: \n%s\n", FormatData(result))
}

func TestOneSignalAPI_EditTagsWithExternalUserID(t *testing.T) {
	// TODO
}

func TestOneSignalAPI_CSVExport(t *testing.T) {
	// TODO
}

func TestOneSignalAPI_ViewNotification(t *testing.T) {
	// TODO
}

func TestOneSignalAPI_ViewNotifications(t *testing.T) {
	// TODO
}

func TestOneSignalAPI_NotificationHistory(t *testing.T) {
	// TODO
}

func TestOneSignalAPI_CreateSegment(t *testing.T) {
	// TODO
}

func TestOneSignalAPI_DeleteSegment(t *testing.T) {
	// TODO
}

func TestOneSignalAPI_ViewOutcomes(t *testing.T) {
	// TODO
}

func TestOneSignalAPI_DeleteUserRecord(t *testing.T) {
	// TODO
}

func FormatData(data interface{}) string {
	result, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		panic("Unexpected error occurred: %s" + err.Error())
	}

	return string(result)
}
