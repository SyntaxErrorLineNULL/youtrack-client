package entity

type Project struct {
	ID             string              `json:"id"`
	StartingNumber int64               `json:"startingNumber"`
	ShortName      string              `json:"shortName"`
	Description    string              `json:"description,omitempty"`
	Leader         User                `json:"leader,omitempty"`
	CreatedBy      User                `json:"createdBy,omitempty"`
	Issues         []Issue             `json:"issues"`
	CustomFields   *ProjectCustomField `json:"customFields"`
	Archived       bool                `json:"archived"`
	FromEmail      string              `json:"fromEmail"`
	ReplyToEmail   string              `json:"replyToEmail,omitempty"`
	Template       bool                `json:"template"`
	IconURL        string              `json:"iconUrl,omitempty"`
	Team           *UserGroup          `json:"team"`
	Name           string              `json:"name,omitempty"`
}
