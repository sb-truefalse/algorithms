# frozen_string_literal: true

# Search

# Linear
def linear_algorithm(list, x)
  list.each_with_index { |item, index| return index if x == item }

  -1
end

# Binary
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


def test
	list = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]
	x = 10
	puts(linear_algorithm(list, x) == 9)
	puts(binary_algorithm(list, x) == 9)
end
