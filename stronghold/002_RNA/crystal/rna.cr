# Transcribes DNA sequence to RNA
def transcribe(dna)
    rna = String.build do |rna|
        dna.each_char do |base|
            if base == 'T' 
                rna << 'U' 
            else 
                rna << base
            end
        end
    end
end

if !STDIN.info.type.pipe?
    puts "not a pipe"
    exit 1
end

seq = STDIN.gets
puts transcribe(seq) if seq
