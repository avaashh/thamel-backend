package dist

import (
	"fmt"
	"thamel/types"

	"encoding/json"
	"io"
	"log"
	"net/http"
)

// Events (sdate string) -> map
// sdate: "yyyymmdd" format date
func LiveEvents(sdate string) types.Response25 {
	client := &http.Client{}

	req, err := http.NewRequest("GET",
		"https://events.grinnell.edu/live/calendar/view/"+
			"day/date/"+sdate+
			"?user_tz=America%2FChicago&"+
			"template_vars=id,latitude,longitude,location,time,href,image_raw,title_link,summary,"+
			"until,is_canceled,is_online,image_src,repeats,is_multi_day,is_first_multi_day,"+
			"multi_day_span,tag_classes,category_classes,has_map"+
			"&syntax=%3Cwidget%20type%3D%22events_calendar%22%3E%3Carg%20id%3D%22mini_cal_heat_map%22%3Etrue%3C%2Farg%3E%3Carg%20id%3D%22thumb_width%22%3E200%3C%2Farg%3E%3Carg%20id%3D%22thumb_height%22%3E200%3C%2Farg%3E%3Carg%20id%3D%22hide_repeats%22%3Etrue%3C%2Farg%3E%3Carg%20id%3D%22show_groups%22%3Etrue%3C%2Farg%3E%3Carg%20id%3D%22show_locations%22%3Efalse%3C%2Farg%3E%3Carg%20id%3D%22show_tags%22%3Etrue%3C%2Farg%3E%3Carg%20id%3D%22month_view_day_limit%22%3E2%3C%2Farg%3E%3Carg%20id%3D%22use_tag_classes%22%3Efalse%3C%2Farg%3E%3Carg%20id%3D%22search_all_events_only%22%3Etrue%3C%2Farg%3E%3Carg%20id%3D%22use_modular_templates%22%3Etrue%3C%2Farg%3E%3Carg%20id%3D%22exclude_group%22%3EAdmin%3C%2Farg%3E%3C%2Fwidget%3E", nil)

	if err != nil {
		log.Fatal(err)
		return types.NullResponse25()
	}

	resp, err := client.Do(req)
	if err != nil {
		return types.NullResponse25()
	}

	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		return types.NullResponse25()
	}

	var eventData types.Response25
	err = json.Unmarshal(bodyText, &eventData)
	if err != nil {
		return types.NullResponse25()
	}

	return eventData
}

// IndivEvent (slug string) -> map
// slug: slug of the event
func IndivEvent(slug string) types.ResponseIE {
	client := &http.Client{}
	req, err := http.NewRequest("GET",
		"https://events.grinnell.edu/live/calendar/view/event/slug/"+slug+
			"?user_tz=America%2FChicago&template_vars=hero_image,title,date_time,location,repeats,until,image,online_url,online_button_label,online_instructions,contact_info,related_content,share_links,summary,description,cost_type,cost,registration,tags_calendar,group,custom_organization,categories_audience,id,ical_download_href,ical_all_series,is_online,is_repeating,location_latitude,has_registration&syntax=%3Cwidget%20type%3D%22events_calendar%22%3E%3Carg%20id%3D%22mini_cal_heat_map%22%3Etrue%3C%2Farg%3E%3Carg%20id%3D%22thumb_width%22%3E200%3C%2Farg%3E%3Carg%20id%3D%22thumb_height%22%3E200%3C%2Farg%3E%3Carg%20id%3D%22hide_repeats%22%3Etrue%3C%2Farg%3E%3Carg%20id%3D%22show_groups%22%3Etrue%3C%2Farg%3E%3Carg%20id%3D%22show_locations%22%3Efalse%3C%2Farg%3E%3Carg%20id%3D%22show_tags%22%3Etrue%3C%2Farg%3E%3Carg%20id%3D%22month_view_day_limit%22%3E2%3C%2Farg%3E%3Carg%20id%3D%22use_tag_classes%22%3Efalse%3C%2Farg%3E%3Carg%20id%3D%22search_all_events_only%22%3Etrue%3C%2Farg%3E%3Carg%20id%3D%22use_modular_templates%22%3Etrue%3C%2Farg%3E%3Carg%20id%3D%22exclude_group%22%3EAdmin%3C%2Farg%3E%3C%2Fwidget%3E", nil)

	if err != nil {
		return types.NullResponseIE()
	}
	resp, err := client.Do(req)
	if err != nil {
		return types.NullResponseIE()
	}
	
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	
	if err != nil {
		return types.NullResponseIE()
	}

	// Define a struct to match the returned JSON structure
	var eventData struct {
		Event types.ResponseIE `json:"event"`
	}

	// Unmarshal JSON into the struct
	err = json.Unmarshal(bodyText, &eventData)
	if err != nil {
		fmt.Println(2, err)
		return types.NullResponseIE()
	}
	return eventData.Event
}