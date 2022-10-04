gem "tinify"

require "tinify"
Tinify.key = "MKqGSH9h2zlNrh0vmB9n4Q939syKpMbx"

for arg in ARGV
   source = Tinify.from_file(arg)
   source.to_file("./output/"+arg.chop.chop.chop.chop+"_optimized.png")
end

