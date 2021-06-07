package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func helpMessage(originalProgramName string) string {
	programName := filepath.Base(originalProgramName)
	return fmt.Sprintf(`%s [OPTIONS...] [NUMBERs...|CSV]  
　OPTIONS  
　　-H,--Hubeny           ヒュベニの公式で距離を算出．デフォルトではharversine公式を使用して算出する．  
　　-x,--same-xcoord      経度が同一直線上に存在する場合の距離算出．ARGUMENTsは二地点の緯度を引数にする．  
　　-y,--same-ycoord      緯度が同一直線上に存在する場合の距離算出．ARGUMENTsは二地点の経度を引数にする．  
　　-r,--radian-method    弧度法を引数にする．デフォルトでは度数法を引数にする．  
　　-h,--help             Usageを表示する．  
    
　ARGUMENTS   
　　NUMBERs...            二地点の緯度経度を引数にする．デフォルトではポイント1経度，ポイント1緯度，ポイント2経度，ポイント2緯度．但しオプション-xの場合，ポイント1緯度，ポイント2緯度，オプション-yの場合，ポイント1経度，ポイント2経度  
　　CSV                   二地点の緯度経度がまとめられたcsvファイルを引数にする．csvファイルはosmで扱うため．デフォルトでは出力ファイルは"ykgeo_output.csv"，変えたい場合は2つ目の引数で設定する．  `, programName, programName)
}

func goMain(args []string) int {
	fmt.Println(helpMessage(args[0]))
	return 0
}

func main() {
	status := goMain(os.Args)
	os.Exit(status)
}
