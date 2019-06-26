module Main

stdin.ReadLine()
|> Seq.map int
|> Seq.sumBy (fun x -> x - int '0')
|> stdout.WriteLine
