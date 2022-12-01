file = File.read("./data.txt")

reindeer_inventories = file.split("\n\n")

reindeer_totals = reindeer_inventories.map do |items_string|
  items_string.split("\n").map(&:to_i).sum
end

puts "Part 1: #{reindeer_totals.max}"
puts "Part 2: #{reindeer_totals.sort.reverse.first(3).sum}"
