//[这里](https://paste.ubuntu.com/p/JbWGy9PVzt/)有一份代码，实现了登录注册、账号密码加密存储的功能。
//不仅写的稀烂，还有很多地方忘了加锁！这会导致程序很不稳定！
//你需要帮他加上足够多的锁，来保证程序运行时的线程安全。

//产生协程的函数：
//c.Save()	含saveUsers函数
//匿名函数X2

package Lv2
//package main

import (
	"bufio"
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
)
const (
	filePath = "./users.data"
	key = "woshifeiwu"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
//				  用户名   密码
type userHash map[string]string

type Checker struct {
	uhLock sync.RWMutex//来把锁
	uh userHash	// 用户信息
	rgLock sync.RWMutex//来把锁
	registerUsers []User // 注册了但未保存的用户
}
//登录
func (c *Checker) SignIn() {
	defer fix()

	fmt.Println("请输入用户名和密码")
	var username, password string
	fmt.Scan(&username, &password)
	c.uhLock.RLock()
	if _, ok := c.uh[username]; !ok {
		fmt.Println("查无此人")
		c.uhLock.RUnlock()
		return
	}
	if c.uh[username] != password {
		fmt.Println("用户名密码错误")
		c.uhLock.RUnlock()
		return
	}
	c.uhLock.RUnlock()

	fmt.Println("登录成功")
}
//注册
func (c *Checker) SignUp() {
	defer fix()

	fmt.Println("请输入用户名")
	var username, password string
	fmt.Scan(&username)
	c.uhLock.RLock()
	if _, ok := c.uh[username]; ok {
		fmt.Println("用户名已被占用")
		c.uhLock.RUnlock()
		return
	}
	c.uhLock.RUnlock()
	fmt.Println("请输入密码")
	for {
		fmt.Scan(&password)
		if len(password) >= 6 {
			break
		}
		fmt.Println("密码长度应大于六位，请重新输入")
	}

	c.rgLock.Lock()
	c.registerUsers = append(c.registerUsers, User{
		Username: username,
		Password: password,
	})
	c.rgLock.Unlock()

	c.rgLock.RLock()
	if len(c.registerUsers) > 10 {
		go c.Save()
	}
	c.rgLock.RUnlock()
	c.uhLock.Lock()
	c.uh[username] = password
	c.uhLock.Unlock()
}
//
func (c *Checker) Save() {
	defer fix()

	fail := saveUsers(c.registerUsers)
	c.rgLock.Lock()
	c.registerUsers = fail//?覆盖?//**\\
	c.rgLock.Unlock()
}

func initUsers() (userHash, error){
	defer fix()
//									如果不存在创建新文件|读写方式打开,0777权限
	f, err := os.OpenFile(filePath, os.O_CREATE | os.O_RDWR, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer f.Close()

	uh := make(userHash)
	var uhLock sync.Mutex
	var wg sync.WaitGroup // WaitGroup的作用是确保所有协程都执行完毕
	reader := bufio.NewReader(f)
	for {
		buf, err := reader.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println(err)
			return nil, err
		}

		wg.Add(1)
		go func(buf []byte) {
			defer wg.Done()

			arr := strings.Split(string(buf), ".")
			sign , err := base64.StdEncoding.DecodeString(arr[1])
			if err != nil {
				fmt.Println(err)
				return
			}

			mac := hmac.New(sha256.New, []byte(key))
			mac.Write([]byte(arr[0]))
			s := mac.Sum(nil)
			if res := bytes.Compare(sign,s); res != 0 {
				fmt.Println("data error")
				return
			}

			u, err := base64.StdEncoding.DecodeString(arr[0])
			if err != nil {
				fmt.Println(err)
				return
			}
			var user User
			err = json.Unmarshal(u, &user)
			if err != nil {
				fmt.Println(err)
				return
			}
			uhLock.Lock()
			uh[user.Username] = user.Password//**\\
			uhLock.Unlock()
		}(buf)

	}
	wg.Wait()
	return uh, nil
}

func saveUsers(users []User) (fail []User){
	var failLock sync.Mutex
	defer fix()

	f, err := os.OpenFile(filePath, os.O_WRONLY | os.O_CREATE | os.O_APPEND, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	var wg sync.WaitGroup // WaitGroup的作用是确保所有协程都执行完毕
	writer := bufio.NewWriter(f)
	for _, user := range users{
		wg.Add(1)
		go func(user User) {
			defer wg.Done()
			buf, err := json.Marshal(user)
			user64 := base64.StdEncoding.EncodeToString(buf)
			if err != nil {
				fmt.Println(err)
				failLock.Lock()
				fail = append(fail, user)
				failLock.Unlock()
				return
			}

			mac := hmac.New(sha256.New, []byte(key))
			mac.Write([]byte(user64))
			s := mac.Sum(nil)
			signature := base64.StdEncoding.EncodeToString(s)

			n, err := writer.Write(append([]byte(user64 + "." + signature), byte('\n')))
			if err != nil {
				fmt.Println(n, err)
				failLock.Lock()
				fail = append(fail, user)
				failLock.Unlock()
				return
			}
		} (user)
	}
	wg.Wait()
	writer.Flush()
	return
}

func showList() {
	fmt.Println("请选择操作：")
	fmt.Println("1、登录")
	fmt.Println("2、注册")
	fmt.Println("3、退出")
}

func main() {
	defer fix()
	checker := Checker{}
	var err error
	checker.uh, err = initUsers()
	if err != nil {
		return
	}


	var opt int
	for {
		showList()
		_, err := fmt.Scanln(&opt)
		if err != nil || opt < 1 || opt > 3 {
			fmt.Println("请输入正确的操作序号")
			continue
		}

		switch opt {
		case 1:
			checker.SignIn()
		case 2:
			checker.SignUp()
		case 3:
			checker.Save()
			return
		}
	}
}

func fix() {
	if err := recover(); err != nil {
		fmt.Println(err)
	}
}