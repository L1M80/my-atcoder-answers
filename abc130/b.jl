parseInt(x) = parse(Int, x)

a = readline() |> 
split |> 
x->map(y->parse(Int, y), x)

println(a)
