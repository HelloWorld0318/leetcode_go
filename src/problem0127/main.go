package problem0127

func ladderLength(beginWord, endWord string, wordList []string) int {
	if len(wordList) == 0 {
		return 0
	}
	graph := make(map[string][]string)
	buildGraph(beginWord, wordList, &graph)
	return bfs(beginWord, endWord, &graph)
}

func bfs(beginWord, endWord string, graph *map[string][]string) int {
	queue := make([]*WordNode, 0)
	queue = append(queue, &WordNode{Word: beginWord, Step: 1})
	hasVisited := make(map[string]bool)
	for len(queue) != 0 {
		wordNode := queue[0]
		queue = queue[1:]
		step := wordNode.Step
		hasVisited[wordNode.Word] = true

		if wordNode.Word == endWord {
			return step
		}
		if wordList, ok := (*graph)[wordNode.Word]; ok {
			for _, word := range wordList {
				if !hasVisited[word] {
					queue = append(queue, &WordNode{Word: word, Step: step + 1})
				}
			}
		}
	}
	return 0
}

func buildGraph(beginWord string, wordList []string, graph *map[string][]string) {
	wordList = append(wordList, beginWord)
	i, j := 0, 0
	for i < len(wordList) {
		j = i + 1
		for j < len(wordList) {
			if isConnection(wordList[i], wordList[j]) {
				(*graph)[wordList[i]] = append((*graph)[wordList[i]], wordList[j])
				(*graph)[wordList[j]] = append((*graph)[wordList[j]], wordList[i])
			}
			j++
		}
		i++
	}
}

func isConnection(word1, word2 string) bool {
	if len(word2) != len(word1) {
		return false
	}
	differentChar := 0
	i := 0
	for i < len(word1) {
		if word1[i] != word2[i] {
			differentChar++
		}
		i++
	}
	return differentChar == 1
}

type WordNode struct {
	Word string
	Step int
}
