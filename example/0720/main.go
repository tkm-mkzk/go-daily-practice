package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

func makeValidatorStep1() func(string) (bool, []string) {
	// 固定のバリデーション（requiredのみ）
	return func(input string) (bool, []string) {
		if input == "" {
			return false, []string{"入力は必須です"}
		}
		return true, []string{}
	}
}

func makeValidatorStep2(rule string) func(string) (bool, []string) {
	return func(input string) (bool, []string) {
		var errors []string

		switch rule {
		case "required":
			if input == "" {
				errors = append(errors, "入力は必須です")
			}
		}

		return len(errors) == 0, errors
	}
}

func makeValidator(rules ...string) func(string) (bool, []string) {
	// クロージャ: rulesを「記憶」した関数を返す
	return func(input string) (bool, []string) {
		var errors []string

		// 各ルールをチェック
		for _, rule := range rules {
			switch {
			case rule == "required":
				if input == "" {
					errors = append(errors, "入力は必須です")
				}

			case strings.HasPrefix(rule, "minLength:"):
				// "minLength:3" から "3" を取り出す
				parts := strings.Split(rule, ":")
				if len(parts) == 2 {
					minLen, err := strconv.Atoi(parts[1])
					if err == nil && utf8.RuneCountInString(input) < minLen {
						errors = append(errors, fmt.Sprintf("最低%d文字必要です", minLen))
					}
				}

			case strings.HasPrefix(rule, "maxLength:"):
				// "maxLength:20" から "20" を取り出す
				parts := strings.Split(rule, ":")
				if len(parts) == 2 {
					maxLen, err := strconv.Atoi(parts[1])
					if err == nil && utf8.RuneCountInString(input) > maxLen {
						errors = append(errors, fmt.Sprintf("最大%d文字までです", maxLen))
					}
				}

			case rule == "email":
				if !strings.Contains(input, "@") {
					errors = append(errors, "有効なメールアドレスではありません")
				}
			}
		}

		return len(errors) == 0, errors
	}
}

func demonstrateSteps() {
	fmt.Println("=== ステップ1: 固定ルール ===")
	validator1 := makeValidatorStep1()
	valid, errs := validator1("")
	fmt.Printf("空文字: valid=%t, errors=%v\n", valid, errs)

	fmt.Println("\n=== ステップ2: 1つのルール ===")
	validator2 := makeValidatorStep2("required")
	valid, errs = validator2("hello")
	fmt.Printf("'hello': valid=%t, errors=%v\n", valid, errs)

	fmt.Println("\n=== ステップ3: 複数のルール ===")
	validator3 := makeValidator("required", "minLength:3")
	valid, errs = validator3("ab")
	fmt.Printf("'ab': valid=%t, errors=%v\n", valid, errs)
}

func main() {
	demonstrateSteps()

	fmt.Println("\n" + strings.Repeat("=", 50))
	fmt.Println("完全版のテスト")
	fmt.Println(strings.Repeat("=", 50))

	validateUsername := makeValidator("required", "minLength:3", "maxLength:20")

	valid, errs := validateUsername("testuser")
	fmt.Printf("'testuser': valid=%t, errors=%v\n", valid, errs)

	validateEmail := makeValidator("required", "email")

	valid, errs = validateEmail("test@example.com")
	fmt.Printf("'test@example.com': valid=%t, errors=%v\n", valid, errs)
}
