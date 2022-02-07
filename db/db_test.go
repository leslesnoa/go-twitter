package db

import (
	"context"
	"testing"
	"time"

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

/* ユーザ作成 */
func TestInsertRegister(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	id, status, err := InsertRegister(testUser, ctx)
	assert.Equal(t, status, true)
	assert.NoError(t, err)
	assert.True(t, status)
	// fmt.Print(uId)

	userID = id
}

/* ログインできること */
func TestTryLogin(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	user, ok := TryLogin(testUser.Email, testUser.Password, ctx)
	assert.True(t, ok)
	assert.Equal(t, testUser.Email, user.Email)
	assert.Equal(t, testUser.Number, user.Number)
	assert.Equal(t, testUser.Name, user.Name)
	assert.Equal(t, testUser.Birth, user.Birth)
}

/* userIdでユーザ情報取得できる */
func TestSearchProfile(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	user, err := SearchProfile(userID, ctx)
	assert.NoError(t, err)
	assert.Equal(t, "test@example.com", user.Email)
	assert.Equal(t, "", user.Password)
	assert.Equal(t, "テスト", user.Number)
	assert.Equal(t, "太郎", user.Name)
	assert.Equal(t, "2012-12-12", user.Birth)
}

/* ユーザ情報を変更できる */
func TestModifyRecord(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	testUser.Number = "hugahuga"
	testUser.Name = "hogehoge"
	testUser.Birth = "2020-11-11"
	// fmt.Println(userID)
	ok, err := ModifyRecord(testUser, userID, ctx)
	assert.Equal(t, ok, true)
	assert.NoError(t, err)
	assert.True(t, ok)
}

/* ユーザ情報が変更されていることの確認 */
func TestSearchProfile2(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	user, err := SearchProfile(userID, ctx)
	assert.NoError(t, err)
	assert.Equal(t, "hugahuga", user.Number)
	assert.Equal(t, "hogehoge", user.Name)
	assert.Equal(t, "2020-11-11", user.Birth)
}

/* 作成したユーザを削除できること */
func TestDeleteUser(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	err := DeleteUser(userID, ctx)
	assert.NoError(t, err)
}
