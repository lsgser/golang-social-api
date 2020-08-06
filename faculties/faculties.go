package faculties

//Faculty struct
type Faculty struct {
	ID        int64  `json:"id,omitempty"`
	SchoolID  int64  `json:"school_id,omitempty"`
	Faculties string `json:"faculty,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
}
