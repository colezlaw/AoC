package main

import (
	"strings"
	"io/ioutil"
	"fmt"
	"os"
	"sort"
)

func input() ([]string, error) {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(content), "\n")
	return lines, nil
}

func verifiedUniqueWords(passphrases []string) []string {
	var valid []string
	for _, phrase := range passphrases{
		words := strings.Fields(phrase)
		unique := map[string]bool{}
		for _, v := range words {
			unique[v] = true
		}

		if len(words) == len(unique){
			valid = append(valid, phrase)
		}
	}

	return valid
}

func verifiedNoAnagrams(passphrases []string) []string {
	var valid []string

	for _, phrase := range passphrases{
		if containsAnagram(phrase) {
			continue
		}
		valid = append(valid, phrase)
	}

	return valid
}

func containsAnagram(phrase string) bool {
	var words []string

	for _, word := range strings.Fields(phrase) {
		runes := strings.Split(word, "")
		sort.Strings(runes)
		word = strings.Join(runes, "")
		if contains(words, word){
			return true
		}
		words = append(words, word)
	}

	return false
}

func contains(list []string, key string) bool {
	for _, value := range list{
		if value == key {
			return true
		}
	}

	return false
}


func main(){
	passphrases, err := input()
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	valid := verifiedUniqueWords(passphrases)

	fmt.Printf("There are %v unique passphrases.\n", len(valid))

	valid = verifiedNoAnagrams(valid)
	fmt.Printf("There are %v passphrases with no anagrams.\n", len(valid))

}