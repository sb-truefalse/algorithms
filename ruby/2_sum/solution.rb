# frozen_string_literal: true

# Algorithm: Two Sum

# Method: Naive
def two_sum(numbers, x)
  (0...numbers.size).each do |i|
    (0...numbers.size).each do |j|
      next if i == j

      return [i, j] if numbers[i] + numbers[j] == x
    end
  end

  return [nil, nil]
end

# Method: 2-pointer
def two_sum_with_2_pointers(numbers, x)
  numbers.sort!

  i = 0
  j = numbers.size  - 1

  while i < j
    sum = numbers[i] + numbers[j]

    if sum == x
      return [i, j]
    elsif sum > x
      j -= 1
    else
      i += 1
    end
  end

  return [nil, nil]
end

# Method: Hash table
def two_sum_with_hash(numbers, x)
  map = {}

  numbers.each_with_index { |number, i| map[x - number] = i }
  numbers.each_with_index { |number, i| return [i, map[number]] if map[number] }

  return [nil, nil]
end
