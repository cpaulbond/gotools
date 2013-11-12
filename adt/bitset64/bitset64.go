package bitset64

type Bitset64 struct {
	bits uint64
}

func (s *Bitset64) Init() {
	s.bits = 0
}

func (s *Bitset64) Get() uint64 {
	return s.bits
}

func (s *Bitset64) Set(bit int) {
	s.bits = s.bits ^ (1 << uint(bit))
}

func (s *Bitset64) Clear(bit int) {
	s.bits = s.bits &^ (1 << uint(bit))
}

func (s *Bitset64) IsSet(bit int) bool {
	return s.bits&(1<<uint(bit)) != 0
}
