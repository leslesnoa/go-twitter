package db

import (
	"fmt"
	"testing"

	"github.com/leslesnoa/go-twitter/models"
	"github.com/mongo-go/testdb"
	"github.com/stretchr/testify/assert"
)

var testDb *testdb.TestDB
var testUser = models.UserInfo{
	Email:    "test@example.com",
	Password: "password",
	Number:   "テスト",
	Name:     "太郎",
	Birth:    "2012-12-12",
}
var userID string

/* 1. ユーザ作成 */
func TestInsertRegister(t *testing.T) {
	id, status, err := InsertRegister(testUser)
	assert.Equal(t, status, true)
	assert.NoError(t, err)
	// fmt.Print(uId)

	userID = id
}

/* 2. 作成したuserIdでユーザ情報取得できる */
func TestSearchProfile(t *testing.T) {
	user, err := SearchProfile(userID)
	assert.NoError(t, err)
	assert.Equal(t, user.Email, "test@example.com")
	assert.Equal(t, user.Password, "")
	assert.Equal(t, user.Number, "テスト")
	assert.Equal(t, user.Name, "太郎")
	assert.Equal(t, user.Birth, "2012-12-12")
}

/* 3. ユーザ情報を変更できること */
func TestModifyRecord(t *testing.T) {
	testUser.Number = "hugahuga"
	testUser.Name = "hogehoge"
	testUser.Birth = "2020-11-11"
	// fmt.Println(userID)
	ok, err := ModifyRecord(testUser, userID)
	assert.Equal(t, ok, true)
	assert.NoError(t, err)
}

/* 4. ユーザ情報が変更されていることの確認 */
func TestSearchProfile2(t *testing.T) {
	user, err := SearchProfile(userID)
	assert.NoError(t, err)
	assert.Equal(t, user.Number, "hugahuga")
	assert.Equal(t, user.Name, "hogehoge")
	assert.Equal(t, user.Birth, "2020-11-11")
}

/* 5. 作成したユーザを削除できること */
func TestDeleteUser(t *testing.T) {
	fmt.Println(userID)
	err := DeleteUser(userID)
	assert.NoError(t, err)
}
