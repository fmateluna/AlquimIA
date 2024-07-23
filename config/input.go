package config

// Estructura que representa el formato del archivo JSON
type TaskDef struct {
	Leng       string `json:"leng"`
	Version    string `json:"version"`
	Info       string `json:"info"`
	SourcePath string `json:"source_path"`
}
