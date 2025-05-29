def complement_of(nuc : Char)
  case nuc
  when 'A'
    return 'T'
  when 'C'
    return 'G'
  when 'G'
    return 'C'
  when 'T'
    return 'A'
  else
    return 'N'
  end
end

def reverse_complement_of(sequence : String)
  rc = String.build do |rc|
    sequence.each_char do |base|
      rc << complement_of(base)
    end
  end
  return rc.reverse
end

if !STDIN.info.type.pipe?
    puts "not a pipe"
    exit 1
end

seq = STDIN.gets
puts reverse_complement_of seq if seq
