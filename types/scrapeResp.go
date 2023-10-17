package types

type Event struct {
	ID                 string      `json:"id,omitempty"`
	TsStart            int64       `json:"ts_start,omitempty"`
	TsEnd              int64       `json:"ts_end,omitempty"`
	Tz                 string      `json:"tz,omitempty"`
	TzAbbrv            string      `json:"tz_abbrv,omitempty"`
	TzOffset           int         `json:"tz_offset,omitempty"`
	Tz2Offset          int         `json:"tz2_offset,omitempty"`
	TzFormat           string      `json:"tz_format,omitempty"`
	Href               string      `json:"href,omitempty"`
	Title              string      `json:"title,omitempty"`
	Status             string      `json:"status,omitempty"`
	Tags               []string    `json:"tags,omitempty"`
	Categories         []string    `json:"categories,omitempty"`
	CategoriesAudience []string    `json:"categories_audience,omitempty"`
	Location           string      `json:"location,omitempty"`
	Latitude           string      `json:"latitude,omitempty"`
	Longitude          string      `json:"longitude,omitempty"`
	Summary            string      `json:"summary,omitempty"`
	IsCanceled         string      `json:"is_canceled,omitempty"`
	ImageSrc           string      `json:"image_src,omitempty"`
	TagClasses         interface{} `json:"tag_classes,omitempty"`
	CategoryClasses    []string    `json:"category_classes,omitempty"`
	IsOnline           string      `json:"is_online,omitempty"`
}

// Response fetched from scraping 25 live will be cast to Response25
type Response25 struct {
	Date         string  `json:"date"`
	Start_date   string  `json:"start_date"`
	End_date     string  `json:"end_date"`
	Prev_date    string  `json:"prev_date"`
	Next_date    string  `json:"next_date"`
	Title        string  `json:"title"`
	Events       []Event `json:"events"`
	Is_this_week bool    `json:"is_this_week"`
}

// Individual Event Response fetched from 25 live will be cast to ResponseIE
type ResponseIE struct {
	Id                  int         `json:"id"`
	Start_date          int         `json:"ts_start"` // in epoch time
	Href                int         `json:"href"`     // in epoch time
	Title               string      `json:"title"`
	Location            string      `json:"location"`
	Description         string      `json:"description"`
	Image               string      `json:"image_src"` // url
	Tags                interface{} `json:"tags_calendar"`
	Contact_info        string      `json:"contact_info"`
	Categories          interface{} `json:"categories"`
	Categories_audience interface{} `json:"categories_audience"`
	Category_classes    interface{} `json:"category_classes"`
	Summary             string      `json:"summary"`
}

func NullResponse25() Response25 {
	return Response25{
		Date:         "",
		Start_date:   "",
		End_date:     "",
		Prev_date:    "",
		Next_date:    "",
		Title:        "",
		Events:       nil,
		Is_this_week: false,
	}
}

func NullResponseIE() ResponseIE {
	return ResponseIE{
		Id:                  0,
		Start_date:          0,
		Href:                0,
		Title:               "",
		Location:            "",
		Description:         "",
		Image:               "",
		Tags:                nil,
		Contact_info:        "",
		Categories:          nil,
		Categories_audience: nil,
		Category_classes:    nil,
		Summary:             "",
	}
}
