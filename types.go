package onesignal

import "time"

type CreateNotificationResult struct {
	Id         string      `json:"id,omitempty"`
	Recipients int         `json:"recipients,omitempty"`
	ExternalId interface{} `json:"external_id,omitempty"`
	Errors     []string    `json:"errors"`
}

type Devices struct {
	TotalCount int      `json:"total_count"`
	Offset     int      `json:"offset"`
	Limit      int      `json:"limit"`
	Players    []Device `json:"players"`
}

type App struct {
	Id                               string    `json:"id"`
	Name                             string    `json:"name"`
	Players                          int       `json:"players,omitempty"`
	MessageablePlayers               int       `json:"messageable_players,omitempty"`
	UpdatedAt                        time.Time `json:"updated_at,omitempty"`
	CreatedAt                        time.Time `json:"created_at,omitempty"`
	GcmKey                           string    `json:"gcm_key,omitempty"`
	ChromeKey                        string    `json:"chrome_key,omitempty"`
	ChromeWebOrigin                  string    `json:"chrome_web_origin,omitempty"`
	ChromeWebGcmSenderId             string    `json:"chrome_web_gcm_sender_id,omitempty"`
	ChromeWebDefaultNotificationIcon string    `json:"chrome_web_default_notification_icon,omitempty"`
	ChromeWebSubDomain               string    `json:"chrome_web_sub_domain,omitempty"`
	ApnsEnv                          string    `json:"apns_env,omitempty"`
	ApnsCertificates                 string    `json:"apns_certificates,omitempty"`
	SafariApnsCertificate            string    `json:"safari_apns_certificate,omitempty"`
	SafariSiteOrigin                 string    `json:"safari_site_origin,omitempty"`
	SafariPushId                     string    `json:"safari_push_id,omitempty"`
	SafariIcon1616                   string    `json:"safari_icon_16_16,omitempty"`
	SafariIcon3232                   string    `json:"safari_icon_32_32,omitempty"`
	SafariIcon6464                   string    `json:"safari_icon_64_64,omitempty"`
	SafariIcon128128                 string    `json:"safari_icon_128_128,omitempty"`
	SafariIcon256256                 string    `json:"safari_icon_256_256,omitempty"`
	SiteName                         string    `json:"site_name,omitempty"`
	BasicAuthKey                     string    `json:"basic_auth_key,omitempty"`
}

type Device struct {
	Id                string            `json:"id"`
	Identifier        string            `json:"identifier"`
	SessionCount      int               `json:"session_count"`
	Language          string            `json:"language"`
	Timezone          int               `json:"timezone"`
	GameVersion       string            `json:"game_version"`
	DeviceOs          string            `json:"device_os"`
	DeviceType        int               `json:"device_type"`
	DeviceModel       string            `json:"device_model"`
	AdId              interface{}       `json:"ad_id"`
	Tags              map[string]string `json:"tags"`
	LastActive        int               `json:"last_active"`
	AmountSpent       float64           `json:"amount_spent"`
	CreatedAt         int               `json:"created_at"`
	InvalidIdentifier bool              `json:"invalid_identifier"`
	BadgeCount        int               `json:"badge_count"`
	Sdk               string            `json:"sdk"`
	TestType          interface{}       `json:"test_type"`
	Ip                string            `json:"ip"`
	ExternalUserId    interface{}       `json:"external_user_id"`
}

type CancelNotificationResult struct {
	Success bool `json:"success"`
}

type AddDeviceResult struct {
	Success bool   `json:"success"`
	Id      string `json:"id"`
}

type EditDeviceResult struct {
	Success bool `json:"success"`
}
