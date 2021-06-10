# PICO Reality

## Build cartgen
    cd ./cwd
    go build
    cd ..

## Build cart
    ./cmd/cartgen blueprint.json > ~/Library/Application\ Support/pico-8/carts/picoreality.p8

## Lens color palette
Color | Hex     | RGB           | Name          | 565    | Crushed RGB
----- | ------- | ------------- | ------------- | ------ | -----------
0	  | #000000 | 0, 0, 0	    | black         | 0x0000 | #000000
1	  | #1D2B53 | 29, 43, 83    | dark-blue     | 0x194a | #182850
130	  | #422136 | 66,33,54      | darker-purple | 0x4106 | #402030
131	  | #125359 | 18,83,89      | blue-green    | 0x128b | #105058
4	  | #AB5236 | 171, 82, 54   | brown         | 0xaa86 | #a85030
5	  | #5F574F | 95, 87, 79    | dark-grey     | 0x5aa9 | #585448
134	  | #A28879 | 162,136,121   | medium-grey   | 0xa44f | #a08878
128	  | #291814 | 41,24,20      | darkest-grey  | 0x28c2 | #281810
129	  | #111D35 | 17,29,53      | darker-blue   | 0x10e6 | #101c30
132	  | #742F29 | 116,47,41     | dark-brown    | 0x7165 | #702c28
133	  | #49333B | 73,51,59      | darker-grey   | 0x4987 | #483038
141	  | #754665 | 117,70,101    | mauve         | 0x722c | #704460
140	  | #065AB5 | 6,90,181      | true-blue     | 0x02d6 | #0058b0
13	  | #83769C | 131, 118, 156 | lavender      | 0x83b3 | #807498
143	  | #FF9D81 | 255,157,129   | peach         | 0xfcf0 | #f89c80
15	  | #FFCCAA | 255, 204, 170 | light-peach   | 0xfe75 | #f8cca8

## Bresenham's Circle

	void plotCircle(int xm, int ym, int r) {
	   int x = -r, y = 0, err = 2-2*r; /* II. Quadrant */ 
	   do {
	      setPixel(xm-x, ym+y); /*   I. Quadrant */
	      setPixel(xm-y, ym-x); /*  II. Quadrant */
	      setPixel(xm+x, ym-y); /* III. Quadrant */
	      setPixel(xm+y, ym+x); /*  IV. Quadrant */
	      r = err;
	      if (r <= y) err += ++y*2+1;
	      if (r > x || err > y) err += ++x*2+1; 
	   } while (x < 0);
	}

## Notes
* Lens is 48x48
* Lens distortion LUT contains 2304 entries
	* x,y offset
* No lens LUT, use dithered flicker

## Dithered flicker
    data={}
    data.lens_pal={0,1,130,131,4,5,134,128,129,132,133,141,140,13,143,15}
    for i=1,16 do
      pal(i-1,data.lens_pal[i],1)
    end
	
    p=0
	cls()
  
	function _update60()
	  p=1-p
	  fillp(23130)
	  if(p>0)fillp(-23131)
	  for i=0,15 do
	    for j=0,15 do
	      rectfill(i*8,j*8,i*8+7,j*8+7,i+j*16)
	    end
	  end
	end

## References
* [http://members.chello.at/~easyfilter/bresenham.html]()
* [https://www.lexaloffle.com/bbs/?pid=f2ba]()
* [http://marcodiiga.github.io/radial-lens-undistortion-filtering]()
