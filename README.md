# PICO Reality

## TODO
* [X] Palette flash to white
* [ ] Convert lens offset LUTs to data strings, unpack in working memory
* [ ] tline for image rotation
* [ ] Better lens center, should zoom slightly
* [ ] Can lens be bigger with newer LUTs?
* [ ] Slow down lens bouncing - total bouncing time should be ~10 seconds
* [ ] Compress orch hit
* [ ] Variable rate audio playback
* [ ] Why doesn't memory based pixel read work properly? (doesn't matter - too slow)

## Build cartgen
    cd ./cwd
    go build
    cd ..

## Build cart
    ./cmd/cartgen blueprint.json > ~/Library/Application\ Support/pico-8/carts/picoreality.p8

## Lens color palette
Color | Hex     | RGB           | Name          
----- | ------- | ------------- | ------------- 
0	  | #000000 | 0, 0, 0	    | black         
1	  | #1D2B53 | 29, 43, 83    | dark-blue     
130	  | #422136 | 66,33,54      | darker-purple 
131	  | #125359 | 18,83,89      | blue-green    
4	  | #AB5236 | 171, 82, 54   | brown         
5	  | #5F574F | 95, 87, 79    | dark-grey     
134	  | #A28879 | 162,136,121   | medium-grey   
128	  | #291814 | 41,24,20      | darkest-grey  
129	  | #111D35 | 17,29,53      | darker-blue   
132	  | #742F29 | 116,47,41     | dark-brown    
133	  | #49333B | 73,51,59      | darker-grey   
141	  | #754665 | 117,70,101    | mauve         
140	  | #065AB5 | 6,90,181      | true-blue     
13	  | #83769C | 131, 118, 156 | lavender      
143	  | #FF9D81 | 255,157,129   | peach         
15	  | #FFCCAA | 255, 204, 170 | light-peach   

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

## Manual palette fade-out/white-out

For each I below, locate in palette index that is darker and enter it as DI

	I:   0  1  2  3  4  5  6  7  8  9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29 30 31
	DI:  0 17 18 19 20 16 22  6 24 25  9 27 28 29 29 31  0  0 16 17 16 16  5  0  2  4  0  3  1 18  2  4

For each I below, locate in palette index that is lighter and enter it as DI

	I:   0  1  2  3  4  5  6  7  8  9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29 30 31
	LI: 17 28 24 27 25 22  7  7  6 10 23 26  6  6 15  7  2  1  2  3  4  5  6  7  8  9 23 26 12 13 15 15  

## References
* [http://members.chello.at/~easyfilter/bresenham.html]()
* [https://www.lexaloffle.com/bbs/?pid=f2ba]()
* [http://marcodiiga.github.io/radial-lens-undistortion-filtering]()

## The Land of Dead Ideas That Didn't Work But That I Don't Want to Forget

### Read sprite pixel
peek((sy<<6)+(sx>>1))>>((sx&1)<<2)&0xf

### Bresenham's Circle

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

### Lens color palette
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
