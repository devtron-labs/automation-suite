package RequestDTOs

type FetchAppsByEnvironmentRequestDTO struct {
	AppNameSearch string `json:"appNameSearch"`
	SortBy        string `json:"sortBy"`
	SortOrder     string `json:"sortOrder"`
	Offset        int    `json:"offset"`
	HOffset       int    `json:"hOffset"`
	Size          int    `json:"size"`
}
