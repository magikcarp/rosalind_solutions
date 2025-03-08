---increments count for base base in table tbl
---@param tbl table
---@param base string
---@return nil
local function increase_base_count (tbl, base)
    if base == "A" then tbl.a = tbl.a + 1
    elseif base == "C" then tbl.c = tbl.c + 1
    elseif base == "G" then tbl.g = tbl.g + 1
    elseif base == "T" then tbl.t = tbl.t + 1
    end
    return nil
end

---produces a table with counts of standard nucleotides (ACGT) in a sequence
---@param sequence string
---@return table
function CountNucleotides(sequence)
    local out = {a = 0, c = 0, g = 0, t = 0}
    for i = 1, #sequence do
        local base = sequence:sub(i, i)
        increase_base_count(out, base)
    end
    return out
end

-- file execution
local dna = io.read("*a"):gsub("%s+", "")
local bc = CountNucleotides(dna) -- short for base counts
print(string.format("%d %d %d %d", bc.a, bc.c, bc.g, bc.t))
