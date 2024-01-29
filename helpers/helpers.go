package helpers

import (
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net"
	"net/smtp"
	"reflect"
)

type TeamJson struct {
	TeamName   string   `json:"teamname"`
	Members    []string `json:"members"`
	Emails     []string `json:"emails"`
	LanguageID int      `json:"languageID"`
	Ai         bool     `json:"ai"`
}

type UserJson struct {
	UserName string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type Competition struct {
	Description string `firestore:"description"`
	EName       string `firestore:"ename"`
}

func GenerateToken() string {
	//Random (10-1 +1) is 0-9 9-0
	//Random (max-min+1)+min
	//35 - 122
	//cant use 39 40 41 44 45 46 58 59 91 - 96
	//65 - 90
	//49 - 57
	var token = ""
	var num = 0
	for i := 0; i < 40; i++ {
		num = rand.Intn(90-49+1) + 49

		if (num < 57 && num > 49) || (num < 90 && num > 65) {
			token = token + string(rune(num)) //rune!!!
		} else {
			i--
		}

	}
	return token
}

func GeneratePassword() string {
	var password = ""
	var num = 0
	for i := 0; i < 15; i++ {
		num = rand.Intn(90-49+1) + 49

		if (num < 57 && num > 49) || (num < 90 && num > 65) {
			password = password + string(rune(num))
		} else {
			i--
		}

	}
	return password
}
func ParseRegisterDataForEmail(data TeamJson) []string {

	var email []string
	email = data.Emails
	return email
}
func EncodeToBase64(code string) string {
	encoded := base64.StdEncoding.EncodeToString([]byte(code))
	return encoded
}
func DecodeFromBase64(code string) string {
	decoded, err := base64.StdEncoding.DecodeString(code)
	if err != nil {
		fmt.Println("error:", err)
		return ""
	}
	return string(decoded)

}
func IpAddress() string {
	ifaces, err := net.Interfaces()
	if err != nil {
		fmt.Println(err)
		return ""
	}

	for _, iface := range ifaces {

		if iface.Flags&net.FlagUp != 0 && iface.Flags&net.FlagLoopback == 0 {
			addrs, err := iface.Addrs()
			if err != nil {
				fmt.Println(err)
				continue
			}

			for _, addr := range addrs {

				if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() && ipnet.IP.To4() != nil {
					fmt.Println(ipnet.IP.String())
					return ipnet.IP.String()
				}
			}
		}
	}

	fmt.Println("IPv4 address not found")
	return "localhost"
}
func SendEmail(email string, password string, name string) {
	//https://www.youtube.com/watch?v=pAPWBHxnFHM
	auth := smtp.PlainAuth("", "vismatestemail@gmail.com", "", "smtp.gmail.com")
	to := []string{email}
	msg := []byte("To: " + email + "\r\n" +
		"Subject: Login\r\n" +
		"\r\n" +
		"This is a your credentials : \n" +
		"name :  " + name + "\n" +
		"password : " + password)

	// Send the email.
	err := smtp.SendMail("smtp.gmail.com:587", auth, "vismatestemail@gmail.com", to, msg)
	if err != nil {
		panic(err)
	}

}

//	func ContainsValue(m map[string]string, value string) bool {
//		for _, v := range m {
//			if v == value {
//				return true
//			}
//		}
//		return false
//	}
func ContainsValue(m map[UserJson]string, value string) bool {
	for _, v := range m {
		if v == value {
			return true
		}
	}
	return false
}

func GetKeyByValue(m map[UserJson]string, value string) UserJson {
	for key, val := range m {
		if val == value {
			return key
		}
	}
	return UserJson{}
}

//	func FindKeyByValue(m map[UserJson]string, value string) string {
//		for k, v := range m {
//			if v == value {
//				return k.UserName
//			}
//		}
//		return ""
//	}
func FindKeyByValue(m map[UserJson]string, value string) UserJson {
	for k, v := range m {
		if v == value {
			return k
		}
	}
	return UserJson{}
}

func ContainsKey(m interface{}, key string) bool {
	v := reflect.ValueOf(m)
	if v.Kind() != reflect.Map {
		return false
	}

	iter := v.MapRange()
	for iter.Next() {
		if reflect.DeepEqual(iter.Key().Interface(), key) {
			return true
		}
	}

	return false
}
func GetUserByToken(tokenMap map[UserJson]string, token string) (UserJson, bool) {
	for user, t := range tokenMap {
		if t == token {
			return user, true
		}
	}
	return UserJson{}, false
}
func ValidateAndAuthorizeAdmin(context *gin.Context, tokenMap map[UserJson]string) error {
	token := context.GetHeader("token")

	user, isInMap := GetUserByToken(tokenMap, token)
	if !isInMap {
		return errors.New("invalid token")
	}
	if len(user.Role) == 0 {
		return errors.New("no role assigned")
	}

	if user.Role != "admin" {
		return errors.New("user is unauthorised to do this")
	}

	return nil
}
