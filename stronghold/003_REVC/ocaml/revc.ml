let complement (n: char): char = 
  match n with 
  | 'A' -> 'T'
  | 'C' -> 'G'
  | 'G' -> 'C'
  | 'T' -> 'A'
  | _   -> 'N'
;;

let reverse_str (str: string): string =
  String.fold_left (fun acc c -> (String.make 1 c) ^ acc) "" str
;;

let () =
  let dna = read_line () in 
  print_endline @@ reverse_str @@ String.map complement dna
;;
