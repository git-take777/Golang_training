package main

import (
	"fmt"
	"todoAPP/app/models"
)
func main() {
	// fmt.Println(config.Config.Port)
	// fmt.Println(config.Config.SQLDriver)
	// fmt.Println(config.Config.DBName)
	// fmt.Println(config.Config.LogFile)
	fmt.Println(models.DB)

	// u := &models.User{}
	// u.Name = "test"
	// u.Email = "test@example.com"
	// u.Password = "testTest"
	// u.CreateUser()
	// // u := &models.User{}
	// u.Name = "test"
	// u.Email = "test@example.com"
	// u.Password = "testTest"
	// u.CreateUser()

	fmt.Println("=== ユーザー取得開始 ===")
	u, err := models.GetUser(1)  // エラーを確認
	if err != nil {
		fmt.Printf("❌ GetUser エラー: %v\n", err)
	} else {
		fmt.Println("✅ ユーザー取得成功")
	}
	fmt.Printf("取得したユーザー: %+v\n", u)

	u.Name = "updatedUser"
	u.Email = "test@example.com"
	fmt.Println('=== ユーザー更新開始 ===')
	u.UpdateUser()

	fmt.Println('=== ユーザー更新完了 ===')
	fmt.Println(u)
}