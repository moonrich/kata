plusTwo = Enum.map [1,2,3,4], &(&1 + 2)
plusTwo
|> IO.inspect


eInspect = Enum.each [1,2,3,4], &(IO.inspect(&1))
eInspect
