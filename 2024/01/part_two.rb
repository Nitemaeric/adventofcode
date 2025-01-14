def array_with_cursor(array, index)
  array.map.with_index { |item, i| i == index ? "{#{item}}" : item }.join(" ")
end

rows = File.read(__dir__ + "/example.txt").split("\n")

left_list = rows.map { |row| row.split(/\s+/)[0].to_i }.sort
right_list = rows.map { |row| row.split(/\s+/)[1].to_i }.sort

cummulative_similarity = 0
similarity_map = {}
right_enumerator = right_list.to_enum.with_index

left_list.each_with_index do |left_value, left_index|
  puts "L: #{array_with_cursor(left_list, left_index)}"

  similarity_map[left_value] ||= begin
    current_multiplier = 0

    loop do
      break if right_enumerator.peek[0] > left_value

      right_value, right_index = right_enumerator.next

      current_multiplier += 1 if right_value == left_value

      puts "R: #{array_with_cursor(right_list, right_index)}"
    end

    current_multiplier
  end

  puts "occurrences = #{similarity_map[left_value]}"
  puts "\n"

  cummulative_similarity += left_value * similarity_map[left_value]
end

puts "#{left_list.map { |item| "(#{item} x #{similarity_map[item]}) #{item * similarity_map[item]}" }.join(" + ")} = #{cummulative_similarity}"
