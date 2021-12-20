package config

//config.yaml中对应的配置的结构体
type Server struct {
	//jwt
	JWT    JWT    `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Zap    Zap    `mapstructure:"zap" json:"zap" yaml:"zap"`
	System System `mapstructure:"system" json:"system" yaml:"system"`
	// gorm
	Mysql  Mysql `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Pgsql  Pgsql `mapstructure:"pgsql" json:"pgsql" yaml:"pgsql"`
	DBList []DB  `mapstructure:"db-list" json:"db-list" yaml:"db-list"`
	// oss
	Local  Local  `mapstructure:"local" json:"local" yaml:"local"`
	Timer  Timer  `mapstructure:"timer" json:"timer" yaml:"timer"`
	Casbin Casbin `mapstructure:"casbin" json:"casbin" yaml:"casbin"`
}
