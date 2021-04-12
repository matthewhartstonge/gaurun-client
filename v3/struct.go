package gaurun

import "github.com/mercari/gaurun/gaurun"

// A Payload contains notifications to be sent to gaurun.
type Payload struct {
	Notifications []*Notification `json:"notifications"`
}

// A Notification has gaurun notification data.
type Notification struct {
	// Common
	Tokens     []string `json:"token"`
	Platform   Platform `json:"platform"`
	Message    string   `json:"message"`
	Identifier string   `json:"identifier,omitempty"`
	// Android
	AndroidSetting
	// iOS
	IOSSetting
	// Metadata
	ID     uint64    `json:"seq_id,omitempty"`
	Extend []*Extend `json:"extend,omitempty"`
}

// A Platform is alias gaurun platform enum.
type Platform int

const (
	// PlatformAndroid is enum for FCM/GCM.
	PlatformAndroid Platform = gaurun.PlatFormAndroid
	// PlatformIOS is enum for APNs.
	PlatformIOS Platform = gaurun.PlatFormIos
)

// An AndroidSetting has setting fields for FCM/GCM.
type AndroidSetting struct {
	CollapseKey    string   `json:"collapse_key,omitempty"`
	DelayWhileIdle bool     `json:"delay_while_idle,omitempty"`
	TimeToLive     int      `json:"time_to_live,omitempty"`
	Priority       Priority `json:"priority,omitempty"`
}

// Priority sets an android specific delivery priority on the push notification.
type Priority string

const (
	// Normal priority.
	// This is the default priority for data messages. Normal priority messages
	// are delivered immediately when the app is in the foreground.
	//
	// When the device is in Doze, delivery may be delayed to conserve battery.
	// For less time-sensitive messages, such as notifications of new email,
	// keeping your UI in sync, or syncing app data in the background, choose
	// normal delivery priority.
	PriorityNormal Priority = "normal"

	// High priority.
	// FCM attempts to deliver high priority messages immediately, allowing the
	// FCM service to wake a sleeping device when necessary and to run some
	// limited processing (including very limited network access). High priority
	// messages generally should result in user interaction with your app or its
	// notifications. If FCM detects a pattern in which they don't, your
	// messages may be de-prioritized.
	// Android P introduced app standby buckets which limit the number of FCM
	// high priority messages you can send to your app that don't result in the
	// user using your app or viewing a notification. If, in response to a high
	// priority message, a notification is displayed in a way that is visible to
	// the user, then your app standby bucket quota will not be consumed by that
	// message.
	PriorityHigh Priority = "high"
)

// An IOSSetting has setting fields for APNs.
type IOSSetting struct {
	Title            string   `json:"title,omitempty"`
	Subtitle         string   `json:"subtitle,omitempty"`
	Badge            int      `json:"badge,omitempty"`
	Category         string   `json:"category,omitempty"`
	Sound            string   `json:"sound,omitempty"`
	ContentAvailable bool     `json:"content_available,omitempty"`
	MutableContent   bool     `json:"mutable_content,omitempty"`
	Expiry           int      `json:"expiry,omitempty"`
	PushType         PushType `json:"push_type,omitempty"`
}

// PushType provides enumerated values for the apns-push-type header.
type PushType string

// The apns-push-type header field has six valid values.
// The descriptions below describe when and how to use these values.
//
// Refer: https://developer.apple.com/documentation/usernotifications/setting_up_a_remote_notification_server/sending_notification_requests_to_apns/
const (
	// Use the alert push type for notifications that trigger a user interaction
	// for example, an alert, badge, or sound.
	//
	// The alert push type is required on watchOS 6 and later.
	// It is recommended on macOS, iOS, tvOS, and iPadOS.
	PushTypeAlert PushType = "alert"

	// Use the background push type for notifications that deliver content in
	// the background, and don’t trigger any user interactions.
	//
	// The background push type is required on watchOS 6 and later.
	// It is recommended on macOS, iOS, tvOS, and iPadOS.
	PushTypeBackground PushType = "background"

	// Use the voip push type for notifications that provide information about
	// an incoming Voice-over-IP (VoIP) call.
	//
	// The voip push type is not available on watchOS.
	// It is recommended on macOS, iOS, tvOS, and iPadOS.
	PushTypeVoip PushType = "voip"

	// Use the complication push type for notifications that contain update
	// information for a watchOS app’s complications.
	//
	// The complication push type is recommended for watchOS and iOS.
	// It is not available on macOS, tvOS, and iPadOS.
	PushTypeComplication PushType = "complication"

	// Use the fileprovider push type to signal changes to a File Provider
	// extension.
	//
	// The fileprovider push type is not available on watchOS.
	// It is recommended on macOS, iOS, tvOS, and iPadOS.
	PushTypeFileProvider PushType = "fileprovider"

	// Use the mdm push type for notifications that tell managed devices to
	// contact the MDM server.
	//
	// The mdm push type is not available on watchOS.
	// It is recommended on macOS, iOS, tvOS, and iPadOS.
	PushTypeMdm PushType = "mdm"
)

// An Extend is alias gaurun.ExtendJSON.
type Extend gaurun.ExtendJSON
