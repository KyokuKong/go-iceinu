package models

type CommandHelp struct {
	IsShown     bool
	CommandName string
	Usage       string
	Description string
	Flags       map[string]string // flag 和 介绍的键值对
}
