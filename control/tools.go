package control

import (
	"github.com/kataras/iris"
	"itTools/config"
	"itTools/model"
	"itTools/service"
	"itTools/utils"
	"math/rand"
	"mime/multipart"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"
)

func HideStr(ctx iris.Context) {
	HideStrJson := model.HideStrJson{}
	err := ctx.ReadJSON(&HideStrJson)
	ctx.Header("Access-Control-Allow-Origin", "*")
	if nil != err {
		ctx.JSON(utils.ResultUtil.Failure(err.Error()))
	} else {
		a, err := service.HideStr(HideStrJson)
		if nil != err {
			ctx.JSON(utils.ResultUtil.Failure(err.Error()))
		} else {
			//ctx.Header("Access-Control-Allow-Origin","*")
			ctx.JSON(utils.ResultUtil.Success(a))
		}

	}
}
func ClientInfo(ctx iris.Context) {
	//优先级X-Real-IP>X-Forwarded-For>ctx.RemoteAddr(tcp数据包IP地址)
	clientInfo := model.ClientInfo{}
	clientInfo.UserAgent = ctx.GetHeader("User-Agent")
	realIp := ctx.GetHeader("X-Real-IP")
	xForwarded := strings.Split(ctx.GetHeader("X-Forwarded-For"), ",")
	if len(realIp) > 0 {
		clientInfo.RemoteIp = realIp
	} else if xForwarded[0] != "" {
		clientInfo.RemoteIp = xForwarded[0]
	} else {
		clientInfo.RemoteIp = ctx.RemoteAddr()
	}
	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.JSON(utils.ResultUtil.Success(clientInfo))
}

func GetPwd(ctx iris.Context) {
	//'小写字母', '大写字母', '数字','特殊字符'
	pwd := model.GenPwd{}
	//var pwd map[string]interface{}
	err := ctx.ReadJSON(&pwd)
	//fmt.Println("hhhh",pwd)
	if err != nil {
		ctx.JSON(utils.ResultUtil.Failure(err.Error()))
	}
	str1 := "+=/-.,?@#$%^&*()"
	str2 := "1234567890"
	str4 := "ABCDEFGHIJKLMNPQRSTUVWXYZ"
	str8 := "abcdefghjkmnpqrstuvwxyz"

	switch pwd.PwdType {
	case 8:
		str08 := []rune(str8)
		ctx.JSON(utils.ResultUtil.Success(genRandomStr(pwd.PwdLength, str08)))
	case 4:
		str04 := []rune(str4)
		ctx.JSON(utils.ResultUtil.Success(genRandomStr(pwd.PwdLength, str04)))
	case 2:
		str02 := []rune(str2)
		ctx.JSON(utils.ResultUtil.Success(genRandomStr(pwd.PwdLength, str02)))
	case 1:
		str01 := []rune(str1)
		ctx.JSON(utils.ResultUtil.Success(genRandomStr(pwd.PwdLength, str01)))
	case 3:
		str3 := []rune(str1 + str2)
		ctx.JSON(utils.ResultUtil.Success(genRandomStr(pwd.PwdLength, str3)))
	case 5:
		str5 := []rune(str1 + str4)
		ctx.JSON(utils.ResultUtil.Success(genRandomStr(pwd.PwdLength, str5)))
	case 6:
		str6 := []rune(str2 + str4)
		ctx.JSON(utils.ResultUtil.Success(genRandomStr(pwd.PwdLength, str6)))
	case 7:
		str7 := []rune(str1 + str2 + str4)
		ctx.JSON(utils.ResultUtil.Success(genRandomStr(pwd.PwdLength, str7)))
	case 9:
		str9 := []rune(str1 + str8)
		ctx.JSON(utils.ResultUtil.Success(genRandomStr(pwd.PwdLength, str9)))
	case 10:
		str10 := []rune(str2 + str8)
		ctx.JSON(utils.ResultUtil.Success(genRandomStr(pwd.PwdLength, str10)))
	case 11:
		str11 := []rune(str1 + str2 + str8)
		ctx.JSON(utils.ResultUtil.Success(genRandomStr(pwd.PwdLength, str11)))
	case 12:
		str12 := []rune(str4 + str8)
		ctx.JSON(utils.ResultUtil.Success(genRandomStr(pwd.PwdLength, str12)))
	case 13:
		str13 := []rune(str1 + str4 + str8)
		ctx.JSON(utils.ResultUtil.Success(genRandomStr(pwd.PwdLength, str13)))
	case 14:
		str14 := []rune(str2 + str4 + str8)
		ctx.JSON(utils.ResultUtil.Success(genRandomStr(pwd.PwdLength, str14)))
	case 15:
		str15 := []rune(str1 + str2 + str4 + str8)
		ctx.JSON(utils.ResultUtil.Success(genRandomStr(pwd.PwdLength, str15)))
	}

}

func In(target string, str_array []string) bool {
	sort.Strings(str_array)
	index := sort.SearchStrings(str_array, target)
	if index < len(str_array) && str_array[index] == target {
		return true
	}
	return false
}

func genRandomStr(n int64, str []rune) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = str[rand.Intn(len(str))]
	}
	return string(b)
}
func UploadFile(ctx iris.Context) {
	saveFilePath := config.Conf.Get("app.upload.allowDir").(string)
	file, info, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(utils.ResultUtil.Failure("Error getting the file"))
		return
	}
	defer file.Close()
	if checkFileIsImage(info) {
		_, err1 := os.Open(saveFilePath)
		if err1 != nil {
			os.Mkdir(saveFilePath, 0660)
		}
		//destFile := saveFilePath + strconv.FormatInt(time.Now().Unix(), 10) + "_" + info.Filename
		destFile := filepath.Join(saveFilePath,strconv.FormatInt(time.Now().Unix(), 10) + "_" + info.Filename)
		//utils.Info(destFile1)
		_, err := ctx.SaveFormFile(info, destFile)
		if err != nil {
			ctx.JSON(utils.ResultUtil.Failure(err.Error()))
		}
		result, err := service.QrDecode(destFile)
		if err != nil {
			ctx.JSON(utils.ResultUtil.Failure("failed"))
		} else {
			ctx.JSON(utils.ResultUtil.Success(result))
		}

	} else {
		ctx.JSON(utils.ResultUtil.Failure("failed"))
	}
}

func checkFileIsImage(info *multipart.FileHeader) bool {
	allowTypeList := []string{"image/jpeg", "image/png"}
	allowType := In(info.Header.Get("Content-Type"), allowTypeList)
	fileNameSplit := strings.Split(info.Filename, ".")
	allowSuffixList := config.Conf.Get("app.upload.allowSuffix").([]interface{})
	var suffixes []string
	for _, elem := range allowSuffixList {
		// 将每个元素转换为 string 类型
		if str, ok := elem.(string); ok {
			suffixes = append(suffixes, str)
		} else {
			utils.Error("Element is not a string:"+elem.(string))
		}
	}
	allowSuffix :=In(fileNameSplit[len(fileNameSplit)-1], suffixes)
	if allowType && allowSuffix{
		return true
	} else {
		return false
	}
}
