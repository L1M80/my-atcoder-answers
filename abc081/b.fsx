// https://atcoder.jp/contests/abc081/tasks/abc081_b
module Main

let _ = stdin.ReadLine() |> int

let rec isEven =
    fun i ->
        match i with
        | i when i % 2L = 0L -> 1L + isEven (i / 2L)
        | _ -> 0L

stdin.ReadLine().Split(' ')
|> Array.map int64
|> Array.map isEven
|> Array.min
|> stdout.WriteLine
