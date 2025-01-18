# frozen_string_literal: true

require_relative './solution'

RSpec.describe "Searching for a list item" do
  let(:list) { [1, 2, 3, 4, 5] }
  let(:x) { 5 }

  describe "#linear_search" do
    subject { linear_search(list, x) }
    it do
      is_expected.to eq(4)
    end
  end

  describe "#linear_search" do
    subject { linear_search(list, x) }

    it do
      is_expected.to eq(4)
    end
  end
end
