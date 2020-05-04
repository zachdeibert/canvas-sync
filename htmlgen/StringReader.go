package htmlgen

import "io"

type stringReader struct {
	data  string
	count int
}

func (s *stringReader) Read(p []byte) (int, error) {
	if s.count >= len(s.data) {
		return 0, io.EOF
	}
	l := copy(p, []byte(s.data[s.count:]))
	s.count += l
	return l, nil
}

func (s *stringReader) ReadRune() (rune, int, error) {
	if s.count >= len(s.data) {
		return 0, 0, io.EOF
	}
	r := s.data[s.count]
	s.count++
	return rune(r), 1, nil
}

func (s *stringReader) UnreadRune() error {
	if s.count <= 0 {
		return io.EOF
	}
	s.count--
	return nil
}
