package functions

import (
	"errors"
	"fmt"
	"log"
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
				return errors.New("")
			}
		} else if ColorPattern.MatchString(args[i]) {
			if color := ColorCheck.FindStringSubmatch(args[i]); color != nil {
				if (Args.ColorFlag.Color != Color{}) {
					return fmt.Errorf("duplicate color flag")
				}
				if RgbPattern.MatchString(color[1]) {
					color, err := RGB(color[1])
					if err != nil {
						log.Fatalln(err)
					}
					Args.ColorFlag.Color = color
				} else if HexPattern.MatchString(color[1]) {
					color, err := HexToRgb(color[1])
					if err != nil {
						log.Fatalln(err)
					}
					Args.ColorFlag.Color = color

				} else {
					color, ok := Colors[strings.ToLower(color[1])]
					if !ok {
						log.Fatalln(errors.New("invalidColor"))
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
				if Args.Align != "" {
					return fmt.Errorf("duplicate flag:%s\n%s", justify[0], "--justify="+Args.Align)
				}
				Args.Align = justify[1]
			} else {
				return errors.New("")
			}
		} else {
			stringArguments = append(stringArguments, args[i])
		}
	}
	if len(stringArguments) == 0 {
		if Args.ColorFlag.ToColor != "" {
			Args.ToDraw = Args.ColorFlag.ToColor
		} else {
			return fmt.Errorf("missing input string")
		}
	} else if len(stringArguments) > 1 {
		return fmt.Errorf("invalid syntax: %v", stringArguments)
	} else {
		if Args.ToDraw == "" {
			Args.ToDraw = stringArguments[0]
			stringArguments = stringArguments[1:]
		}
	}
	// fmt.Println(Arguments)
	fmt.Println(stringArguments)
	return nil
}

// // FlagChecker checks if the provided flags are valid
// func FlagChecker() {
// 	if len(Params.Args) < 1 || len(Params.Args) > 4 {
// 		Params.Err = errors.New("color")
// 		return
// 	}
// 	if len(Params.Args) == 1 {
// 		if OutputPattern.MatchString(Params.Args[0]) {
// 			Params.Err = errors.New("output")
// 		} else if ColorPattern.MatchString(Params.Args[0]) {
// 			Params.Err = errors.New("color")
// 		}
// 		return
// 	}

// 	if OutputPattern.MatchString(Params.Args[0]) {
// 		if len(Params.Args) > 3 {
// 			Params.Err = errors.New("output")
// 			return
// 		}
// 		if output := OutputCheck.FindStringSubmatch(Params.Args[0]); output != nil {
// 			Params.OutputFile = output[1]
// 			Params.Args = Params.Args[1:]
// 			return
// 		}
// 		Params.Err = errors.New("output")
// 		return
// 	}

// 	if ColorPattern.MatchString(Params.Args[0]) {
// 		if color := ColorCheck.FindStringSubmatch(Params.Args[0]); color != nil {
// 			if strings.ToLower(color[1]) == "random" {
// 				RandomColor()
// 			}
// 			if RgbPattern.MatchString(color[1]) {
// 				RGB(color[1])
// 			} else if HexPattern.MatchString(color[1]) {
// 				HexToRgb(color[1])
// 			} else {
// 				colorCode, ok := Colors[strings.ToLower(color[1])]
// 				if !ok {
// 					Params.Err = errors.New("invalidColor")
// 					return
// 				}
// 				Params.Color = "\033[38;2;" + colorCode + "m"
// 			}

// 			Params.Args = Params.Args[1:]

// 			if len(Params.Args) == 2 {
// 				if !ValidBanner.MatchString(Params.Args[1]) {
// 					Params.ToColor = Params.Args[0]
// 					Params.Args = Params.Args[1:]
// 					return
// 				}
// 				Params.ToColor = Params.Args[0]
// 				return
// 			}

// 			if len(Params.Args) == 1 {
// 				Params.ToColor = Params.Args[0]
// 				return
// 			}

// 			Params.ToColor = Params.Args[0]
// 			Params.Args = Params.Args[1:]
// 			return
// 		}
// 		Params.Err = errors.New("color")
// 		return
// 	}
// }

// // ArgsChecker checks if the required arguments are provided
// func ArgsChecker() {
// 	if len(Params.Args) == 3 {
// 		Params.Err = errors.New("color")
// 	} else if len(Params.Args) == 1 {
// 		Params.Input = Params.Args[0]
// 		Params.Banner = "standard.txt"
// 	} else {
// 		Params.Input = Params.Args[0]
// 		if !regexp.MustCompile(`\.txt$`).MatchString(Params.Args[1]) {
// 			Params.Banner = Params.Args[1] + ".txt"
// 		} else {
// 			Params.Banner = Params.Args[1]
// 		}
// 	}
// }
