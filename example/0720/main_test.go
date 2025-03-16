package main

import (
	"testing"
)

// エラースライスを比較するヘルパー関数
func errorsEqual(got, want []string) bool {
	if len(got) != len(want) {
		return false
	}
	for i := range got {
		if got[i] != want[i] {
			return false
		}
	}
	return true
}

// TestValidateUsername - ユーザー名のバリデーションテスト
func TestValidateUsername(t *testing.T) {
	validateUsername := makeValidator("required", "minLength:3", "maxLength:20")

	tests := []struct {
		name       string
		input      string
		wantValid  bool
		wantErrors []string
	}{
		{
			name:       "有効なユーザー名",
			input:      "validuser",
			wantValid:  true,
			wantErrors: nil, // エラーなしはnilを期待
		},
		{
			name:       "短いユーザー名",
			input:      "ab",
			wantValid:  false,
			wantErrors: []string{"最低3文字必要です"},
		},
		{
			name:       "空のユーザー名",
			input:      "",
			wantValid:  false,
			wantErrors: []string{"入力は必須です", "最低3文字必要です"}, // 複数エラー
		},
		{
			name:       "長すぎるユーザー名",
			input:      "verylongusernamethatistoolong",
			wantValid:  false,
			wantErrors: []string{"最大20文字までです"},
		},
		{
			name:       "境界値テスト - 最小長",
			input:      "abc",
			wantValid:  true,
			wantErrors: nil,
		},
		{
			name:       "境界値テスト - 最大長",
			input:      "abcdefghijklmnopqrst", // 20文字
			wantValid:  true,
			wantErrors: nil,
		},
		{
			name:       "日本語ユーザー名 - 有効",
			input:      "山田太郎",
			wantValid:  true,
			wantErrors: nil,
		},
		{
			name:       "日本語ユーザー名 - 短い",
			input:      "山田",
			wantValid:  false,
			wantErrors: []string{"最低3文字必要です"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotValid, gotErrors := validateUsername(tt.input)

			if gotValid != tt.wantValid {
				t.Errorf("validateUsername(%q) valid = %v, want %v", tt.input, gotValid, tt.wantValid)
			}

			// エラーがない場合の特別な処理
			if tt.wantErrors == nil {
				if len(gotErrors) != 0 {
					t.Errorf("validateUsername(%q) errors = %v, want no errors", tt.input, gotErrors)
				}
			} else {
				if !errorsEqual(gotErrors, tt.wantErrors) {
					t.Errorf("validateUsername(%q) errors = %v, want %v", tt.input, gotErrors, tt.wantErrors)
				}
			}
		})
	}
}

// TestValidateEmail - メールアドレスのバリデーションテスト
func TestValidateEmail(t *testing.T) {
	validateEmail := makeValidator("required", "email")

	tests := []struct {
		name       string
		input      string
		wantValid  bool
		wantErrors []string
	}{
		{
			name:       "有効なメール",
			input:      "test@example.com",
			wantValid:  true,
			wantErrors: nil,
		},
		{
			name:       "無効なメール - @なし",
			input:      "invalid-email",
			wantValid:  false,
			wantErrors: []string{"有効なメールアドレスではありません"},
		},
		{
			name:       "空のメール",
			input:      "",
			wantValid:  false,
			wantErrors: []string{"入力は必須です", "有効なメールアドレスではありません"}, // 複数エラー
		},
		{
			name:       "日本語を含むメール",
			input:      "山田@example.com",
			wantValid:  true,
			wantErrors: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotValid, gotErrors := validateEmail(tt.input)

			if gotValid != tt.wantValid {
				t.Errorf("validateEmail(%q) valid = %v, want %v", tt.input, gotValid, tt.wantValid)
			}

			// エラーがない場合の特別な処理
			if tt.wantErrors == nil {
				if len(gotErrors) != 0 {
					t.Errorf("validateEmail(%q) errors = %v, want no errors", tt.input, gotErrors)
				}
			} else {
				if !errorsEqual(gotErrors, tt.wantErrors) {
					t.Errorf("validateEmail(%q) errors = %v, want %v", tt.input, gotErrors, tt.wantErrors)
				}
			}
		})
	}
}

// TestValidatePassword - パスワードのバリデーションテスト（複数ルール）
func TestValidatePassword(t *testing.T) {
	validatePassword := makeValidator("required", "minLength:8", "maxLength:50")

	tests := []struct {
		name       string
		input      string
		wantValid  bool
		wantErrors []string
	}{
		{
			name:       "有効なパスワード",
			input:      "password123",
			wantValid:  true,
			wantErrors: nil,
		},
		{
			name:       "短いパスワード",
			input:      "pass",
			wantValid:  false,
			wantErrors: []string{"最低8文字必要です"},
		},
		{
			name:       "空のパスワード",
			input:      "",
			wantValid:  false,
			wantErrors: []string{"入力は必須です", "最低8文字必要です"}, // 複数エラー
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotValid, gotErrors := validatePassword(tt.input)

			if gotValid != tt.wantValid {
				t.Errorf("validatePassword(%q) valid = %v, want %v", tt.input, gotValid, tt.wantValid)
			}

			// エラーがない場合の特別な処理
			if tt.wantErrors == nil {
				if len(gotErrors) != 0 {
					t.Errorf("validatePassword(%q) errors = %v, want no errors", tt.input, gotErrors)
				}
			} else {
				if !errorsEqual(gotErrors, tt.wantErrors) {
					t.Errorf("validatePassword(%q) errors = %v, want %v", tt.input, gotErrors, tt.wantErrors)
				}
			}
		})
	}
}

// TestMakeValidator - makeValidator関数自体のテスト
func TestMakeValidator(t *testing.T) {
	t.Run("複数のエラーが同時に発生", func(t *testing.T) {
		validator := makeValidator("required", "minLength:5", "email")

		gotValid, gotErrors := validator("")
		wantValid := false
		wantErrors := []string{"入力は必須です", "最低5文字必要です", "有効なメールアドレスではありません"}

		if gotValid != wantValid {
			t.Errorf("validator('') valid = %v, want %v", gotValid, wantValid)
		}

		if !errorsEqual(gotErrors, wantErrors) {
			t.Errorf("validator('') errors = %v, want %v", gotErrors, wantErrors)
		}
	})

	t.Run("無効な数値のルール", func(t *testing.T) {
		validator := makeValidator("minLength:invalid")

		gotValid, gotErrors := validator("test")
		wantValid := true // 無効なルールは無視される

		if gotValid != wantValid {
			t.Errorf("validator('test') valid = %v, want %v", gotValid, wantValid)
		}

		if len(gotErrors) != 0 {
			t.Errorf("validator('test') errors = %v, want no errors", gotErrors)
		}
	})

	t.Run("ルールなしのバリデーター", func(t *testing.T) {
		validator := makeValidator() // ルールなし

		gotValid, gotErrors := validator("anything")
		wantValid := true

		if gotValid != wantValid {
			t.Errorf("validator('anything') valid = %v, want %v", gotValid, wantValid)
		}

		if len(gotErrors) != 0 {
			t.Errorf("validator('anything') errors = %v, want no errors", gotErrors)
		}
	})
}

// TestEdgeCases - エッジケースのテスト
func TestEdgeCases(t *testing.T) {
	t.Run("境界値 - 文字数ちょうど", func(t *testing.T) {
		validator := makeValidator("minLength:5", "maxLength:5")

		gotValid, gotErrors := validator("12345")
		wantValid := true

		if gotValid != wantValid {
			t.Errorf("validator('12345') valid = %v, want %v", gotValid, wantValid)
		}

		if len(gotErrors) != 0 {
			t.Errorf("validator('12345') errors = %v, want no errors", gotErrors)
		}
	})

	t.Run("絵文字を含む文字列", func(t *testing.T) {
		validator := makeValidator("minLength:3")

		gotValid, gotErrors := validator("🚀🌟💻") // 3文字の絵文字
		wantValid := true

		if gotValid != wantValid {
			t.Errorf("validator('🚀🌟💻') valid = %v, want %v", gotValid, wantValid)
		}

		if len(gotErrors) != 0 {
			t.Errorf("validator('🚀🌟💻') errors = %v, want no errors", gotErrors)
		}
	})
}

// Benchmark - パフォーマンステスト
func BenchmarkValidator(b *testing.B) {
	validator := makeValidator("required", "minLength:3", "maxLength:20", "email")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validator("test@example.com")
	}
}
