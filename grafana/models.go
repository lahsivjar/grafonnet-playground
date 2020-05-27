package grafana

// CreateRequest models parameters required for creating/updating dashboard
type CreateRequest struct {
	Dashboard interface{} `json:"dashboard"`
	Overwrite bool        `json:"overwrite"`
	FolderID  int         `json:"folderId,omitempty"`
}

// CreateResponse models response from grafana for creating/updating dashboard
type CreateResponse struct {
	ID      int    `json:"id"`
	UID     string `json:"uid"`
	URL     string `json:"url"`
	Status  string `json:"status"`
	Version int    `json:"version"`
}

// NewCreateRequest returns a new create request model
func NewCreateRequest(folderID int) *CreateRequest {
	return &CreateRequest{
		Overwrite: true,
		FolderID:  folderID,
	}
}
