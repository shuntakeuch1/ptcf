package main

import (
    "testing"
    "flag"
)

func TestGetFile(t *testing.T) {
    expected := "test.puml"
    flag.CommandLine.Set("f", "test.puml")
    
    actual := getFile()
    if expected != actual {
        t.Error(actual)
    }
}

func TestCheckLanguage(t *testing.T) {
    expected := true
    actual := checkLanguage()
    if expected != actual {
        t.Error("No actual check")
    }
}
