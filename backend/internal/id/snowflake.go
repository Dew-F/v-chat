package id

import (
	"sync"

	"github.com/bwmarrin/snowflake"
)

var (
	node *snowflake.Node
	once sync.Once
)

func Init(workerID int64) error {
	var err error

	once.Do(func() {
		node, err = snowflake.NewNode(workerID)
	})

	return err
}

func New() int64 {
	return node.Generate().Int64()
}
