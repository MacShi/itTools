package utils
import (
	"encoding/hex"
	"fmt"
	"strings"
)

var unicodeMap = map[string]string{
	"0": "e280ac",
	"1": "e280ad",
}

func HidenInfo(info string) (string, error) {
	if info == "" {
		return "", nil
	}
	binaryStr := convertStrToBinaryStr(info)
	ret := ""
	for _, c := range binaryStr {
		val, err := unicodeMap[string(c)]
		if !err {
			return "", fmt.Errorf("invalid character in binary string: %d", c)
		}
		ret += convertHexStrToStr(val)
	}
	return ret, nil
}

func ShowInfo(info string) (string, error) {
	if info == "" {
		return "", nil
	}

	hexStr := convertStrToHexStr(info)

	indexStr0, ok0 := unicodeMap["0"]
	indexStr1, ok1 := unicodeMap["1"]
	if !ok0 || !ok1 {
		return "", fmt.Errorf("invalid unicode map")
	}

	binaryStr := ""
	var  beginIndex int
	var  endIndex int
	if strings.Index(hexStr,indexStr0)>strings.Index(hexStr,indexStr1){
		beginIndex =strings.Index(hexStr,unicodeMap["1"])
	}else {
		beginIndex =strings.Index(hexStr,unicodeMap["0"])
	}
	if strings.LastIndex(hexStr,unicodeMap["0"])>strings.LastIndex(hexStr,unicodeMap["1"]){
		endIndex =strings.LastIndex(hexStr,unicodeMap["0"])
	}else {
		endIndex =strings.LastIndex(hexStr,unicodeMap["1"])
	}

	for i := beginIndex; i <= endIndex; i += 6 {
		subStr := hexStr[i : i+6]
		switch subStr {
		case indexStr0:
			binaryStr += "0"
		case indexStr1:
			binaryStr += "1"
		default:
			return "", fmt.Errorf("invalid hex string: %s", subStr)
		}
	}

	return convertBinaryStrToStr(binaryStr), nil
}

func convertStrToBinaryStr(str string) string {
	byteInfo := []byte(str)

	var binaryStr strings.Builder
	for _, b := range byteInfo {
		biStr := fmt.Sprintf("%08b", b)
		binaryStr.WriteString(biStr)
	}

	return binaryStr.String()
}

func convertBinaryStrToStr(str string) string {
	if str == "" || len(str) <= 0 {
		return ""
	}

	size := len(str) / 8
	baKeyword := make([]uint8, size)

	for index := 0; index < size; index++ {
		b, err := convertBinaryToByte(str[index*8 : (index+1)*8])
		if err != nil {
			panic(err)
		}
		baKeyword[index] = b
	}

	return string(baKeyword)
}

func convertBinaryToByte(str string) (uint8, error) {
	var b uint8
	for i := 0; i < len(str); i++ {
		if str[i] == '1' {
			b |= 1 << (7 - i)
		} else if str[i] != '0' {
			return 0, fmt.Errorf("invalid binary string")
		}
	}
	return b, nil
}

func convertStrToHexStr(str string) string {
	byteInfo := []byte(str)

	hexStr := hex.EncodeToString(byteInfo)

	return hexStr
}

func convertHexStrToStr(str string) string {
	if str == "" || len(str) <= 0 {
		return ""
	}

	hexStr := strings.ReplaceAll(str, " ", "")
	byteInfo, err := hex.DecodeString(hexStr)
	if err != nil {
		panic(err)
	}

	return string(byteInfo)
}


