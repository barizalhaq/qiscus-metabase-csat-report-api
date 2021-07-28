package viewmodels

type MetabaseDataRequest struct {
	IgnoreCache bool                `json:"ignore_cache"`
	Parameters  []MetabaseParameter `json:"parameters"`
}

type MetabaseParameter struct {
	Type   string        `json:"type"`
	Target []interface{} `json:"target"`
	Value  string        `json:"value"`
}

type JSONRequest struct {
	On        string `json:"on" binding:"excluded_with=StartDate EndDate"`
	StartDate string `json:"start_date" binding:"required_with=EndDate"`
	EndDate   string `json:"end_date" binding:"required_with=StartDate"`
	Limit     int    `json:"limit" binding:"required"`
}
