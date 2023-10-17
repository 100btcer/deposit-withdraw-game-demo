package config

import (
	"log"
	"os"

	"github.com/naoina/toml"
	"github.com/naoina/toml/ast"
	"github.com/spf13/viper"
)

const (
	Env          = "GACHA_ENV"
	EnvTest      = "test"
	EnvPro       = "pro"
	ConfFileName = "starland"
)

// CommonConfig Common
type CommonConfig struct {
	Version   string
	IsDebug   bool
	LogLevel  string
	LogPath   string
	StartTime string
}

// ServerConf config struct
type ServerConf struct {
	Addr string
}

type RedisConf struct {
	RedisHost    string
	SignPrefix   string
	ImportPrefix string
	Password     string
}

// MySQLConf mysql配置
type MySQLConf struct {
	User               string `toml:"user" json:"user"`                                                                        // 用户
	Password           string `toml:"password" json:"password"`                                                                // 密码
	Host               string `toml:"host" json:"host"`                                                                        // 地址
	Port               int    `toml:"port" json:"port"`                                                                        // 端口
	Database           string `toml:"database" json:"database"`                                                                // 数据库
	MaxIdleConns       int    `toml:"max_idle_conns" mapstructure:"max_idle_conns" json:"max_idle_conns"`                      // 最大空闲连接数
	MaxOpenConns       int    `toml:"max_open_conns" mapstructure:"max_open_conns" json:"max_open_conns"`                      // 最大打开连接数
	MaxConnMaxLifetime int64  `toml:"max_conn_max_lifetime" mapstructure:"max_conn_max_lifetime" json:"max_conn_max_lifetime"` // 连接复用时间
	LogLevel           string `toml:"log_level" mapstructure:"log_level" json:"log_level"`                                     // 日志级别，枚举（info、warn、error和silent）
}

type EtcdConf struct {
	Host1 string `toml:"host1" json:"host1"` //节点1
	Host2 string `toml:"host2" json:"host2"` //节点2
	Host3 string `toml:"host3" json:"host3"` //节点3
}

type AwsConf struct {
	Bucket    string `toml:"bucket" json:"bucket"`
	AccessKey string `toml:"access_key" json:"access_key"`
	SecretKey string `toml:"secret_key" json:"secret_key"`
	Region    string `toml:"region" json:"region"`     //
	Endpoint  string `toml:"endpoint" json:"endpoint"` //
}

type JwtConf struct {
	Issuer         string `toml:"issuer" json:"issuer"`
	SecretKey      string `toml:"secret_key" json:"secret_key"`
	ExpirationTime int64  `toml:"expiration_time" json:"expiration_time"`
}

type TwitterConf struct {
	TweetsEnpoint   string `toml:"tweets_endpoint" json:"tweets_endpoint"`
	UserEndpoints   string `toml:"user_endpoint" json:"user_endpoint"`
	TaskRetweetedBy string `toml:"task_retweeted_by" json:"task_retweeted_by"`
	TaskLikingUsers string `toml:"task_liking_users" json:"task_liking_users"`
	TaskFollowers   string `toml:"task_followers" json:"task_followers"`
	AccessKey       string `toml:"access_key" json:"access_key"`
}

type EmailConf struct {
	Subject  string `toml:"subject" json:"subject"`
	Host     string `toml:"host" json:"host"`
	Port     int    `toml:"port" json:"port"`
	FromMail string `toml:"from_email" json:"from_email"`
	Password string `toml:"password" json:"password"`
	Content  string `toml:"content" json:"content"`
}

type CrontabConf struct {
	SixSecondsTask string `toml:"six_seconds_task" json:"six_seconds_task"`
}

type RangersConf struct {
	ChainId    int64  `toml:"chain_id" json:"chain_id"`
	RPC        string `toml:"rpc" json:"rpc"`
	MM         string `toml:"mm" json:"mm"`
	PrivateKey string `toml:"private_key" json:"private_key"`
	Ticket     string `toml:"ticket" json:"ticket"`
}

type ContractConf struct {
	Imx    string `toml:"imx" json:"imx"`
	Rpg    string `toml:"rpg" json:"rpg"`
	Usdt   string `toml:"usdt" json:"usdt"`
	Coupon string `toml:"coupon" json:"coupon"`
}

// Config ...
type Config struct {
	Common    *CommonConfig
	ServerC   *ServerConf
	MySQLC    *MySQLConf
	RedisC    *RedisConf
	EtcdC     *EtcdConf
	AwsC      *AwsConf
	JwtC      *JwtConf
	TwitterC  *TwitterConf
	EmailC    *EmailConf
	CrontabC  *CrontabConf
	RangersC  *RangersConf
	ContractC *ContractConf
}

// Conf ...
var Conf = &Config{}

// LoadConfig ...
func LoadConfig() {
	// init the new config params
	initConf()

	env := GetEnv()
	configFileName := ConfFileName
	if env != "" {
		configFileName += "_" + env + ".toml"
	} else {
		configFileName += ".toml"
	}
	contents, err := os.ReadFile(configFileName)
	if err != nil {
		log.Fatal("[FATAL] load "+configFileName+": ", err)
	}
	tbl, err := toml.Parse(contents)
	if err != nil {
		log.Fatal("[FATAL] parse starland.toml: ", err)
	}
	// parse common config
	parseCommon(tbl)
	// init log
	InitLogger()
	// parse server config
	parseServer(tbl)

	//parse mysql config
	parseMsq(tbl)

	//parse redis config
	parseReds(tbl)

	parseEtcd(tbl)

	parseAws(tbl)

	parseJwt(tbl)

	parseTwitter(tbl)

	//parse emial config
	parseEmail(tbl)

	parseCrontab(tbl)

	parseRangers(tbl)

	parseContract(tbl)
}

func initConf() {
	Conf = &Config{
		Common:    &CommonConfig{},
		ServerC:   &ServerConf{},
		MySQLC:    &MySQLConf{},
		RedisC:    &RedisConf{},
		EtcdC:     &EtcdConf{},
		AwsC:      &AwsConf{},
		JwtC:      &JwtConf{},
		TwitterC:  &TwitterConf{},
		EmailC:    &EmailConf{},
		CrontabC:  &CrontabConf{},
		RangersC:  &RangersConf{},
		ContractC: &ContractConf{},
	}
}

func parseCommon(tbl *ast.Table) {
	if val, ok := tbl.Fields["common"]; ok {
		subTbl, ok := val.(*ast.Table)
		if !ok {
			log.Fatalln("[FATAL] : ", subTbl)
		}

		err := toml.UnmarshalTable(subTbl, Conf.Common)
		if err != nil {
			log.Fatalln("[FATAL] parseCommon: ", err, subTbl)
		}
	}
}

func parseServer(tbl *ast.Table) {
	if val, ok := tbl.Fields["ser"]; ok {
		subTbl, ok := val.(*ast.Table)
		if !ok {
			log.Fatalln("[FATAL] : ", subTbl)
		}

		err := toml.UnmarshalTable(subTbl, Conf.ServerC)
		if err != nil {
			log.Fatalln("[FATAL] parseServer: ", err, subTbl)
		}
	}
}

func parseMsq(tbl *ast.Table) {
	if val, ok := tbl.Fields["mysql"]; ok {
		subTbl, ok := val.(*ast.Table)
		if !ok {
			log.Fatalln("[FATAL] : ", subTbl)
		}

		err := toml.UnmarshalTable(subTbl, Conf.MySQLC)
		if err != nil {
			log.Fatalln("[FATAL] parseMySQL: ", err, subTbl)
		}
	}
}

func parseReds(tbl *ast.Table) {
	if val, ok := tbl.Fields["redis"]; ok {
		subTbl, ok := val.(*ast.Table)
		if !ok {
			log.Fatalln("[FATAL] : ", subTbl)
		}

		err := toml.UnmarshalTable(subTbl, Conf.RedisC)
		if err != nil {
			log.Fatalln("[FATAL] parseReds: ", err, subTbl)
		}
	}
}

func parseEtcd(tbl *ast.Table) {
	if val, ok := tbl.Fields["etcd"]; ok {
		subTbl, ok := val.(*ast.Table)
		if !ok {
			log.Fatalln("[FATAL] : ", subTbl)
		}

		err := toml.UnmarshalTable(subTbl, Conf.EtcdC)
		if err != nil {
			log.Fatalln("[FATAL] parseEtcd: ", err, subTbl)
		}
	}
}

func parseAws(tbl *ast.Table) {
	if val, ok := tbl.Fields["aws"]; ok {
		subTbl, ok := val.(*ast.Table)
		if !ok {
			log.Fatalln("[FATAL] : ", subTbl)
		}

		err := toml.UnmarshalTable(subTbl, Conf.AwsC)
		if err != nil {
			log.Fatalln("[FATAL] parseAws: ", err, subTbl)
		}
	}
}

func parseJwt(tbl *ast.Table) {
	if val, ok := tbl.Fields["jwt"]; ok {
		subTbl, ok := val.(*ast.Table)
		if !ok {
			log.Fatalln("[FATAL] : ", subTbl)
		}

		err := toml.UnmarshalTable(subTbl, Conf.JwtC)
		if err != nil {
			log.Fatalln("[FATAL] parseAws: ", err, subTbl)
		}
	}
}

func parseTwitter(tbl *ast.Table) {
	if val, ok := tbl.Fields["twitter"]; ok {
		subTbl, ok := val.(*ast.Table)
		if !ok {
			log.Fatalln("[FATAL] : ", subTbl)
		}

		err := toml.UnmarshalTable(subTbl, Conf.TwitterC)
		if err != nil {
			log.Fatalln("[FATAL] parseAws: ", err, subTbl)
		}
	}
}

func parseEmail(tbl *ast.Table) {
	if val, ok := tbl.Fields["email"]; ok {
		subTbl, ok := val.(*ast.Table)
		if !ok {
			log.Fatalln("[FATAL] : ", subTbl)
		}

		err := toml.UnmarshalTable(subTbl, Conf.EmailC)
		if err != nil {
			log.Fatalln("[FATAL] parseEamil: ", err, subTbl)
		}
	}
}

func parseCrontab(tbl *ast.Table) {
	if val, ok := tbl.Fields["crontab"]; ok {
		subTbl, ok := val.(*ast.Table)
		if !ok {
			log.Fatalln("[FATAL] : ", subTbl)
		}

		err := toml.UnmarshalTable(subTbl, Conf.CrontabC)
		if err != nil {
			log.Fatalln("[FATAL] parse crontab: ", err, subTbl)
		}
	}
}

func parseRangers(tbl *ast.Table) {
	if val, ok := tbl.Fields["rangers"]; ok {
		subTbl, ok := val.(*ast.Table)
		if !ok {
			log.Fatalln("[FATAL] : ", subTbl)
		}

		err := toml.UnmarshalTable(subTbl, Conf.RangersC)
		if err != nil {
			log.Fatalln("[FATAL] parse rangers: ", err, subTbl)
		}
	}
}

func parseContract(tbl *ast.Table) {
	if val, ok := tbl.Fields["contract"]; ok {
		subTbl, ok := val.(*ast.Table)
		if !ok {
			log.Fatalln("[FATAL] : ", subTbl)
		}

		err := toml.UnmarshalTable(subTbl, Conf.ContractC)
		if err != nil {
			log.Fatalln("[FATAL] parse contract: ", err, subTbl)
		}
	}
}

func UnmarshaCOnfig(configFilePath string) (*Config, error) {
	viper.SetConfigFile(configFilePath)
	viper.SetConfigType("toml")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}
	config, err := DefaultConfig()
	if err != nil {
		return nil, err
	}

	if err := viper.Unmarshal(config); err != nil {
		return nil, err
	}
	return config, nil
}

func DefaultConfig() (*Config, error) {
	return &Config{}, nil
}

func GetEnv() string {
	env := os.Getenv(Env)
	log.Println(Env, " : ", env)
	return env
}
