package internals

type StorageClient interface {
	AllowToken(token string, limit, ttl int) (bool, error)
	AllowIp(ip string, limit, ttl int) (bool, error)
}
