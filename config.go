package stringsplit

import (
	"errors"
)

var (
	ErrNotFoundSection = errors.New("Not found section")
)

type Configuration struct {
	Delimiter string
	Sections  Sections
}

func NewConfiguration(delimiter string) Configuration {
	return Configuration{
		Delimiter: delimiter,
		Sections:  Sections{},
	}
}

func (c *Configuration) Append(begin, end string) {
	c.Sections = append(c.Sections, NewSectionString(begin, end))
}

func (c *Configuration) FindSectionByBeginString(s string) (*Section, error) {
	for _, item := range c.Sections {
		if !item.EqualBeginString(s) {
			continue
		}

		return item, nil
	}

	return nil, ErrNotFoundSection
}

func (c *Configuration) GetBeginStrings() []string {
	ret := make([]string, 0, len(c.Sections))

	for _, item := range c.Sections {
		ret = append(ret, (*item).Begin)
	}

	return ret
}
