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

type Client struct {
	userAuthKey string

	appID      string
	restAPIKey string

	client          HTTPClient
	shutdownChannel chan interface{}

	apiEndpoint string
}

func NewClient(userAuthKey, appID, restAPIKey string) (*Client, error) {
	return NewCustomClient(userAuthKey, appID, restAPIKey, APIEndpoint, &http.Client{})
}

func NewCustomClient(userAuthKey, appID, restAPIKey, apiEndpoint string, httpClient HTTPClient) (*Client, error) {
	client := &Client{
		userAuthKey: userAuthKey,

		appID:      appID,
		restAPIKey: restAPIKey,

		client:          httpClient,
		shutdownChannel: make(chan interface{}),

		apiEndpoint: apiEndpoint,
	}

	//_, err := oneSignalClient.ViewApp()

	return client, nil
}

// SetAPIEndpoint changes the OneSignal API endpoint used by the instance.
func (c *Client) SetAPIEndpoint(apiEndpoint string) *Client {
	c.apiEndpoint = apiEndpoint
	return c
}

// MakeRequest makes a request to a specific endpoint with our apiKey.
func (c *Client) MakeRequest(endpoint string, params Params) ([]byte, error) {
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

func (c Client) wrapRequest(req *http.Request, params Params) {
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

func (c *Client) getAuthorizationKey(params Params) string {
	switch params.KeyType() {
	case typeAuth:
		return c.userAuthKey
	case typeRestAPI:
		return c.restAPIKey
	}
	return ""
}

func (c *Client) decodeApiError(data []byte) error {
	var apiErr ApiError

	err := json.Unmarshal(data, &apiErr)
	if err != nil {
		return err
	}

	return apiErr
}

func (c *Client) CreateNotification(config *CreateNotificationConfig) (*CreateNotificationResult, error) {
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

func (c *Client) CancelNotification(config *CancelNotificationConfig) (*CancelNotificationResult, error) {
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

func (c *Client) ViewApps() ([]App, error) {
	resp, err := c.MakeRequest("apps", NewParamsWithMethod(http.MethodGet).SetKeyType(typeAuth))
	if err != nil {
		return nil, err
	}

	var apps []App
	err = json.Unmarshal(resp, &apps)

	return apps, err
}

func (c *Client) ViewApp() (*App, error) {
	resp, err := c.MakeRequest("apps/"+c.appID, NewParamsWithMethod(http.MethodGet).SetKeyType(typeRestAPI))
	if err != nil {
		return nil, err
	}

	var apps App
	err = json.Unmarshal(resp, &apps)

	return &apps, err
}

func (c *Client) CreateApp(config *CreateAppConfig) (*App, error) {
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
func (c *Client) UpdateApp() {
	panic("implement me")
}

func (c *Client) ViewDevices(config *ViewDevicesConfig) (*Devices, error) {
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

func (c *Client) ViewDevice(config *ViewDeviceConfig) (*Device, error) {
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

func (c *Client) AddDevice(config *AddDeviceConfig) (*AddDeviceResult, error) {
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

func (c *Client) EditDevice(config *EditDeviceConfig) (*EditDeviceResult, error) {
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
func (c *Client) EditTagsWithExternalUserID() {
	panic("implement me")
}

// CSVExport is now Disable
func (c *Client) CSVExport() {
	panic("implement me")
}

// ViewNotification is now Disable
func (c *Client) ViewNotification() {
	panic("implement me")
}

// ViewNotifications is now Disable
func (c *Client) ViewNotifications() {
	panic("implement me")
}

// NotificationHistory is now Disable
func (c *Client) NotificationHistory() {
	panic("implement me")
}

// CreateSegment is now Disable
func (c *Client) CreateSegment() {
	panic("implement me")
}

// DeleteSegment is now Disable
func (c *Client) DeleteSegment() {
	panic("implement me")
}

// ViewOutcomes is now Disable
func (c *Client) ViewOutcomes() {
	panic("implement me")
}

// DeleteUserRecord is now Disable
func (c *Client) DeleteUserRecord() {
	panic("implement me")
}
