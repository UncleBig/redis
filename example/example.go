package main

import (
	"charles/redisCli"

	"flag"
	"fmt"
	"os"

	"github.com/dlintw/goconf"
)

func main() {
	var err error
	conf_file := flag.String("config", "./config.ini", "set redis config file.")
	flag.Parse()

	l_conf, err := goconf.ReadConfigFile(*conf_file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "LoadConfiguration: Error: Could not open %q for reading: %s\n", conf_file, err)
		os.Exit(1)
	}
	err = redisCli.InitRedis(l_conf)
	if err != nil {
		fmt.Println(err)
	}
	defer redisCli.Cli.Close()
	//	v, err := c.Do("SET", "name", "red")
	//	if err != nil {
	//		fmt.Println(err)
	//		return
	//	}
	//	fmt.Println(v)
	var v string
	v, err = redisCli.String(redisCli.Cli.Do("GET", "name"))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(v)
	return
}
