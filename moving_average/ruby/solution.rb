# frozen_string_literal: true

# Moving average

# Naive
def simple_algorithm(timeseries, k)
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

# 2-pointer
def tpm_optimized_algorithm(timeseries, k)
  result = []

  sum = 0
  (0...k).each { |i| sum += timeseries[i] }
  result << sum / k

  (k...timeseries.size).each do |i|
    sum -= timeseries[i - k]
    sum += timeseries[i]

    result << sum / k
  end

  result
end

def test
  data = [1, 2, 3, 4, 5, 6]
  puts simple_algorithm(data, 3)
  puts tpm_optimized_algorithm(data, 3)
end
