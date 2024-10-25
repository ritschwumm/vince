// Code generated by ingernal/gen/main.go. DO NOT EDIT.
package models

type Field byte

const (
	Field_unknown           Field = 0
	Field_timestamp         Field = 1
	Field_id                Field = 2
	Field_bounce            Field = 3
	Field_duration          Field = 4
	Field_city              Field = 5
	Field_view              Field = 6
	Field_session           Field = 7
	Field_browser           Field = 8
	Field_browser_version   Field = 9
	Field_country           Field = 10
	Field_device            Field = 11
	Field_domain            Field = 12
	Field_entry_page        Field = 13
	Field_event             Field = 14
	Field_exit_page         Field = 15
	Field_host              Field = 16
	Field_os                Field = 17
	Field_os_version        Field = 18
	Field_page              Field = 19
	Field_referrer          Field = 20
	Field_source            Field = 21
	Field_utm_campaign      Field = 22
	Field_utm_content       Field = 23
	Field_utm_medium        Field = 24
	Field_utm_source        Field = 25
	Field_utm_term          Field = 26
	Field_subdivision1_code Field = 27
	Field_subdivision2_code Field = 28
)

var (
	Field_name = map[Field]string{
		Field_unknown:           "unknown",
		Field_timestamp:         "timestamp",
		Field_id:                "id",
		Field_bounce:            "bounce",
		Field_duration:          "duration",
		Field_city:              "city",
		Field_view:              "view",
		Field_session:           "session",
		Field_browser:           "browser",
		Field_browser_version:   "browser_version",
		Field_country:           "country",
		Field_device:            "device",
		Field_domain:            "domain",
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
	}
	Field_value = map[string]Field{
		"unknown":           Field_unknown,
		"timestamp":         Field_timestamp,
		"id":                Field_id,
		"bounce":            Field_bounce,
		"duration":          Field_duration,
		"city":              Field_city,
		"view":              Field_view,
		"session":           Field_session,
		"browser":           Field_browser,
		"browser_version":   Field_browser_version,
		"country":           Field_country,
		"device":            Field_device,
		"domain":            Field_domain,
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
	}
)

func (f Field) String() string { return Field_name[f] }
