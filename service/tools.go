package service

import (
	"errors"
	"github.com/tuotoo/qrcode"
	"itTools/model"
	"itTools/utils"
	"os"
)

func HideStr(HideStr model.HideStrJson) (model.HideResult, error) {
	hideResult := model.HideResult{}
	if HideStr.Optiontype == "1" {
		hideStr := HideStr.HideStr
		carrierStr := HideStr.CarrierStr
		carrier := []rune(carrierStr)
		hide, err := utils.HidenInfo(hideStr)
		if err != nil {
			return hideResult, err
		}
		hideResult.HideResultStr = string(carrier[0]) + hide + string(carrier[1:])
		hideResult.HideResultLen = len(string(carrier[0]) + hide + string(carrier[1:]))
		return hideResult, nil
	} else if HideStr.Optiontype == "2" {
		hideResultStr := HideStr.HideResultStr
		hidestr,err:=utils.ShowInfo(hideResultStr)
		hideResult.HideStr = hidestr
		hideResult.HideResultLen = len([]rune(hidestr))
		if err!=nil{
			return hideResult,err
		}
		return hideResult,nil
	} else {
		return hideResult,errors.New("错误的操作类型")
	}

}

func QrDecode(filePath string) (string,error)  {
	// 从文件中读取 QR 码图像
	fi, err := os.Open(filePath)
	if err != nil{
		utils.Error(err.Error())
		return "",err
	}
	defer fi.Close()
	qrmatrix, err := qrcode.Decode(fi)
	if err != nil{
		utils.Error(err.Error())
		 return "",err
	}
	return qrmatrix.Content,nil
}
