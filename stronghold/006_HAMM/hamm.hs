module Main where

main :: IO()
main = do
    seq1 <- getLine
    seq2 <- getLine

    print (sum $ detectDiff seq1 seq2)

detectDiff :: String -> String -> [Int]
detectDiff [] [] = []
detectDiff (x:xs) (y:ys)
    | x == y = 0 : detectDiff xs ys
    | x /= y = 1 : detectDiff xs ys
