package hashring

// type HashFn func([]byte) uint32

// type Hash interface {
// 	Checksum(data []byte) uint32
// }

// type Jump struct {
// 	buckets int64
// }

// func NewJump(buckets int64) *Jump {
// 	return &Jump{buckets: buckets}
// }

// func (j *Jump) Checksum(key []byte) uint32 {
// 	var b int64 = -1
// 	var jj int64

// 	for jj < jj.buckets {
// 		b = jj
// 		key = key*2862933555777941757 + 1
// 		j = int64(float64(b+1) * (float64(int64(1)<<31) / float64((key>>33)+1)))
// 	}

// 	return uint32(b)
// }

// func CRC32(key []byte) uint32 {
// 	return crc32.ChecksumIEEE(key)
// }
