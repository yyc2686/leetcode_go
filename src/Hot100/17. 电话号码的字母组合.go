package Hot100

import (
	"reflect"
)

//17. 电话号码的字母组合
//给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。答案可以按 任意顺序 返回。
//
//给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。
//
//
//
//
//
//示例 1：
//
//输入：digits = "23"
//输出：["ad","ae","af","bd","be","bf","cd","ce","cf"]
//示例 2：
//
//输入：digits = ""
//输出：[]
//示例 3：
//
//输入：digits = "2"
//输出：["a","b","c"]
//
//
//提示：
//
//0 <= digits.length <= 4
//digits[i] 是范围 ['2', '9'] 的一个数字。

// 生成数字和字符串的对应关系map
var m = map[byte]string{
	'2': "abc",
	'3': "def",
	'4': "ghi",
	'5': "jkl",
	'6': "mno",
	'7': "pqrs",
	'8': "tuv",
	'9': "wxyz",
}

var ret []string

// 思路1：模拟？
func letterCombinations1(digits string) []string {
	// 初始化结果slice
	ret := make([]string, 0)

	// 生成数字和字符串的对应关系map
	m := make(map[byte][]byte, 0)
	m['2'] = []byte{'a', 'b', 'c'}
	m['3'] = []byte{'d', 'e', 'f'}
	m['4'] = []byte{'g', 'h', 'i'}
	m['5'] = []byte{'j', 'k', 'l'}
	m['6'] = []byte{'m', 'n', 'o'}
	m['7'] = []byte{'p', 'q', 'r', 's'}
	m['8'] = []byte{'t', 'u', 'v'}
	m['9'] = []byte{'w', 'x', 'y', 'z'}

	// 依次遍历digits，生成所有可能结果
	for i := 0; i < len(digits); i++ {
		runes, _ := m[digits[i]]
		tmp := make([]string, 0)

		if i == 0 {
			for _, v := range runes {
				tmp = append(tmp, string(v))
			}
		} else {
			for _, str := range ret {
				for _, v := range runes {
					tmp = append(tmp, str+string(v))
				}
			}
		}
		ret = tmp
	}
	return ret
}

// 思路1：模拟（优化map）
func letterCombinations2(digits string) []string {
	// 初始化结果slice
	ret := make([]string, 0)

	// 依次遍历digits，生成所有可能结果
	for i := 0; i < len(digits); i++ {
		str, _ := m[digits[i]]
		tmp := make([]string, 0)

		if i == 0 {
			for j := 0; j < len(str); j++ {
				tmp = append(tmp, string(str[j]))
			}
		} else {
			for _, v := range ret {
				for j := 0; j < len(str); j++ {
					tmp = append(tmp, v+string(str[j]))
				}
			}
		}
		ret = tmp
	}
	return ret
}

// 思路2：回溯
func backtrace(condition string, index int, selected string) {
	if index == len(condition) {
		if selected != "" {
			ret = append(ret, selected)
		}
		return
	}

	runes, _ := m[condition[index]]
	for i := 0; i < len(runes); i++ {
		backtrace(condition, index+1, selected+string(runes[i]))
	}

}

func letterCombinations3(digits string) []string {
	// 回溯
	ret = make([]string, 0)
	backtrace(digits, 0, "")
	return ret
}

// 思路3：队列
type Element interface{}

type sliceEntry struct {
	element []Element
}

func NewQueue() *sliceEntry {
	return &sliceEntry{}
}

type Queue interface {
	Offer(e Element) //向队列中添加元素
	Poll() Element   //移除队列中最前面的元素
	Clear() bool     //清空队列
	Size() int       //获取队列的元素个数
	IsEmpty() bool   //判断队列是否是空
}

//向队列中添加元素
func (entry *sliceEntry) Offer(e Element) {
	entry.element = append(entry.element, e)
}

//移除队列中最前面的额元素
func (entry *sliceEntry) Poll() Element {
	if entry.IsEmpty() {
		//fmt.Println("queue is empty!")
		return nil
	}

	firstElement := entry.element[0]
	entry.element = entry.element[1:]
	return firstElement
}

func (entry *sliceEntry) Clear() bool {
	if entry.IsEmpty() {
		//fmt.Println("queue is empty!")
		return false
	}
	for i := 0; i < entry.Size(); i++ {
		entry.element[i] = nil
	}
	entry.element = nil
	return true
}

func (entry *sliceEntry) Size() int {
	return len(entry.element)
}

func (entry *sliceEntry) IsEmpty() bool {
	if len(entry.element) == 0 {
		return true
	}
	return false
}

func (entry *sliceEntry) Contains(e Element) bool {
	for _, element := range entry.element {
		if Compare(&e, &element) {
			return true
		}
	}
	return false
}

func Compare(p1, p2 *Element) bool {
	if reflect.DeepEqual(p1, p2) {
		return true
	} else {
		return false
	}
}

func letterCombinations(digits string) []string {
	// 回溯
	ret = make([]string, 0)

	queue := NewQueue()
	queue.Offer("")
	for i := 0; i < len(digits); i++ {
		size := queue.Size()
		for j := 0; j < size; j++ {
			tmp := queue.Poll().(string)
			for _, str := range m[digits[i]] {
				queue.Offer(tmp + string(str))
			}
		}
	}

	for {
		if queue.IsEmpty() {
			break
		}
		tmp := queue.Poll().(string)
		if tmp != "" {
			ret = append(ret, tmp)
		}
	}
	return ret
}
