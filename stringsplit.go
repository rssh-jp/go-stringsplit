package stringsplit

import (
	"strings"
	"time"
)

type Sections []Section

func (sections Sections) IsIn(index int) bool {
	for _, section := range sections {
		if section.BeginIndex <= index && index <= section.EndIndex {
			return true
		}
	}

	return false
}

func Exec(str, delimiter, begin, end string) []string {
	c := NewConfiguration(delimiter)

	c.Append(begin, end)

	return Execute(str, c)
}

func Execute(str string, config Configuration) []string {
	secs := Sections{}

	begins := config.GetBeginStrings()

	workindex := 0

	for workindex < len(str) {
		sec, err := func(str string) (*Section, error) {
			beginindex, s := firstIndex(str, begins)
			if beginindex < 0 {
				return nil, nil
			}

			section, err := config.FindSectionByBeginString(s)
			if err != nil {
				return nil, err
			}

			endindex, _ := firstIndex(str[beginindex+1:], []string{section.End})
			if endindex < 0 {
				return nil, nil
			}

			return &Section{BeginIndex: beginindex, EndIndex: beginindex + 1 + endindex}, nil
		}(str[workindex:])

		if err != nil {
			return nil
		}

		if sec == nil {
			break
		}

		sec.BeginIndex += workindex
		sec.EndIndex += workindex

		secs = append(secs, *sec)

		workindex = sec.EndIndex + 1
	}

	ret := make([]string, 0, 8)

	workindex = 0
	sindex := 0

	for workindex < len(str) {
		index := strings.Index(str[workindex:], config.Delimiter)
		if index < 0 {
			ret = append(ret, str[workindex:])
			break
		}

		endindex := workindex + index

		if secs.IsIn(endindex) {
			workindex = endindex + 1
			continue
		}

		s := string(str[sindex:endindex])

		ret = append(ret, s)

		workindex = endindex + 1
		sindex = workindex

		time.Sleep(time.Millisecond * 100)
	}

	return ret
}

func firstIndex(str string, substrings []string) (int, string) {
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
