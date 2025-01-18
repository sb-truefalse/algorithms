# frozen_string_literal: true

# Algorithm: Moving average

# Method: Naive
def moving_avrg(timeseries, k)
  result = []

  ((k - 1)...timeseries.size).each do |i|
    temp = 0
    ((i + 1 - k)..i).each do |j|
      temp += timeseries[j]
    end
    result << (1.0 * temp) / k
  end

  result
end

# Method: 2-pointer
def moving_avrg_with_2_pointers(timeseries, k)
  result = []

  sum = 0
  (0...k).each { |i| sum += timeseries[i] }
  result << (sum * 1.0) / k

  (k...timeseries.size).each do |i|
    sum -= timeseries[i - k]
    sum += timeseries[i]

    result << (1.0 * sum) / k
  end

  result
end
