package functions

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func RandomColor() {
	source := rand.NewSource(time.Now().UnixNano())
	generator := rand.New(source)
	Colors["random"] = fmt.Sprintf("%d;%d;%d", generator.Intn(255), generator.Intn(255), generator.Intn(255))
}

func ToColorIndexes() {
	if Arguments.color.toColor == "" {
		Arguments.color.toColorIndices = nil
		return
	}
	for i := 0; i < len(Arguments.ToDraw)-len(Arguments.color.toColor)+1; i++ {
		if Arguments.ToDraw[i:i+len(Arguments.color.toColor)] == Arguments.color.toColor {
			Arguments.color.toColorIndices = append(Arguments.color.toColorIndices, []int{i, i + len(Arguments.color.toColor) - 1})
			i += len(Arguments.color.toColor) - 1
		}
	}
}

func InRange(index int) bool {
	for _, pair := range Arguments.color.toColorIndices {
		if index >= pair[0] && index <= pair[1] {
			return true
		}
	}
	return false
}

func HexToRgb(hexColor string) (string, error) {
	if HexCheck.MatchString(hexColor) {
		r, err := strconv.ParseInt(hexColor[1:3], 16, 64)
		if err != nil {
			return "", err
		}
		g, err := strconv.ParseInt(hexColor[3:5], 16, 64)
		if err != nil {
			return "", err
		}
		b, err := strconv.ParseInt(hexColor[5:7], 16, 64)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("\033[38;2;%d;%d;%dm", r, g, b), nil

	}
	return "", errors.New("invalidhex")
}

func RGB(color string) (string, error) {
	if match := RgbCheck.FindStringSubmatch(color); match != nil {
		r, err := strconv.Atoi(match[1])
		if err != nil {
			return "", err
		}
		g, err := strconv.Atoi(match[2])
		if err != nil {
			return "", err
		}
		b, err := strconv.Atoi(match[3])
		if err != nil {
			return "", err
		}
		if r > 255 || g > 255 || b > 255 {
			return "", errors.New("InvalidRgbValue")

		}
		return fmt.Sprintf("\033[38;2;%d;%d;%dm", r, g, b), nil
	}
	return "", errors.New("rgbFormat")
}
