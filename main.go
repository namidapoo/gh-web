package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

func main() {
	// サブコマンドの取得
	if len(os.Args) < 2 {
		// サブコマンドがない場合にはただの `gh browse` を実行
		runGHCommand("gh", "browse")
		return
	}
	command := os.Args[1]

	// サブコマンドに基づいて処理を分岐
	switch command {
	case "i", "issue":
		// イシューの一覧表示
		runGHCommand("gh", "issue", "list", "--limit", "10")

		// ユーザーにイシュー番号を入力してもらう
		fmt.Print("\n\033[32m?\033[0m Please enter the number of the Issue you want to open ... ")
		issueNumber := readInput()

		// 入力が数字かどうかを確認
		if !isNumeric(issueNumber) {
			fmt.Println("Please enter a valid Issue number.")
			os.Exit(1)
		}

		// イシューが存在するか確認
		if !ghIssueExists(issueNumber) {
			fmt.Println("The specified Issue number does not exist.")
			os.Exit(1)
		}

		// 指定されたイシューをブラウザで開く
		runGHCommand("gh", "issue", "view", issueNumber, "--web")

	case "p", "pr":
		// プルリクエストの一覧表示
		runGHCommand("gh", "pr", "list", "--limit", "10")

		// ユーザーにPR番号を入力してもらう
		fmt.Print("\n\033[32m?\033[0m Please enter the number of the PR you want to open ... ")
		prNumber := readInput()

		// 入力が数字かどうかを確認
		if !isNumeric(prNumber) {
			fmt.Println("Please enter a valid PR number.")
			os.Exit(1)
		}

		// PRが存在するか確認
		if !ghPRExists(prNumber) {
			fmt.Println("The specified PR number does not exist.")
			os.Exit(1)
		}

		// 指定されたPRをブラウザで開く
		runGHCommand("gh", "pr", "view", prNumber, "--web")

	default:
		args := []string{"browse"}
		args = append(args, os.Args[1:]...)
		runGHCommand("gh", args...)
	}
}

// ghコマンドを実行する関数
func runGHCommand(name string, args ...string) {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error running command %s: %v\n", name, err)
		os.Exit(1)
	}
}

// ユーザー入力を読み取る関数
func readInput() string {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		os.Exit(1)
	}
	return strings.TrimSpace(input)
}

// 入力が数字かどうかを確認する関数
func isNumeric(input string) bool {
	regex := regexp.MustCompile(`^[0-9]+$`)
	return regex.MatchString(input)
}

// イシューが存在するか確認する関数
func ghIssueExists(issueNumber string) bool {
	cmd := exec.Command("gh", "issue", "view", issueNumber)
	err := cmd.Run()
	return err == nil
}

// PRが存在するか確認する関数
func ghPRExists(prNumber string) bool {
	cmd := exec.Command("gh", "pr", "view", prNumber)
	err := cmd.Run()
	return err == nil
}
