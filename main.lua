function _init()
  palt(0,false)
  poke(0x5f36,8)
  update=nil
  draw=title_draw
  t=0
  fading=0
end
   
function _update60()
  if (update) update()
  t+=1
  if t>120 then
    draw=fadeout_draw
  end
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
