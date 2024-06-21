package redis

import (
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

type IpTest struct {
	Ip    string
	Limit int
	Ttl   int
}

type TokenTest struct {
	Token string
	Limit int
	Ttl   int
}

type RateLimiterTestSuite struct {
	suite.Suite
	testeIp1    IpTest
	testeIp2    IpTest
	testeToken1 TokenTest
	testeToken2 TokenTest
	redisClient *RedisClient
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(RateLimiterTestSuite))
}

func (suite *RateLimiterTestSuite) SetupTest() {
	suite.redisClient = NewRedisClient("localhost:6379", "", "redispass", 10)
	suite.testeIp1 = IpTest{Ip: "192.168.0.1", Limit: 10, Ttl: 1}
	suite.testeIp2 = IpTest{Ip: "133.101.101.10", Limit: 10, Ttl: 2}
	suite.testeToken1 = TokenTest{Token: "token1", Limit: 100, Ttl: 1}
	suite.testeToken2 = TokenTest{Token: "token2", Limit: 100, Ttl: 2}
}

func (suite *RateLimiterTestSuite) TestRateLimiter_AllowIp() {
	res, err := suite.redisClient.AllowIp(suite.testeIp1.Ip, suite.testeIp1.Limit, suite.testeIp1.Ttl)
	if err != nil {
		suite.Fail("Error on AllowIp: ", err.Error())
	}
	suite.True(res)
	time.Sleep(time.Millisecond * 1200)
}

func (suite *RateLimiterTestSuite) TestRateLimiter_AllowToken() {
	res, err := suite.redisClient.AllowToken(suite.testeToken1.Token, suite.testeToken1.Limit, suite.testeToken1.Ttl)
	if err != nil {
		suite.Fail("Error on AllowToken: ", err.Error())
	}
	suite.True(res)
	time.Sleep(time.Millisecond * 1200)
}

func (suite *RateLimiterTestSuite) TestRateLimiter_AllowIpLimit() {
	for i := 0; i < suite.testeIp1.Limit; i++ {
		res, err := suite.redisClient.AllowIp(suite.testeIp1.Ip, suite.testeIp1.Limit, suite.testeIp1.Ttl)
		if err != nil {
			suite.Fail("Error on AllowIp: ", err.Error())
		}
		suite.True(res)
	}
	res, err := suite.redisClient.AllowIp(suite.testeIp1.Ip, suite.testeIp1.Limit, suite.testeIp1.Ttl)
	if err != nil {
		suite.Fail("Error on AllowIp: ", err.Error())
	}
	suite.False(res)
	time.Sleep(time.Millisecond * 1200)
}

func (suite *RateLimiterTestSuite) TestRateLimiter_AllowTokenLimit() {
	for i := 0; i < suite.testeToken1.Limit; i++ {
		res, err := suite.redisClient.AllowToken(suite.testeToken1.Token, suite.testeToken1.Limit, suite.testeToken1.Ttl)
		if err != nil {
			suite.Fail("Error on AllowToken: ", err.Error())
		}
		suite.True(res)
	}
	res, err := suite.redisClient.AllowToken(suite.testeToken1.Token, suite.testeToken1.Limit, suite.testeToken1.Ttl)
	if err != nil {
		suite.Fail("Error on AllowToken: ", err.Error())
	}
	suite.False(res)
}
