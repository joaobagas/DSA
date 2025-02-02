package dstr

// https://blog.stackademic.com/crafting-a-custom-hashmap-a-golang-odyssey-00fdfde9a775

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
