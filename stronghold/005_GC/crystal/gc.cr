class FastaSequence
  getter name : String
  getter sequence : String
  getter gc_content : Float64

  def initialize(@name : String, @sequence : String)
    gc_count = @sequence.count do |base|
      base == 'C' || base == 'G'
    end
    @gc_content = gc_count / @sequence.size
  end

  def to_s
    ">" + @name + "\n" + @sequence
  end
end

class FastaParser
  def parse(contents : String) 
    fa_raw_entries = contents.split '>'
    fa_entries : Array(FastaSequence | Nil) = fa_raw_entries.map do |entry|
      entry_contents = entry.split '\n'
      entry_name = entry_contents.first
      entry_sequence = entry_contents[1..-1].join("")

      next if entry_name.empty? || entry_sequence.empty?
      FastaSequence.new name:entry_name, sequence:entry_sequence
    end
    fa_entries.compact
  end
end

if !STDIN.info.type.pipe?
    puts "not a pipe"
    exit 1
end

fa_contents = STDIN.gets_to_end
fa_parser = FastaParser.new
fa_entries = fa_parser.parse fa_contents

max_gc = fa_entries.max_by{|entry| entry.gc_content}
puts max_gc.name + "\n" + max_gc.gc_content.to_s
