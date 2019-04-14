package main

import (
    "fmt"
    "flag"
)
var (
    f = flag.String("f", "example.puml", "string flag")
)
func getFile() string {
    // 該当のpumlを受け取る
    flag.Parse()
    return *f
}

func checkLanguage() bool {
    // 第一引数を受け取って対応言語か判定
    return true
}

func pumlConversionJson() {    
}

func main(){

    // pumlをjsonに変換

    // mkdir -p みたいに再帰的なディレクトリの作成 ログ出力をする

    // 初期フォーマット + ファイル名でファイルに記述

    // 終了コードをいれる
    // q,r := div(19, 7)
    // fmt.Printf("商=%d 剰余=%d\n", q,r)
    fmt.Println(getFile())
}
