package triedbopts

type Option interface {
	triedbopt()
}

func AtomicReferenceAfterUpdate() Option {
	return atomicUpdateAndRef{}
}

type atomicUpdateAndRef struct{ Option }

func ShouldAtomicReferenceAfterUpdate(opts []Option) bool {
	for _, o := range opts {
		if _, ok := o.(atomicUpdateAndRef); ok {
			return true
		}
	}
	return false
}
