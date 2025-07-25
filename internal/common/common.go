package common

import (
	"sync"

	"github.com/bwmarrin/snowflake"
)

type TokenKey string

const (
	BearerToken     TokenKey = "BearerToken"
	TokenInspectRes TokenKey = "TokenInspectRes"
)

var (
	snowOnce      sync.Once
	snowflakeNode *snowflake.Node
)

func init() {
	snowOnce.Do(func() {
		node, err := snowflake.NewNode(1)
		if err != nil {
			panic(err)
		}
		snowflakeNode = node
	})
}

func GetSnowflakeID() string {
	return snowflakeNode.Generate().String()
}
