package stringsplit

import (
	"errors"
)

var (
	ErrNotFoundSection = errors.New("Not found section")
)

type Configuration struct {
	Delimiter string
	Sections  []Section
}

type Section struct {
	Begin string
	End   string

	BeginIndex int
	EndIndex   int
}

func NewConfiguration(delimiter string) Configuration {
	return Configuration{
		Delimiter: delimiter,
		Sections:  make([]Section, 0, 8),
	}
}

func (c *Configuration) Append(begin, end string) {
	c.Sections = append(c.Sections, Section{Begin: begin, End: end})
}

func (c *Configuration) FindSectionByBeginString(s string) (Section, error) {
	for _, item := range c.Sections {
		if s != item.Begin {
			continue
		}

		return item, nil
	}

	return Section{}, ErrNotFoundSection
}

func (c *Configuration) GetBeginStrings() []string {
	ret := make([]string, 0, len(c.Sections))

	for _, item := range c.Sections {
		ret = append(ret, item.Begin)
	}

	return ret
}
