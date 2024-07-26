# frozen_string_literal: true

# Скользящее среднее

# Наивный алгоритм
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

# Метод двух указателей
def optimized_algorithm(timeseries, k)
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
