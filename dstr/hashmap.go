package dstr

type Pair[T any] struct {
	Key   string
	Value T
}

type HashMap[T any] struct {
	BucketSize int
	FilledSize int
	Bucket     [][]Pair[T]
	LoadFactor int

	HashFunc func(string) uint32
}

func NewHashMap[T any]() HashMap[T] {
	return HashMap[T]{}
}

func Hash[T any]() string {
	return "t"
}
