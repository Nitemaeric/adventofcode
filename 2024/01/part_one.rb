rows = File.read(__dir__ + "/data.txt").split("\n")

left_list = rows.map { |row| row.split(/\s+/)[0].to_i }.sort
right_list = rows.map { |row| row.split(/\s+/)[1].to_i }.sort

cummulative_distance = 0

left_list.each_with_index do |item, index|
  cummulative_distance += (right_list[index] - item).abs
end

p cummulative_distance
