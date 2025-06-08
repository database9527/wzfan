package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port string `mapstructure:"port"`
	} `mapstructure:"server"`
	Redis struct {
		Host         string `mapstructure:"host"`
		Port         string `mapstructure:"port"`
		Password     string `mapstructure:"password"`
		DB           int    `mapstructure:"db"`
		PageCacheDB  int    `mapstructure:"page_cache_db"`
		LocalDataDB  int    `mapstructure:"local_data_db"`
		ArticleDB    int    `mapstructure:"article_db"`
	} `mapstructure:"redis"`
	PageCache struct {
		TTLSeconds int  `mapstructure:"ttl_seconds"`
		Enabled    bool `mapstructure:"enabled"`
	} `mapstructure:"page_cache"`
	Site struct {
		Title       string `mapstructure:"title"`
		Keywords    string `mapstructure:"keywords"`
		Description string `mapstructure:"description"`
	} `mapstructure:"site"`
	AccessControl AccessControlConfig `mapstructure:"access_control"`
}

// AccessControlConfig holds settings for request filtering and access rules.
type AccessControlConfig struct {
	Enabled      bool                `mapstructure:"enabled"`
	PreviewParam string              `mapstructure:"preview_param"`
	BlacklistMode BlacklistModeConfig `mapstructure:"blacklist_mode"`
	SpiderMode   SpiderModeConfig    `mapstructure:"spider_mode"`
	// IntegrityChecks IntegrityChecksConfig `mapstructure:"integrity_checks"` // Optional
	// Constants       ConstantsConfig       `mapstructure:"constants"`        // Optional
}

// BlacklistModeConfig defines IP blacklisting rules.
type BlacklistModeConfig struct {
	Enabled bool     `mapstructure:"enabled"`
	IPs     []string `mapstructure:"ips"` // List of regex patterns for IPs
	Page    string   `mapstructure:"page"`  // Path to the HTML page to show blacklisted users
}

// SpiderModeConfig defines rules for identifying and handling spiders/bots.
type SpiderModeConfig struct {
	Enabled            bool     `mapstructure:"enabled"`
	AllowedSpiders     []string `mapstructure:"allowed_spiders"`      // Keywords to identify allowed spiders (e.g., "Googlebot")
	VerifiedIPs        []string `mapstructure:"verified_ips"`         // Regex list for IPs of verified spiders
	VerifyByHostname   bool     `mapstructure:"verify_by_hostname"`   // Whether to perform reverse DNS lookup for verification
	DefaultPcPage      string   `mapstructure:"default_pc_page"`      // Page to show non-verified PC spiders
	DefaultMobPage     string   `mapstructure:"default_mob_page"`     // Page to show non-verified mobile spiders
	MobileKeywords     []string `mapstructure:"mobile_keywords"`      // Keywords to detect mobile user agents
}

// IntegrityChecksConfig (Optional)
// type IntegrityChecksConfig struct {
// 	ConfigMd5 string `mapstructure:"config_md5"`
// 	TplMd5    string `mapstructure:"tpl_md5"`
// }

// ConstantsConfig (Optional)
// type ConstantsConfig struct {
// 	SiteID string `mapstructure:"SITE_ID"`
// }

func LoadConfig() (*Config, error) {
	viper.AddConfigPath("./config") // Look for config in the "config" directory
	viper.SetConfigName("config")   // Name of config file (without extension)
	viper.SetConfigType("yaml")     // Type of config file

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("Error reading config file: %s", err)
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("Unable to decode into struct: %s", err)
	}

	return &config, nil
}
