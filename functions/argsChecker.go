package functions

import (
	"errors"
	"fmt"
	"strings"
)

func ArgsChecker(args []string) error {
	var stringArguments []string
	if len(args) > 1 && ValidBanner.MatchString(args[len(args)-1]) {
		if !strings.HasSuffix(args[len(args)-1], ".txt") {
			Args.Banner = args[len(args)-1] + ".txt"
		} else {
			Args.Banner = args[len(args)-1]
		}
		args = args[:len(args)-1]
	} else {
		Args.Banner = "standard.txt"
	}
	for i := 0; i < len(args); i++ {
		if OutputPattern.MatchString(args[i]) {
			if output := OutputCheck.FindStringSubmatch(args[i]); output != nil {
				if Args.FileName != "" {
					return fmt.Errorf("duplicate flag:%s\n%s", output[0], "--output="+Args.FileName)
				}
				Args.FileName = output[1]
			} else {
				return errors.New("output")
			}
		} else if ColorPattern.MatchString(args[i]) {
			if color := ColorCheck.FindStringSubmatch(args[i]); color != nil {
				if (Args.ColorFlag.Color != Color{}) {
					return fmt.Errorf("duplicate color flag")
				}
				if RgbPattern.MatchString(color[1]) {
					color, err := RGB(color[1])
					if err != nil {
						ErrHandler(err)
					}
					Args.ColorFlag.Color = color
				} else if HexPattern.MatchString(color[1]) {
					color, err := HexToRgb(color[1])
					if err != nil {
						ErrHandler(err)
					}
					Args.ColorFlag.Color = color

				} else {
					color, ok := Colors[strings.ToLower(color[1])]
					if !ok {
						ErrHandler(errors.New("invalidColor"))
					}
					Args.ColorFlag.Color = color
				}

			}
			if i+1 < len(args) {
				validOutput := OutputCheck.MatchString(args[i+1])
				validBanner := ValidBanner.MatchString(args[i+1])
				validJustify := JustifyCheck.MatchString(args[i+1])
				if !validBanner && !validOutput && !validJustify {
					i++
					Args.ColorFlag.ToColor = args[i]
				} else {
					continue
				}
			}
		} else if JustifyPattern.MatchString(args[i]) {
			if justify := JustifyCheck.FindStringSubmatch(args[i]); justify != nil {
				if Args.AlignFlag.Align != "" {
					return fmt.Errorf("duplicate justify flag")
				}
				Args.AlignFlag.Align = justify[1]
			} else {
				return errors.New("justify")
			}
		} else {
			stringArguments = append(stringArguments, args[i])
		}
	}
	if len(stringArguments) == 0 {
		if Args.ColorFlag.ToColor != "" {
			Args.ToDraw = Args.ColorFlag.ToColor
		} else {
			return fmt.Errorf("global")
		}
	} else if len(stringArguments) == 1 {
		if Args.ToDraw == "" {
			Args.ToDraw = stringArguments[0]
		}
	} else if len(stringArguments) >= 2 {
		return fmt.Errorf("global")
	}
	return nil
}
