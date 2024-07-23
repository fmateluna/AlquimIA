package bot

type ResponseIA struct {
	Path        string `json:"Path,omitempty"`
	FileName    string `json:"FileName,omitempty"`
	Content     string `json:"Content,omitempty"`
	Description string `json:"Description,omitempty"`
}
