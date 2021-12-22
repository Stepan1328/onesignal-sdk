package onesignal

import "time"

type Filters []interface{}

func NewFilters() Filters {
	return []interface{}{}
}

func (f *Filters) AddFilter(filter Filter) *Filters {
	*f = append(*f, filter)
	return f
}

func (f *Filters) AddOperatorOr() *Filters {
	*f = append(*f, OR)
	return f
}

func (f *Filters) AddOperatorAnd() *Filters {
	*f = append(*f, AND)
	return f
}

type Filter struct {
	Field Field `json:"field"`

	Key      string   `json:"key"`
	Relation Relation `json:"relation"`
	Value    string   `json:"value"`
	HoursAgo string   `json:"hours_ago"`

	Radius string `json:"radius"`
	Lat    string `json:"lat"`
	Long   string `json:"long"`
}

type Operator struct {
	Operator string `json:"operator"`
}

var (
	OR  = Operator{"OR"}
	AND = Operator{"AND"}
)

type Field string

const (
	LastSession  = Field("last_session")
	FirstSession = Field("first_session")
	SessionCount = Field("session_count")
	SessionTime  = Field("session_time")
	AmountSpent  = Field("amount_spent")
	BoughtSKU    = Field("bought_sku")
	Tag          = Field("tag")
	Language     = Field("language")
	AppVersion   = Field("app_version")
	Location     = Field("location")
	EmailField   = Field("email")
	Country      = Field("country")
)

type Relation string

const (
	More       = Relation(">")
	Less       = Relation("<")
	Equally    = Relation("=")
	NotEqually = Relation("!=")
	Exists     = Relation("exists")
	NotExists  = Relation("not_exists")
)

func NewFilter(field Field) *Filter {
	return &Filter{
		Field: field,
	}
}

func (f *Filter) AddRelation(relation Relation) *Filter {
	f.Relation = relation
	return f
}

func (f *Filter) AddValue(value string) *Filter {
	f.Value = value
	return f
}

type SpecificDevices struct {
	Devices                IDDevices `json:"include_player_ids,omitempty"`
	IncludeExternalUserIDs IDDevices `json:"include_external_user_ids,omitempty"`
}

type IDDevices []string

func NewDevices() IDDevices {
	return IDDevices{}
}

func (d *IDDevices) AddDevice(id string) *IDDevices {
	*d = append(*d, id)
	return d
}

type ContentAndLanguage struct {
	Names    string           `json:"names,omitempty"`
	Contents LocalizedContent `json:"contents,omitempty"`
	Headings LocalizedContent `json:"headings,omitempty"`
}

type LocalizedContent map[string]string

func NewLocalizedContent() LocalizedContent {
	return make(map[string]string)
}

func (c LocalizedContent) Add(location, content string) {
	c[location] = content
}

type Content struct {
	Location string
	Message  string
	Title    string
}

func English(message, title string) Content {
	return Content{
		Location: "en",
		Message:  message,
		Title:    title,
	}
}

func Russian(message, title string) Content {
	return Content{
		Location: "en",
		Message:  message,
		Title:    title,
	}
}

type Attachments struct {
	Data interface{} `json:"data,omitempty"`
}

type ActionButtons struct {
	Buttons     []Button `json:"buttons,omitempty"`
	WebButtons  []Button `json:"web_buttons"`
	IOSCategory string   `json:"ios_category"`
	IconType    string   `json:"icon_type"`
}

type Button struct {
	ID   string `json:"id,omitempty"`
	Text string `json:"text,omitempty"`
	Icon string `json:"icon,omitempty"`
}

type Delivery struct {
	SendAfter            *time.Time `json:"send_after,omitempty"`
	DelayedOption        string     `json:"delayed_option,omitempty"`
	DeliveryTimeOfDay    string     `json:"delivery_time_of_day,omitempty"`
	TTL                  int        `json:"ttl,omitempty"`
	Priority             int        `json:"priority,omitempty"`
	ApnsPushTypeOverride string     `json:"apns_push_type_override,omitempty"`
}

type GroupingAndCollapsing struct {
	AndroidGroup LocalizedContent `json:"android_group,omitempty"`
}
