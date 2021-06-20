package main

import (
	"fmt"

	flag "github.com/spf13/pflag"
)

type options struct {
	args   []string
	help   bool
	Hubeny bool
	samex  bool
	samey  bool
	radian bool
}

func parseArgs(args []string) (*options, error) {
	opts := &options{}
	flags := flag.NewFlagSet("edkd", flag.ContinueOnError)
	flags.Usage = func() { fmt.Println(helpMessage(args[0])) }
	flags.BoolVarP(&opts.Hubeny, "Hubeny", "H", false, "ヒュベニの公式を使用します．")
	flags.BoolVarP(&opts.samex, "samex", "x", false, "緯度が同じ場合です．")
	flags.BoolVarP(&opts.samey, "samey", "y", false, "経度が同じ場合です．")
	flags.BoolVarP(&opts.radian, "radian", "r", false, "弧度法で受け取ります．")
	flags.BoolVarP(&opts.help, "help", "h", false, "このメッセージを出力します．")
	if err := flags.Parse(args); err != nil {
		return nil, err
	}
	opts.args = flags.Args()[1:]
	return opts, nil
}
