module Main where

translate :: String -> String
translate = map (\n -> if n == 'T' then 'U' else n)

main :: IO ()
main = do
    seq <- getLine
    putStrLn $ translate seq
