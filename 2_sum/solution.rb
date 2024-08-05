# frozen_string_literal: true

# 2-SUM

# Наивный алгоритм
def simple_algorithm(numbers, x)
  (0...numbers.size).each do |i|
    (0...numbers.size).each do |j|
      next if i == j

      return [i, j] if numbers[i] + numbers[j] == x
    end
  end

  return [nil, nil]
end

# Метод двух указателей
def tpm_optimized_algorithm(numbers, x)
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

# Доп. память (хеш-таблица)
def map_optimized_algorithm(numbers, x)
  map = {}

  numbers.each_with_index { |number, i| map[x - number] = i }
  numbers.each_with_index { |number, i| return [i, map[number]] if map[number] }

  return [nil, nil]
end

def test
	list = [1, 2, 3, 4, 5, 6, 7, 9, 10]
	x = 5
	puts simple_algorithm(list, x).join(' ')
	puts tpm_optimized_algorithm(list, x).join(' ')
	puts map_optimized_algorithm(list, x).join(' ')
end
