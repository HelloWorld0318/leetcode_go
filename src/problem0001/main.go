package main

import (
	"crypto/md5"
	"fmt"
	"regexp"
	"sort"
	"strings"
	"time"
)

const _timeStamp20170701 = 1498838400

func main() {
	fmt.Println("======================")

	fmt.Println(GetHash(31110987))

	fmt.Println(time.Now().Unix())
	fmt.Println("======================")

	arr := []int64{12,323}

	arrCombine := getAllGroupCombile(arr)
	for _, item := range arrCombine {
		fmt.Println(item)
	}
	fmt.Println("======================")

	str := "#小豆芽#额分罚恶风#额发#"
	fmt.Println("======================", len(str))
	fmt.Println(string(str[2]))
	topics, onlyHasTopic, textLen, err := DetectTopicsFromDynamicDesc(str)
	fmt.Println(topics)
	fmt.Println(onlyHasTopic)
	fmt.Println(textLen)
	fmt.Println(err)
}

type Stu struct {
	Name string `json:"name"`
	Age  string `json:"age"`
}

func twoSum(nums []int, target int) []int {
	hash := make(map[int]int)

	for index, v := range nums {
		value, ok := hash[target-v]
		if ok {
			return []int{value, index}
		}
		hash[v] = index
	}
	return []int{}
}

func shardDynId(dynamicID uint64) (shardInfo *ShardInfo, err error) {
	ts := (dynamicID >> 32) + _timeStamp20170701
	tm := time.Unix(int64(ts), 0)
	fmt.Println("timestamp:", tm)
	shardInfo = &ShardInfo{
		Year:  int32(tm.Year()),
		Month: int32(tm.Month()),
		Rand:  (dynamicID & 0xffffffff) >> 20,
	}
	return
}

type ShardInfo struct {
	Year  int32
	Month int32
	Rand  uint64
}

type ExpiredTs struct {
	ExpiredTs int64 `json:"ex_ts"`
}

func GetHash(mid int64) int32 {
	midStr := fmt.Sprintf("%d", mid)
	bys := md5.Sum([]byte(midStr))
	fmt.Println(len(bys))

	h := uint32(0)
	for _, by := range bys {
		h *= 16777619
		h ^= uint32(by)
	}
	return int32(h % 10)
}


func getAllGroupCombile(groupTopicIds []int64) (groupIdCombine []string) {
	//两个话题关联
	groupIdCombine = make([]string, 0)
	if len(groupTopicIds) == 0 {
		return
	}
	sort.Slice(groupTopicIds, func(i, j int) bool { return groupTopicIds[i] < groupTopicIds[j] })
	//三个话题关联
	groupIdCombine = make([]string, 0)
	if len(groupTopicIds) >= 2 {
		i, j := 0, 0
		for i < len(groupTopicIds) {
			j = i + 1
			for j < len(groupTopicIds) {
				groupIdCombine = append(groupIdCombine, strings.Join([]string{fmt.Sprintf("%d", groupTopicIds[i]),
					fmt.Sprintf("%d", groupTopicIds[j])}, ","))
				j++
			}
			i++
		}
	}
	if len(groupTopicIds) >= 3 {
		i, j, k := 0, 0, 0
		for i < len(groupTopicIds) {
			j = i + 1
			for j < len(groupTopicIds) {
				k = j + 1
				for k < len(groupTopicIds) {
					groupIdCombine = append(groupIdCombine, strings.Join([]string{fmt.Sprintf("%d", groupTopicIds[i]),
						fmt.Sprintf("%d", groupTopicIds[j]),
						fmt.Sprintf("%d", groupTopicIds[k])},
						","))
					k++
				}
				j++
			}
			i++
		}
	}
	return
}



const (
	DETECT_STATE_RESTART = iota
	DETECT_STATE_START_DETECT
	DETECT_STATE_IN_TOPIC
	DETECT_STATE_END_DETECT
)

func DetectTopicsFromDynamicDesc(desc string) (topics []string, onlyHasTopic bool,textLen int32, err error) {
	ts, onlyHasTopic, textLen, err := scan(desc)
	if err != nil {
		return
	}
	// 去重
	unique := make(map[string]struct{})
	for _, topic := range ts {
		unique[topic] = struct{}{}
	}
	for topic := range unique {
		topics = append(topics, topic)
	}
	return
}

func scan(desc string) (topics []string, onlyHasTopic bool, textLen int32, err error) {
	var (
		start  int
		end    int
		length int
		state  = DETECT_STATE_RESTART

		textStart int
	)
	onlyHasTopic = true
	desc = removeUrl(desc)

	textArrExcluedTag := make([]string, 0)
	for i := 0; i < len(desc); {
		c := desc[i]
		state = detect(c, state, &onlyHasTopic)

		if state == DETECT_STATE_START_DETECT {
			start = i + 1
			length = 0
			if textStart < i-1 {
				textArrExcluedTag = append(textArrExcluedTag, trim(desc[textStart:i]))
			}
		}
		if state == DETECT_STATE_IN_TOPIC {
			length++
			if length > 32 {
				state = DETECT_STATE_RESTART
				textStart = i
			}
		}
		if state == DETECT_STATE_END_DETECT {
			end = i
			candidate := trim(desc[start:end])
			if len(candidate) > 0 {
				topics = append(topics, candidate)
			} else {
				state = DETECT_STATE_START_DETECT
				start = i + 1
				length = 0
			}
			textStart = i + 1
		}
		if state == DETECT_STATE_RESTART {
			if i == len(desc)-1 {
				textArrExcluedTag = append(textArrExcluedTag, trim(desc[textStart:]))
			}
		}
		if c <= 0x7f {
			i += 1
		} else if c < 0xc0 {
			return
		} else if c < 0xe0 {
			i += 2
		} else if c < 0xf0 {
			i += 3
		} else {
			if len(desc) < i+4 {
				return
			}
			isEmoji := false
			// emoji in [\xF0\x9F\x98\x81, \xF0\x9F\x98\xBF] or [\xF0\x9F\x99\x80, \xF0\x9F\x99\x8F]
			if c == 0xf0 && desc[i+1] == 0x9f {
				if desc[i+2] == 0x98 {
					if desc[i+3] >= 0x81 && desc[i+3] <= 0xbf {
						isEmoji = true
					}
				} else if desc[i+2] == 0x99 {
					if desc[i+3] >= 0x80 && desc[i+3] <= 0x8f {
						isEmoji = true
					}
				}
			}
			if isEmoji {
				state = DETECT_STATE_RESTART
			}
			i += 4
		}
		if i >= len(desc)-1 {
			textArrExcluedTag = append(textArrExcluedTag, trim(desc[textStart:]))
		}
	}

	for _, item := range textArrExcluedTag {
		words := ([]rune)(item)
		textLen += int32(len(words))
	}
	return
}

func detect(c uint8, enterState int, onlyHasTopic *bool) (exitState int) {
	nextState := DETECT_STATE_RESTART
	if (enterState == DETECT_STATE_RESTART || enterState == DETECT_STATE_END_DETECT) &&
		c != 0x23 && c != 0x20 && c != 0x0a {
		*onlyHasTopic = false
	}
	if c == 0x40 || c == 0x0a {
		return DETECT_STATE_RESTART
	}
	switch enterState {
	case DETECT_STATE_RESTART:
		if c == 0x23 {
			nextState = DETECT_STATE_START_DETECT
		} else {
			nextState = DETECT_STATE_RESTART
		}
	case DETECT_STATE_START_DETECT:
		if c == 0x23 {
			nextState = DETECT_STATE_START_DETECT
		} else {
			nextState = DETECT_STATE_IN_TOPIC
		}
	case DETECT_STATE_IN_TOPIC:
		if c == 0x23 {
			nextState = DETECT_STATE_END_DETECT
		} else {
			nextState = DETECT_STATE_IN_TOPIC
		}
	case DETECT_STATE_END_DETECT:
		if c == 0x23 {
			nextState = DETECT_STATE_START_DETECT
		} else {
			nextState = DETECT_STATE_RESTART
		}
	default:

	}
	return nextState
}

func trim(content string) string {
	if len(content) == 0 {
		return ""
	}
	var (
		start           = 0
		totalLen        = len(content)
		end             = totalLen - 1
		subLenFromBegin = 0
	)
	for i := start; i <= end; i++ {
		c := content[i]
		if c == ' ' {
			subLenFromBegin++
			continue
		} else if (c == 0xe3) && ((i + 2) < totalLen) {
			c1 := content[i+1]
			c2 := content[i+2]
			if c1 == 0x80 && c2 == 0x80 {
				i += 2
				subLenFromBegin += 3
				continue
			}
		}
		break
	}
	content = content[subLenFromBegin:]
	subLenFromBegin = 0
	totalLen = len(content)
	end = totalLen - 1
	for i := end; i >= 0; i-- {
		c := content[i]
		if c == ' ' {
			subLenFromBegin++
			continue
		}
		if (c == 0x80) && ((i - 2) >= 0) {
			c2 := content[i-2]
			c1 := content[i-1]
			if c2 == 0xe3 && c1 == 0x80 {
				i -= 2
				continue
			}
		}
		break
	}
	return content[:totalLen-subLenFromBegin]
}

func removeUrl(subject string) string {
	reg, _ := regexp.Compile(`https?://[0-9a-zA-Z!@#$&*=./?~_%-]+`)
	//subject := `#sdf#http://a #你好# http://a.bb`
	sb := strings.Builder{}
	matchIndex := reg.FindAllStringIndex(subject, len(subject))
	cur := 0
	for _, matchRange := range matchIndex {
		l := matchRange[0]
		r := matchRange[1]
		sb.WriteString(subject[cur:l])
		cur = r
	}
	if cur < len(subject) {
		sb.WriteString(subject[cur:])
	}
	return sb.String()
}
