package main

import (
	"bufio"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"io/fs"
	"os"
	// "reflect"
)

type User struct {
	UserName   string
	ActualPass string
	Password   string //hash(actualpass+name)
	BellaLevel int
	BibaLevel  int
}

func createHash(key string) string {
	hasher := md5.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}

var AuthUsers = []User{
	{UserName: "neha", ActualPass: "neha@123", Password: "0d5d6947297c942c3ad3105e7d40da63", BellaLevel: 1, BibaLevel: 4},
	{UserName: "pooja", ActualPass: "pooja@123", Password: "ba291bccdd8da41025994c915a67f8e4", BellaLevel: 3, BibaLevel: 1},
	{UserName: "sweety", ActualPass: "sweety@123", Password: "f062ebbe981b8f61ac3c25bb49f46f94", BellaLevel: 3, BibaLevel: 2},
	{UserName: "rina", ActualPass: "rina@123", Password: "44cc551727c5e5a842d3354f272fabad", BellaLevel: 2, BibaLevel: 3},
}

func getUser(username string) User {
	for _, user := range AuthUsers {
		if username == user.UserName {
			return user
		}
	}
	return User{}
}

func validateLogin(user User, password string) bool {
	emptyuser := User{}
	if user != emptyuser {
		if createHash(password+user.UserName) == user.Password {
			return true
		} else {
			return false
		}
	}
	return false
}

type File struct {
	name       string
	BellaLevel int //confidentiality level
	BibaLevel  int
}

//higher level value means higher level of confidentiality
var file1 = File{name: "file1.txt", BellaLevel: 3, BibaLevel: 1}
var file2 = File{name: "file2.txt", BellaLevel: 1, BibaLevel: 3}
var file3 = File{name: "file3.txt", BellaLevel: 2, BibaLevel: 3}
var metadata = []File{file1, file2, file3}

func encrypt(data []byte, passphrase string) []byte {
	block, _ := aes.NewCipher([]byte(createHash(passphrase)))
	gcm, _ := cipher.NewGCM(block)
	nonce := make([]byte, gcm.NonceSize())
	io.ReadFull(rand.Reader, nonce)
	cipherText := gcm.Seal(nonce, nonce, data, nil)
	return cipherText
}

func decrypt(data []byte, passphrase string) []byte {
	key := []byte(createHash(passphrase))
	block, _ := aes.NewCipher(key)
	gcm, _ := cipher.NewGCM(block)
	nonceSize := gcm.NonceSize()
	nonce := data[:nonceSize]
	cipherText := data[nonceSize:]
	plainText, err := gcm.Open(nil, nonce, cipherText, nil)
	if err != nil {
		fmt.Println("err in decryption", err)
	}
	return plainText
}

func WriteToFile(filename string, fileMode os.FileMode, data string) error {
	file, err := os.OpenFile("./data/"+filename, int(fileMode), 0777)
	if err != nil {
		return err
	}
	defer file.Close()
	if _, err := file.Write([]byte(encrypt([]byte(data), "hello"))); err != nil {
		return err
	}
	return nil
}

func ReadFromFile(filename string) (string, error) {
	data, err := os.ReadFile("./data/" + filename)
	// fmt.Println(string(data))
	if err != nil {
		return "", err
	}
	// fmt.Println("data type is", reflect.TypeOf(data))
	return string(decrypt([]byte(data), "hello")), nil
}

func checkControl(userBellaLevel int, userBibaLevel int, operation string) []string {
	var filesAlllowedToread []string
	var filesAlllowedTowrite []string
	for _, file := range metadata {
		if operation == "read" {
			if (file.BibaLevel > userBibaLevel) && (file.BellaLevel < userBellaLevel) {
				filesAlllowedToread = append(filesAlllowedToread, file.name)

			}
		} else if operation == "write" {
			if (file.BibaLevel < userBibaLevel) && (file.BellaLevel > userBellaLevel) {
				filesAlllowedTowrite = append(filesAlllowedTowrite, file.name)
			}

		} else {
			fmt.Println("invalid operation")
		}
	}
	if len(filesAlllowedToread) > 0 {
		return filesAlllowedToread
	} else if len(filesAlllowedTowrite) > 0 {
		return filesAlllowedTowrite
	}
	return nil
}
func main() {
	var username string
	var password string
	var operation string
	var filename string
	fmt.Println("Enter your name-")
	fmt.Scanln(&username)
	fmt.Println("Enter your password")
	fmt.Scanln(&password)
	currentUser := getUser(username)

	if validateLogin(currentUser, password) {
		fmt.Println("--------------login succeeded-------------")
		fmt.Println("Enter the operation read or write")
		fmt.Scanln(&operation)
		userBellaLevel := currentUser.BellaLevel
		userBibaLevel := currentUser.BibaLevel
		file := checkControl(userBellaLevel, userBibaLevel, operation)
		if operation == "read" {
			fmt.Println("You Are Having An Access To Read Following Files .Choose one of them")
			fmt.Println(file)
			fmt.Scanln(&filename)
			data, err := ReadFromFile(filename)
			if err != nil {
				fmt.Println("error in reading file", err)
			}
			fmt.Println(data)
		} else if operation == "write" {
			fmt.Println("You Are Having An Access To write Following Files .Choose one of them")
			fmt.Println(file)
			fmt.Scanln(&filename)
			scanner1 := bufio.NewScanner(os.Stdin)
			fmt.Println("Enter data to write-")
			scanner1.Scan()
			data := scanner1.Text()

			// fmt.Println("ur entered data is ", data)
			err := WriteToFile(filename, fs.FileMode(os.O_APPEND), data)
			if err != nil {
				fmt.Println("error in writing to file-", err)
			}
		}

	} else {
		fmt.Println("plz enter valid crediantials")
	}
}
