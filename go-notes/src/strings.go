package main

import (
    "fmt"
    "strings"
    "sort"
    "os"
    "log"
    "io/ioutil"
    //"strconv"
)

func main() {
    //egStrings()
    //egSort()
    egOs()
}

func egStrings() {
    testString:="Test string"
    csvString:="1,2,3,4,5,6,7,8,9,0"

    fmt.Println(strings.Contains(testString, "string"))
    fmt.Println(strings.Index(testString, "string"))
    fmt.Println(strings.Count(testString, "st"))
    fmt.Println(strings.Replace(testString, "t", "yo", -1))

    fmt.Println(strings.Split(csvString, ","))
}

func egSort() {
    unsortedLetters:=[]string{"c","a","b"}

    sort.Strings(unsortedLetters)

    fmt.Println(unsortedLetters)
}

func egOs() {
    file, err := os.Create("test.file")

    if err != nil {
        log.Fatal(err)
    }

    file.WriteString("random text")
    file.Close()

    stream, err := ioutil.ReadFile("test.file")

    if err != nil {
        log.Fatal(err)
    }

    readString := string(stream)

    fmt.Println(readString)
}
