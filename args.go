package main

import (
	"fmt"
	"os"

	flag "github.com/spf13/pflag"
)

type options struct {
	/*path     string
	number   bool
	nonblank bool
	squeeze  bool
	table    bool*/
	help   bool
	args   []string
	Hubeny bool
	samex  bool
	samey  bool
	radian bool
}

func ExistFile(path string) bool {
	stat, err := os.Stat(path)
	return err == nil && stat.Mode().IsRegular()
}

func parseArgs(args []string) (*options, error) {
	opts := &options{}
	flags := flag.NewFlagSet("edkd", flag.ContinueOnError)
	flags.Usage = func() { fmt.Println(helpMessage(args[0])) }
	/*flags.StringVarP(&opts.path, "path", "p", "", "ファイルパス．")
	flags.BoolVarP(&opts.number, "number", "n", false, "行番号を表示する．")
	flags.BoolVarP(&opts.nonblank, "number-nonblank", "b", false, "行番号を表示する．ただし空白行には付けない．")
	flags.BoolVarP(&opts.squeeze, "squeeze", "s", false, "連続した空行を1行にする．")
	flags.BoolVarP(&opts.table, "table", "t", false, "2つ目以降のファイルでは先頭行を無視する．")*/
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
