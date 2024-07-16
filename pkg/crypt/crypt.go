package crypt

import (
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func CreatePassword(password []byte) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return hash, nil
}

func ComparePassword(hash, password []byte) error {
	err := bcrypt.CompareHashAndPassword(hash, password)
	return err
}

func save(fullexecpath string, password []byte) error {
	fileName := filepath.Join(fullexecpath, "password.txt")
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Printf("Error creating file: %s\n", err)
		return err
	}
	defer file.Close() // 确保在函数结束时关闭文件
	_, err = file.Write(password)
	if err != nil {
		fmt.Printf("Error writing to file: %s\n", err)
		return err
	}
	return nil
}

func SavePassword(fullexecpath string, password []byte) ([]byte, error) {
	hash, err := CreatePassword(password)
	if err != nil {
		return hash, errors.New(fmt.Sprintf("password hash create failed %v", err))
	}
	//glog.Debug(string(password), string(hash))
	return hash, save(fullexecpath, hash)
}
func GetPassword() []byte {
	fullexecpath, err := os.Executable()
	if err != nil {
		return nil
	}

	dir, _ := filepath.Split(fullexecpath)
	fileName := filepath.Join(dir, "password.txt")
	// 读取整个文件内容
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil
	}
	return content
}

func IsHashOk(password []byte) bool {
	content := GetPassword()
	if strings.Compare(string(password), string(content)) == 0 {
		return true
	}
	return false
}

func IsPasswordOk(password []byte) (bool, []byte) {
	//fullexecpath, err := os.Executable()
	//if err != nil {
	//	return false, nil
	//}
	//
	//dir, _ := filepath.Split(fullexecpath)
	//fileName := filepath.Join(dir, "password.txt")
	//// 读取整个文件内容
	//content, err := ioutil.ReadFile(fileName)
	//if err != nil {
	//	return false, nil
	//}

	content := GetPassword()
	err := ComparePassword(content, password)
	if err == nil {
		return true, content
	}
	return false, nil

}
