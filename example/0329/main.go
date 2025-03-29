package main

type User struct {
	ID       int
	Name     string
	Email    string
	IsActive bool
}

type UserManager struct {
	users  []User
	nextID int
}

// 1. 新しいUserManagerを作成する関数
func NewUserManager() *UserManager {
	// newUserManager := UserManager{}
	// return &newUserManager
	return &UserManager{nextID: 1}
}

// 2. 新しいユーザーを追加する関数（IDは自動採番）
func (um *UserManager) AddUser(name, email string) *User {
	// ここを実装
	// - IDは自動で採番（1から開始）
	// - IsActiveはtrueで初期化
	// - 追加したユーザーのポインタを返す
	// newUser := User{ID: len(um.users) + 1, Name: name, Email: email, IsActive: true}
	// return &newUser
	newUser := User{
		ID:       um.nextID,
		Name:     name,
		Email:    email,
		IsActive: true,
	}

	um.users = append(um.users, newUser)
	um.nextID++

	return &um.users[len(um.users)-1]
}

// 3. IDでユーザーを検索する関数
func (um *UserManager) FindUserByID(id int) *User {
	// ここを実装
	// - 見つからない場合はnilを返す
	for i, user := range um.users {
		if user.ID == id {
			return &um.users[i]
		}
	}
	return nil
}

// 4. ユーザーを非アクティブにする関数
func (um *UserManager) DeactivateUser(id int) bool {
	// ここを実装
	// - 成功時true、失敗時falseを返す
	for i, user := range um.users {
		if user.ID == id {
			um.users[i].IsActive = false
			return true
		}
	}
	return false
}

// 5. アクティブなユーザーのみを取得する関数
func (um *UserManager) GetActiveUsers() []User {
	// ここを実装
	// - 新しいスライスを作成して返す（元のデータは変更しない）
	activeUsers := []User{}
	for _, user := range um.users {
		if user.IsActive {
			activeUsers = append(activeUsers, user)
		}
	}

	return activeUsers
}
