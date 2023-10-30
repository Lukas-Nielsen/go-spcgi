package spcgi

import "time"

type AlertingcenterAlerts struct {
	AlertID     string    `json:"alert_id,omitempty"`
	Severity    string    `json:"severity,omitempty"`
	Time        time.Time `json:"time,omitempty"`
	Collector   string    `json:"collector,omitempty"`
	Program     string    `json:"program,omitempty"`
	UserMessage string    `json:"user_message,omitempty"`
	Content     []struct {
		Attribute string `json:"attribute,omitempty"`
		Value     string `json:"value,omitempty"`
	} `json:"content,omitempty"`
}

type AlertingcenterAlertsProgram struct {
	Program           string `json:"program,omitempty"`
	ProgramTranslated string `json:"program_translated,omitempty"`
}

type AlertingcenterSeverity struct {
	Priority int      `json:"priority,omitempty"`
	Severity string   `json:"severity,omitempty"`
	Emitter  []string `json:"emitter,omitempty"`
}

type AlertingcenterSyslogPattern struct {
	PatternID      int    `json:"pattern_id,omitempty"`
	PatternName    string `json:"pattern_name,omitempty"`
	GroupID        int    `json:"group_id,omitempty"`
	PatternMessage string `json:"pattern_message,omitempty"`
}

type AlertingcenterSyslogPatterngroup struct {
	GroupID       int    `json:"group_id,omitempty"`
	GroupName     string `json:"group_name,omitempty"`
	GroupSeverity string `json:"group_severity,omitempty"`
	GroupMessage  string `json:"group_message,omitempty"`
}
