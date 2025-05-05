package team_member

type InviteTokenInput struct {
	TeamID int    `json:"team_id"`
	Mail   string `json:"mail"`
}
