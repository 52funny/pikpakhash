package pikpakhash

import (
	"crypto/sha1"
	"fmt"
	"io"
	"math"
	"os"
)

const (
	DefaultCalcRoutines = 1 << 4  // 16
	DefaultBufferSize   = 1 << 18 // 256KB
)

// segmentOffset is an offset into a segment of a file.
type segmentOffset struct {
	id     int64
	offset int64
}

// segmentHash is a hash of a segment of a file.
type segmentHash struct {
	id   int64
	hash []byte
}

// PikPakHash contains routine memeber and buf arrary,
// routine value represents how many go routines will compute the hash.
// buf value represents
type PikPakHash struct {
	routine int
	buffer  int
}

// Default return instance.
// routine and buffer size are default value.
func Default() PikPakHash {
	return PikPakHash{
		routine: DefaultCalcRoutines,
		buffer:  DefaultBufferSize,
	}
}

// NewPikPakHashWithBuf return instance,
// it contains routine and buffer size
func NewPikPakHash(routine int, buffer int) PikPakHash {
	return PikPakHash{
		routine,
		buffer,
	}
}

// HashFromPath returns the hash of the file at the given path
func (ph *PikPakHash) HashFromPath(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()
	return ph.HashFromFile(file)
}

// HashFromFile return the hash of the file by given the file interface
func (ph *PikPakHash) HashFromFile(file *os.File) (string, error) {
	resHash := sha1.New()

	state, _ := file.Stat()
	chunk := getChunkSize(state.Size())

	// segment length
	segmentLen := int64(math.Ceil(float64(state.Size()) / float64(chunk)))

	inCh := make(chan segmentOffset, segmentLen)
	outCh := make(chan segmentHash, ph.routine)

	// start go routine
	for i := 0; i < ph.routine; i++ {
		go ph.calculate(file, chunk, inCh, outCh)
	}

	// send info data to go routine
	for i := int64(0); i < segmentLen; i++ {
		inCh <- segmentOffset{id: i, offset: chunk * i}
	}
	close(inCh)

	// heap structure
	h := newHeap()

	for i := int64(0); i < segmentLen; i++ {
		h.Push(<-outCh)
	}

	// close the channel
	close(outCh)

	// fmt.Printf("%d %x\n", result[1].id, result[1].data)
	// fmt.Printf("%d %x\n", result[2].id, result[2].data)
	// calculate file sha1
	for h.Len() > 0 {
		resHash.Write(h.Pop().hash)
	}
	checksum := fmt.Sprintf("%x", resHash.Sum(nil))
	return checksum, nil
}

func getChunkSize(size int64) int64 {
	if size > 0 && size < 0x8000000 {
		return 0x40000
	}
	if size >= 0x8000000 && size < 0x10000000 {
		return 0x80000
	}
	if size <= 0x10000000 || size > 0x20000000 {
		return 0x200000
	}
	return 0x100000
}

func (ph *PikPakHash) calculate(file *os.File, chunkSize int64, ch chan segmentOffset, outCh chan segmentHash) {
	buf := make([]byte, ph.buffer)
	for {
		info, ok := <-ch
		// fmt.Println(info, ok)
		if !ok {
			break
		}
		partHash := sha1.New()
		// io.CopyN(partHash, file, chunkSize)
		offset := info.offset
		total := int64(0)
		for {

			// ReadAt not similar with Read
			n, err := file.ReadAt(buf, offset)
			offset += int64(n)
			total += int64(n)

			// write buf to hash to calculate the sha1
			partHash.Write(buf[:n])

			if total >= chunkSize {
				break
			}

			if err != nil && err == io.EOF {
				break
			}
		}
		outCh <- segmentHash{id: info.id, hash: partHash.Sum(nil)}
	}
}
