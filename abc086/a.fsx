// https://atcoder.jp/contests/abc086/tasks/abc086_a
module Main

stdin.ReadLine().Split(' ')
|> Array.map int
|> fun integers ->
    if (integers.[0] % 2) = 0 || (integers.[1] % 2) = 0 then "Even"
    else "Odd"
|> stdout.WriteLine
