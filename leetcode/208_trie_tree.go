// A trie (pronounced as "try") or prefix tree is a tree data structure used to efficiently store and retrieve keys in a dataset of strings. There are various applications of this data structure, such as autocomplete and spellchecker.

// Implement the Trie class:

// Trie() Initializes the trie object.
// void insert(String word) Inserts the string word into the trie.
// boolean search(String word) Returns true if the string word is in the trie (i.e., was inserted before), and false otherwise.
// boolean startsWith(String prefix) Returns true if there is a previously inserted string word that has the prefix prefix, and false otherwise.
 

// Example 1:

// Input
// ["Trie", "insert", "search", "search", "startsWith", "insert", "search"]
// [[], ["apple"], ["apple"], ["app"], ["app"], ["app"], ["app"]]
// Output
// [null, null, true, false, true, null, true]

// Explanation
// Trie trie = new Trie();
// trie.insert("apple");
// trie.search("apple");   // return True
// trie.search("app");     // return False
// trie.startsWith("app"); // return True
// trie.insert("app");
// trie.search("app");     // return True
 

// Constraints:

// 1 <= word.length, prefix.length <= 2000
// word and prefix consist only of lowercase English letters.
// At most 3 * 104 calls in total will be made to insert, search, and startsWith.


package main

import (
	"fmt"
)

func main () {
	obj := Constructor()
	obj.Insert("apple")
	fmt.Println(obj.Search("apple"))
	fmt.Println(obj.Search("app"))
	fmt.Println(obj.StartsWith("app"))
	obj.Insert("app")
	fmt.Println(obj.Search("app"))
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

 type Node struct {
	val rune
	children []*Node
	flag bool
}
type Trie struct {
    root *Node
}


func Constructor() Trie {
    n := &Node{
		' ',
		make([]*Node, 26),
		false,
	}

	t := &Trie {
		n,
	}

	return *t
}


func (this *Trie) Insert(word string)  {
	findPath := this.root.children
    for i:=0; i<len(word); i++ {
		idx := int(word[i] - 'a') //'a'
		if findPath != nil {
			if findPath[idx] == nil { //[0]
				findPath[idx] = &Node{
                    rune(word[i]),
					make([]*Node, 26),
					false,
				}
			}
			if i == len(word) -1 {
				findPath[idx].flag = true
				break
			} else {
				findPath = findPath[idx].children
			}
		}
	}
}


func (this *Trie) Search(word string) bool {
	findPath := this.root.children
    for i:=0; i < len(word); i++ {
		idx := int(word[i] - 'a')
		if findPath == nil || findPath[idx] == nil {
			return false
		} 
		if findPath[idx] != nil {
			if i == len(word) - 1 {
				return findPath[idx].flag
			}
			findPath = findPath[idx].children
		}
	} 
	return false
}


func (this *Trie) StartsWith(prefix string) bool {
    findPath := this.root.children
    for i:=0; i < len(prefix); i++ {
		idx := int(prefix[i] - 'a')
		if findPath == nil || findPath[idx] == nil {
			return false
		} else if findPath[idx] != nil {
			findPath = findPath[idx].children
		}
	} 
	return true
}


/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */