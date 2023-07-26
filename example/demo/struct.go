package rpc

import (
	"bytes"
	"errors"
	"io"
	"net"
	"net/http"
	"sync"
)

type SS io.ByteReader
type SP *io.ByteReader
type SF func()
type SM map[string]int
type State uint

const (
	S0 State = 1
	S1
	S3 = iota + 10*10
	S4
	S5, s6 = 7, 9
)

func abc() (int, int) {
	return 1, 2
}

const (
	X0 = 0
	X1
	X3 = -1 + iota + 10*10
	X4
)
const SX0, SX1 = uint32(100), int32(99)
const S64 uint64 = 100
const S = "ko"
const NotFound = http.StatusNotFound

var (
	V0     uint32 = 0
	V1            = 1
	V2, V3        = 2, 3 + 1
	V4            = http.StateNew
	Err           = errors.New("abc")
	X5, X6        = abc()
	X7            = &X6
)

var ErrK0 = errors.New(`ko`)
var P0 = &sync.Pool{}
var P1 = sync.Pool{New: func() any {
	b := make([]byte, 8192)
	return &b
}}

type Callback func(data string) bool
type KO interface {
}
type Options interface {
	KO
	io.WriteCloser
	Addr() net.Addr
}
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

func (s *Server) Serve() (e error) {
	s.abc()
	return
}
func (s Server) abc() (e error) {
	return
}

type animal struct{}

func (a *animal) Eat() {}
