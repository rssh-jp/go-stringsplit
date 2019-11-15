package stringsplit

type Section struct {
	Begin string
	End   string

	BeginIndex int
	EndIndex   int
}

func NewSectionString(begin, end string) *Section {
	return &Section{
		Begin: begin,
		End:   end,
	}
}

func NewSectionIndex(begin, end int) *Section {
	return &Section{
		BeginIndex: begin,
		EndIndex:   end,
	}
}

func (s Section) EqualBeginString(str string) bool {
	return s.Begin == str
}

func (s Section) EqualEndString(str string) bool {
	return s.End == str
}

func (s Section) IsInIndex(index int) bool {
	return s.BeginIndex <= index && index <= s.EndIndex
}

type Sections []*Section

func (sections Sections) IsInIndex(index int) bool {
	for _, section := range sections {
		if section.IsInIndex(index) {
			return true
		}
	}

	return false
}
