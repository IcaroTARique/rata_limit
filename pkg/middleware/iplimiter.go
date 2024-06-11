package middleware

import (
	"github.com/IcaroTARique/rate_limit/configs"
	"github.com/IcaroTARique/rate_limit/internals/redis"
	"net/http"
	"strings"
	"time"
)

type RateLimiterIp struct {
	ammount int
	time    time.Time
}

func IPRateLimit(next http.Handler, conf configs.Conf) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		redisClient := redis.NewRedisClient(conf.RedisAddr, "", conf.RedisPassword, conf.RedisDB)

		ip := strings.Split(r.RemoteAddr, ":")[0]
		allowed, err := redisClient.AllowIp(ip)
		if err != nil || !allowed {
			http.Error(w, http.StatusText(http.StatusTooManyRequests), http.StatusTooManyRequests)
			return
		}
		next.ServeHTTP(w, r)
	})
}
