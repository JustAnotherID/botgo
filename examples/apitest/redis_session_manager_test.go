package apitest

import (
	"log"
	"testing"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/sealdice/botgo"
	"github.com/sealdice/botgo/dto"
	"github.com/sealdice/botgo/event"
	"github.com/sealdice/botgo/sessions/remote"
)

func Test_redisSessionManager(t *testing.T) {
	ws, err := api.WS(ctx, nil, "")
	log.Printf("%+v, err:%v", ws, err)

	conn := redis.NewClient(
		&redis.Options{
			Addr:         "localhost:6379",
			DialTimeout:  800 * time.Millisecond,
			ReadTimeout:  3 * time.Second,
			WriteTimeout: 3 * time.Second,
		},
	)

	botgo.SetSessionManager(remote.New(conn, remote.WithClusterKey("abcccc")))

	t.Run(
		"at message", func(t *testing.T) {
			var message event.ATMessageEventHandler = func(event *dto.WSPayload, data *dto.WSATMessageData) error {
				log.Println(event, data)
				return nil
			}
			intent := event.RegisterHandlers(message)
			ws.Shards = 2
			botgo.NewSessionManager().Start(ws, botToken, &intent)
		},
	)
}
