package wecom

import (
	"fmt"

	"github.com/ArtisanCloud/PowerSocialite/v3/src/providers"

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
		CallbackURL: config.callbackUrl, // 内部应用的场景回调设置
		OAuth:       config.oauth,       // 内部应用的app oauth url
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
func (c *Config) GetQrConnectURL(state string) (string, error) {
	c.work.OAuth.Provider.WithState(state)
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

// GetUserDetail 获取用户敏感信息
func (c *Config) GetUserDetail(userTicket string) (*weCom.ResponseGetUserDetail, error) {
	accessToken, err := c.work.AccessToken.GetToken(false)
	if err != nil {
		return nil, err
	}
	if accessToken.ErrCode != 0 {
		return nil, fmt.Errorf(accessToken.ErrMsg)
	}
	userDetail, err := c.work.OAuth.Provider.WithApiAccessToken(accessToken.AccessToken).GetUserDetail(userTicket)
	if err != nil {
		return nil, err
	}
	if userDetail.ErrCode != 0 {
		return nil, fmt.Errorf(userDetail.ErrMSG)
	}
	return userDetail, nil
}

// ContactFromCode 根据code获取企业用户信息 (注意：/user/getuserinfo为旧接口 auth/getuserinfo为新接口)
func (c *Config) ContactFromCode(code string) (*providers.User, error) {
	return c.work.OAuth.Provider.Detailed().ContactFromCode(code)
}

// GetUserInfo 获取访问用户身份
func (c *Config) GetUserInfo(code string) (*weCom.ResponseGetUserInfo, error) {
	userInfo, err := c.work.OAuth.Provider.GetUserInfo(code)
	if err != nil {
		return nil, err
	}
	if userInfo.ErrCode != 0 {
		return nil, fmt.Errorf(userInfo.ErrMSG)
	}
	return userInfo, nil
}

// GetUserByID 根据用户id获取用户信息
func (c *Config) GetUserByID(userID string) (*weCom.ResponseGetUserByID, error) {
	userByID, err := c.work.OAuth.Provider.GetUserByID(userID)
	if err != nil {
		return nil, err
	}
	if userByID.ErrCode != 0 {
		return nil, fmt.Errorf(userByID.ErrMSG)
	}
	return userByID, nil
}
