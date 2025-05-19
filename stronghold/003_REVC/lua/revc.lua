local function complement(nucleotide)
  if nucleotide == "A" then
    return "T"
  elseif nucleotide == "C" then
    return "G"
  elseif nucleotide == "G" then
    return "C"
  elseif nucleotide == "T" then
    return "A"
  else
    return "N"
  end
end

local function reverse_complement(sequence)
  local reversed_sequence = sequence:gsub("%s*", ""):reverse()
  local complemented_sequence = ""

  for i = 1, reversed_sequence:len() do
    local nuc = reversed_sequence:sub(i, i)
    complemented_sequence = complemented_sequence .. complement(nuc)
  end
  
  return complemented_sequence
end

-- file execution
local sequence = io.read("*a")
local comp_seq = reverse_complement(sequence)
print(comp_seq)
