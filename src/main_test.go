package main

import (
    "testing"
    "flag"
    "fmt"
)

func TestGetFile(t *testing.T) {
    expected := "test.puml"
    flag.CommandLine.Set("f", "test.puml")
    
    actual := getFile()
    if expected != actual {
        t.Error("No actual main " + actual)
    }
}

func TestCheckLanguage(t *testing.T) {
    actual := checkLanguage("go")
    if (!actual) {
        t.Error("No actual checkLanguage " + fmt.Sprint(actual))
    }
}
