package beer

import (
	"errors"
	"fmt"
	"strings"
)

func Verse(verse int) (string, error) {
	if verse > 99 || verse < 0 {
		return "", errors.New("error")
	} else if verse == 2 {
		return "2 bottles of beer on the wall, 2 bottles of beer.\nTake one down and pass it around, 1 bottle of beer on the wall.\n", nil
	} else if verse == 1 {
		return "1 bottle of beer on the wall, 1 bottle of beer.\nTake it down and pass it around, no more bottles of beer on the wall.\n", nil
	} else if verse == 0 {
		return "No more bottles of beer on the wall, no more bottles of beer.\nGo to the store and buy some more, 99 bottles of beer on the wall.\n", nil
	} else {
		return fmt.Sprintf("%d bottles of beer on the wall, %d bottles of beer.\nTake one down and pass it around, %d bottles of beer on the wall.\n", verse, verse, verse - 1), nil
	}
}

func Verses(upperBound int, lowerBound int) (string, error) {
	if upperBound > 99 || upperBound < 0 || lowerBound > 99 || lowerBound < 0 || upperBound < lowerBound{
		return "", errors.New("error")
	}
	text := make([]string, 0)
	for ; upperBound >= lowerBound; upperBound-- {
		s, err := Verse(upperBound)
		if err == nil {
			text = append(text, s + "\n")
		} else {
			return "", errors.New("error")
		}
	}
	return strings.Join(text, ""), nil
}

func Song() string {
	res, err := Verses(99, 0)
	if err == nil {
		return res
	} else {
		return "error"
	}
}