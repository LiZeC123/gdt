package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func init() {

}

type Header struct {
	Start byte
	PNG   [3]byte
	CRLF  [2]byte
	EOF   byte
	LF    byte
}

type Chunk struct {
	Length int32
	Type   [4]byte
	Data   []byte
	CRC    int32
}

type PNGFile struct {
	Header
	Chunks []*Chunk
}

var IEND = [4]byte{'I', 'E', 'N', 'D'}
var IMSG = [4]byte{'r', 'M', 'S', 't'}

func LoadPNG(filename string) *PNGFile {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	img := &PNGFile{}
	_ = binary.Read(file, binary.BigEndian, &img.Header)

	chunkType := [4]byte{}
	for chunkType != IEND {
		chunk := ReadChunk(file)
		img.Chunks = append(img.Chunks, chunk)
		chunkType = chunk.Type
	}

	err = file.Close()
	if err != nil {
		panic(err)
	}

	return img
}

func ReadChunk(r io.Reader) *Chunk {
	chunk := &Chunk{}
	_ = binary.Read(r, binary.BigEndian, &chunk.Length)
	chunk.Data = make([]byte, chunk.Length)

	_ = binary.Read(r, binary.BigEndian, &chunk.Type)
	_ = binary.Read(r, binary.BigEndian, &chunk.Data)
	_ = binary.Read(r, binary.BigEndian, &chunk.CRC)

	return chunk
}

func (img *PNGFile) Embed() {
	msg, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	L := len(img.Chunks) - 1
	EOFChunk := img.Chunks[L]
	img.Chunks = img.Chunks[:L]
	chunk := Chunk{Length: int32(len(msg)), Type: IMSG, Data: msg, CRC: 0}
	img.Chunks = append(img.Chunks, &chunk, EOFChunk)
}

func (img *PNGFile) Extract() {
	for _, chunk := range img.Chunks {
		chunkType := string(chunk.Type[:])
		if chunkType == "rMSt" {
			_, err := os.Stdout.Write(chunk.Data)
			if err != nil {
				panic(err)
			}
		}
	}
}

func (img *PNGFile) Save(outFile string, ignore [4]byte) {
	file, err := os.Create(outFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	err = binary.Write(file, binary.BigEndian, img.Header)
	if err != nil {
		panic(err)
	}

	for _, chunk := range img.Chunks {
		if chunk.Type == ignore {
			continue
		}

		_ = binary.Write(file, binary.BigEndian, chunk.Length)
		_ = binary.Write(file, binary.BigEndian, chunk.Type)
		_ = binary.Write(file, binary.BigEndian, chunk.Data)
		_ = binary.Write(file, binary.BigEndian, chunk.CRC)
	}
}

func (img *PNGFile) Show() {
	h := img.Header
	fmt.Printf("Header: %x %x %x %x %x\n", h.Start, h.PNG, h.CRLF, h.EOF, h.LF)

	for _, chunk := range img.Chunks {
		fmt.Printf("Type: %v Length %d\n", string(chunk.Type[:]), chunk.Length)
	}
}

func Embed(pngFile string) {
	img := LoadPNG(pngFile)
	img.Embed()
	img.Save(pngFile, [4]byte{})
}

func Extract(pngFile string) {
	img := LoadPNG(pngFile)
	img.Extract()
}

func Show(pngFile string) {
	img := LoadPNG(pngFile)
	img.Show()
}

func Clear(pngFile string) {
	img := LoadPNG(pngFile)
	img.Save(pngFile, IMSG)
}
