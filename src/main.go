package main

import (
    "fmt"
    "flag"
    "os"
    "bufio"
    "strings"
)
var (
    f = flag.String("f", "example.puml", "file flag")
    l = flag.String("l", "php", "language flag")
    format = flag.String("format", "default", "format flag")
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

func readFile(file string){
    // ファイル読み込み処理
    fp, err := os.Open(file)
    if err != nil {
        fmt.Fprintf(os.Stdout,"Can not read file:%s \n", file)
    }
    defer fp.Close()

    scanner := bufio.NewScanner(fp)
    all := map[string][]string{}
    var package_slice []string
    var class_slice []string
    for scanner.Scan(){
        record := scanner.Text()
        if strings.Contains(record,"package"){
            r := strings.NewReader(record)
            var (
                title string
                word string
            )
            n,err := fmt.Fscanf(r,"%s %s",&title,&word)
            if err != nil {
                fmt.Println(n)
            }
            package_slice = append(package_slice,word)
            // fmt.Println(word)
            if len(package_slice) != 1{
                all[package_slice[len(package_slice) - 1]] = class_slice
            }
            class_slice = nil // 初期化
        }else if strings.Contains(record,"class"){
            r := strings.NewReader(record)
            var (
                title string
                word string
            )
            n,err := fmt.Fscanf(r,"%s %s",&title,&word)
            if err != nil {
                fmt.Println(n)
            }
            class_slice = append(class_slice,word)
        }
    }
    // fmt.Println(package_slice)
    // fmt.Println(class_slice)
    fmt.Println(all)
    // 失敗読み込み業の出力
    if err = scanner.Err(); err != nil {
        fmt.Fprintf(os.Stdout,"Can not read line:%s \n", file)
    }
}

func main(){
    flag.Parse()
    file := getFile()
    if !checkLanguage(*l) {
        fmt.Fprintf(os.Stdout,"This language is not supported :%s \n", *l)
        return
    }
    readFile(file)
    // mkdir -p みたいに再帰的なディレクトリの作成 ログ出力をする

    // 初期フォーマット + ファイル名でファイルに記述

    // 終了コードをいれる
    // q,r := div(19, 7)
    // fmt.Printf("商=%d 剰余=%d\n", q,r)
}
