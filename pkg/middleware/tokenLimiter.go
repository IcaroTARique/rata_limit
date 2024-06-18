package middleware

import (
	"github.com/IcaroTARique/rate_limit/configs"
	"github.com/IcaroTARique/rate_limit/internals/redis"
	"net/http"
)

func TokenRateLimit(next http.Handler, conf configs.Conf) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		redisClient := redis.NewRedisClient(conf.RedisAddr, "", conf.RedisPassword, conf.RedisDB)
		token := r.Header.Get("API_KEY")

		allowed, err := redisClient.AllowToken(token, conf.LimitReqToken, conf.TimeLimitToken)
		if err != nil || !allowed {
			http.Error(
				w,
				"you have reached the maximum number of requests or actions allowed for this API_KEY within a certain time frame",
				http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	})
}
