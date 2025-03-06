(* Translates DNA to RNA *)
let translate = fun dna -> 
  String.map (fun n -> if n = 'T' then 'U' else n) dna ;;

let () = 
  let dna = read_line () in 
  let rna = translate dna in
  print_endline rna ;;
