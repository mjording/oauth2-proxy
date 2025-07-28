package sessions

import (
	"fmt"

	"github.com/mjording/oauth2-proxy/v7/pkg/apis/options"
	"github.com/mjording/oauth2-proxy/v7/pkg/apis/sessions"
	"github.com/mjording/oauth2-proxy/v7/pkg/sessions/cookie"
	"github.com/mjording/oauth2-proxy/v7/pkg/sessions/redis"
)

// NewSessionStore creates a SessionStore from the provided configuration
func NewSessionStore(opts *options.SessionOptions, cookieOpts *options.Cookie) (sessions.SessionStore, error) {
	switch opts.Type {
	case options.CookieSessionStoreType:
		return cookie.NewCookieSessionStore(opts, cookieOpts)
	case options.RedisSessionStoreType:
		return redis.NewRedisSessionStore(opts, cookieOpts)
	default:
		return nil, fmt.Errorf("unknown session store type '%s'", opts.Type)
	}
}
