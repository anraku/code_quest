// CODEQUESTの毒沼の試練をクリアするために作ったプログラム
package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var (
	traceMap [][]int
	dokunuma [][]string
	HP       int
)

func main() {
	// ファイルから呪文を読み込む
	v, err := input("./map.txt")
	if err != nil {
		fmt.Printf("ファイル入力でエラーになりました: %s", err)
		return
	}
	// Startの地点を探す
	x, y, _ := searchWord(v, "S")

	//グローバル変数の初期化
	dokunuma = v
	HP = 36
	traceMap = initMap(len(dokunuma), len(dokunuma[0]))
	// 探索開始
	walk(x, y)
}

// 迷路を再帰的に探索する
func walk(x int, y int) {
	if x < 0 || y < 0 || len(traceMap) <= y || len(traceMap[0]) <= x || traceMap[y][x] == 1 {
		return
	}
	switch dokunuma[y][x] {
	case "-1":
		HP--
		if HP < 33 {
			return
		}
	case "1":
		HP++
	case "L":
		return
	case "G":
		traceMap[y][x] = 1
		battle()
		traceMap[y][x] = 0
		return
	}

	traceMap[y][x] = 1
	walk(x+1, y)
	back(x+1, y)

	walk(x, y+1)
	back(x, y+1)

	walk(x-1, y)
	back(x-1, y)

	walk(x, y-1)
	back(x, y-1)
	traceMap[y][x] = 0
}

// 一歩進んだ先から一歩戻る
func back(x int, y int) {
	if x < 0 || y < 0 || len(traceMap) <= y || len(traceMap[0]) <= x || traceMap[y][x] == 1 {
		return
	}
	// 一歩先の毒沼や回復マスを踏んだ分を戻す
	switch dokunuma[y][x] {
	case "-1":
		HP++
	case "1":
		HP--
	}
}

// 魔王とバトル
func battle() {
	if HP >= 50 {
		win()
	}
	if HP >= 40 {
		output()
	}
}

func win() {
	fmt.Println("問題の答え: ")
	for y := range traceMap {
		fmt.Println(traceMap[y])
	}
	os.Exit(0)
}

func output() {
	fmt.Printf("ゴールしました HP: %v\n ", HP)
	for y := range traceMap {
		fmt.Println(traceMap[y])
	}
}

// ファイルからデータ読み込み
func input(filepath string) ([][]string, error) {
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	line_arr := strings.Split(string(data), "\n")
	line_arr = line_arr[:len(line_arr)-1] // 最後に空の配列が入ってしまうので削除
	var result [][]string
	for _, line := range line_arr {
		v_arr := strings.Split(line, ",")
		result = append(result, v_arr)
	}
	return result, nil
}

// mからsに指定した文字列を検索する
// 最初に見つけた文字列のx, yの要素を返す
func searchWord(m [][]string, s string) (int, int, error) {
	for y := range m {
		for x := range m[y] {
			if m[y][x] == s {
				return x, y, nil
			}
		}
	}
	return 0, 0, errors.New("文字が見つかりません")
}

// 毒沼の迷路と同じ大きさの配列し0で初期化された状態で返す
func initMap(ysize int, xsize int) [][]int {
	var m [][]int
	for i := 0; i < ysize; i++ {
		arr := make([]int, xsize, xsize)
		m = append(m, arr)
	}
	return m
}
