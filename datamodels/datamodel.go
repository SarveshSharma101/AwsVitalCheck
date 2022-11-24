package datamodels

type EdgeCamera struct {
	Name       string
	Resolution string
}
type EdgeStats struct {
	Camera           []EdgeCamera
	FolderUpdateTime string
}
type Stats struct {
	Stats []map[string]EdgeStats
}
type YamlConfig struct {
	EdgesParent string `yaml:"EdgesParent"`
	LatestImg   string `yaml:"LatestImg"`
	LogFile     string `yaml:"LogFile"`
	StatFile    string `yaml:"StatFile"`
}
