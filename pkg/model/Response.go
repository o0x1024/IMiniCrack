package model

type Sensitive struct {
	Desc     string
	MatchStr []string
	LineNo   string
	Path     string
}

type Regex struct {
	Id     string `yaml:"id"`
	Desc   string `yaml:"desc"`
	Record string `yaml:"record"` //正则
	Status bool   `yaml:"status"`
}

type Response struct {
	Code       int
	Sensitives []Sensitive
	Err        string
	Data       string
	Regexs     []Regex
	FileList   []string
	Msg        string
}
