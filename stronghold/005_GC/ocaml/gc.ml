(* Function to read FASTA from stdin and parse sequences *)
let read_fasta () =
  let rec aux lines current_header current_seq acc =
    match lines with
    | [] -> (match current_header with
        | Some h -> List.rev ((h, String.concat "" (List.rev current_seq)) :: acc)
        | None -> List.rev acc)
    | line :: rest ->
      if String.length line > 0 && line.[0] = '>' then
        (match current_header with
        | Some h -> aux rest (Some line) [] ((h, String.concat "" (List.rev current_seq)) :: acc)
        | None -> aux rest (Some line) [] acc)
      else
        aux rest current_header (line :: current_seq) acc
  in
  let lines = In_channel.input_lines stdin in
  aux lines None [] []
;;

(* calculates GC content of a DNA/RNA sequence *)
let gc_content (s:string): float =
  let seq_len = String.length s in
  let gc_count = String.fold_left (fun acc base ->
    if base == 'G' || base == 'C' then acc + 1 else acc
  ) 0 s in
  if seq_len == 0 then 0.0
  else (float_of_int gc_count) /. (float_of_int seq_len)
;;
    
let () =
  let seqs = read_fasta () in
  match seqs with
  | [] -> print_endline "No sequences in input"
  | _  ->
    let best_seq = List.fold_left (fun best (header, seq) ->
      let gc = gc_content seq in
      match best with
      | None -> Some(header, gc)
      | Some (_, best_gc) -> if gc > best_gc then Some (header, gc) else best
    ) None seqs in
    (match best_seq with
     | Some (header, gc) -> Printf.printf "%s\n%.6f\n" header (gc *. 100.)
     | None -> print_endline "No valid sequences found.")
 ;;
