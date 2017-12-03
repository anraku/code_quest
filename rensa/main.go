// CODEQUESTの連鎖の試練をクリアするために作ったプログラム
package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

var (
	ariaWords []string        // 詠唱する呪文を先頭から詰めた配列
	spellList []string        // 呪文のリスト
	usedFlag  map[string]bool // すでに使った呪文を管理
)

func main() {
	// ファイルから呪文を読み込む
	v, err := input("./input.txt")
	if err != nil {
		fmt.Printf("ファイル入力でエラーになりました: %s", err)
		return
	}
	spellList = strings.Split(v, ",")

	usedFlag = make(map[string]bool)
	// 詠唱開始！
	for i := 0; i < len(spellList); i++ {
		aria(i)
	}
}

// ファイルからデータ読み込み
func input(filepath string) (string, error) {
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// しりとりのmapを使って再帰的に呪文の連鎖を作って標準出力する
func aria(i int) {
	if _, exist := usedFlag[spellList[i]]; exist {
		return
	}
	usedFlag[spellList[i]] = true
	ariaWords = append(ariaWords, spellList[i])
	if len(ariaWords) >= 20 { // 呪文が20個繋げられたら標準出力する
		output(ariaWords)
	}
	for j := 0; j < len(spellList); j++ { // 全呪文のリストを操作して次の呪文を探す
		lastSpell := []rune(ariaWords[len(ariaWords)-1])
		iSpell := []rune(spellList[j])
		if lastSpell[len(lastSpell)-1] == iSpell[0] { // 語尾と語頭が同じ呪文
			aria(j)
		}
	}
	delete(usedFlag, spellList[i])
	ariaWords = ariaWords[:len(ariaWords)-1]
}

// 呪文のリストを問題回答用のテキストに変換
func output(list []string) {
	fmt.Printf("出力結果: %v\n", strings.Join(list, " "))
}
