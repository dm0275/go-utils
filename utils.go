package utils

import (
	"io/ioutil"
	"bufio"
	"os"
	"fmt"
	//"strings"
	"regexp"
	//"reflect"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func ReadFile(file string) string {
	creds,err := ioutil.ReadFile(file)
	check(err)
	return string(creds)
}

func GetAcctNames(awsCreds string) []string {
	awsAccounts := []string{}
	if (awsCreds == "") {
		return awsAccounts
	}
	re := regexp.MustCompile(`\[.*\]`)
	matches := re.FindAllString(string(awsCreds), -1)
	for _, acct := range matches {
		awsAccounts = append(awsAccounts, string(acct[1:len(acct)-1]))
	}
	return awsAccounts
}

func GetAwsAcctMatches(awsAccts []string, awsCreds string) {
	fmt.Println(awsCreds)
	if (len(awsAccts) > 0) {
		for _, acct := range awsAccts {
			re := regexp.MustCompile(`(?m)^.*\[` + acct + `]*\]\n(([a-z].+\n)*)`)
			match := re.FindAllString(string(awsCreds), -1)
			fmt.Println(match)
		}
	}
}

func ReadFile2(file string) {
	creds,err := os.Open(file)
	
	if err != nil {
        panic(err)
    }
	
	scanner := bufio.NewScanner(creds)
    for scanner.Scan() {
		// if !strings.Contains(scanner.Text(), "[") {
		// 	fmt.Println(scanner.Text())
		// }

		re := regexp.MustCompile(`\[[a-zA-Z-\d]*\]`)
		match := re.FindString(scanner.Text())
		fmt.Println(match)

		// matched, err := regexp.MatchString(`\[[a-zA-Z]*\]`, scanner.Text())
		// if err == nil {
		// 	// fmt.Println(err)	
		// 	// fmt.Println(matched)	
		// }
		// if matched {
			
		// }
	}

	re2 := regexp.MustCompile(`(?m)^.*\[[a-zA-Z-\d]*\]\n(([a-z].+\n)*)`)
	creds2,err2 := ioutil.ReadFile(file)
	if err2 != nil {
		
	}
	match2 := re2.FindAllString(string(creds2), -1)
	fmt.Println(match2[0])
}

