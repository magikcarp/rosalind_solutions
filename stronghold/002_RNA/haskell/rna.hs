module Main where

transcribe :: String -> String
transcribe = map (\n -> if n == 'T' then 'U' else n)

main :: IO ()
main = do
    seq <- getLine
    putStrLn $ transcribe seq
