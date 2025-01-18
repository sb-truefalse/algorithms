# frozen_string_literal: true

require_relative './solution'

RSpec.describe "Moving average" do
  let(:timeseries) { [1, 2, 3, 4, 5] }
  let(:k) { 2 }

  describe "#moving_avrg" do
    subject { moving_avrg(timeseries, k) }
    it do
      is_expected.to eq([1.5, 2.5, 3.5, 4.5])
    end
  end

  describe "#moving_avrg_with_2_pointers" do
    subject { moving_avrg_with_2_pointers(timeseries, k) }

    it do
      is_expected.to eq([1.5, 2.5, 3.5, 4.5])
    end
  end
end
