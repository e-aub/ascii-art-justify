package functions

import (
	"errors"
	"strconv"
)

func ToColorIndexes() {
	if Args.ColorFlag.ToColor == "" {
		Args.ColorFlag.ToColorIndices = append(Args.ColorFlag.ToColorIndices, []int{0, len(Args.ToDraw) - 1})
		return
	}
	for i := 0; i < len(Args.ToDraw)-len(Args.ColorFlag.ToColor)+1; i++ {
		if Args.ToDraw[i:i+len(Args.ColorFlag.ToColor)] == Args.ColorFlag.ToColor {
			Args.ColorFlag.ToColorIndices = append(Args.ColorFlag.ToColorIndices, []int{i, i + len(Args.ColorFlag.ToColor) - 1})
			i += len(Args.ColorFlag.ToColor) - 1
		}
	}
}

func InRange(index int) bool {
	for _, pair := range Args.ColorFlag.ToColorIndices {
		if index >= pair[0] && index <= pair[1] {
			return true
		}
	}
	return false
}

func HexToRgb(hexColor string) (Color, error) {
	if HexCheck.MatchString(hexColor) {
		r, err := strconv.ParseInt(hexColor[1:3], 16, 64)
		if err != nil {
			return Color{}, err
		}
		g, err := strconv.ParseInt(hexColor[3:5], 16, 64)
		if err != nil {
			return Color{}, err
		}
		b, err := strconv.ParseInt(hexColor[5:7], 16, 64)
		if err != nil {
			return Color{}, err
		}
		return Color{R: int(r), G: int(g), B: int(b)}, nil
		// return fmt.Sprintf("\033[38;2;%d;%d;%dm", r, g, b), nil

	}
	return Color{}, errors.New("invalidhex")
}

func RGB(color string) (Color, error) {
	if match := RgbCheck.FindStringSubmatch(color); match != nil {
		r, err := strconv.Atoi(match[1])
		if err != nil {
			return Color{}, err
		}
		g, err := strconv.Atoi(match[2])
		if err != nil {
			return Color{}, err
		}
		b, err := strconv.Atoi(match[3])
		if err != nil {
			return Color{}, err
		}
		if r > 255 || g > 255 || b > 255 {
			return Color{}, errors.New("InvalidRgbValue")
		}
		// return fmt.Sprintf("\033[38;2;%d;%d;%dm", r, g, b), nil
		return Color{R: r, G: g, B: b}, nil
	}
	return Color{}, errors.New("rgbFormat")
}
