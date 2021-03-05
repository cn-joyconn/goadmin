package global
// import(
// 	// "bytes"
// 	// "crypto/aes"
// 	// "crypto/cipher"
// 	// "crypto/rand"
// 	// "encoding/base64"
// 	// "encoding/hex"
// 	// "strings"

// 	// encrypt "joyconn-server-manage/lib/encrypt"
// 	sonyflake	"github.com/sony/sonyflake"
// )

// type LoginTokenID struct {
// 	token_ekey string  
//     uid string
// 	pwd string
// 	method string
// 	sonyflakeID  string
	
//     sonyflakeObj sonyflake.Sonyflake
// }
// // // func getkey() []byte{
// // // 	cookieEnKey,_:=config.String("joyconn.loginToken.cookieEnKey")	
// // // 	return  []byte(cookieEnKey) // 加密的密钥
// // // }
// // func Parse(token string,key string) LoginTokenID{
// // 	decrypted  := encrypt.AesDecryptCBC( []byte(token), []byte(key))
// // 	decryptStr :=	base64.StdEncoding.EncodeToString(decrypted )
// // 	ss :=strings.Split(decryptStr, "\f")
// // 	for val := range ss {
// // 		newMap[key] = value
// // 	} 

// // }
// // func (l *LoginTokenID)Creat(token string,key string) string{
// // 	ncrypted := encrypt.AesEncryptCBC( []byte(token), []byte(key))
// // }
// // flake := sonyflake.NewSonyflake(sonyflake.Settings{})
// //     id, err := flake.NextID()

