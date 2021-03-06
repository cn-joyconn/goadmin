package joyCaptcha

import (
	"image/color"

	"github.com/cn-joyconn/goadmin/models/global"
	gocache "github.com/cn-joyconn/gocache"
	"github.com/mojocn/base64Captcha"
)

//configJsonBody json request body.
type CaptchaParam struct {
	CaptchaType string //audio、string、math、chinese、digit

	Length int //Length random string length.

	//audio
	Language string // Language possible values for lang are "en", "ja", "ru", "zh".

	//string
	Height          int         // Height png height in pixel.
	Width           int         // Width Captcha png width in pixel.
	NoiseCount      int         //NoiseCount text noise count.
	ShowLineOptions int         //  OptionShowHollowLine=2 | OptionShowSlimeLine=4 | OptionShowSineLine=8 .
	Source          string      //Source is a unicode which is the rand string from.
	BgColor         *color.RGBA //BgColor captcha image background color (optional)
	// Fonts           []string    //Fonts loads by name see fonts.go's comment

	// Digit
	MaxSkew  float64 // MaxSkew max absolute skew factor of a single digit.
	DotCount int     // DotCount Number of background circles.
}

type JoyCaptchaStore struct {
	StoreCacheObj *gocache.Cache
}

func (store *JoyCaptchaStore) getCacheKey(id string) string {
	return "Captcha_" + id
}
func (store *JoyCaptchaStore) Set(id string, value string) {
	store.StoreCacheObj.Put(store.getCacheKey(id), value, 1000*60*30) //30分钟有效
}
func (store *JoyCaptchaStore) Get(id string, clear bool) string {
	var value string
	store.StoreCacheObj.Get(store.getCacheKey(id), &value)
	if clear {
		store.StoreCacheObj.Delete(store.getCacheKey(id))
	}
	return value
}
func (store *JoyCaptchaStore) Verify(id, answer string, clear bool) bool {
	var value string
	store.StoreCacheObj.Get(store.getCacheKey(id), &value)
	if clear {
		store.StoreCacheObj.Delete(store.getCacheKey(id))
	}
	return value == answer
}
func InitCaptcha() {
	store.StoreCacheObj = &gocache.Cache{
		Catalog:   global.AdminCatalog,
		CacheName: global.AdminCacheName,
	}
}

var store = &JoyCaptchaStore{}

// base64Captcha create http handler
func GenerateCaptcha(param *CaptchaParam) (string, string, error) {
	//parse request parameters
	var driver base64Captcha.Driver
	//id, content, answer := c.Driver.GenerateIdQuestionAnswer()
	//create base64 encoding captcha
	switch param.CaptchaType {
	case "audio":
		driver = base64Captcha.NewDriverAudio(param.Length, param.Language)
	case "string":
		_fonts := []string{"ApothecaryFont.ttf", "3Dumb.ttf", "DENNEthree-dee.ttf", "Flim-Flam.ttf", "RitaSmith.ttf", "actionj.ttf"}
		// _fonts := []string{"chromohv.ttf"}
		driver = base64Captcha.NewDriverString(param.Height, param.Width, param.NoiseCount, param.ShowLineOptions, param.Length, param.Source, param.BgColor, _fonts)
	case "math":
		driver = base64Captcha.NewDriverMath(param.Height, param.Width, param.NoiseCount, param.ShowLineOptions, param.BgColor, []string{"3Dumb.ttf"})
	case "chinese":
		driver = base64Captcha.NewDriverChinese(param.Height, param.Width, param.NoiseCount, param.ShowLineOptions, param.Length, param.Source, param.BgColor, []string{"wqy-microhei.ttf"})
	default:
		driver = base64Captcha.NewDriverDigit(param.Height, param.Width, param.Length, param.MaxSkew, param.DotCount)

	}
	c := base64Captcha.NewCaptcha(driver, store)
	id, b64s, err := c.Generate()
	return id, b64s, err
}

// base64Captcha verify http handler
func CaptchaVerify(id string, code string) bool {
	return store.Verify(id, code, true)
}
