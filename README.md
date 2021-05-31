[![CI](https://github.com/YKatsuy/edkd/actions/workflows/blank.yml/badge.svg)](https://github.com/YKatsuy/edkd/actions/workflows/blank.yml)
[![GitHub](https://img.shields.io/github/license/YKatsuy/edkd)](https://img.shields.io/github/license/YKatsuy/edkd)
[![Coverage Status](https://coveralls.io/repos/github/YKatsuy/edkd/badge.svg?branch=main)](https://coveralls.io/github/YKatsuy/edkd?branch=main)
[![CSV Validation](https://csvlint.io/validation/60b3b4ac053bd1000400001e.svg)](https://csvlint.io/validation/60b3b4ac053bd1000400001e)

# edkd
二地点の緯度経度から距離の算出を行う．

## icon
![edkd](edkd.svg)
所得サイト:https://freesvg.org/restaurant-map-location

## Description
私は地理情報を用いた研究を行っている．その中で二点間の緯度経度から距離の算出を行うことがあった．しかし，距離の算出方法は一つではない．私が行った研究内では，harversine公式を使用し，距離を算出した．  
本ソフトウェアでは二点間の緯度経度から距離を算出するが，harversine公式以外にヒュベニの公式での距離算出を行えるソフトウェアの開発を行う．またosm(Open Street Map)データを想定したファイルの読み込み，距離を算出するファイルを出力するなどの機能を加える．

## Usage
```
edkd [OPTIONS...] [NUMBERs...|CSV]  
　OPTIONS:  
　　-H,--Hubeny           ヒュベニの公式で距離を算出．デフォルトではharversine公式を使用して算出する．  
　　-x,--same-xcoord      経度が同一直線上に存在する場合の距離算出．ARGUMENTsは二地点の緯度を引数にする．  
　　-y,--same-ycoord      緯度が同一直線上に存在する場合の距離算出．ARGUMENTsは二地点の経度を引数にする．  
　　-r,--radian-method    弧度法を引数にする．デフォルトでは度数法を引数にする．  
　　-h,--help             Usageを表示する．  
    
　ARGUMENTS   
　　NUMBERs...            二地点の緯度経度を引数にする．デフォルトではポイント1経度，ポイント1緯度，ポイント2経度，ポイント2緯度．但しオプション-xの場合，ポイント1緯度，ポイント2緯度，オプション-yの場合，ポイント1経度，ポイント2経度  
　　CSV                   二地点の緯度経度がまとめられたcsvファイルを引数にする．csvファイルはosmで扱うため．デフォルトでは出力ファイルは"ykgeo_output.csv"，変えたい場合は2つ目の引数で設定する．  
