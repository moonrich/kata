fun1 = fn -> fn -> "Hello" end end

fun1.()
|> IO.inspect

fun1.().()
|> IO.inspect
