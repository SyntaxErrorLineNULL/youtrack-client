package main

// IncorrectField This article examines an error that might happen when you try to update
// a custom field in an issue using YouTrack REST API.
// You can face this error while making requests to different endpoints.
// This topic describes the most common scenario when you can receive this error message.
type IncorrectField struct {
	CustomFields []Fields `json:"customFields"`
}

type Fields struct {
	Name  string `json:"name,omitempty"`
	Type  string `json:"$type,omitempty"`
	Value struct {
		Name string `json:"name,omitempty"`
	}
}
