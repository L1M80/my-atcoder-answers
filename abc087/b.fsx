// https://atcoder.jp/contests/abc087/tasks/abc087_b
module Main

let a = stdin.ReadLine() |> int
let b = stdin.ReadLine() |> int
let c = stdin.ReadLine() |> int
let x = stdin.ReadLine() |> int

seq {
    for i = 0 to a do
        for j = 0 to b do
            for k = 0 to c do
                yield i * 500 + j * 100 + k * 50
}
|> Seq.sumBy (fun v ->
       if v = x then 1
       else 0)
|> printfn "%d"
