package common

import (
	"log"
	"strings"

	"github.com/spf13/viper"

	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

type Nacos_Info struct {
	NacosIp     string
	NacosPort   uint64
	NacosDataId string
	NacosGroup  string
}

type DiskSecret struct {
	MailPwd string
	Bucket  struct {
		TencentSecretId  string
		TencentSecretKey string
	}
}

var Cfg *DiskSecret

func GetConfig() {
	n := &Nacos_Info{
		NacosIp:     "",
		NacosPort:   8848,
		NacosDataId: "cherry_disk",
		NacosGroup:  "disk",
	}

	clientConfig := constant.ClientConfig{
		TimeoutMs:           5000,
		NamespaceId:         "",
		CacheDir:            "cache",
		NotLoadCacheAtStart: false,
		LogDir:              "log",
		LogLevel:            "debug",
	}

	serverConfigs := []constant.ServerConfig{
		{
			IpAddr:      n.NacosIp,
			ContextPath: "/nacos",
			Port:        n.NacosPort,
			Scheme:      "http",
		},
	}

	configClient, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &clientConfig,
			ServerConfigs: serverConfigs,
		},
	)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	config, err := configClient.GetConfig(vo.ConfigParam{
		DataId: n.NacosDataId,
		Group:  n.NacosGroup,
	})

	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	defaultConfg := viper.New()
	defaultConfg.SetConfigType("yaml")
	defaultConfg.ReadConfig(strings.NewReader(config))

	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	Cfg = new(DiskSecret)
	err = defaultConfg.Unmarshal(&Cfg)

	if err != nil {
		log.Fatal(err)
		panic(err)
	}
}
