package rpc

import (
	"bytes"
	"io"
)

type Options struct{}
type Server struct {
	Options

	B    string
	Bp   *string
	Bpp  **string
	Bb   bytes.Buffer
	Bi   io.Writer
	Bbp  *bytes.Buffer
	Bip  *io.Writer
	Bbpp **bytes.Buffer
	Bipp **io.Writer

	A    []string
	PA   *[]string
	A2   [][]string
	A3   [][]*[]string
	PA2  *[]*[]string
	Ab   []bytes.Buffer
	Abp2 [][]*bytes.Buffer

	M   map[string]bool
	MA  map[int][]string
	AM  []map[int][]string
	AMP []*map[int][]string

	F  func()
	FM map[int]func(int, int) (string, int)
	FP *func(a, b int) (c string, d int)

	FR  func() int
	FRN func() (id int)
}
