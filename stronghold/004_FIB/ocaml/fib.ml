let rabbit_reproduction (n: int) (k: int): int = 
  let real_n = n - 1 in 
  let rec aux i j n k = 
    match n with 
    | 0 | 1 -> i
    | _     -> aux (i+(j*k)) i (n-1) k
  in aux 1 1 real_n k
;;

let int_from_input (s: string): int list =
  s
  |> String.split_on_char ' '
  |> List.filter (fun x -> x <> "")
  |> List.map int_of_string

let sum (ls: int list): int =
  let rec aux acc ls = 
    match ls with 
    | [] -> acc
    | x :: tail -> aux (x + acc) tail 
  in aux 0 ls
;;

let () = 
  let input = read_line () in 
  let nums = int_from_input input in 
  let result = rabbit_reproduction (List.nth nums 0) (List.nth nums 1) in
  Printf.printf "%d\n" result 
;;
