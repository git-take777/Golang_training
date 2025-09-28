package models


import (
	"time"
	"fmt"
	"log"
)
type User struct {
	ID int
	UUID string
	Name string
	Email string
	Password string
	CreatedAt time.Time
}

// func (u * User) CreateUser() (err error) {
// 	cmd := `insert into users (id, name, email, password, created_at) values (?, ?, ?, ?, ?)`
// 	_, err = DB.Exec(cmd, createUUID(), u.Name, u.Email, Encrypt(u.Password), time.Now())
// 	return err
// }

func (u *User) CreateUser() (err error) {
    fmt.Println("🔧 CreateUser関数開始")
    
    cmd := `insert into users (uuid, name, email, password, created_at) values (?, ?, ?, ?, ?)`
    
    // 暗号化されたパスワードを確認
    encryptedPassword := Encrypt(u.Password)
    fmt.Printf("🔐 暗号化されたパスワード: %s\n", encryptedPassword)
    
    // 現在時刻を確認
    now := time.Now()
    fmt.Printf("⏰ 現在時刻: %s\n", now)
    
    fmt.Printf("📝 実行するSQL: %s\n", cmd)
    fmt.Printf("📝 パラメータ: Name=%s, Email=%s, Password=%s, Time=%s\n", 
        u.Name, u.Email, encryptedPassword, now)
    
    result, err := DB.Exec(cmd, createUUID(), u.Name, u.Email, encryptedPassword, time.Now())
    
    if err != nil {
        fmt.Printf("❌ SQL実行エラー: %v\n", err)
        return err
    }
    
    // 挿入された行数を確認
    rowsAffected, err := result.RowsAffected()
    if err != nil {
        fmt.Printf("⚠️ RowsAffected取得エラー: %v\n", err)
    } else {
        fmt.Printf("✅ 挿入された行数: %d\n", rowsAffected)
    }
    
    // 最後に挿入されたIDを確認
    lastID, err := result.LastInsertId()
    if err != nil {
        fmt.Printf("⚠️ LastInsertId取得エラー: %v\n", err)
    } else {
        fmt.Printf("🆔 挿入されたID: %d\n", lastID)
    }
    
    fmt.Println("🔧 CreateUser関数終了")
    return nil
}

func GetUser(id int) (user User, err error) {
	user = User{}
	cmd := `select * from users where id = ?`
	err = DB.QueryRow(cmd, id).Scan(
		&user.ID,
		&user.UUID,
		&user.Name,
		&user.Email,
		&user.Password, 
		&user.CreatedAt,
	)
	return user, err
}

func (u * User) UpdateUser() (err error) {
	cmd := `update users set name = ?, email = ? where id = ?`
	_, err = DB.Exec(cmd, u.Name, u.Email, u.ID)

	if err != nil {
		log.Fatalln(err)
	}
	return err
}