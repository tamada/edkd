package main

//import "math"

func Example_help() {
	goMain([]string{"/some/path/of/edkd", "-h"})
	// Output:
	// edkd [OPTIONS...] [NUMBERs...|CSV]
	//　OPTIONS:
	//　　-H,--Hubeny           ヒュベニの公式で距離を算出．デフォルトではharversine公式を使用して算出する．
	//　　-x,--same-xcoord      経度が同一直線上に存在する場合の距離算出．ARGUMENTsは二地点の緯度を引数にする．
	//　　-y,--same-ycoord      緯度が同一直線上に存在する場合の距離算出．ARGUMENTsは二地点の経度を引数にする．
	//　　-r,--radian-method    弧度法を引数にする．デフォルトでは度数法を引数にする．
	//　　-h,--help             Usageを表示する．

	//　ARGUMENTS
	//　NUMBERs...            二地点の緯度経度を引数にする．デフォルトではポイント1経度，ポイント1緯度，ポイント2経度，ポイント2緯度．但しオプション-xの場合，ポイント1緯度，ポイント2緯度，オプション-yの場合，ポイント1経度，ポイント2経度
	//　CSV                   二地点の緯度経度がまとめられたcsvファイルを引数にする．csvファイルはosmで扱うため．出力ファイルは"edkd.csv"という名前で出力する．}
}
func Example_edkd() {
	goMain([]string{"/some/path/of/edkd", "139.74472", "35.65500", "140.09111", "36.10056"})
	// Output:
	// 58.55427749503559
}

func Example_edkd_Hubeny() {
	goMain([]string{"/some/path/of/edkd", "-H", "139.74472", "35.65500", "140.09111", "36.10056"})
	// Output:
	// 58.50245893181839
}

func Example_edkd_file() {
	goMain([]string{"/some/path/of/edkd", "./testdata/result_utf.csv"})
	// Output:
	//

}
