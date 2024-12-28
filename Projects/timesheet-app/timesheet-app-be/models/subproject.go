package models

type SubProject struct {
	SubProjectID   int    `json:"sub_project_id"`
	SubProjectName string `json:"sub_project_name"`
	ProjectID      int    `json:"project_id"`
}
