package entity

// DuplicateVote  represents a vote for duplicates of the issue.
// This table describes attributes of the DuplicateVote entity.
// https://www.jetbrains.com/help/youtrack/devportal/api-entity-DuplicateVote.html
type DuplicateVote struct {
	ID    string `json:"id"`
	Issue Issue  `json:"issue,omitempty"`
	User  User   `json:"user,omitempty"`
}
