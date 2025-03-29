package main

import "testing"

func TestNewUserManager(t *testing.T) {
	um := NewUserManager()

	if um == nil {
		t.Error("NewUserManager() returned nil")
	} else if um.nextID != 1 {
		t.Errorf("Expected nextID to be 1, got %d", um.nextID)
	} else if len(um.users) != 0 {
		t.Errorf("Expected empty users slice, got length %d", len(um.users))
	}
}

func TestAddUser(t *testing.T) {
	um := NewUserManager()

	user := um.AddUser("Alice", "alice@example.com")

	if user == nil {
		t.Error("AddUser() returned nil")
	} else if user.ID != 1 {
		t.Errorf("Expected user ID to be 1, got %d", user.ID)
	} else if user.Name != "Alice" {
		t.Errorf("Expected user name to be 'Alice', got '%s'", user.Name)
	} else if user.Email != "alice@example.com" {
		t.Errorf("Expected user email to be 'alice@example.com', got '%s'", user.Email)
	} else if !user.IsActive {
		t.Error("Expected user to be active")
	} else if len(um.users) != 1 {
		t.Errorf("Expected users slice length to be 1, got %d", len(um.users))
	}
}

func TestFindUserByID(t *testing.T) {
	um := NewUserManager()
	um.AddUser("Alice", "alice@example.com")
	um.AddUser("Bob", "bob@example.com")

	// 存在するユーザーを検索
	user := um.FindUserByID(1)
	if user == nil {
		t.Error("FindUserByID(1) returned nil")
	} else if user.Name != "Alice" {
		t.Errorf("Expected user name to be 'Alice', got '%s'", user.Name)
	}

	// 存在しないユーザーを検索
	user = um.FindUserByID(999)
	if user != nil {
		t.Error("FindUserByID(999) should return nil")
	}
}

func TestDeactivateUser(t *testing.T) {
	um := NewUserManager()
	um.AddUser("Alice", "alice@example.com")

	// 存在するユーザーを非アクティブ化
	success := um.DeactivateUser(1)
	if !success {
		t.Error("DeactivateUser(1) should return true")
	}

	user := um.FindUserByID(1)
	if user.IsActive {
		t.Error("User should be inactive after deactivation")
	}

	// 存在しないユーザーを非アクティブ化
	success = um.DeactivateUser(999)
	if success {
		t.Error("DeactivateUser(999) should return false")
	}
}

func TestGetActiveUsers(t *testing.T) {
	um := NewUserManager()
	um.AddUser("Alice", "alice@example.com")
	um.AddUser("Bob", "bob@example.com")
	um.AddUser("Charlie", "charlie@example.com")

	// 全員アクティブの状態
	activeUsers := um.GetActiveUsers()
	if len(activeUsers) != 3 {
		t.Errorf("Expected 3 active users, got %d", len(activeUsers))
	}

	// 1人を非アクティブ化
	um.DeactivateUser(1)
	activeUsers = um.GetActiveUsers()
	if len(activeUsers) != 2 {
		t.Errorf("Expected 2 active users after deactivation, got %d", len(activeUsers))
	}

	// アクティブユーザーの名前をチェック
	names := make(map[string]bool)
	for _, user := range activeUsers {
		names[user.Name] = true
	}

	if !names["Bob"] || !names["Charlie"] {
		t.Error("Active users should be Bob and Charlie")
	} else if names["Alice"] {
		t.Error("Alice should not be in active users")
	}
}
