package onesignal

import "net/http"

// OneSignal constants
const (
	// APIEndpoint is the endpoint for all API methods,
	// with formatting for Sprintf.
	APIEndpoint = "https://onesignal.com/api/v1/%s"
)

type CreateNotificationConfig struct {
	// AppID is your OneSignal Application ID, which can be found in Keys & IDs.
	AppID string `json:"app_id"`
	// Filters work the same as segments without needing to create the segment first.
	// Filters can be combined to form advanced, highly precise user targeting. The
	// 'filters' parameter targets notification recipients using an array of JSON
	// objects containing field conditions to check
	Filters Filters `json:"filters,omitempty"`
	// SpecificDevices can target specific devices. Targeting devices is typically used in
	// two ways:
	// 1) For transactional notifications that target individual users.
	// 2) For apps that wish to manage their own segments, such as tracking a user's
	// followers and sending notifications to them when that user posts.
	SpecificDevices
	// ContentAndLanguage can help to send push notifications, used the following
	// parameters
	// Read more: https://documentation.onesignal.com/docs/language-localization#supported-languages
	ContentAndLanguage
	// Attachments used for Push notifications only
	//
	// These are additional content attached to push notifications, primarily images.
	Attachments
	// ActionButtons used for Push notifications only
	//
	// These add buttons to push notifications, allowing the user to take more than
	// one action on a notification.
	// Read more: https://documentation.onesignal.com/docs/action-buttons
	ActionButtons
	// Delivery will help you send notifications at a certain point in time in the
	// future
	Delivery
	// GroupingAndCollapsing used for Push notifications only
	//
	// Grouping lets you combine multiple notifications into a single notification to
	// improve the user experience. Collapsing lets you dismiss old notifications in
	// favor of newer ones.
	GroupingAndCollapsing
}

func (c *CreateNotificationConfig) params() (Params, error) {
	params := NewParamsWithMethod(c.method())

	params.AddJSONPayload(c)

	return params, nil
}

func (c *CreateNotificationConfig) method() string {
	return http.MethodPost
}

func (c *CreateNotificationConfig) endpoint() string {
	return "notifications"
}

type CancelNotificationConfig struct {
	AppID          string `json:"app_id"`
	NotificationID string `json:"id"`
}

func (c *CancelNotificationConfig) params() (Params, error) {
	params := NewParamsWithMethod(c.method())

	return params, nil
}

func (c *CancelNotificationConfig) method() string {
	return http.MethodDelete
}

func (c *CancelNotificationConfig) endpoint() string {
	return "notifications/" + c.NotificationID
}

// CreateAppConfig can create a new OneSignal app
type CreateAppConfig struct {
	// Name is the name of your new app, as displayed on your apps list on the
	// dashboard. This can be renamed later.
	Name string `json:"name"`
	// ApnsEnv is Either sandbox or production
	//
	// iOS optional
	ApnsEnv string `json:"apns_env"`
	// ApnsP12 is Your apple push notification p12 certificate file, converted to a
	// string and Base64 encoded.
	//
	// iOS optional
	ApnsP12 string `json:"apns_p12"`
	// ApnsP12Password is Password for the apns_p12 file
	//
	// iOS required if adding p12 certificate
	ApnsP12Password string `json:"apns_p12_password"`
	// GcmKey is Your FCM Google Push Server Auth Key
	//
	// Android optional
	GcmKey string `json:"gcm_key"`
	// AndroidGGcmSenderID is Your FCM Google Project number. Also known as Sender ID.
	//
	// Android optional
	AndroidGGcmSenderID string `json:"android_g_gcm_sender_id"`
	// ChromeWebOrigin is The URL to your website. This field is required if you wish
	// to enable web push and specify other web push parameters.
	//
	// Recommended for Chrome (All Browsers except Safari)
	ChromeWebOrigin string `json:"chrome_web_origin"`
	// ChromeWebDefaultNotificationIcon is Your default notification icon. Should be
	// 256x256 pixels, min 80x80.
	//
	// Chrome (All Browsers except Safari) optional
	ChromeWebDefaultNotificationIcon string `json:"chrome_web_default_notification_icon"`
	// ChromeWebSubDomain is A subdomain of your choice in order to support Web Push
	// on non-HTTPS websites. This field must be set in order for the
	// chrome_web_gcm_sender_id property to be processed.
	//
	// Chrome (All Browsers except Safari) optional
	ChromeWebSubDomain string `json:"chrome_web_sub_domain"`
	// SiteName is The Site name. Requires both chrome_web_origin and
	// safari_site_origin to be set to add or update it.
	//
	// Recommended for All Browsers
	SiteName string `json:"site_name"`
	// SafariSiteOrigin is The hostname to your website including http(s)://
	//
	// Recommended for Safari
	SafariSiteOrigin string `json:"safari_site_origin"`
	// SafariApnsP12 is Your apple push notification p12 certificate file for Safari
	// Push Notifications, converted to a string and Base64 encoded.
	//
	// Safari optional
	SafariApnsP12 string `json:"safari_apns_p_12"`
	// SafariApnsP12Password is Password for safari_apns_p12 file
	//
	// Safari optional
	SafariApnsP12Password string `json:"safari_apns_p_12_password"`
	// SafariIcon256256 is A url for a 256x256 png notification icon. This is the
	// only Safari icon URL you need to provide.
	//
	// Safari optional
	SafariIcon256256 string `json:"safari_icon_256256"`
	// ChromeKey is Not for web push
	// Your Google Push Messaging Auth Key if you use Chrome Apps / Extensions.
	ChromeKey string `json:"chrome_key"`
	// AdditionalDataIsRootPayload is Notification data (additional data) values will
	// be added to the root of the apns payload when sent to the device.
	// Ignore if you're not using any other plugins or not using OneSignal SDK
	// methods to read the payload.
	//
	// iOS optional
	AdditionalDataIsRootPayload bool `json:"additional_data_is_root_payload"`
	// OrganizationId is The ID of the Organization you would like to add this app to.
	OrganizationId string `json:"organization_id"`
}

func (c *CreateAppConfig) params() (Params, error) {
	params := NewParamsWithMethod(c.method())

	params.AddJSONPayload(c)

	return params, nil
}

func (c *CreateAppConfig) method() string {
	return http.MethodPost
}

func (c *CreateAppConfig) endpoint() string {
	return "apps"
}

// ViewDevicesConfig can view the details of multiple devices in one of your OneSignal apps.
// !Unavailable for Apps Over 80,000 Users!
type ViewDevicesConfig struct {
	// AppID is the app ID that you want to view devices from.
	AppID string `json:"app_id"`
	// Limit how many devices to return. Max is 300. Default is 300.
	Limit int `json:"limit"`
	// Offset is result offset. Default is 0. Results are sorted by id.
	Offset int `json:"offset"`
}

func (c *ViewDevicesConfig) params() (Params, error) {
	params := NewParamsWithMethod(c.method())

	params.AddURLNonZero("limit", c.Limit)
	params.AddURLNonZero("offset", c.Offset)

	return params, nil
}

func (c *ViewDevicesConfig) method() string {
	return http.MethodGet
}

func (c *ViewDevicesConfig) endpoint() string {
	return "players"
}

// ViewDeviceConfig can view the details of an existing device in one of your OneSignal apps.
type ViewDeviceConfig struct {
	// ID is OneSignal Player ID.
	ID string `json:"id"`
	// AppID is the AppId that contains the Player ID.
	AppID string `json:"app_id"`
	// EmailAuthHash is only required if you have enabled Identity Verification
	// (https://documentation.onesignal.com/docs/identity-verification) and
	// device_type is email (11).
	//
	// optional
	EmailAuthHash string `json:"email_auth_hash"`
}

func (c *ViewDeviceConfig) params() (Params, error) {
	params := NewParamsWithMethod(c.method())

	return params, nil
}

func (c *ViewDeviceConfig) method() string {
	return http.MethodGet
}

func (c *ViewDeviceConfig) endpoint() string {
	return "players/" + c.ID
}

// AddDeviceConfig is used to register a new device with OneSignal.
//
// If a device is already registered with the specified identifier,
// then this will update the existing device record instead of creating a new one.
//
// The returned ID is for the OneSignal Device Channel Record. If you set
// device_type = 11 the returned ID is for the Email Channel associated with the
// device. device_type = 14 is the SMS Channel. All other device_type correspond
// to the Push Channel. It is recommended to include an external_user_id to
// associate all Device Channel Records with your own User ID.
type AddDeviceConfig struct {
	// AppID Required Your OneSignal AppID found in Keys & IDs (https://documentation.onesignal.com/docs/accounts-and-keys).
	AppID string `json:"app_id"`
	// DeviceType is device's platform
	DeviceType DeviceType `json:"device_type"`
	// For Push Notifications, this is the Push Token Identifier from Google or Apple.
	// For Apple Push identifiers, you must strip all non-alphanumeric characters.
	//
	// recommended
	Identifier string `json:"identifier,omitempty"`
	// Only required if you have enabled Identity Verification and device_type is 11 (Email) or 14 SMS (coming soon).
	//
	// optional
	IdentifierAuthHash string `json:"identifier_auth_hash,omitempty"`
	// This is used in deciding whether to use your iOS Sandbox or Production push
	// certificate when sending a push when both have been uploaded. Set to the iOS
	// provisioning profile that was used to build your app.
	// 1 = Development
	// 2 = Ad-Hoc
	// Omit this field for App Store builds.
	//
	// optional
	TestType int `json:"test_type,omitempty"`
	// Language code. Typically, lower case two letters, except for Chinese where it
	// must be one of zh-Hans or zh-Hant. Example: en
	//
	// recommended
	Language string `json:"language,omitempty"`
	// Number of seconds away from UTC. Example: -28800
	//
	// recommended
	Timezone int `json:"timezone,omitempty"`
	// Version of your app. Example: 1.1
	//
	// recommended
	GameVersion string `json:"game_version,omitempty"`
	// Device make and model. Example: iPhone5,1
	//
	// recommended
	DeviceModel string `json:"device_model,omitempty"`
	// Device operating system version. Example: 7.0.4
	//
	// recommended
	DeviceOS string `json:"device_os,omitempty"`
	// The ad id for the device's platform:
	// Android = Advertising ID
	// iOS = identifierForVendor
	// WP8.1 = AdvertisingId
	//
	// optional
	AdID string `json:"ad_id,omitempty"`
	// Name and version of the plugin that's calling this API method (if any)
	//
	// recommended
	SDK string `json:"sdk,omitempty"`
	// Number of times the user has played the game, defaults to 1
	//
	// optional
	SessionCount int `json:"session_count,omitempty"`
	// Custom tags for the player. Only support string key value pairs.
	// Does not support arrays or other nested objects.
	// Example: {"foo":"bar","this":"that"}
	//
	// optional
	Tags map[string]string `json:"tags,omitempty"`
	// Amount the user has spent in USD, up to two decimal places
	//
	// optional
	AmountSpent string `json:"amount_spent,omitempty"`
	// Unix timestamp in seconds indicating date and time when the device downloaded
	// the app or subscribed to the website.
	//
	// optional
	CreatedAt int `json:"created_at,omitempty"`
	// Seconds player was running your app.
	//
	// optional
	Playtime int `json:"playtime,omitempty"`
	// Unix timestamp in seconds indicating date and time when the device last used the app or website.
	//
	// optional
	LastActive int `json:"last_active,omitempty"`
	// iOS - These values are set each time the user opens the app from the SDK. Use
	// the SDK function set Subscription instead.
	//
	// Android - You may set this, but you can no longer use the SDK method
	// setSubscription later in your app as it will create synchronization issues.
	//
	// optional
	NotificationTypes NotificationTypes `json:"notification_types,omitempty"`
	// Longitude of the device, used for geotagging to segment on.
	//
	// optional
	Long float64 `json:"long,omitempty"`
	// Latitude of the device, used for geotagging to segment on.
	//
	// optional
	Lat float64 `json:"lat,omitempty"`
	// Country code in the ISO 3166-1 Alpha 2 format
	//
	// optional
	Country string `json:"country,omitempty"`
	// A custom user ID
	//
	// optional
	ExternalUserID string `json:"external_user_id,omitempty"`
	// Only required if you have enabled Identity Verification (https://documentation.onesignal.com/docs/identity-verification).
	//
	// optional
	ExternalUserIDAuthHash string `json:"external_user_id_auth_hash,omitempty"`
}

type DeviceType int

const (
	IOS DeviceType = iota
	Android
	Amazon
	WindowsPhone
	ChromeAppsAndExtensions
	ChromeWebPush
	WindowsWNS
	Safari
	Firefox
	MacOS
	Alexa
	Email
	_
	HuaweiApp
	SMS
)

type NotificationTypes int

const (
	Subscribed   NotificationTypes = 1
	Unsubscribed NotificationTypes = -2
)

func (c *AddDeviceConfig) params() (Params, error) {
	params := NewParamsWithMethod(c.method())

	params.AddJSONPayload(c)

	return params, nil
}

func (c *AddDeviceConfig) method() string {
	return http.MethodPost
}

func (c *AddDeviceConfig) endpoint() string {
	return "players"
}

type EditDeviceConfig struct {
	// ID is the device's OneSignal ID
	ID string `json:"id"`
	// AppID is your OneSignal AppId found in Keys & IDs
	AppID string `json:"app_id"`
	// ExternalUserID is a custom user ID
	ExternalUserID string `json:"external_user_id"`
}

func (c *EditDeviceConfig) params() (Params, error) {
	params := NewParamsWithMethod(c.method())

	params.AddJSONPayload(c)

	return params, nil
}

func (c *EditDeviceConfig) method() string {
	return http.MethodPut
}

func (c *EditDeviceConfig) endpoint() string {
	return "players/" + c.ID
}
