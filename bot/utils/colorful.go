package utils

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

// RGBToANSI 将RGB颜色代码转换成ANSI颜色代码
func RGBToANSI(rgb string) (int, int, int) {
	// 如果RGB字符串以#开头则去掉
	if strings.HasPrefix(rgb, "#") {
		rgb = rgb[1:]
	}

	// 将RGB字符串转换为整数
	r, err := strconv.ParseInt(rgb[0:2], 16, 64)
	if err != nil {
		panic("Invalid red value")
	}
	g, err := strconv.ParseInt(rgb[2:4], 16, 64)
	if err != nil {
		panic("Invalid green value")
	}
	b, err := strconv.ParseInt(rgb[4:6], 16, 64)
	if err != nil {
		panic("Invalid blue value")
	}

	return int(r), int(g), int(b)
}

// GenerateGradientString 生成横向渐变颜色字符串
func GenerateGradientString(s string, colors ...string) string {
	if len(colors) < 2 {
		panic("At least two colors are required")
	}

	// 解析所有颜色
	rgbColors := make([][3]int, len(colors))
	for i, color := range colors {
		r, g, b := RGBToANSI(color)
		rgbColors[i] = [3]int{r, g, b}
	}

	// 计算渐变
	n := utf8.RuneCountInString(s)
	segmentLength := n / (len(colors) - 1)
	result := ""

	runes := []rune(s)
	for i := 0; i < n; i++ {
		segmentIndex := i / segmentLength
		if segmentIndex >= len(colors)-1 {
			segmentIndex = len(colors) - 2
		}

		startColor := rgbColors[segmentIndex]
		endColor := rgbColors[segmentIndex+1]

		r := interpolate(startColor[0], endColor[0], segmentLength, i%segmentLength)
		g := interpolate(startColor[1], endColor[1], segmentLength, i%segmentLength)
		b := interpolate(startColor[2], endColor[2], segmentLength, i%segmentLength)

		result += fmt.Sprintf("\033[38;2;%d;%d;%dm%c", r, g, b, runes[i])
	}

	return result + "\033[0m"
}

// interpolate 计算渐变中间的颜色值
func interpolate(start, end, length, position int) int {
	return start + (end-start)*position/length
}

// RGBTextColor 设置字体颜色
func RGBTextColor(rgb string) string {
	// 复用函数
	r, g, b := RGBToANSI(rgb)

	// 返回ANSI颜色代码
	return fmt.Sprintf("\033[38;2;%d;%d;%dm", r, g, b)
}

// 定义一系列ANSI颜色代码

var ResetColor = "\033[0m"
var HihglightColor = "\033[1m"
var UnderlineColor = "\033[4m"
var BlinkColor = "\033[5m"
var ReverseColor = "\033[7m"
var HiddenColor = "\033[8m"

var ClearScreen = "\033[2J"
var ClearLine = "\033[K"

var Red = "\033[31m"
var RedBackground = "\033[41m"
var Orange = "\033[33m"
var OrangeBackground = "\033[43m"
var Yellow = "\033[93m"
var YellowBackground = "\033[103m"
var Green = "\033[32m"
var GreenBackground = "\033[42m"
var Cyan = "\033[36m"
var CyanBackground = "\033[46m"
var Blue = "\033[34m"
var BlueBackground = "\033[44m"
var Purple = "\033[35m"
var PurpleBackground = "\033[45m"
var White = "\033[37m"
var WhiteBackground = "\033[47m"
var Black = "\033[30m"
var BlackBackground = "\033[40m"
var Gray = "\033[90m"
var GrayBackground = "\033[100m"
var LightRed = "\033[91m"
var LightRedBackground = "\033[101m"
var LightGreen = "\033[92m"
var LightGreenBackground = "\033[102m"
var LightYellow = "\033[93m"
var LightYellowBackground = "\033[103m"
var LightBlue = "\033[94m"
var LightBlueBackground = "\033[104m"
var LightPurple = "\033[95m"
var LightPurpleBackground = "\033[105m"
var LightCyan = "\033[96m"
var LightCyanBackground = "\033[106m"
var LightGray = "\033[97m"
var LightGrayBackground = "\033[107m"
var DarkGray = "\033[90m"
var DarkGrayBackground = "\033[100m"
var DarkRed = "\033[31m"
var DarkRedBackground = "\033[41m"
var DarkGreen = "\033[32m"
var DarkGreenBackground = "\033[42m"
var DarkYellow = "\033[33m"
var DarkYellowBackground = "\033[43m"
var DarkBlue = "\033[34m"
var DarkBlueBackground = "\033[44m"
var DarkPurple = "\033[35m"
var DarkPurpleBackground = "\033[45m"
var DarkCyan = "\033[36m"
var DarkCyanBackground = "\033[46m"
var Pink = "\033[95m"
var PinkBackground = "\033[105m"
