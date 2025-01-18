# frozen_string_literal: true

require_relative './solution'

RSpec.describe 'Two Sum' do
  let(:list) { [1, 2, 3, 4, 5] }
  let(:x) { 5 }

  describe '#two_sum' do
    subject { two_sum(list, x) }
    it do
      is_expected.to eq([0, 3])
    end
  end

  describe '#two_sum_with_2_pointers' do
    subject { two_sum_with_2_pointers(list, x) }

    it do
      is_expected.to eq([0, 3])
    end
  end

  describe '#two_sum_with_hash' do
    subject { two_sum_with_hash(list, x) }

    it do
      is_expected.to eq([0, 3])
    end
  end
end
