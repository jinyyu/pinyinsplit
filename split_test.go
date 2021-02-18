package pinyinsplit

import "testing"

func TestSplit(t *testing.T) {
	tree := buildTrie()
	for _, pinyin := range pingyinData {
		ok := tree.search(pinyin)
		if !ok {
			t.Errorf("error")
		}
	}

	str := ""
	ok := tree.search(str)
	if ok {
		t.Errorf("must not found")
	}

	str = "kugouyinyue"
	ok = tree.search(str)
	if ok {
		t.Errorf("must not found")
	}

	str = "we"
	ok = tree.search(str)
	if ok {
		t.Errorf("must not found")
	}
}
