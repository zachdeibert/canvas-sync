package task

// Alignment of some text
type Alignment int

const (
	// AlignLeft alignment
	AlignLeft Alignment = 0
	// AlignCenter alignment
	AlignCenter Alignment = 1
	// AlignRight alignment
	AlignRight Alignment = 2
)

// Section represents either a header or footer
type Section struct {
	text      [][]string
	listeners []func(*Section, [][]string)
}

// CreateSection creates a new Section
func CreateSection() *Section {
	return &Section{
		text:      [][]string{},
		listeners: []func(*Section, [][]string){},
	}
}

// SetSize sets the size of the section
func (s *Section) SetSize(rows int) {
	new := make([][]string, rows)
	copied := copy(new, s.text)
	for i := copied; i < len(new); i++ {
		new[i] = make([]string, 3)
	}
	s.text = new
	s.dispatch()
}

// SetText sets the text in a cell
func (s *Section) SetText(row int, alignment Alignment, text string) {
	s.text[row][alignment] = text
	s.dispatch()
}

// GetText gets the text in the section
func (s *Section) GetText() [][]string {
	return s.text
}

func (s *Section) dispatch() {
	for _, l := range s.listeners {
		l(s, s.text)
	}
}

// AddListener adds a new listener that's fired every time the text changes
func (s *Section) AddListener(listener func(*Section, [][]string)) {
	s.listeners = append(s.listeners, listener)
}
