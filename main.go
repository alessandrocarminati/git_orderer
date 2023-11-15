package main

import (
	"fmt"
	"io/ioutil"
        "os/exec"
	"os"
	"bufio"
	"sort"
	"strings"
)
const filename = "/tmp/git_orderere_cache"

func gitFetchHashHistory(repoPath string) ([]byte, error){

        cmd := exec.Command("git", "-C", repoPath, "log", "--pretty=%H")
        logOutput, err := cmd.CombinedOutput()
        if err != nil {
                return nil, err
        }

        return logOutput, nil
}

func order(toBeordered, reference []string) []string {
	indexMap := make(map[string]int)
	for i, val := range reference {
		indexMap[val] = i
	}

	var filtered []string
	for _, val := range toBeordered {
		if _, ok := indexMap[val]; ok {
			filtered = append(filtered, val)
		} else {
			return nil
		}
	}

	sort.Slice(filtered, func(i, j int) bool {
		return indexMap[filtered[i]] < indexMap[filtered[j]]
	})

	return filtered
}

func loadTextFromFile() ([]byte, error) {
	// Check if the file exists
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return nil, fmt.Errorf("file %s does not exist", filename)
	}

	// Read the content of the file
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	return content, nil
}

func main() {
	var lines []string

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}

	gitHistory, err := loadTextFromFile()
	if err != nil {
		gitHistory, err = gitFetchHashHistory(".")
		if err != nil {
			os.Exit(-1)
		}
		err = ioutil.WriteFile(filename, gitHistory, 0644)
		if err != nil {
			os.Exit(-1)
		}
	}

	s := order(lines, strings.Split(string(gitHistory),"\n"))
	for i := len(s) - 1; i >= 0; i-- {
		fmt.Println(s[i])
	}
}
