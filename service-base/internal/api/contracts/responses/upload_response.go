package responses

type UploadResponse struct {
	Body uploadUrls `json:"body"`
}

type uploadUrls struct {
	UploadURL string `json:"uploadUrl" doc:"PUT the file here directly, with the same Content-Type used to request this URL"`
	PublicURL string `json:"publicUrl" doc:"Where the file will be reachable once the upload completes"`
}

func ToUploadResponse(uploadURL, publicURL string) *UploadResponse {
	return &UploadResponse{
		Body: uploadUrls{
			UploadURL: uploadURL,
			PublicURL: publicURL,
		},
	}
}
