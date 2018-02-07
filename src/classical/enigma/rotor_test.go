package enigma_test



import (
	"testing"
	"github.com/stretchr/testify/assert"
	"../../classical/enigma"
)


/* ********* SHIFT ********* */

func TestShift_1(t *testing.T) {
	k := 'a'
	a := "abcdefghijklmnopqrstuvwxyz"
	r := enigma.NewRotor([]rune(a), k)

	assert.Equal(t, 0, r.Offset)
}

func TestShift_2(t *testing.T) {
	k := 'a'
	a := "abcdefghijklmnopqrstuvwxyz"
	r := enigma.NewRotor([]rune(a), k)

	r.Shift()

	assert.Equal(t, 1, r.Offset)
}

func TestShift_3(t *testing.T) {
	k := 'a'
	a := "abcdefghijklmnopqrstuvwxyz"
	r := enigma.NewRotor([]rune(a), k)

	r.Shift()
	r.Shift()

	assert.Equal(t, 2, r.Offset)
}


/* ********* UNSHIFT ********* */

func TestUnShift_2(t *testing.T) {
	k := 'a'
	a := "abcdefghijklmnopqrstuvwxyz"
	r := enigma.NewRotor([]rune(a), k)

	r.UnShift()

	assert.Equal(t, 25, r.Offset)
}


func TestUnShift_3(t *testing.T) {
	k := 'a'
	a := "abcdefghijklmnopqrstuvwxyz"
	r := enigma.NewRotor([]rune(a), k)

	r.UnShift()
	r.UnShift()

	assert.Equal(t, 24, r.Offset)
}


/* ********* ENCODE ********* */

func TestEncodeLetter_a(t *testing.T) {
	ptext := 'a'
	k := 'a'
	a := "abcdefghijklmnopqrstuvwxyz"
	r := enigma.NewRotor([]rune(a), k)

	assert.Equal(t, 'a', r.EncodeLetter(ptext))
}

func TestEncodeLetter_ab(t *testing.T) {
	ptext := 'a'
	k := 'a'
	a := "abcdefghijklmnopqrstuvwxyz"
	r := enigma.NewRotor([]rune(a), k)

	assert.Equal(t, 'a', r.EncodeLetter(ptext))
	assert.Equal(t, 'b', r.EncodeLetter(ptext))
}

func TestEncodeLetter_ab2(t *testing.T) {
	ptext := 'a'
	k := 'b'
	a := "abcdefghijklmnopqrstuvwxyz"
	r := enigma.NewRotor([]rune(a), k)

	assert.Equal(t, 'b', r.EncodeLetter(ptext))
	assert.Equal(t, 'c', r.EncodeLetter(ptext))
}

func TestEncodeLetter_acb2(t *testing.T) {
	ptext := 'a'
	k := 'b'
	a := "abcdefghijklmnopqrstuvwxyz"
	r := enigma.NewRotor([]rune(a), k)

	assert.Equal(t, 'b', r.EncodeLetter(ptext))
	assert.Equal(t, 'c', r.EncodeLetter(ptext))
	assert.Equal(t, 'd', r.EncodeLetter(ptext))
}

func TestEncodeLetter_za(t *testing.T) {
	ptext := 'z'
	k := 'a'
	a := "abcdefghijklmnopqrstuvwxyz"
	r := enigma.NewRotor([]rune(a), k)

	assert.Equal(t, 'z', r.EncodeLetter(ptext))
}

func TestEncodeLetter_zb(t *testing.T) {
	ptext := 'z'
	k := 'b'
	a := "abcdefghijklmnopqrstuvwxyz"
	r := enigma.NewRotor([]rune(a), k)

	assert.Equal(t, 'a', r.EncodeLetter(ptext))
}

func TestEncodeLetter_yc(t *testing.T) {
	ptext := 'y'
	k := 'c'
	a := "abcdefghijklmnopqrstuvwxyz"
	r := enigma.NewRotor([]rune(a), k)

	assert.Equal(t, 'a', r.EncodeLetter(ptext))
}


/* ********* DECODE ********* */

func TestDecodeLetter_a1(t *testing.T) {
	ctext := 'a'
	k := 'b'
	a := "abcdefghijklmnopqrstuvwxyz"
	r := enigma.NewRotor([]rune(a), k)

	assert.Equal(t, 'a', r.DecodeLetter(ctext))
	assert.Equal(t, 'a', r.GetKey())
}

func TestDecodeLetter_b2(t *testing.T) {
	ctext := 'b'
	k := 'c'
	a := "abcdefghijklmnopqrstuvwxyz"
	r := enigma.NewRotor([]rune(a), k)

	assert.Equal(t, 'a', r.DecodeLetter(ctext))
	assert.Equal(t, 'b', r.GetKey())
}

func TestDecodeLetter_a3(t *testing.T) {
	ctext := 'd'
	k := 'd'
	a := "abcdefghijklmnopqrstuvwxyz"
	r := enigma.NewRotor([]rune(a), k)

	assert.Equal(t, 'b', r.DecodeLetter(ctext))
	assert.Equal(t, 'c', r.GetKey())

	assert.Equal(t, 'a', r.DecodeLetter('b'))
	assert.Equal(t, 'b', r.GetKey())
}

func TestDecodeLetter_za(t *testing.T) {
	ctext := 'z'
	k := 'a'
	a := "abcdefghijklmnopqrstuvwxyz"
	r := enigma.NewRotor([]rune(a), k)

	assert.Equal(t, 'a', r.DecodeLetter(ctext))
	assert.Equal(t, 'z', r.GetKey())
}

func TestDecodeLetter_zb(t *testing.T) {
	ctext := 'a'
	k := 'c'
	a := "abcdefghijklmnopqrstuvwxyz"
	r := enigma.NewRotor([]rune(a), k)

	assert.Equal(t, 'z', r.DecodeLetter(ctext))
	assert.Equal(t, 'b', r.GetKey())
}

func TestDecodeLetter_yc(t *testing.T) {
	ctext := 'a'
	k := 'd'
	a := "abcdefghijklmnopqrstuvwxyz"
	r := enigma.NewRotor([]rune(a), k)

	assert.Equal(t, 'y', r.DecodeLetter(ctext))
	assert.Equal(t, 'c', r.GetKey())
}


/* ********* DECODE ********* */

func TestEncodeDecodeLetter_abc(t *testing.T) {
	k := 'a'
	a := "abcdefghijklmnopqrstuvwxyz"
	ptext := 'a'
	r := enigma.NewRotor([]rune(a), k)

	ctext := r.EncodeLetter(ptext)
	assert.Equal(t, ptext, r.DecodeLetter(ctext))
}

func TestEncodeDecodeLetter_aa(t *testing.T) {
	k := 'a'
	a := "abcdefghijklmnopqrstuvwxyz"
	ptext := 'a'
	r := enigma.NewRotor([]rune(a), k)

	ctext := r.EncodeLetter(ptext)
	ctext = r.EncodeLetter(ctext)
	ctext = r.EncodeLetter(ctext)

	tmp := r.DecodeLetter(ctext)
	tmp = r.DecodeLetter(tmp)
	tmp = r.DecodeLetter(tmp)

	assert.Equal(t, ptext, tmp)
}

func TestEncodeDecodeLetter_za(t *testing.T) {
	k := 'a'
	a := "abcdefghijklmnopqrstuvwxyz"
	ptext := 'z'
	r := enigma.NewRotor([]rune(a), k)

	ctext := r.EncodeLetter(ptext)
	ctext = r.EncodeLetter(ctext)
	ctext = r.EncodeLetter(ctext)

	tmp := r.DecodeLetter(ctext)
	tmp = r.DecodeLetter(tmp)
	tmp = r.DecodeLetter(tmp)

	assert.Equal(t, ptext, tmp)
}

func TestEncodeDecodeLetter_yd(t *testing.T) {
	k := 'd'
	a := "abcdefghijklmnopqrstuvwxyz"
	ptext := 'y'
	r := enigma.NewRotor([]rune(a), k)

	ctext := r.EncodeLetter(ptext)
	ctext = r.EncodeLetter(ctext)
	ctext = r.EncodeLetter(ctext)

	tmp := r.DecodeLetter(ctext)
	tmp = r.DecodeLetter(tmp)
	tmp = r.DecodeLetter(tmp)

	assert.Equal(t, ptext, tmp)
}


