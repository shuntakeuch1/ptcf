package main

import (
    "testing"
)

func TestGetFile(t *testing.T) {
    expected := "Hello World"
    actual := getFile()
    if expected != actual {
        t.Error("No actual main")
    }
}

func TestCheckLanguage(t *testing.T) {
    expected := true
    actual := checkLanguage()
    if expected != actual {
        t.Error("No actual check")
    }
}
