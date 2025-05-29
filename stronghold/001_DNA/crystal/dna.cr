base_counts = Hash{'A' => 0, 'C' => 0, 'G' => 0, 'T' => 0,}

if !STDIN.info.type.pipe?
    puts "not a pipe"
    exit 1
end

seq = STDIN.gets
if seq
    seq.each_char do |base|
        base_counts[base] += 1
    end
end

puts "#{base_counts['A']} #{base_counts['C']} #{base_counts['G']} #{base_counts['T']}"
