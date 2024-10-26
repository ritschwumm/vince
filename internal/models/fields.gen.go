// Code generated by ingernal/gen/main.go. DO NOT EDIT.
package models

type Field byte

const (
	Field_domain            Field = 0
	Field_browser           Field = 1
	Field_browser_version   Field = 2
	Field_country           Field = 3
	Field_device            Field = 4
	Field_entry_page        Field = 5
	Field_event             Field = 6
	Field_exit_page         Field = 7
	Field_host              Field = 8
	Field_os                Field = 9
	Field_os_version        Field = 10
	Field_page              Field = 11
	Field_referrer          Field = 12
	Field_source            Field = 13
	Field_utm_campaign      Field = 14
	Field_utm_content       Field = 15
	Field_utm_medium        Field = 16
	Field_utm_source        Field = 17
	Field_utm_term          Field = 18
	Field_subdivision1_code Field = 19
	Field_subdivision2_code Field = 20
	Field_city              Field = 21
	Field_view              Field = 22
	Field_session           Field = 23
	Field_timestamp         Field = 24
	Field_id                Field = 25
	Field_bounce            Field = 26
	Field_duration          Field = 27
)

var (
	Field_name = map[Field]string{
		Field_domain:            "domain",
		Field_browser:           "browser",
		Field_browser_version:   "browser_version",
		Field_country:           "country",
		Field_device:            "device",
		Field_entry_page:        "entry_page",
		Field_event:             "event",
		Field_exit_page:         "exit_page",
		Field_host:              "host",
		Field_os:                "os",
		Field_os_version:        "os_version",
		Field_page:              "page",
		Field_referrer:          "referrer",
		Field_source:            "source",
		Field_utm_campaign:      "utm_campaign",
		Field_utm_content:       "utm_content",
		Field_utm_medium:        "utm_medium",
		Field_utm_source:        "utm_source",
		Field_utm_term:          "utm_term",
		Field_subdivision1_code: "subdivision1_code",
		Field_subdivision2_code: "subdivision2_code",
		Field_city:              "city",
		Field_view:              "view",
		Field_session:           "session",
		Field_timestamp:         "timestamp",
		Field_id:                "id",
		Field_bounce:            "bounce",
		Field_duration:          "duration",
	}
	Field_value = map[string]Field{
		"domain":            Field_domain,
		"browser":           Field_browser,
		"browser_version":   Field_browser_version,
		"country":           Field_country,
		"device":            Field_device,
		"entry_page":        Field_entry_page,
		"event":             Field_event,
		"exit_page":         Field_exit_page,
		"host":              Field_host,
		"os":                Field_os,
		"os_version":        Field_os_version,
		"page":              Field_page,
		"referrer":          Field_referrer,
		"source":            Field_source,
		"utm_campaign":      Field_utm_campaign,
		"utm_content":       Field_utm_content,
		"utm_medium":        Field_utm_medium,
		"utm_source":        Field_utm_source,
		"utm_term":          Field_utm_term,
		"subdivision1_code": Field_subdivision1_code,
		"subdivision2_code": Field_subdivision2_code,
		"city":              Field_city,
		"view":              Field_view,
		"session":           Field_session,
		"timestamp":         Field_timestamp,
		"id":                Field_id,
		"bounce":            Field_bounce,
		"duration":          Field_duration,
	}
)

func (f Field) String() string { return Field_name[f] }
