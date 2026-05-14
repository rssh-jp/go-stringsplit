package stringsplit

import (
	"strings"
)

// ExecuteSimple is execute 1 begin-end sections configure
func ExecuteSimple(str, delimiter, begin, end string) ([]string, error) {
	c := NewConfiguration(delimiter)

	c.Append(begin, end)

	return Execute(str, c)
}

// Execute splits str by config.Delimiter, ignoring delimiters that appear
// inside any registered section (begin..end pairs).
func Execute(str string, config Configuration) ([]string, error) {
	secs := Sections{}

	workindex := 0

	for workindex < len(str) {
		sec, err := findSection(str[workindex:], config)
		if err != nil {
			return nil, err
		}

		if sec == nil {
			break
		}

		sec.BeginIndex += workindex
		sec.EndIndex += workindex

		secs = append(secs, sec)

		workindex = sec.EndIndex + 1
	}

	ret := make([]string, 0, 8)

	workindex = 0
	sindex := 0

	for workindex < len(str) {
		index := strings.Index(str[workindex:], config.Delimiter)
		if index < 0 {
			break
		}

		endindex := workindex + index

		if secs.IsInIndex(endindex) {
			workindex = endindex + 1
			continue
		}

		ret = append(ret, str[sindex:endindex])

		workindex = endindex + len(config.Delimiter)
		sindex = workindex
	}

	ret = append(ret, str[sindex:])

	return ret, nil
}

func findSection(str string, config Configuration) (*Section, error) {
	beginindex, s := findFirstIndex(str, config.GetBeginStrings())
	if beginindex < 0 {
		return nil, nil
	}

	section, err := config.FindSectionByBeginString(s)
	if err != nil {
		return nil, err
	}

	beginLen := len(section.Begin)
	endindex, _ := findFirstIndex(str[beginindex+beginLen:], []string{section.End})
	if endindex < 0 {
		return nil, nil
	}

	// EndIndex は end 文字列の最後の文字位置（inclusive）
	return NewSectionIndex(beginindex, beginindex+beginLen+endindex+len(section.End)-1), nil
}

func findFirstIndex(str string, substrings []string) (int, string) {
	indexes := make([]int, len(substrings))
	for i, substring := range substrings {
		index := strings.Index(str, substring)
		indexes[i] = index
	}

	var workIndex int = -1
	var workSubstring string
	for i, item := range indexes {
		if item < 0 {
			continue
		}

		if workIndex != -1 && item >= workIndex {
			continue
		}

		workIndex = item
		workSubstring = substrings[i]
	}

	return workIndex, workSubstring
}
