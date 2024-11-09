package config

type Configs struct {
	BanditThreshoulds	`yaml:"bandit"`
	Database 			`yaml:"mysql"` 
}

type BanditThreshoulds struct {
	Confidence struct {
		High      int   `yaml:"high"`
		Medium    int 	`yaml:"medium"`
		Low       int   `yaml:"low"`
		Undefined int	`yaml:"undefined"`
	}`yaml:"confidence"`
	Severity      struct {
		High      int   `yaml:"high"`
		Medium    int 	`yaml:"medium"`
		Low       int   `yaml:"low"`
		Undefined int	`yaml:"undefined"`
	} `yaml:"severity"`
}  

type Database struct {
	Username string `yaml:"username"`
	Schema   string `yaml:"schema"`
	Hostname string `yaml:"hostname"`
	Password string `yaml:"password"`
} 