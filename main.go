package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"path/filepath"
	"strconv"
)

func helpMessage(originalProgramName string) string {
	name := filepath.Base(originalProgramName)
	return fmt.Sprintf(`%s [OPTIONS...] [NUMBERs...|CSV]
	　OPTIONS:  
	　　-H,--Hubeny           ヒュベニの公式で距離を算出．デフォルトではharversine公式を使用して算出する．  
	　　-x,--same-xcoord      経度が同一直線上に存在する場合の距離算出．ARGUMENTsは二地点の緯度を引数にする．  
	　　-y,--same-ycoord      緯度が同一直線上に存在する場合の距離算出．ARGUMENTsは二地点の経度を引数にする．  
	　　-r,--radian-method    弧度法を引数にする．デフォルトでは度数法を引数にする．  
	　　-h,--help             Usageを表示する．  
		
	　ARGUMENTS   
	　　NUMBERs...            二地点の緯度経度を引数にする．デフォルトではポイント1経度，ポイント1緯度，ポイント2経度，ポイント2緯度．但しオプション-xの場合，ポイント1緯度，ポイント2緯度，オプション-yの場合，ポイント1経度，ポイント2経度  
	　　CSV                   二地点の緯度経度がまとめられたcsvファイルを引数にする．csvファイルはosmで扱うため．出力ファイルは"edkd.csv"という名前で出力する．`, name)
}

func stout(filename string, opts *options) {
	var (
		kd2   float64
		ed2   float64
		kdstr string
		edstr string
		dis   float64
	)

	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	rd := csv.NewReader(bufio.NewReader(file))

	file, err = os.Create(`edkd.csv`)
	if err != nil {
		log.Fatal(err)
	}

	writer := csv.NewWriter(file)

	for {
		record, err := rd.Read()
		if err == io.EOF {
			break
		}

		kd1, _ := strconv.ParseFloat(record[2], 64)
		ed1, _ := strconv.ParseFloat(record[3], 64)
		if !opts.radian {
			kd1 = degree_radian(kd1)
			ed1 = degree_radian(ed1)
		}

		if !((kdstr != "0") && (edstr != "0") && (kd2 == 0) && (ed2 == 0)) {
			if opts.Hubeny {
				dis = Hubeny(kd1, ed1, kd2, ed2)
			} else {
				dis = harversine(kd1, ed1, kd2, ed2)
			}
			disstr := strconv.FormatFloat(dis, 'f', 10, 64)
			writer.Write([]string{disstr})
		}
		kd2 = kd1
		ed2 = ed1
		kdstr = record[2]
		edstr = record[3]
	}
	writer.Flush()
	file.Close()
}

func stout_loop(filename []string, opts *options) {
	for _, file := range filename {
		stout(file, opts)
	}
}

func harversine(kd1 float64, ed1 float64, kd2 float64, ed2 float64) float64 {

	abx := math.Abs(kd1 - kd2)
	aby := math.Abs(ed1 - ed2)
	dis := 0.0

	dis = math.Pow(math.Sin(aby/2), 2) + math.Cos(ed1)*math.Cos(ed2)*math.Pow(math.Sin(abx/2), 2)
	dis = 2 * math.Asin(math.Sqrt(dis)) * 6371.007177356707

	return dis

}

func Hubeny(kd1 float64, ed1 float64, kd2 float64, ed2 float64) float64 {

	abx := math.Abs(kd1 - kd2)
	aby := math.Abs(ed1 - ed2)
	p := (ed1 + ed2) / 2

	rx := 6378.137000000
	ry := 6356.752314245
	e := math.Sqrt((math.Pow(rx, 2) - math.Pow(ry, 2)) / math.Pow(rx, 2))
	w := math.Sqrt(1 - math.Pow(e, 2)*math.Pow(math.Sin(p), 2))
	m := rx * (1 - math.Pow(e, 2)) / math.Pow(w, 3)
	n := rx / w
	dis := math.Sqrt(math.Pow(aby*m, 2) + math.Pow(abx*n*math.Cos(p), 2))

	return dis
}
func degree_radian(num float64) float64 {
	return num * math.Pi / 180.0
}

func perform(opts *options) int {

	var (
		kd1 float64
		ed1 float64
		kd2 float64
		ed2 float64
		dis float64
	)

	if len(opts.args) == 1 {
		stout_loop(opts.args, opts)
		return 0

	} else if len(opts.args) == 2 && (opts.samex || opts.samey) {
		if !opts.samey && opts.samex {
			kd1, _ = strconv.ParseFloat(opts.args[0], 64)
			kd2, _ = strconv.ParseFloat(opts.args[1], 64)
		} else if !opts.samex && opts.samey {
			ed1, _ = strconv.ParseFloat(opts.args[0], 64)
			ed2, _ = strconv.ParseFloat(opts.args[1], 64)
		}

	} else if len(opts.args) == 4 {

		kd1, _ = strconv.ParseFloat(opts.args[0], 64)
		ed1, _ = strconv.ParseFloat(opts.args[1], 64)
		kd2, _ = strconv.ParseFloat(opts.args[2], 64)
		ed2, _ = strconv.ParseFloat(opts.args[3], 64)

	} else {
		return 1
	}
	if !opts.radian {
		kd1 = degree_radian(kd1)
		ed1 = degree_radian(ed1)
		kd2 = degree_radian(kd2)
		ed2 = degree_radian(ed2)
	}

	if opts.Hubeny {
		dis = Hubeny(kd1, ed1, kd2, ed2)
	} else {
		dis = harversine(kd1, ed1, kd2, ed2)
	}
	fmt.Printf("%g\n", dis)

	return 0
}

func goMain(args []string) int {
	opts, err := parseArgs(args)
	if err != nil {
		fmt.Printf("parsing args fail: %s\n", err.Error())
		fmt.Println(helpMessage(filepath.Base(args[0])))
		return 1
	}
	if opts.help || len(opts.args) < 1 {
		fmt.Println(helpMessage(filepath.Base(args[0])))
		return 0
	}
	return perform(opts)
}

func main() {
	status := goMain(os.Args)
	os.Exit(status)
}
