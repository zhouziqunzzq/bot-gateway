package main

var config = globalConfig{}

type globalConfig struct {
	BufferSize      uint   `toml:"buffer_size"`
	ChannelLifeTime string `toml:"channel_lifetime"`
	GCInterval      string `toml:"garbage_collection_interval"`
	LogLevel        string `toml:"log_level"`
	PluginDir       string `toml:"plugin_dir"`
	PluginConfDir   string `toml:"plugin_conf_dir"`
}
