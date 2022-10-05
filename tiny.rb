gem "tinify"

require "tinify"
Tinify.key = "3Jkz9YskWc7gb4qgj6shsh86MHMTKPTK"

for arg in ARGV
   source = Tinify.from_file(arg)
   source.to_file("./output/"+arg.chop.chop.chop.chop+"_optimized.png")
end

compressions_this_month = Tinify.compression_count
File.write('./count.txt', compressions_this_month)
