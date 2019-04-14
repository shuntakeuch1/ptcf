package main

import (
    "fmt"
    "flag"
    "os"
)
var (
    f = flag.String("f", "example.puml", "file flag")
    l = flag.String("l", "php", "language flag")
)
func getFile() string {
    // 該当のpumlを受け取る
    return *f
}

func checkLanguage(lang string) bool {
    // 第2引数を受け取って対応言語か判定
    if lang == "php" {
        return true        
    }
    return false
}

func pumlConversionJson() {    
    // pumlをjsonに変換
}

func main(){
    flag.Parse()
    file := getFile()
    if !checkLanguage(*l) {
        fmt.Fprintf(os.Stdout,"This language is not supported :%s \n", *l)
        return 
    }

    // mkdir -p みたいに再帰的なディレクトリの作成 ログ出力をする

    // 初期フォーマット + ファイル名でファイルに記述

    // 終了コードをいれる
    // q,r := div(19, 7)
    // fmt.Printf("商=%d 剰余=%d\n", q,r)
    fmt.Println(file)
}
