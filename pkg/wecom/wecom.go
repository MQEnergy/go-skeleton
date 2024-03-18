package wecom

import (
	"fmt"

	"github.com/ArtisanCloud/PowerSocialite/v3/src/response/weCom"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work"
)

type Config struct {
	corpID      string
	agentID     int
	secret      string
	callbackUrl string
	oauth       work.OAuth
	httpDebug   bool
	cache       kernel.CacheInterface
	work        *work.Work
}

type Options interface {
	apply(cfg *Config)
}

type OptionFunc func(cfg *Config)

func New(options ...Options) (*Config, error) { // 生成配置
	config := &Config{}
	for _, option := range options {
		option.apply(config)
	}
	newWork, err := work.NewWork(&work.UserConfig{
		CorpID:      config.corpID,
		AgentID:     config.agentID,
		Secret:      config.secret,
		CallbackURL: config.callbackUrl,
		OAuth:       config.oauth,
		HttpDebug:   config.httpDebug,
		Cache:       config.cache,
	})
	if err != nil {
		return nil, err
	}
	config.work = newWork
	return config, nil
}

func (o OptionFunc) apply(cfg *Config) {
	o(cfg)
}

func WithCorpID(corpID string) OptionFunc {
	return func(cfg *Config) {
		cfg.corpID = corpID
	}
}

func WithAgentID(agentID int) OptionFunc {
	return func(cfg *Config) {
		cfg.agentID = agentID
	}
}

func WithSecret(secret string) OptionFunc {
	return func(cfg *Config) {
		cfg.secret = secret
	}
}

func WithCallbackURL(callbackURL string) OptionFunc {
	return func(cfg *Config) {
		cfg.callbackUrl = callbackURL
	}
}

func WithOAuth(oauth work.OAuth) OptionFunc {
	return func(cfg *Config) {
		cfg.oauth = oauth
	}
}

func WithHttpDebug(debug bool) OptionFunc {
	return func(cfg *Config) {
		cfg.httpDebug = debug
	}
}

func WithCache(cache kernel.CacheInterface) OptionFunc {
	return func(cfg *Config) {
		cfg.cache = cache
	}
}

// GetQrConnectURL 扫码授权登录
func (c *Config) GetQrConnectURL() (string, error) {
	c.work.OAuth.Provider.WithRedirectURL(c.callbackUrl)
	return c.work.OAuth.Provider.GetQrConnectURL()
}

// GetAuthURL 网页授权登录
func (c *Config) GetAuthURL() (string, error) {
	c.work.OAuth.Provider.WithRedirectURL(c.callbackUrl)
	return c.work.OAuth.Provider.GetAuthURL()
}

// GetOAuthUrl 获取授权url地址
func (c *Config) GetOAuthUrl() string {
	c.work.OAuth.Provider.WithRedirectURL(c.callbackUrl)
	return c.work.OAuth.Provider.GetOAuthURL()
}

// GetUserInfo 获取用户信息
func (c *Config) GetUserDetail(code string) (*weCom.ResponseGetUserDetail, error) {
	userInfo, err := c.work.OAuth.Provider.GetUserInfo(code)
	if err != nil {
		return nil, err
	}
	if userInfo.ErrCode != 0 {
		return nil, fmt.Errorf(userInfo.ErrMSG)
	}
	detail, err := c.work.OAuth.Provider.GetUserDetail(userInfo.UserTicket)
	if err != nil {
		return nil, err
	}
	if detail.ErrCode != 0 {
		return nil, fmt.Errorf(detail.ErrMSG)
	}
	return detail, nil
}
