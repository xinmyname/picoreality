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
    { t=175, init=lens_init, draw=lens_draw, update=lens_update}
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
      c=peek(24336+i)
      if (c>=128) c-=112
      p=fade[c]
      if (p>=16) p+=112
      pal(i,p,1)
    end
    if fading==7*5+1 then
      cls()
      pal()
      fading=-1
    end
  end
end

function lens_init()
  for i=1,16 do
    pal(i-1,data.monster_pal[i],1)
  end
end

function lens_update()
end

function lens_draw()
 cls()
	for y=0,3 do
		for x=0,3 do
			c=y*4+x
      rectfill(x*16,y*16,x*16+15,y*16+15, c)
		end
	end
end
