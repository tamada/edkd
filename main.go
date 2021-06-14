package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func helpMessage(originalProgramName string) string {
	name := filepath.Base(originalProgramName)
	return fmt.Sprintf(`%s [OPTIONS...] [NUMBERs...|CSV]
OPTIONS
     -H,--Hubeny           ヒュベニの公式で距離を算出．デフォルトではharversine公式を使用して算出する．
     -x,--same-xcoord      経度が同一直線上に存在する場合の距離算出．ARGUMENTsは二地点の緯度を引数にする．
     -y,--same-ycoord      緯度が同一直線上に存在する場合の距離算出．ARGUMENTsは二地点の経度を引数にする．
     -r,--radian-method    弧度法を引数にする．デフォルトでは度数法を引数にする．
     -h,--help             Usageを表示する．
ARGUMENTS
     NUMBERs...            二地点の緯度経度を引数にする．デフォルトではポイント1経度，ポイント1緯度，ポイント2経度，ポイント2緯度．但しオプション-xの場合，ポイント1緯度，ポイント2緯度，オプション-yの場合，ポイント1経度，ポイント2経度
     CSV                   二地点の緯度経度がまとめられたcsvファイルを引数にする．csvファイルはosmで扱うため．デフォルトでは出力ファイルは"ykgeo_output.csv"，変えたい場合は2つ目の引数で設定する.`, name)
}

func stin() {
	scanner := bufio.NewScanner(os.Stdin)

	// 標準入力から受け取ったテキストを出力
	for scanner.Scan() {
		fmt.Fprintf(os.Stdout, "%s\n", scanner.Text())
	}
}

func stout(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	rd := bufio.NewReader(file)
	for {
		s, err := rd.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Print(s)
	}
	file.Close()
}

func stout_loop(filename []string) {
	for _, file := range filename {
		stout(file)
	}
}

func perform(opts *options) int {
	if len(opts.args) < 1 {
		stin()
	} else {
		stout_loop(opts.args)
	}
	return 0
}

func goMain(args []string) int {
	opts, err := parseArgs(args)
	if err != nil {
		fmt.Printf("parsing args fail: %s\n", err.Error())
		fmt.Println(helpMessage(filepath.Base(args[0])))
		return 1
	}
	if opts.help {
		fmt.Println(helpMessage(filepath.Base(args[0])))
		return 0
	}
	return perform(opts)
}

func main() {
	status := goMain(os.Args)
	os.Exit(status)
}
