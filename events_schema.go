// DO NOT EDIT Code generated by schema/make_schema.go
package vince

import (
	schemav2pb "github.com/polarsignals/frostdb/gen/proto/go/frostdb/schema/v1alpha2"
)

const EventTable = "event"
const SessionTable = "session"
const LabelTable = "label"
const ImportedVisitorTable = "imported_visitor"
const ImportedSourcesTable = "imported_sources"
const ImportedPagesTable = "imported_pages"
const ImportedEntryPagesTable = "imported_entry_pages"
const ImportedExitPagesTable = "imported_exit_pages"
const ImportedLocationsTable = "imported_locations"
const ImportedDevicesTable = "imported_devices"
const ImportedBrowserTable = "imported_browser"
const ImportedOperatingSystemTable = "imported_operating_system"

var EventSchema = &schemav2pb.Schema{
	Root: &schemav2pb.Group{
		Name: "event",
		Nodes: []*schemav2pb.Node{
			{
				Type: &schemav2pb.Node_Leaf{
					Leaf: &schemav2pb.Leaf{
						Name: "name",
						StorageLayout: &schemav2pb.StorageLayout{
							Type: schemav2pb.StorageLayout_TYPE_STRING,
						},
					},
				},
			},
			{
				Type: &schemav2pb.Node_Leaf{
					Leaf: &schemav2pb.Leaf{
						Name: "domain",
						StorageLayout: &schemav2pb.StorageLayout{
							Type: schemav2pb.StorageLayout_TYPE_STRING,
						},
					},
				},
			},
			{
				Type: &schemav2pb.Node_Leaf{
					Leaf: &schemav2pb.Leaf{
						Name: "user_id",
						StorageLayout: &schemav2pb.StorageLayout{
							Type: schemav2pb.StorageLayout_TYPE_INT64,
						},
					},
				},
			},
			{
				Type: &schemav2pb.Node_Leaf{
					Leaf: &schemav2pb.Leaf{
						Name: "session_id",
						StorageLayout: &schemav2pb.StorageLayout{
							Type: schemav2pb.StorageLayout_TYPE_INT64,
						},
					},
				},
			},
			{
				Type: &schemav2pb.Node_Leaf{
					Leaf: &schemav2pb.Leaf{
						Name: "hostname",
						StorageLayout: &schemav2pb.StorageLayout{
							Type: schemav2pb.StorageLayout_TYPE_STRING,
						},
					},
				},
			},
			{
				Type: &schemav2pb.Node_Leaf{
					Leaf: &schemav2pb.Leaf{
						Name: "pathname",
						StorageLayout: &schemav2pb.StorageLayout{
							Type: schemav2pb.StorageLayout_TYPE_STRING,
						},
					},
				},
			},
			{
				Type: &schemav2pb.Node_Leaf{
					Leaf: &schemav2pb.Leaf{
						Name: "referrer",
						StorageLayout: &schemav2pb.StorageLayout{
							Type: schemav2pb.StorageLayout_TYPE_STRING,
						},
					},
				},
			},
			{
				Type: &schemav2pb.Node_Leaf{
					Leaf: &schemav2pb.Leaf{
						Name: "referrer_source",
						StorageLayout: &schemav2pb.StorageLayout{
							Type: schemav2pb.StorageLayout_TYPE_STRING,
						},
					},
				},
			},
			{
				Type: &schemav2pb.Node_Leaf{
					Leaf: &schemav2pb.Leaf{
						Name: "country_code",
						StorageLayout: &schemav2pb.StorageLayout{
							Type: schemav2pb.StorageLayout_TYPE_STRING,
						},
					},
				},
			},
			{
				Type: &schemav2pb.Node_Leaf{
					Leaf: &schemav2pb.Leaf{
						Name: "screen_size",
						StorageLayout: &schemav2pb.StorageLayout{
							Type: schemav2pb.StorageLayout_TYPE_STRING,
						},
					},
				},
			},
			{
				Type: &schemav2pb.Node_Leaf{
					Leaf: &schemav2pb.Leaf{
						Name: "operating_system",
						StorageLayout: &schemav2pb.StorageLayout{
							Type: schemav2pb.StorageLayout_TYPE_STRING,
						},
					},
				},
			},
			{
				Type: &schemav2pb.Node_Leaf{
					Leaf: &schemav2pb.Leaf{
						Name: "browser",
						StorageLayout: &schemav2pb.StorageLayout{
							Type: schemav2pb.StorageLayout_TYPE_STRING,
						},
					},
				},
			},
			{
				Type: &schemav2pb.Node_Leaf{
					Leaf: &schemav2pb.Leaf{
						Name: "utm_medium",
						StorageLayout: &schemav2pb.StorageLayout{
							Type: schemav2pb.StorageLayout_TYPE_STRING,
						},
					},
				},
			},
			{
				Type: &schemav2pb.Node_Leaf{
					Leaf: &schemav2pb.Leaf{
						Name: "utm_source",
						StorageLayout: &schemav2pb.StorageLayout{
							Type: schemav2pb.StorageLayout_TYPE_STRING,
						},
					},
				},
			},
			{
				Type: &schemav2pb.Node_Leaf{
					Leaf: &schemav2pb.Leaf{
						Name: "utm_campaign",
						StorageLayout: &schemav2pb.StorageLayout{
							Type: schemav2pb.StorageLayout_TYPE_STRING,
						},
					},
				},
			},
			{
				Type: &schemav2pb.Node_Leaf{
					Leaf: &schemav2pb.Leaf{
						Name: "browser_version",
						StorageLayout: &schemav2pb.StorageLayout{
							Type: schemav2pb.StorageLayout_TYPE_STRING,
						},
					},
				},
			},
			{
				Type: &schemav2pb.Node_Leaf{
					Leaf: &schemav2pb.Leaf{
						Name: "operating_system_version",
						StorageLayout: &schemav2pb.StorageLayout{
							Type: schemav2pb.StorageLayout_TYPE_STRING,
						},
					},
				},
			},
			{
				Type: &schemav2pb.Node_Leaf{
					Leaf: &schemav2pb.Leaf{
						Name: "subdivision1_code",
						StorageLayout: &schemav2pb.StorageLayout{
							Type: schemav2pb.StorageLayout_TYPE_STRING,
						},
					},
				},
			},
			{
				Type: &schemav2pb.Node_Leaf{
					Leaf: &schemav2pb.Leaf{
						Name: "subdivision2_code",
						StorageLayout: &schemav2pb.StorageLayout{
							Type: schemav2pb.StorageLayout_TYPE_STRING,
						},
					},
				},
			},
			{
				Type: &schemav2pb.Node_Leaf{
					Leaf: &schemav2pb.Leaf{
						Name: "city_geoname_id",
						StorageLayout: &schemav2pb.StorageLayout{
							Type: schemav2pb.StorageLayout_TYPE_INT64,
						},
					},
				},
			},
			{
				Type: &schemav2pb.Node_Leaf{
					Leaf: &schemav2pb.Leaf{
						Name: "utm_content",
						StorageLayout: &schemav2pb.StorageLayout{
							Type: schemav2pb.StorageLayout_TYPE_STRING,
						},
					},
				},
			},
			{
				Type: &schemav2pb.Node_Leaf{
					Leaf: &schemav2pb.Leaf{
						Name: "utm_term",
						StorageLayout: &schemav2pb.StorageLayout{
							Type: schemav2pb.StorageLayout_TYPE_STRING,
						},
					},
				},
			},
			{
				Type: &schemav2pb.Node_Leaf{
					Leaf: &schemav2pb.Leaf{
						Name: "transferred_from",
						StorageLayout: &schemav2pb.StorageLayout{
							Type: schemav2pb.StorageLayout_TYPE_STRING,
						},
					},
				},
			},

			{
				Type: &schemav2pb.Node_Group{
					Group: &schemav2pb.Group{
						Name:  "labels",
						Nodes: []*schemav2pb.Node{},
					},
				},
			},
			{
				Type: &schemav2pb.Node_Leaf{
					Leaf: &schemav2pb.Leaf{
						Name: "timestamp",
						StorageLayout: &schemav2pb.StorageLayout{
							Type:        schemav2pb.StorageLayout_TYPE_INT64,
							Encoding:    schemav2pb.StorageLayout_ENCODING_DELTA_BINARY_PACKED,
							Compression: schemav2pb.StorageLayout_COMPRESSION_ZSTD,
						},
					},
				},
			},
		},
	},
}
var SessionSchema = &schemav2pb.Schema{
	Root: &schemav2pb.Group{
		Name: "session",
		Nodes: []*schemav2pb.Node{
			{
				Type: &schemav2pb.Node_Leaf{
					Leaf: &schemav2pb.Leaf{
						Name: "session_id",
						StorageLayout: &schemav2pb.StorageLayout{
							Type: schemav2pb.StorageLayout_TYPE_INT64,
						},
					},
				},
			},
			{
				Type: &schemav2pb.Node_Leaf{
					Leaf: &schemav2pb.Leaf{
						Name: "sign",
						StorageLayout: &schemav2pb.StorageLayout{
							Type: schemav2pb.StorageLayout_TYPE_INT64,
						},
					},
				},
			},
			{
				Type: &schemav2pb.Node_Leaf{
					Leaf: &schemav2pb.Leaf{
						Name: "domain",
						StorageLayout: &schemav2pb.StorageLayout{
							Type: schemav2pb.StorageLayout_TYPE_STRING,
						},
					},
				},
			},
			{
				Type: &schemav2pb.Node_Leaf{
					Leaf: &schemav2pb.Leaf{
						Name: "user_id",
						StorageLayout: &schemav2pb.StorageLayout{
							Type: schemav2pb.StorageLayout_TYPE_INT64,
						},
					},
				},
			},
			{
				Type: &schemav2pb.Node_Leaf{
					Leaf: &schemav2pb.Leaf{
						Name: "hostname",
						StorageLayout: &schemav2pb.StorageLayout{
							Type: schemav2pb.StorageLayout_TYPE_STRING,
						},
					},
				},
			},
			{
				Type: &schemav2pb.Node_Leaf{
					Leaf: &schemav2pb.Leaf{
						Name: "is_bounce",
						StorageLayout: &schemav2pb.StorageLayout{
							Type: schemav2pb.StorageLayout_TYPE_BOOL,
						},
					},
				},
			},
			{
				Type: &schemav2pb.Node_Leaf{
					Leaf: &schemav2pb.Leaf{
						Name: "entry_page",
						StorageLayout: &schemav2pb.StorageLayout{
							Type: schemav2pb.StorageLayout_TYPE_STRING,
						},
					},
				},
			},
			{
				Type: &schemav2pb.Node_Leaf{
					Leaf: &schemav2pb.Leaf{
						Name: "exit_page",
						StorageLayout: &schemav2pb.StorageLayout{
							Type: schemav2pb.StorageLayout_TYPE_STRING,
						},
					},
				},
			},
			{
				Type: &schemav2pb.Node_Leaf{
					Leaf: &schemav2pb.Leaf{
						Name: "page_views",
						StorageLayout: &schemav2pb.StorageLayout{
							Type: schemav2pb.StorageLayout_TYPE_INT64,
						},
					},
				},
			},
			{
				Type: &schemav2pb.Node_Leaf{
					Leaf: &schemav2pb.Leaf{
						Name: "events",
						StorageLayout: &schemav2pb.StorageLayout{
							Type: schemav2pb.StorageLayout_TYPE_INT64,
						},
					},
				},
			},
			{
				Type: &schemav2pb.Node_Leaf{
					Leaf: &schemav2pb.Leaf{
						Name: "duration",
						StorageLayout: &schemav2pb.StorageLayout{
							Type:     schemav2pb.StorageLayout_TYPE_INT64,
							Encoding: schemav2pb.StorageLayout_ENCODING_RLE_DICTIONARY,
						},
					},
				},
			},
			{
				Type: &schemav2pb.Node_Leaf{
					Leaf: &schemav2pb.Leaf{
						Name: "referrer",
						StorageLayout: &schemav2pb.StorageLayout{
							Type: schemav2pb.StorageLayout_TYPE_STRING,
						},
					},
				},
			},
			{
				Type: &schemav2pb.Node_Leaf{
					Leaf: &schemav2pb.Leaf{
						Name: "referrer_source",
						StorageLayout: &schemav2pb.StorageLayout{
							Type: schemav2pb.StorageLayout_TYPE_STRING,
						},
					},
				},
			},
			{
				Type: &schemav2pb.Node_Leaf{
					Leaf: &schemav2pb.Leaf{
						Name: "country_code",
						StorageLayout: &schemav2pb.StorageLayout{
							Type: schemav2pb.StorageLayout_TYPE_STRING,
						},
					},
				},
			},
			{
				Type: &schemav2pb.Node_Leaf{
					Leaf: &schemav2pb.Leaf{
						Name: "operating_system",
						StorageLayout: &schemav2pb.StorageLayout{
							Type: schemav2pb.StorageLayout_TYPE_STRING,
						},
					},
				},
			},
			{
				Type: &schemav2pb.Node_Leaf{
					Leaf: &schemav2pb.Leaf{
						Name: "browser",
						StorageLayout: &schemav2pb.StorageLayout{
							Type: schemav2pb.StorageLayout_TYPE_STRING,
						},
					},
				},
			},
			{
				Type: &schemav2pb.Node_Leaf{
					Leaf: &schemav2pb.Leaf{
						Name: "utm_medium",
						StorageLayout: &schemav2pb.StorageLayout{
							Type: schemav2pb.StorageLayout_TYPE_STRING,
						},
					},
				},
			},
			{
				Type: &schemav2pb.Node_Leaf{
					Leaf: &schemav2pb.Leaf{
						Name: "utm_source",
						StorageLayout: &schemav2pb.StorageLayout{
							Type: schemav2pb.StorageLayout_TYPE_STRING,
						},
					},
				},
			},
			{
				Type: &schemav2pb.Node_Leaf{
					Leaf: &schemav2pb.Leaf{
						Name: "utm_campaign",
						StorageLayout: &schemav2pb.StorageLayout{
							Type: schemav2pb.StorageLayout_TYPE_STRING,
						},
					},
				},
			},
			{
				Type: &schemav2pb.Node_Leaf{
					Leaf: &schemav2pb.Leaf{
						Name: "browser_version",
						StorageLayout: &schemav2pb.StorageLayout{
							Type: schemav2pb.StorageLayout_TYPE_STRING,
						},
					},
				},
			},
			{
				Type: &schemav2pb.Node_Leaf{
					Leaf: &schemav2pb.Leaf{
						Name: "operating_system_version",
						StorageLayout: &schemav2pb.StorageLayout{
							Type: schemav2pb.StorageLayout_TYPE_STRING,
						},
					},
				},
			},
			{
				Type: &schemav2pb.Node_Leaf{
					Leaf: &schemav2pb.Leaf{
						Name: "subdivision1_code",
						StorageLayout: &schemav2pb.StorageLayout{
							Type: schemav2pb.StorageLayout_TYPE_STRING,
						},
					},
				},
			},
			{
				Type: &schemav2pb.Node_Leaf{
					Leaf: &schemav2pb.Leaf{
						Name: "subdivision2_code",
						StorageLayout: &schemav2pb.StorageLayout{
							Type: schemav2pb.StorageLayout_TYPE_STRING,
						},
					},
				},
			},
			{
				Type: &schemav2pb.Node_Leaf{
					Leaf: &schemav2pb.Leaf{
						Name: "city_geoname_id",
						StorageLayout: &schemav2pb.StorageLayout{
							Type: schemav2pb.StorageLayout_TYPE_INT64,
						},
					},
				},
			},
			{
				Type: &schemav2pb.Node_Leaf{
					Leaf: &schemav2pb.Leaf{
						Name: "utm_content",
						StorageLayout: &schemav2pb.StorageLayout{
							Type: schemav2pb.StorageLayout_TYPE_STRING,
						},
					},
				},
			},
			{
				Type: &schemav2pb.Node_Leaf{
					Leaf: &schemav2pb.Leaf{
						Name: "utm_term",
						StorageLayout: &schemav2pb.StorageLayout{
							Type: schemav2pb.StorageLayout_TYPE_STRING,
						},
					},
				},
			},
			{
				Type: &schemav2pb.Node_Leaf{
					Leaf: &schemav2pb.Leaf{
						Name: "transferred_from",
						StorageLayout: &schemav2pb.StorageLayout{
							Type: schemav2pb.StorageLayout_TYPE_STRING,
						},
					},
				},
			},
			{
				Type: &schemav2pb.Node_Leaf{
					Leaf: &schemav2pb.Leaf{
						Name: "screen_size",
						StorageLayout: &schemav2pb.StorageLayout{
							Type: schemav2pb.StorageLayout_TYPE_STRING,
						},
					},
				},
			},

			{
				Type: &schemav2pb.Node_Group{
					Group: &schemav2pb.Group{
						Name:  "labels",
						Nodes: []*schemav2pb.Node{},
					},
				},
			},
			{
				Type: &schemav2pb.Node_Leaf{
					Leaf: &schemav2pb.Leaf{
						Name: "start",
						StorageLayout: &schemav2pb.StorageLayout{
							Type:        schemav2pb.StorageLayout_TYPE_INT64,
							Encoding:    schemav2pb.StorageLayout_ENCODING_DELTA_BINARY_PACKED,
							Compression: schemav2pb.StorageLayout_COMPRESSION_ZSTD,
						},
					},
				},
			},
			{
				Type: &schemav2pb.Node_Leaf{
					Leaf: &schemav2pb.Leaf{
						Name: "timestamp",
						StorageLayout: &schemav2pb.StorageLayout{
							Type:        schemav2pb.StorageLayout_TYPE_INT64,
							Encoding:    schemav2pb.StorageLayout_ENCODING_DELTA_BINARY_PACKED,
							Compression: schemav2pb.StorageLayout_COMPRESSION_ZSTD,
						},
					},
				},
			},
		},
	},
}
var LabelSchema = &schemav2pb.Schema{
	Root: &schemav2pb.Group{
		Name: "label",
		Nodes: []*schemav2pb.Node{
			{
				Type: &schemav2pb.Node_Leaf{
					Leaf: &schemav2pb.Leaf{
						Name: "name",
						StorageLayout: &schemav2pb.StorageLayout{
							Type: schemav2pb.StorageLayout_TYPE_STRING,
						},
					},
				},
			},
			{
				Type: &schemav2pb.Node_Leaf{
					Leaf: &schemav2pb.Leaf{
						Name: "value",
						StorageLayout: &schemav2pb.StorageLayout{
							Type: schemav2pb.StorageLayout_TYPE_STRING,
						},
					},
				},
			},
		},
	},
}
var ImportedVisitorSchema = &schemav2pb.Schema{
	Root: &schemav2pb.Group{
		Name: "imported_visitor",
		Nodes: []*schemav2pb.Node{
			{
				Type: &schemav2pb.Node_Leaf{
					Leaf: &schemav2pb.Leaf{
						Name: "site_id",
						StorageLayout: &schemav2pb.StorageLayout{
							Type: schemav2pb.StorageLayout_TYPE_INT64,
						},
					},
				},
			},
			{
				Type: &schemav2pb.Node_Leaf{
					Leaf: &schemav2pb.Leaf{
						Name: "date",
						StorageLayout: &schemav2pb.StorageLayout{
							Type:        schemav2pb.StorageLayout_TYPE_INT64,
							Encoding:    schemav2pb.StorageLayout_ENCODING_DELTA_BINARY_PACKED,
							Compression: schemav2pb.StorageLayout_COMPRESSION_ZSTD,
						},
					},
				},
			},
			{
				Type: &schemav2pb.Node_Leaf{
					Leaf: &schemav2pb.Leaf{
						Name: "visitors",
						StorageLayout: &schemav2pb.StorageLayout{
							Type: schemav2pb.StorageLayout_TYPE_INT64,
						},
					},
				},
			},
			{
				Type: &schemav2pb.Node_Leaf{
					Leaf: &schemav2pb.Leaf{
						Name: "pageviews",
						StorageLayout: &schemav2pb.StorageLayout{
							Type: schemav2pb.StorageLayout_TYPE_INT64,
						},
					},
				},
			},
			{
				Type: &schemav2pb.Node_Leaf{
					Leaf: &schemav2pb.Leaf{
						Name: "bounces",
						StorageLayout: &schemav2pb.StorageLayout{
							Type: schemav2pb.StorageLayout_TYPE_INT64,
						},
					},
				},
			},
			{
				Type: &schemav2pb.Node_Leaf{
					Leaf: &schemav2pb.Leaf{
						Name: "visits",
						StorageLayout: &schemav2pb.StorageLayout{
							Type: schemav2pb.StorageLayout_TYPE_INT64,
						},
					},
				},
			},
			{
				Type: &schemav2pb.Node_Leaf{
					Leaf: &schemav2pb.Leaf{
						Name: "visit_duration",
						StorageLayout: &schemav2pb.StorageLayout{
							Type: schemav2pb.StorageLayout_TYPE_INT64,
						},
					},
				},
			},
		},
	},
}
var ImportedSourcesSchema = &schemav2pb.Schema{
	Root: &schemav2pb.Group{
		Name: "imported_sources",
		Nodes: []*schemav2pb.Node{
			{
				Type: &schemav2pb.Node_Leaf{
					Leaf: &schemav2pb.Leaf{
						Name: "site_id",
						StorageLayout: &schemav2pb.StorageLayout{
							Type: schemav2pb.StorageLayout_TYPE_INT64,
						},
					},
				},
			},
			{
				Type: &schemav2pb.Node_Leaf{
					Leaf: &schemav2pb.Leaf{
						Name: "date",
						StorageLayout: &schemav2pb.StorageLayout{
							Type:        schemav2pb.StorageLayout_TYPE_INT64,
							Encoding:    schemav2pb.StorageLayout_ENCODING_DELTA_BINARY_PACKED,
							Compression: schemav2pb.StorageLayout_COMPRESSION_ZSTD,
						},
					},
				},
			},
			{
				Type: &schemav2pb.Node_Leaf{
					Leaf: &schemav2pb.Leaf{
						Name: "source",
						StorageLayout: &schemav2pb.StorageLayout{
							Type: schemav2pb.StorageLayout_TYPE_STRING,
						},
					},
				},
			},
			{
				Type: &schemav2pb.Node_Leaf{
					Leaf: &schemav2pb.Leaf{
						Name: "utm_medium",
						StorageLayout: &schemav2pb.StorageLayout{
							Type: schemav2pb.StorageLayout_TYPE_STRING,
						},
					},
				},
			},
			{
				Type: &schemav2pb.Node_Leaf{
					Leaf: &schemav2pb.Leaf{
						Name: "utm_campaign",
						StorageLayout: &schemav2pb.StorageLayout{
							Type: schemav2pb.StorageLayout_TYPE_STRING,
						},
					},
				},
			},
			{
				Type: &schemav2pb.Node_Leaf{
					Leaf: &schemav2pb.Leaf{
						Name: "utm_content",
						StorageLayout: &schemav2pb.StorageLayout{
							Type: schemav2pb.StorageLayout_TYPE_STRING,
						},
					},
				},
			},
			{
				Type: &schemav2pb.Node_Leaf{
					Leaf: &schemav2pb.Leaf{
						Name: "utm_term",
						StorageLayout: &schemav2pb.StorageLayout{
							Type: schemav2pb.StorageLayout_TYPE_STRING,
						},
					},
				},
			},
			{
				Type: &schemav2pb.Node_Leaf{
					Leaf: &schemav2pb.Leaf{
						Name: "visitors",
						StorageLayout: &schemav2pb.StorageLayout{
							Type: schemav2pb.StorageLayout_TYPE_INT64,
						},
					},
				},
			},
			{
				Type: &schemav2pb.Node_Leaf{
					Leaf: &schemav2pb.Leaf{
						Name: "visits",
						StorageLayout: &schemav2pb.StorageLayout{
							Type: schemav2pb.StorageLayout_TYPE_INT64,
						},
					},
				},
			},
			{
				Type: &schemav2pb.Node_Leaf{
					Leaf: &schemav2pb.Leaf{
						Name: "visit_duration",
						StorageLayout: &schemav2pb.StorageLayout{
							Type: schemav2pb.StorageLayout_TYPE_INT64,
						},
					},
				},
			},
			{
				Type: &schemav2pb.Node_Leaf{
					Leaf: &schemav2pb.Leaf{
						Name: "bounces",
						StorageLayout: &schemav2pb.StorageLayout{
							Type: schemav2pb.StorageLayout_TYPE_INT64,
						},
					},
				},
			},
		},
	},
}
var ImportedPagesSchema = &schemav2pb.Schema{
	Root: &schemav2pb.Group{
		Name: "imported_pages",
		Nodes: []*schemav2pb.Node{
			{
				Type: &schemav2pb.Node_Leaf{
					Leaf: &schemav2pb.Leaf{
						Name: "site_id",
						StorageLayout: &schemav2pb.StorageLayout{
							Type: schemav2pb.StorageLayout_TYPE_INT64,
						},
					},
				},
			},
			{
				Type: &schemav2pb.Node_Leaf{
					Leaf: &schemav2pb.Leaf{
						Name: "date",
						StorageLayout: &schemav2pb.StorageLayout{
							Type:        schemav2pb.StorageLayout_TYPE_INT64,
							Encoding:    schemav2pb.StorageLayout_ENCODING_DELTA_BINARY_PACKED,
							Compression: schemav2pb.StorageLayout_COMPRESSION_ZSTD,
						},
					},
				},
			},
			{
				Type: &schemav2pb.Node_Leaf{
					Leaf: &schemav2pb.Leaf{
						Name: "hostname",
						StorageLayout: &schemav2pb.StorageLayout{
							Type: schemav2pb.StorageLayout_TYPE_STRING,
						},
					},
				},
			},
			{
				Type: &schemav2pb.Node_Leaf{
					Leaf: &schemav2pb.Leaf{
						Name: "page",
						StorageLayout: &schemav2pb.StorageLayout{
							Type: schemav2pb.StorageLayout_TYPE_STRING,
						},
					},
				},
			},
			{
				Type: &schemav2pb.Node_Leaf{
					Leaf: &schemav2pb.Leaf{
						Name: "visitors",
						StorageLayout: &schemav2pb.StorageLayout{
							Type: schemav2pb.StorageLayout_TYPE_INT64,
						},
					},
				},
			},
			{
				Type: &schemav2pb.Node_Leaf{
					Leaf: &schemav2pb.Leaf{
						Name: "pagevies",
						StorageLayout: &schemav2pb.StorageLayout{
							Type: schemav2pb.StorageLayout_TYPE_INT64,
						},
					},
				},
			},
			{
				Type: &schemav2pb.Node_Leaf{
					Leaf: &schemav2pb.Leaf{
						Name: "exits",
						StorageLayout: &schemav2pb.StorageLayout{
							Type: schemav2pb.StorageLayout_TYPE_INT64,
						},
					},
				},
			},
			{
				Type: &schemav2pb.Node_Leaf{
					Leaf: &schemav2pb.Leaf{
						Name: "time_on_page",
						StorageLayout: &schemav2pb.StorageLayout{
							Type: schemav2pb.StorageLayout_TYPE_INT64,
						},
					},
				},
			},
		},
	},
}
var ImportedEntryPagesSchema = &schemav2pb.Schema{
	Root: &schemav2pb.Group{
		Name: "imported_entry_pages",
		Nodes: []*schemav2pb.Node{
			{
				Type: &schemav2pb.Node_Leaf{
					Leaf: &schemav2pb.Leaf{
						Name: "site_id",
						StorageLayout: &schemav2pb.StorageLayout{
							Type: schemav2pb.StorageLayout_TYPE_INT64,
						},
					},
				},
			},
			{
				Type: &schemav2pb.Node_Leaf{
					Leaf: &schemav2pb.Leaf{
						Name: "date",
						StorageLayout: &schemav2pb.StorageLayout{
							Type:        schemav2pb.StorageLayout_TYPE_INT64,
							Encoding:    schemav2pb.StorageLayout_ENCODING_DELTA_BINARY_PACKED,
							Compression: schemav2pb.StorageLayout_COMPRESSION_ZSTD,
						},
					},
				},
			},
			{
				Type: &schemav2pb.Node_Leaf{
					Leaf: &schemav2pb.Leaf{
						Name: "entry_page",
						StorageLayout: &schemav2pb.StorageLayout{
							Type: schemav2pb.StorageLayout_TYPE_STRING,
						},
					},
				},
			},
			{
				Type: &schemav2pb.Node_Leaf{
					Leaf: &schemav2pb.Leaf{
						Name: "visitors",
						StorageLayout: &schemav2pb.StorageLayout{
							Type: schemav2pb.StorageLayout_TYPE_INT64,
						},
					},
				},
			},
			{
				Type: &schemav2pb.Node_Leaf{
					Leaf: &schemav2pb.Leaf{
						Name: "entrances",
						StorageLayout: &schemav2pb.StorageLayout{
							Type: schemav2pb.StorageLayout_TYPE_INT64,
						},
					},
				},
			},
			{
				Type: &schemav2pb.Node_Leaf{
					Leaf: &schemav2pb.Leaf{
						Name: "visit_duration",
						StorageLayout: &schemav2pb.StorageLayout{
							Type: schemav2pb.StorageLayout_TYPE_INT64,
						},
					},
				},
			},
			{
				Type: &schemav2pb.Node_Leaf{
					Leaf: &schemav2pb.Leaf{
						Name: "bounces",
						StorageLayout: &schemav2pb.StorageLayout{
							Type: schemav2pb.StorageLayout_TYPE_INT64,
						},
					},
				},
			},
		},
	},
}
var ImportedExitPagesSchema = &schemav2pb.Schema{
	Root: &schemav2pb.Group{
		Name: "imported_exit_pages",
		Nodes: []*schemav2pb.Node{
			{
				Type: &schemav2pb.Node_Leaf{
					Leaf: &schemav2pb.Leaf{
						Name: "site_id",
						StorageLayout: &schemav2pb.StorageLayout{
							Type: schemav2pb.StorageLayout_TYPE_INT64,
						},
					},
				},
			},
			{
				Type: &schemav2pb.Node_Leaf{
					Leaf: &schemav2pb.Leaf{
						Name: "date",
						StorageLayout: &schemav2pb.StorageLayout{
							Type:        schemav2pb.StorageLayout_TYPE_INT64,
							Encoding:    schemav2pb.StorageLayout_ENCODING_DELTA_BINARY_PACKED,
							Compression: schemav2pb.StorageLayout_COMPRESSION_ZSTD,
						},
					},
				},
			},
			{
				Type: &schemav2pb.Node_Leaf{
					Leaf: &schemav2pb.Leaf{
						Name: "exit_page",
						StorageLayout: &schemav2pb.StorageLayout{
							Type: schemav2pb.StorageLayout_TYPE_STRING,
						},
					},
				},
			},
			{
				Type: &schemav2pb.Node_Leaf{
					Leaf: &schemav2pb.Leaf{
						Name: "visitors",
						StorageLayout: &schemav2pb.StorageLayout{
							Type: schemav2pb.StorageLayout_TYPE_INT64,
						},
					},
				},
			},
			{
				Type: &schemav2pb.Node_Leaf{
					Leaf: &schemav2pb.Leaf{
						Name: "exits",
						StorageLayout: &schemav2pb.StorageLayout{
							Type: schemav2pb.StorageLayout_TYPE_INT64,
						},
					},
				},
			},
		},
	},
}
var ImportedLocationsSchema = &schemav2pb.Schema{
	Root: &schemav2pb.Group{
		Name: "imported_locations",
		Nodes: []*schemav2pb.Node{
			{
				Type: &schemav2pb.Node_Leaf{
					Leaf: &schemav2pb.Leaf{
						Name: "site_id",
						StorageLayout: &schemav2pb.StorageLayout{
							Type: schemav2pb.StorageLayout_TYPE_INT64,
						},
					},
				},
			},
			{
				Type: &schemav2pb.Node_Leaf{
					Leaf: &schemav2pb.Leaf{
						Name: "date",
						StorageLayout: &schemav2pb.StorageLayout{
							Type:        schemav2pb.StorageLayout_TYPE_INT64,
							Encoding:    schemav2pb.StorageLayout_ENCODING_DELTA_BINARY_PACKED,
							Compression: schemav2pb.StorageLayout_COMPRESSION_ZSTD,
						},
					},
				},
			},
			{
				Type: &schemav2pb.Node_Leaf{
					Leaf: &schemav2pb.Leaf{
						Name: "country",
						StorageLayout: &schemav2pb.StorageLayout{
							Type: schemav2pb.StorageLayout_TYPE_STRING,
						},
					},
				},
			},
			{
				Type: &schemav2pb.Node_Leaf{
					Leaf: &schemav2pb.Leaf{
						Name: "region",
						StorageLayout: &schemav2pb.StorageLayout{
							Type: schemav2pb.StorageLayout_TYPE_STRING,
						},
					},
				},
			},
			{
				Type: &schemav2pb.Node_Leaf{
					Leaf: &schemav2pb.Leaf{
						Name: "city",
						StorageLayout: &schemav2pb.StorageLayout{
							Type: schemav2pb.StorageLayout_TYPE_INT64,
						},
					},
				},
			},
			{
				Type: &schemav2pb.Node_Leaf{
					Leaf: &schemav2pb.Leaf{
						Name: "visitors",
						StorageLayout: &schemav2pb.StorageLayout{
							Type: schemav2pb.StorageLayout_TYPE_INT64,
						},
					},
				},
			},
			{
				Type: &schemav2pb.Node_Leaf{
					Leaf: &schemav2pb.Leaf{
						Name: "visits",
						StorageLayout: &schemav2pb.StorageLayout{
							Type: schemav2pb.StorageLayout_TYPE_INT64,
						},
					},
				},
			},
			{
				Type: &schemav2pb.Node_Leaf{
					Leaf: &schemav2pb.Leaf{
						Name: "visit_duration",
						StorageLayout: &schemav2pb.StorageLayout{
							Type: schemav2pb.StorageLayout_TYPE_INT64,
						},
					},
				},
			},
			{
				Type: &schemav2pb.Node_Leaf{
					Leaf: &schemav2pb.Leaf{
						Name: "bounces",
						StorageLayout: &schemav2pb.StorageLayout{
							Type: schemav2pb.StorageLayout_TYPE_INT64,
						},
					},
				},
			},
		},
	},
}
var ImportedDevicesSchema = &schemav2pb.Schema{
	Root: &schemav2pb.Group{
		Name: "imported_devices",
		Nodes: []*schemav2pb.Node{
			{
				Type: &schemav2pb.Node_Leaf{
					Leaf: &schemav2pb.Leaf{
						Name: "site_id",
						StorageLayout: &schemav2pb.StorageLayout{
							Type: schemav2pb.StorageLayout_TYPE_INT64,
						},
					},
				},
			},
			{
				Type: &schemav2pb.Node_Leaf{
					Leaf: &schemav2pb.Leaf{
						Name: "date",
						StorageLayout: &schemav2pb.StorageLayout{
							Type:        schemav2pb.StorageLayout_TYPE_INT64,
							Encoding:    schemav2pb.StorageLayout_ENCODING_DELTA_BINARY_PACKED,
							Compression: schemav2pb.StorageLayout_COMPRESSION_ZSTD,
						},
					},
				},
			},
			{
				Type: &schemav2pb.Node_Leaf{
					Leaf: &schemav2pb.Leaf{
						Name: "devise",
						StorageLayout: &schemav2pb.StorageLayout{
							Type: schemav2pb.StorageLayout_TYPE_STRING,
						},
					},
				},
			},
			{
				Type: &schemav2pb.Node_Leaf{
					Leaf: &schemav2pb.Leaf{
						Name: "visitors",
						StorageLayout: &schemav2pb.StorageLayout{
							Type: schemav2pb.StorageLayout_TYPE_INT64,
						},
					},
				},
			},
			{
				Type: &schemav2pb.Node_Leaf{
					Leaf: &schemav2pb.Leaf{
						Name: "visits",
						StorageLayout: &schemav2pb.StorageLayout{
							Type: schemav2pb.StorageLayout_TYPE_INT64,
						},
					},
				},
			},
			{
				Type: &schemav2pb.Node_Leaf{
					Leaf: &schemav2pb.Leaf{
						Name: "visit_duration",
						StorageLayout: &schemav2pb.StorageLayout{
							Type: schemav2pb.StorageLayout_TYPE_INT64,
						},
					},
				},
			},
			{
				Type: &schemav2pb.Node_Leaf{
					Leaf: &schemav2pb.Leaf{
						Name: "bounces",
						StorageLayout: &schemav2pb.StorageLayout{
							Type: schemav2pb.StorageLayout_TYPE_INT64,
						},
					},
				},
			},
		},
	},
}
var ImportedBrowserSchema = &schemav2pb.Schema{
	Root: &schemav2pb.Group{
		Name: "imported_browser",
		Nodes: []*schemav2pb.Node{
			{
				Type: &schemav2pb.Node_Leaf{
					Leaf: &schemav2pb.Leaf{
						Name: "site_id",
						StorageLayout: &schemav2pb.StorageLayout{
							Type: schemav2pb.StorageLayout_TYPE_INT64,
						},
					},
				},
			},
			{
				Type: &schemav2pb.Node_Leaf{
					Leaf: &schemav2pb.Leaf{
						Name: "date",
						StorageLayout: &schemav2pb.StorageLayout{
							Type:        schemav2pb.StorageLayout_TYPE_INT64,
							Encoding:    schemav2pb.StorageLayout_ENCODING_DELTA_BINARY_PACKED,
							Compression: schemav2pb.StorageLayout_COMPRESSION_ZSTD,
						},
					},
				},
			},
			{
				Type: &schemav2pb.Node_Leaf{
					Leaf: &schemav2pb.Leaf{
						Name: "browser",
						StorageLayout: &schemav2pb.StorageLayout{
							Type: schemav2pb.StorageLayout_TYPE_STRING,
						},
					},
				},
			},
			{
				Type: &schemav2pb.Node_Leaf{
					Leaf: &schemav2pb.Leaf{
						Name: "visitors",
						StorageLayout: &schemav2pb.StorageLayout{
							Type: schemav2pb.StorageLayout_TYPE_INT64,
						},
					},
				},
			},
			{
				Type: &schemav2pb.Node_Leaf{
					Leaf: &schemav2pb.Leaf{
						Name: "visits",
						StorageLayout: &schemav2pb.StorageLayout{
							Type: schemav2pb.StorageLayout_TYPE_INT64,
						},
					},
				},
			},
			{
				Type: &schemav2pb.Node_Leaf{
					Leaf: &schemav2pb.Leaf{
						Name: "visit_duration",
						StorageLayout: &schemav2pb.StorageLayout{
							Type: schemav2pb.StorageLayout_TYPE_INT64,
						},
					},
				},
			},
			{
				Type: &schemav2pb.Node_Leaf{
					Leaf: &schemav2pb.Leaf{
						Name: "bounces",
						StorageLayout: &schemav2pb.StorageLayout{
							Type: schemav2pb.StorageLayout_TYPE_INT64,
						},
					},
				},
			},
		},
	},
}
var ImportedOperatingSystemSchema = &schemav2pb.Schema{
	Root: &schemav2pb.Group{
		Name: "imported_operating_system",
		Nodes: []*schemav2pb.Node{
			{
				Type: &schemav2pb.Node_Leaf{
					Leaf: &schemav2pb.Leaf{
						Name: "site_id",
						StorageLayout: &schemav2pb.StorageLayout{
							Type: schemav2pb.StorageLayout_TYPE_INT64,
						},
					},
				},
			},
			{
				Type: &schemav2pb.Node_Leaf{
					Leaf: &schemav2pb.Leaf{
						Name: "date",
						StorageLayout: &schemav2pb.StorageLayout{
							Type:        schemav2pb.StorageLayout_TYPE_INT64,
							Encoding:    schemav2pb.StorageLayout_ENCODING_DELTA_BINARY_PACKED,
							Compression: schemav2pb.StorageLayout_COMPRESSION_ZSTD,
						},
					},
				},
			},
			{
				Type: &schemav2pb.Node_Leaf{
					Leaf: &schemav2pb.Leaf{
						Name: "browser",
						StorageLayout: &schemav2pb.StorageLayout{
							Type: schemav2pb.StorageLayout_TYPE_STRING,
						},
					},
				},
			},
			{
				Type: &schemav2pb.Node_Leaf{
					Leaf: &schemav2pb.Leaf{
						Name: "visitors",
						StorageLayout: &schemav2pb.StorageLayout{
							Type: schemav2pb.StorageLayout_TYPE_INT64,
						},
					},
				},
			},
			{
				Type: &schemav2pb.Node_Leaf{
					Leaf: &schemav2pb.Leaf{
						Name: "visits",
						StorageLayout: &schemav2pb.StorageLayout{
							Type: schemav2pb.StorageLayout_TYPE_INT64,
						},
					},
				},
			},
			{
				Type: &schemav2pb.Node_Leaf{
					Leaf: &schemav2pb.Leaf{
						Name: "visit_duration",
						StorageLayout: &schemav2pb.StorageLayout{
							Type: schemav2pb.StorageLayout_TYPE_INT64,
						},
					},
				},
			},
			{
				Type: &schemav2pb.Node_Leaf{
					Leaf: &schemav2pb.Leaf{
						Name: "bounces",
						StorageLayout: &schemav2pb.StorageLayout{
							Type: schemav2pb.StorageLayout_TYPE_INT64,
						},
					},
				},
			},
		},
	},
}
