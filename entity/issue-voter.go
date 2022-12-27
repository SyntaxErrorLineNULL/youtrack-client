package entity

type IssueVoter struct {
	ID        string          `json:"id"`
	HasVote   bool            `json:"hasVote"`
	Original  []User          `json:"original"`
	Duplicate []DuplicateVote `json:"duplicate"`
}
