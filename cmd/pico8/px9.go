package pico8

import (
	"bytes"
	"image"
)

func Px9Compress(img image.PalettedImage) []byte {

	cmp := compressor{
		bit: 1,
		w:   img.Bounds().Dx(),
		h:   img.Bounds().Dy(),
	}

	el := []int{}
	found := make(map[int]bool)
	highest := 0

	for y := 0; y < cmp.h; y += 1 {
		for x := 0; x < cmp.w; x += 1 {
			c := int(img.ColorIndexAt(x, y))
			if !found[c] {
				found[c] = true
				el = append(el, c)
				highest = max(highest, c)
			}
		}
	}

	bits := 1
	for ; highest >= 1<<bits; bits += 1 {
	}

	cmp.putnum(cmp.w - 1)
	cmp.putnum(cmp.h - 1)
	cmp.putnum(bits - 1)
	cmp.putnum(len(el) - 1)

	for i := 0; i < len(el); i += 1 {
		cmp.putval(el[i], bits)
	}

	pr := make(map[int][]int)
	dat := []int{}

	for y := 0; y < cmp.h; y += 1 {
		for x := 0; x < cmp.w; x += 1 {
			v := int(img.ColorIndexAt(x, y))
			a := 0
			if y > 0 {
				a += int(img.ColorIndexAt(x, y-1))
			}

			l, found := pr[a]
			if !found {
				l = make([]int, len(el))
				copy(l, el)
				pr[a] = l
			}

			dat = append(dat, vlist_val(l, v))
			vlist_val(el, v)
		}
	}

	nopredict := false

	for pos := 0; pos < len(dat); {
		pos0 := pos

		if nopredict {
			for ; pos < len(dat) && dat[pos] != 0; pos += 1 {
			}
		} else {
			for ; pos < len(dat) && dat[pos] == 0; pos += 1 {
			}
		}

		splen := pos - pos0
		cmp.putnum(splen - 1)

		if nopredict {
			for ; pos0 < pos; pos0 += 1 {
				cmp.putnum(dat[pos0] - 2)
			}
		}

		nopredict = !nopredict
	}

	if cmp.bit != 1 {
		cmp.dest.WriteByte(byte(cmp.byt))
	}

	// Pad with nuls
	cmp.dest.WriteByte(0)
	cmp.dest.WriteByte(0)

	return cmp.dest.Bytes()
}

func vlist_val(l []int, val int) int {
	var v, i int
	for v, i = l[0], 0; v != val; i += 1 {
		v, l[i] = l[i], v
	}
	l[0] = val
	return i
}

/*
func print_vlist(l []int, x int) {
	fmt.Printf("%d: ", len(l))
	for i, e := range l {
		if i > 0 {
			fmt.Printf(",")
		}
		fmt.Printf("%d", e)
	}
	fmt.Printf(" -- %d\n", x)
}
*/
//16: 0,7,5,13,11,8,1,12,3,9,4,6,14,10,15,2
//16: 1,0,7,5,13,11,8,12,3,9,4,6,14,10,15,2
//16: 2,1,0,7,5,13,11,8,12,3,9,4,6,14,10,15

type compressor struct {
	dest bytes.Buffer
	w    int
	h    int
	byt  int
	bit  int
}

func (cmp *compressor) putbit(bval bool) {
	if bval {
		cmp.byt += cmp.bit
	}
	cmp.bit = cmp.bit << 1
	if cmp.bit == 256 {
		cmp.dest.WriteByte(byte(cmp.byt))
		cmp.bit = 1
		cmp.byt = 0
	}
}

func (cmp *compressor) putval(val int, bits int) {
	for i := 0; i < bits; i += 1 {
		cmp.putbit((val & (1 << i)) > 0)
	}
}

func (cmp *compressor) putnum(val int) {
	bits := 0
	done := false
	mx := 0
	vv := 0
	for !done {
		bits += 1
		mx = (1 << bits) - 1
		vv = min(val, mx)
		cmp.putval(vv, bits)
		val -= vv
		done = vv < mx
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
