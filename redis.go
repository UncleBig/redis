package redisCli

import (
	"errors"
	"fmt"

	"github.com/dlintw/goconf"
	"github.com/garyburd/redigo/redis"
)

var (
	Cli redis.Conn
)

func InitRedis(conf *goconf.ConfigFile) (err error) {

	var dbHost, _ = conf.GetString("redis", "db_host")
	var dbPass, _ = conf.GetString("redis", "db_pass")
	var network, _ = conf.GetString("redis", "network")
	if dbPass != "" {
		options := redis.DialPassword(dbPass)
		Cli, err = redis.Dial(network, dbHost, options)
		if err != nil {
			err = errors.New(fmt.Sprintf("redisCli.Dail(%s, %s,%s), error: %s", network, dbHost, dbPass, err))
			return
		}
	} else {
		Cli, err = redis.Dial(network, dbHost)
		if err != nil {
			err = errors.New(fmt.Sprintf("redisCli.Dail(%s, %s,%s), error: %s", network, dbHost, err))
			return
		}
	}
	return
}

func String(reply interface{}, err1 error) (value string, err2 error) {
	value, err2 = redis.String(reply, err1)
	return
}
