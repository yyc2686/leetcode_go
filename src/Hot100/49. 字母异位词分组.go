package Hot100

import (
	"sort"
)

//49. 字母异位词分组
//给定一个字符串数组，将字母异位词组合在一起。字母异位词指字母相同，但排列不同的字符串。
//
//示例:
//
//输入: ["eat", "tea", "tan", "ate", "nat", "bat"]
//输出:
//[
//["ate","eat","tea"],
//["nat","tan"],
//["bat"]
//]
//说明：
//
//所有输入均为小写字母。
//不考虑答案输出的顺序。

// 输入全为小写字母，可直接使用长度为26的数组作为key
func groupAnagrams(strs []string) [][]string {
	ret := [][]string{}
	//stringMap := map[[26]int][]string{}
	stringMap := make(map[[26]int][]string, 0)

	for id, str := range strs {
		key := [26]int{}
		for _, ch := range str {
			key[ch-'a']++
		}

		if _, ok := stringMap[key]; ok {
			stringMap[key] = append(stringMap[key], strs[id])
		} else {
			stringMap[key] = []string{strs[id]}
		}
	}

	for _, val := range stringMap {
		//sort.Strings(val)
		ret = append(ret, val)
	}

	return ret
}

//互为字母异位词的两个字符串包含的字母相同
//排序之后的字符串作为哈希表的键

// map中存str（减少时间消耗）
func groupAnagrams2(strs []string) [][]string {
	ret := [][]string{}
	//stringMap := map[string][]string{}
	stringMap := make(map[string][]string, 0)

	for id, str := range strs {

		s := []byte(str)
		sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })
		sortedString := string(s)
		stringMap[sortedString] = append(stringMap[sortedString], strs[id])
	}

	for _, val := range stringMap {
		sort.Strings(val)
		ret = append(ret, val)
	}

	return ret
}

// map中存id（降低空间消耗）
func groupAnagrams1(strs []string) [][]string {
	ret := [][]string{}
	stringMap := map[string][]int{}

	for id, str := range strs {
		//sortedString := sortString(str)

		s := []byte(str)
		sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })
		sortedString := string(s)

		if v, ok := stringMap[sortedString]; ok {
			stringMap[sortedString] = append(v, id)
		} else {
			stringMap[sortedString] = []int{id}
		}
	}

	for _, strIds := range stringMap {
		tmp := []string{}
		for _, strId := range strIds {
			tmp = append(tmp, strs[strId])
		}
		sort.Strings(tmp)
		ret = append(ret, tmp)
	}

	return ret
}
