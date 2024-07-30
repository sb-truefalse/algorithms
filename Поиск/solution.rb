# frozen_string_literal: true

# Поиск

# Линейный поиск
def linear_algorithm(list, x)
  list.each_with_index { |item, index| return index if x == item }

  -1
end

# Бинарный поиск
def binary_algorithm(list, x)
  l, r = 0, (list.size - 1)
  list.sort!

  while l <= r
    mid = (r + l) / 2

    if list[mid] == x
      return mid
    elsif list[mid] > x
      r = mid - 1
    else
      l = mid + 1
    end
  end

  -1
end
