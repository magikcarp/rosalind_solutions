(* Count nucleotides in a DNA sequence *)
let count_nucs = fun dna -> 
    let counts = Hashtbl.create 4 in
    (* Initializes the counts for ACGT *)
    List.iter (fun c -> Hashtbl.add counts c 0) ['A'; 'C'; 'G'; 'T'];

    String.iter (fun c -> 
        if Hashtbl.mem counts c then 
            Hashtbl.replace counts c ((Hashtbl.find counts c) + 1)
        ) dna; 

    (
        Hashtbl.find counts 'A', 
        Hashtbl.find counts 'C', 
        Hashtbl.find counts 'G', 
        Hashtbl.find counts 'T'
    ) ;;

let () = 
    let seq = read_line () in
    let (a, c, g, t) = count_nucs seq in 
    Printf.printf "%d %d %d %d\n" a c g t ;;
