package RequestDTOs

type SaveTeamRequestDto struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Active bool   `json:"active"`
}
