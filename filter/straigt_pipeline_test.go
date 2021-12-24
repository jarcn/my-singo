package filter

import "testing"

func TestFilterPipeline(t *testing.T) {
	splitFilter := NewSplitFilter(",")
	toIntFilter := NewToIntFilter()
	sumFilter := NewSumFilter()
	//Demo 需要强制作filter添加顺序
	sp := NewStraigthPipeline("chenjia", splitFilter, toIntFilter, sumFilter)
	ret, err := sp.Process("1,2,3,5")
	if err != nil {
		t.Fatal(err)
	}
	if ret != 6 {
		t.Fatalf("the expected is 6,but the actual is %d", ret)
	}
}
