package pinyinsplit

import (
	"strings"
	"testing"
)

func Test_match(t *testing.T) {
	tree := buildTrie(pingYinTable)
	for _, pinyin := range pingYinTable {
		ok := tree.match(pinyin)
		if !ok {
			t.Errorf("error")
		}
	}

	type Test struct {
		spell string
		ok    bool
	}

	tests := []Test{
		{
			"", //空串
			false,
		},
		{
			"test", //没有这个节点
			false,
		},
		{
			"we", //非结构节点
			false,
		},
	}

	for i, test := range tests {
		ok := tree.match(test.spell)
		if ok != test.ok {
			t.Errorf("test error %d, spell: %s", i, test.spell)
		}
	}
}

func Test_searchLongest(t *testing.T) {
	tree := buildTrie([]string{
		"abc",
		"abcdefg",
		"abcf",
		"ak",
		"b",
	})

	type Test struct {
		spell    string
		lastChar byte //最后匹配的字母
		deep     int  //匹配的树的深度
	}

	tests := []Test{
		{
			"b",
			'b',
			1,
		},
		{
			"abc",
			'c',
			3,
		},
		{
			"abcg",
			'c',
			3,
		},
		{
			"abcdefgllllllll", //到达根节点
			'g',
			7,
		},
	}

	for i, test := range tests {
		node := tree.searchLongest([]byte(test.spell))
		if node == nil {
			t.Errorf("searchLongest error %d, %s", i, test.spell)
		}
	}

	node := tree.searchLongest([]byte(""))
	if node != nil {
		t.Errorf("searchLongest error")
	}

	node = tree.searchLongest([]byte("s"))
	if node != nil {
		t.Errorf("searchLongest error")
	}
}

func Test_Split(t *testing.T) {
	type Test struct {
		input  string //原始拼音
		output string //分割之后的拼音
	}
	tests := []Test{
		{
			"shenrulijiejisuanjixitong",
			"shen ru li jie ji suan ji xi tong",
		},
		{
			"suanfadaolun",
			"suan fa dao lun",
		},
		{
			"suan     fadao     lun", // 忽略空格，及非拼音字符
			"suan fa dao lun",
		},
		{
			"suan我fadaolun", // 忽略空格，及非拼音字符
			"suan fa dao lun",
		},
		{
			"xianshirenminzhengfu",
			"xian shi ren min zheng fu", //目前的算法不能正确处理'西安'这种拼音
		},
	}

	for _, test := range tests {
		words := Split(test.input)
		output := strings.Join(words, " ")
		if output != test.output {
			t.Errorf("split error input = %s, output=%s", test.input, output)
		}
	}

}
