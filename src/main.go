package main

import (
    "fmt"
)

func div(a, b int) (int, int){
    q := a / b
    r := a % b
    return q,r
}

func main(){
    // 該当のpumlを受け取る

    // 第一引数を受け取って対応言語か判定

    // pumlをjsonに変換

    // mkdir -p みたいに再帰的なディレクトリの作成 ログ出力をする

    // 初期フォーマット + ファイル名でファイルに記述

    // 終了コードをいれる
    q,r := div(19, 7)
    fmt.Printf("商=%d 剰余=%d\n", q,r)
}
