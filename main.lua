function _init()
	-- switch upper rom to ram, render sprite 0 opaque
	poke(0x5f36,0x18)
	palt(0,false)
	update=nil
	draw=nil
	t=0
	fading=0
	script_index=1
	script={
		{ t=0, draw=title_draw },
		{ t=120, draw=fadeout_draw },
		{ t=175, init=lens_init, draw=lens_draw, update=lens_update},
		{ t=475, init=monster_init, draw=monster_draw, update=monster_update},
		{ t=575, draw=whiteout_draw },
		{ t=630, init=spin_init, draw=spin_draw, update=spin_update}
	}
end
	 
function _update60()
	if (script_index>#script) return
	if t==script[script_index].t then
		draw=script[script_index].draw
		update=script[script_index].update
		if (script[script_index].init) script[script_index].init()
		script_index+=1
	end
	if (update) update()
	t+=1
end
	 
function _draw()
	draw()
end

function title_draw()
	cls()
	map(0,0,0,0,16,16)
	for i=1,16 do
		pal(i-1,data.title_pal[i],1)
	end
end

function fadeout_draw()
	local fade,c,p={[0]=0,17,18,19,20,16,22,6,24,25,9,27,28,29,29,31,0,0,16,17,16,16,5,0,2,4,0,3,1,18,2,4}
	fading+=1
	if fading%5==1 then
		for i=0,15 do
			c=peek(0x5f10+i)
			if (c>=128) c-=112
			p=fade[c]
			if (p>=16) p+=112
			pal(i,p,1)
		end
		if fading==7*5+1 then
			fading=-1
		end
	end
end

function lens_init()
	for i=1,16 do
		pal(i-1,data.monster_pal[i],1)
	end
	px9_decomp(0,0,data.monster_image,sget,sset)
	lens_pat=0b1010010110100101.1	
 xm=0
 ym=0
	lx=24
	ly=1
	lr=21
	vx=1
	vy=2
	l=0
	r=128-(lr<<1)
	b=128-(lr<<1) 
	bounces=0		
end

function lens_update()
	vy+=0.1
	
	lx+=vx
	ly+=vy

	if ly>=b then
		ly=b
		vx=vx*0.95
		vy=vy*-0.85
		bounces+=1
		if (bounces>3) b=32767
	end	

	if lx>r or lx<l then
		vx=vx*-0.85
	end
end

function lens_draw()
	map(0,0,0,0,16,16)
	for y=1,42 do
		xlut=data.xoffsets[y]
		ylut=data.yoffsets[y]
		for x=1,42 do
			pset(x+lx,y+ly,sget(x+lx+xlut[x],y+ly+ylut[x]))
		end
	end

	lens_pat^^=0xffff
	fillp(lens_pat)
	circfill(lx+lr,ly+lr,lr,12)
end

function monster_init()
	cls()
	audio={}
	audio.base=0x4300
	audio.len=#data.atomic_audio
	audio.pos=0
	audio.buffered=0
	audio.playing=true
end

function monster_update()

	if (not audio.playing) return

	audio.buffered = stat(108)

	if (audio.buffered<512) then
		local size = audio.pos - audio.len
		
		if (size < -256) then
			size = 256
		else
			size *= -1
			audio.playing=false
		end		
		
		for i=1,size do
			poke(audio.base+(i-1), ord(sub(data.atomic_audio,i+audio.pos,i+audio.pos)))
		end
		
		serial(0x808,audio.base,size)
		
		if (audio.playing) then
			audio.pos += size
		else
			audio.pos = 0
		end
	end


end

function monster_draw()
	map(0,0,0,0,16,16)
end

function whiteout_draw()
		local fade,c,p={[0]=17,28,24,27,25,22,7,7,6,10,23,26,6,6,15,7,2,1,2,3,4,5,6,7,8,9,23,26,12,13,15,15}
		fading+=1
		if fading%3==1 then
			for i=0,15 do
				c=peek(0x5f10+i)
				if (c>=128) c-=112
				p=fade[c]
				if (p>=16) p+=112
				pal(i,p,1)
			end
			if fading==7*5+1 then
				fading=-1
			end
		end
	end
	
function spin_init()
 cls()
	for i=1,16 do
		pal(i-1,data.monster_pal[i],1)
	end
end

function spin_update()
end

function spin_draw()
	map(0,0,0,0,16,16)
end
