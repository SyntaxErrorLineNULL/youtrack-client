package entity

// Issue  represents an issue in YouTrack.
// This table describes attributes of the Issue entity.
// https://www.jetbrains.com/help/youtrack/devportal/api-entity-Issue.html
type Issue struct {
	ID                  string             `json:"id"`
	IDReadable          string             `json:"idReadable"`
	Created             int64              `json:"created,omitempty"`
	Updated             int64              `json:"updated,omitempty"`
	Resolved            int64              `json:"resolved"`
	Project             Project            `json:"project,omitempty"`
	Summary             string             `json:"summary,omitempty"`
	Description         string             `json:"description,omitempty"`
	WikifiedDescription string             `json:"wikifiedDescription"`
	Reporter            User               `json:"reporter,omitempty"`
	Updater             User               `json:"updater,omitempty"`
	DraftOwner          User               `json:"draftOwner,omitempty"`
	IsDraft             bool               `json:"isDraft,omitempty"`
	Visibility          any                `json:"visibility,omitempty"`
	Votes               int64              `json:"votes"`
	Comments            []IssueComment     `json:"comments"`
	CommentsCount       int64              `json:"commentsCount"`
	Tags                []IssueTag         `json:"tags"`
	Links               []IssueLink        `json:"links"`
	ExternalIssue       ExternalIssue      `json:"externalIssue,omitempty"`
	CustomFields        []IssueCustomField `json:"customFields"`
	Voters              []IssueVoter       `json:"voters"`
	Watchers            []IssueWatcher     `json:"watchers"`
	Attachments         []IssueAttachment  `json:"attachments"`
	SubTasks            IssueLink          `json:"subtasks"`
	Parent              IssueLink          `json:"parent,omitempty"`
}
