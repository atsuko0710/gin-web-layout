package snowflake

import (
	sf "github.com/bwmarrin/snowflake"
)

// var node *sf.Node

func GetId() int64 {
	node, _ := sf.NewNode(64)
	return node.Generate().Int64()
}