import           Text.Printf

main :: IO ()

isCenter w h x y | x == w / 2 && y == h / 2 = True
                 | otherwise                = False

main = do
    [w, h, x, y] <- map (read :: String -> Double) . words <$> getLine
    let c = isCenter w h x y
    printf "%.6f\n" $ w * h / 2
    if c then print 1 else print 0
