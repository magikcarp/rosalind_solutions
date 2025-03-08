---Parses a FASTA file and returns a table with sequence information
---@param fasta_content string
---@return table|nil: table of fasta sequences and associated data. 
local function parse_fasta(fasta_content)
    local fasta_data = {}
    local curr_head = nil
    local curr_seq = ""

    if fasta_content == "" then return nil end

    for line in fasta_content:gmatch("([^\n]*)\n?") do
        -- Trim leading and trailing whitespace
        line = line:match("^%s*(.-)%s*$")

        if line:sub(1, 1) == ">" then
            if curr_head then
                fasta_data[curr_head] = curr_seq
            end
            -- start a new sequence
            curr_head = line:sub(2)
            curr_seq = ""
        else
            curr_seq = curr_seq .. line
        end
    end
    -- add the last sequence
    if curr_head then
        fasta_data[curr_head] = curr_seq
    end

    return fasta_data
end

local function gc_content(sequence)
    local seq = sequence:upper()
    local _, c = seq:gsub("C", "")
    local _, g = seq:gsub("G", "")
    return (g + c) / #seq
end

local function greatest_gc(seqs)
    local highest_gc = -1
    local highest_gc_id = nil

    for id, seq in pairs(seqs) do
        local gc = gc_content(seq)
        if gc > highest_gc then 
            highest_gc = gc 
            highest_gc_id = id
        end
    end

    return highest_gc_id, highest_gc
end

-- file execution
local fasta_data = io.read("*a")
local fasta_seqs = parse_fasta(fasta_data)
if fasta_seqs ~= nil then
    local gc_id, gc = greatest_gc(fasta_seqs)
    print(gc_id)
    print(gc*100)
else
    print("No FASTA sequences provided.")
end
