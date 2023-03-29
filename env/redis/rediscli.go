package redis

import (
	"context"
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/MrBruin-ctrl/runcli/internal/driver"
	"github.com/MrBruin-ctrl/runcli/internal/registry"
	"github.com/MrBruin-ctrl/runcli/internal/utils"
	"github.com/c-bata/go-prompt"
	"github.com/gomodule/redigo/redis"
	"strings"
	"time"
)

type RedisCliDriver struct {
	RedisConf
	conn redis.Conn
}

func (cli *RedisCliDriver) InitConfig() {
	InitConfig()
}

func (cli *RedisCliDriver) Completer(d prompt.Document) []prompt.Suggest {
	currentArg := d.GetWordBeforeCursor()
	if currentArg == "" {
		return []prompt.Suggest{}
	}
	args := strings.Split(d.Text, " ")
	if len(args) > 1 {
		return []prompt.Suggest{}
	}
	return prompt.FilterHasPrefix(redisCMDSuggest, args[0], true)
}

func (cli *RedisCliDriver) LivePrefix() string {
	return cli.Host + "@" + cli.Name + ">>"
}

func (cli *RedisCliDriver) getConn() redis.Conn {
	return cli.conn
}

func (cli *RedisCliDriver) SurveyConfig() {
	configs := GetConfigIfExist()
	if len(configs) != 0 {
		options := make([]string, 0)
		for key, _ := range configs {
			options = append(options, key)
		}
		elems := "new connect"
		options = append(options, elems)
		ans := ""
		selects := &survey.Select{
			Message: "choose a connect",
			Options: options,
			Description: func(value string, index int) string {
				conf, ok := configs[value]
				if ok {
					return fmt.Sprintf("address[%s:%s]", conf.Host, conf.Port)
				}
				return ""
			},
		}
		utils.AskOne(selects, &ans)
		if ans != elems {
			cli.RedisConf = configs[ans]
			return
		}
	}
	redisConfig := New()
	utils.Ask(BaseQs, &redisConfig)
	//保存配置，如果名称一样，则覆盖配置
	SaveConfig(redisConfig)
	cli.RedisConf = redisConfig
}

func (cli *RedisCliDriver) Exit() {
	cli.getConn().Close()
}

func (cli *RedisCliDriver) Conn(ctx context.Context) error {
	conn, err := redis.DialContext(ctx, "tcp", cli.Host+":"+cli.Port,
		redis.DialPassword(cli.Password),
		redis.DialConnectTimeout(60*time.Second))
	if err != nil {
		return err
	}
	cli.conn = conn
	return nil
}

func (cli *RedisCliDriver) Executor(cmd string) {
	splitCmd := strings.Split(strings.TrimSpace(cmd), " ")
	if len(splitCmd) == 0 {
		fmt.Println("")
		return
	}
	conn := cli.getConn()
	strings := splitCmd[1:]

	args := make([]interface{}, 0)
	for _, arg := range strings {
		args = append(args, arg)
	}
	s, err := redis.Values(conn.Do(splitCmd[0], args...))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(s)
}

func (cli *RedisCliDriver) CliName() string {
	return "redis"
}

func init() {
	registry.RegisterCli(func() driver.CliDriver {
		return &RedisCliDriver{}
	})
}
