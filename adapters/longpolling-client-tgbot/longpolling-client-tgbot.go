package main

import (
	"github.com/BurntSushi/toml"
	"github.com/projectriri/bot-gateway/router"
	"github.com/projectriri/bot-gateway/types"
	"github.com/projectriri/bot-gateway/utils"
	log "github.com/sirupsen/logrus"
)

var (
	BuildTag      string
	BuildDate     string
	GitCommitSHA1 string
	GitTag        string
)

type Plugin struct{}

var manifest = types.Manifest{
	BasicInfo: types.BasicInfo{
		Name:    "longpolling-client-tgbot",
		Author:  "Project Riri Staff",
		Version: "v0.1",
		License: "MIT",
		URL:     "https://github.com/projectriri/bot-gateway/adapters/longpolling-client-tgbot",
	},
	BuildInfo: types.BuildInfo{
		BuildTag:      BuildTag,
		BuildDate:     BuildDate,
		GitCommitSHA1: GitCommitSHA1,
		GitTag:        GitTag,
	},
}

func (p *Plugin) GetManifest() types.Manifest {
	return manifest
}

func (p *Plugin) Init(filename string, configPath string) {
	// load toml config
	_, err := toml.DecodeFile(configPath+"/"+filename+".toml", &config)
	if err != nil {
		panic(err)
	}
	updateConfig.Limit = config.Limit
	updateConfig.Timeout = config.Timeout
}

func (p *Plugin) Start() {
	log.Infof("[longpolling-client-tgbot] registering producer channel %v", config.ChannelUUID)
	pc := router.RegisterProducerChannel(config.ChannelUUID, false)
	log.Infof("[longpolling-client-tgbot] registered producer channel %v", pc.UUID)
	log.Info("[longpolling-client-tgbot] start polling from Telegram-Bot-API via LongPolling")
	for {
		data := getUpdates()
		if data != nil {
			pc.Produce(types.Packet{
				Head: types.Head{
					From: config.AdaptorName,
					UUID: utils.GenerateUUID(),
					Format: types.Format{
						API:      "Telegram-Bot-API",
						Version:  "latest",
						Method:   "Update",
						Protocol: "HTTP",
					},
				},
				Body: data,
			})
		}
	}
}

var PluginInstance types.Adapter = &Plugin{}
