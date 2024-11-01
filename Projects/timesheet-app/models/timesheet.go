package models

type TimesheetDetail struct {
	ProjectID       int    `json:"ProjectID"`
	SubProjectID    int    `json:"SubProjectID"`
	JiraSnowID      string `json:"JiraSnowID"`
	TaskDescription string `json:"TaskDescription"`
	HoursSpent      int    `json:"HoursSpent"`
	Comments        string `json:"Comments"`
}
