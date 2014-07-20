package main 

import (
    "fmt"
    "strings"
    "os"
    "flag"
    "path/filepath"
    "io/ioutil"
    "regexp"
    "strconv"
)

func main() {
	regularExpression := ""
    flag.Parse()
    argument := flag.Arg(0)
    for i := 0; i < len(argument); i++ {
    	regularExpression += "[" + strings.ToLower(string(argument[i])) + "+" + strings.ToUpper(string(argument[i])) + "]"
    }
    fmt.Println(regularExpression)
    r, _ := regexp.Compile(regularExpression)
    fmt.Println("Searching root directory and subdirectories ...")
    fmt.Println("Replacing the word " + argument + " with " + strings.ToUpper(argument))
    fmt.Println("Changes          Files")
    slice := make([]string, 0)
    filepath.Walk(".", func (path string, fileInfo os.FileInfo, err error) error {
        bs, err := ioutil.ReadFile(path)
    	if err != nil {
    		if(fileInfo.IsDir()) {
    			slice = append(slice, strconv.Itoa(0) + "|" + path + "|" + " <This is a directory>")
    		} else {
    			slice = append(slice, strconv.Itoa(0) + "|" + path + "|" + " <Can't open this file>")
        		}
        	return nil
    	}
    	fileContent := string(bs)
    	file, err := os.OpenFile(path, os.O_WRONLY, 0666)
    	if err != nil {
        	return nil
    	}
    	defer file.Close()
    	file.WriteString(r.ReplaceAllString(fileContent, strings.ToUpper(argument)))
    	numberOfChanges := len(r.FindAllString(fileContent, -1))
    	slice = append(slice, strconv.Itoa(numberOfChanges) + "|" + path + "|" + " ")
        return nil
    })
    maxChanges := 0
    for i := 0; i < len(slice); i++ {
    	changes, err := strconv.Atoi(strings.Split(slice[i], "|")[0])
    	if(err != nil) {
    		fmt.Printf("error occured")
    	}
    	if(maxChanges < changes) {
    		maxChanges = changes
    	}
    }
    for j := maxChanges; j >= 0; j-- {
    	for i := 0; i < len(slice); i++ {
    		changes, err := strconv.Atoi(strings.Split(slice[i], "|")[0])
    		if(err != nil) {
    			fmt.Printf("error occured")
    		}
    		if(changes == j) {
    			fmt.Printf("%s			%s%s\n", strings.Split(slice[i], "|")[0], strings.Split(slice[i], "|")[1], strings.Split(slice[i], "|")[2])
    		}
    	}
    }
}