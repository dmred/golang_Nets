package parser

import (
	"github.com/chzyer/readline"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	App = kingpin.New("SSH-client console application", "[conn] [close] [ssh] [cmd] [exit]")

	Debug = App.Command("debug", "Debug connection")

	AddConn         = App.Command("conn", "Adds a new connection to ssh-server")
	AddConnHost     = AddConn.Arg("Host", "Address to SSH-server").Required().String()
	AddConnName     = AddConn.Arg("Name", "Your login").Default("dmred").String()
	AddConnWithPass = AddConn.Flag("pass", "Should we use password instead ssh keyes").Default("true").Bool()
	AddConnPass     = AddConn.Arg("Pass", "Password to establish connection").Default("").String()

	CloseConn     = App.Command("close", "Closes all existing connections")
	CloseConnHost = CloseConn.Arg("Host", "Address to SSH-server").Required().String()

	RecreateSsh = App.Command("ssh", "Recreates existing SSH key")

	StartTransmitting = App.Command("cmd",
		"Exit mode")

	Exit = App.Command("exit", "Exit application")

	completer *readline.PrefixCompleter
)
