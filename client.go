package onesignal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// HTTPClient is the type needed for the bot to perform HTTP requests.
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type OneSignalAPI struct {
	userAuthKey string

	appID      string
	restAPIKey string

	client          HTTPClient
	shutdownChannel chan interface{}

	apiEndpoint string
}

func NewOneSignalClient(userAuthKey, appID, restAPIKey string) (*OneSignalAPI, error) {
	return NewOneSignalClientWithClient(userAuthKey, appID, restAPIKey, APIEndpoint, &http.Client{})
}

func NewOneSignalClientWithClient(userAuthKey, appID, restAPIKey, apiEndpoint string, client HTTPClient) (*OneSignalAPI, error) {
	oneSignalClient := &OneSignalAPI{
		userAuthKey: userAuthKey,

		appID:      appID,
		restAPIKey: restAPIKey,

		client:          client,
		shutdownChannel: make(chan interface{}),

		apiEndpoint: apiEndpoint,
	}

	//_, err := oneSignalClient.ViewApp()

	return oneSignalClient, nil
}

// SetAPIEndpoint changes the OneSignal API endpoint used by the instance.
func (c *OneSignalAPI) SetAPIEndpoint(apiEndpoint string) *OneSignalAPI {
	c.apiEndpoint = apiEndpoint
	return c
}

// MakeRequest makes a request to a specific endpoint with our apiKey.
func (c *OneSignalAPI) MakeRequest(endpoint string, params Params) ([]byte, error) {
	endpointURL := fmt.Sprintf(c.apiEndpoint, endpoint)

	req, err := http.NewRequest(params.Method(), endpointURL, bytes.NewBuffer(params.GetPayload()))
	if err != nil {
		return nil, err
	}
	c.wrapRequest(req, params)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, c.decodeApiError(data)
	}

	return data, nil
}

func (c OneSignalAPI) wrapRequest(req *http.Request, params Params) {
	values := url.Values{}
	for key, param := range params {
		if param.URLValue == "" {
			continue
		}

		values.Set(key, param.URLValue)
	}

	req.URL.RawQuery = values.Encode()

	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Authorization", "Basic "+c.getAuthorizationKey(params))
}

func (c *OneSignalAPI) getAuthorizationKey(params Params) string {
	switch params.KeyType() {
	case typeAuth:
		return c.userAuthKey
	case typeRestAPI:
		return c.restAPIKey
	}
	return ""
}

func (c *OneSignalAPI) decodeApiError(data []byte) error {
	var apiErr ApiError

	fmt.Println(string(data))
	err := json.Unmarshal(data, &apiErr)
	if err != nil {
		return err
	}

	return apiErr
}

func (c *OneSignalAPI) CreateNotification(config *CreateNotificationConfig) (*CreateNotificationResult, error) {
	params, _ := config.params()
	params.AddURLNonEmpty("app_id", c.appID)

	resp, err := c.MakeRequest(config.endpoint(), params.SetKeyType(typeRestAPI))
	if err != nil {
		return nil, err
	}

	var result CreateNotificationResult
	err = json.Unmarshal(resp, &result)

	return &result, err
}

func (c *OneSignalAPI) CancelNotification(config *CancelNotificationConfig) (*CancelNotificationResult, error) {
	params, _ := config.params()
	params.AddURLNonEmpty("app_id", c.appID)

	resp, err := c.MakeRequest(config.endpoint(), params.SetKeyType(typeRestAPI))
	if err != nil {
		return nil, err
	}

	var result CancelNotificationResult
	err = json.Unmarshal(resp, &result)

	return &result, err
}

func (c *OneSignalAPI) ViewApps() (*Apps, error) {
	resp, err := c.MakeRequest("apps", NewParamsWithMethod(http.MethodGet).SetKeyType(typeAuth))
	if err != nil {
		return nil, err
	}

	var apps Apps
	err = json.Unmarshal(resp, &apps)

	return &apps, err
}

func (c *OneSignalAPI) ViewApp() (*App, error) {
	resp, err := c.MakeRequest("apps/"+c.appID, NewParamsWithMethod(http.MethodGet).SetKeyType(typeRestAPI))
	if err != nil {
		return nil, err
	}

	var apps App
	err = json.Unmarshal(resp, &apps)

	return &apps, err
}

func (c *OneSignalAPI) CreateApp(config *CreateAppConfig) (*App, error) {
	params, _ := config.params()
	params.AddURLNonEmpty("app_id", c.appID)

	resp, err := c.MakeRequest(config.endpoint(), params.SetKeyType(typeAuth))
	if err != nil {
		return nil, err
	}

	var app App
	err = json.Unmarshal(resp, &app)

	return &app, err
}

// UpdateApp is now Disable
func (c *OneSignalAPI) UpdateApp() {
	panic("implement me")
}

func (c *OneSignalAPI) ViewDevices(config *ViewDevicesConfig) (*Devices, error) {
	params, _ := config.params()
	params.AddURLNonEmpty("app_id", c.appID)

	resp, err := c.MakeRequest(config.endpoint(), params.SetKeyType(typeRestAPI))
	if err != nil {
		return nil, err
	}

	var devices Devices
	err = json.Unmarshal(resp, &devices)

	return &devices, err
}

func (c *OneSignalAPI) ViewDevice(config *ViewDeviceConfig) (*Device, error) {
	params, _ := config.params()
	params.AddURLNonEmpty("app_id", c.appID)

	resp, err := c.MakeRequest(config.endpoint(), params.SetKeyType(typeRestAPI))
	if err != nil {
		return nil, err
	}

	var device Device
	err = json.Unmarshal(resp, &device)

	return &device, err
}

func (c *OneSignalAPI) AddDevice(config *AddDeviceConfig) (*AddDeviceResult, error) {
	config.AppID = c.appID
	params, _ := config.params()

	resp, err := c.MakeRequest(config.endpoint(), params.SetKeyType(typeRestAPI))
	if err != nil {
		return nil, err
	}

	var result AddDeviceResult
	err = json.Unmarshal(resp, &result)

	return &result, err
}

func (c *OneSignalAPI) EditDevice(config *EditDeviceConfig) (*EditDeviceResult, error) {
	config.AppID = c.appID
	params, _ := config.params()

	resp, err := c.MakeRequest(config.endpoint(), params.SetKeyType(typeRestAPI))
	if err != nil {
		return nil, err
	}

	var result EditDeviceResult
	err = json.Unmarshal(resp, &result)

	return &result, err
}

// EditTagsWithExternalUserID is now Disable
func (c *OneSignalAPI) EditTagsWithExternalUserID() {
	panic("implement me")
}

// CSVExport is now Disable
func (c *OneSignalAPI) CSVExport() {
	panic("implement me")
}

// ViewNotification is now Disable
func (c *OneSignalAPI) ViewNotification() {
	panic("implement me")
}

// ViewNotifications is now Disable
func (c *OneSignalAPI) ViewNotifications() {
	panic("implement me")
}

// NotificationHistory is now Disable
func (c *OneSignalAPI) NotificationHistory() {
	panic("implement me")
}

// CreateSegment is now Disable
func (c *OneSignalAPI) CreateSegment() {
	panic("implement me")
}

// DeleteSegment is now Disable
func (c *OneSignalAPI) DeleteSegment() {
	panic("implement me")
}

// ViewOutcomes is now Disable
func (c *OneSignalAPI) ViewOutcomes() {
	panic("implement me")
}

// DeleteUserRecord is now Disable
func (c *OneSignalAPI) DeleteUserRecord() {
	panic("implement me")
}
