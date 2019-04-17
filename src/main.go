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
    return *f
}

func checkLanguage(lang string) bool {
    // 第2引数を受け取って対応言語か判定
    languages := []string{"php","go","ruby"}
    for _,l := range languages {
        if lang == l {
            return true
        }
    }
    return false
}

func pumlConversionJson() {
    // pumlをjsonに変換
}

func readFile(file string) map[string][]string{
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
                all[package_slice[len(package_slice) - 2]] = class_slice
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
    // 最後だけ漏れてしまうので追加
    all[package_slice[len(package_slice) - 1]] = class_slice
    // fmt.Println(all)
    // 失敗読み込み業の出力
    if err = scanner.Err(); err != nil {
        fmt.Fprintf(os.Stdout,"Can not read line:%s \n", file)
    }
    return all
}
func makeFile(class_map map[string][]string,lang string){
    // mkdir -p みたいに再帰的なディレクトリの作成 ログ出力をする
    for p,clasies := range class_map {
        // unix 依存 ディレクトリ文字列の作成
        p :=strings.Replace(p, ".", "/", -1)
        // ディレクトリ作成
        if err := os.MkdirAll(p, 0777); err != nil {
            fmt.Println(err)
        }
        fmt.Println("directory create ",p)

        for _, m := range clasies {
            // 初期フォーマット + ファイル名でファイルに記述
            filename := p + "/" + m + "." + lang
            _,err := os.Create(filename)
            if err != nil {
                os.Exit(1)
            }
            fmt.Println("create ",filename)
        }
    }
}
func main(){
    flag.Parse()
    file := getFile()
    if !checkLanguage(*l) {
        fmt.Fprintf(os.Stdout,"This language is not supported :%s \n", *l)
        return
    }
    class_map := readFile(file)

    makeFile(class_map,*l)

    os.Exit(0)
}

