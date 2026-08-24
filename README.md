# Algorithms

![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)

Algorithmic training - solving problems and taking notes.

[![LeetCode](https://img.shields.io/badge/LeetCode-FFA116?style=for-the-badge&logo=leetcode&logoColor=black)](https://leetcode.com/u/venexene/)

## Structure

Each problem lives in its own directory with:

- `solution.go` — the implementation
- `notes.md` — solution idea, complexity, and mistakes encountered

## Solved

**Solved: 111 problems · 16 topics**

**Difficulty: 42 Easy · 64 Medium · 5 Hard**

| Topic | Solved |
|------|--------|
| Arrays & Strings | 13 |
| Hash Tables | 8 |
| Two Pointers | 6 |
| Sliding Window | 3 |
| Stack & Queue | 6 |
| Binary Search | 8 |
| Linked Lists | 8 |
| Trees | 14 |
| Graphs | 9 |
| Dynamic Programming | 10 |
| Greedy | 3 |
| Cache | 2 |
| Heap | 6 |
| Trie | 3 |
| Backtracking | 6 |
| Intervals | 6 |

## Arrays & Strings (13)

- [x] [Two Sum](./arrs-and-strs/two-sum/) - hash map, two pointers, brute force
- [x] [Contains Duplicate](./arrs-and-strs/contains-duplicate/) - hash set, sorting
- [x] [Valid Anagram](./arrs-and-strs/valid-anagram/) - frequency array, hash map, sorting
- [x] [Valid Palindrome](./arrs-and-strs/valid-palindrome/) - two pointers, filter + compare
- [x] [Remove Duplicates from Sorted Array](./arrs-and-strs/remove-duplicates-from-sorted-array/) - two pointers (read/write), in-place
- [x] [Best Time to Buy and Sell Stock](./arrs-and-strs/best-time-to-buy-and-sell-stock/) - one pass, track min
- [x] [Group Anagrams](./arrs-and-strs/group-anagrams/) - hash map, sort key, frequency array key
- [x] [Longest Substring Without Repeating Characters](./arrs-and-strs/longest-substring-without-repeating-characters/) - sliding window, map of positions
- [x] [Longest Palindromic Substring](./arrs-and-strs/longest-palindromic-substring/) - expand around center
- [x] [Majority Element](./arrs-and-strs/majority-element/) - Boyer-Moore voting, hash map
- [x] [Add Binary](./arrs-and-strs/add-binary/) - digit-by-digit, carry, from the end
- [x] [Spiral Matrix](./arrs-and-strs/spiral-matrix/) - four boundaries, spiral traversal
- [x] [Sort Colors](./arrs-and-strs/sort-colors/) - Dutch National Flag, three pointers

## Hash Tables (8)

- [x] [Intersection of Two Arrays](./hash-tables/intersection-of-two-arrays/) - hash set, delete for uniqueness
- [x] [Jewels and Stones](./hash-tables/jewels-and-stones/) - hash set lookup
- [x] [Ransom Note](./hash-tables/ransom-note/) - frequency map, frequency array
- [x] [First Unique Character in a String](./hash-tables/first-unique-character-in-a-string/) - frequency map, two passes
- [x] [Subarray Sum Equals K](./hash-tables/subarray-sum-equals-k/) - prefix sum + hash map
- [x] [Top K Frequent Elements](./hash-tables/top-k-frequent-elements/) - min-heap, bucket sort
- [x] [Longest Consecutive Sequence](./hash-tables/longest-consecutive-sequence/) - hash set, sequence start detection
- [x] [Longest Palindrome](./hash-tables/longest-palindrome/) - frequency set, even/odd counts

## Two Pointers (6)

- [x] [Two Sum II](./two-pointers/two-sum-2/) - two pointers on sorted array
- [x] [Remove Element](./two-pointers/remove-element/) - write pointer, in-place
- [x] [Merge Sorted Array](./two-pointers/merge-sorted-array/) - three pointers from the end
- [x] [3Sum](./two-pointers/sum-3/) - sort + two pointers, hash map
- [x] [Container With Most Water](./two-pointers/container-with-most-water/) - two pointers greedy
- [x] [Trapping Rain Water](./two-pointers/trapping-rain-water/) - two pointers, leftMax/rightMax

## Sliding Window (3)

- [x] [Maximum Average Subarray I](./sliding-window/maximum-average-sum-i/) - fixed-size window
- [x] [Minimum Size Subarray Sum](./sliding-window/minimum-size-subarray-sum/) - variable-size window
- [x] [Permutation in String](./sliding-window/permutation-in-string/) - fixed window + frequency counter

## Stack & Queue (6)

- [x] [Valid Parentheses](./stack-and-queue/valid-parentheses/) - stack
- [x] [Implement Queue using Stacks](./stack-and-queue/implement-queue-using-stacks/) - two stacks, amortized O(1)
- [x] [Baseball Game](./stack-and-queue/baseball-game/) - stack of ints
- [x] [Min Stack](./stack-and-queue/min-stack/) - two stacks, history of mins
- [x] [Evaluate Reverse Polish Notation](./stack-and-queue/evaluate-reverse-polish-notation/) - stack
- [x] [Daily Temperatures](./stack-and-queue/daily-temperatures/) - monotonic decreasing stack

## Binary Search (8)

- [x] [Binary Search](./binary-search/binary-search/) - classic binary search
- [x] [Search Insert Position](./binary-search/search-insert-position/) - binary search, return low
- [x] [Sqrt(x)](./binary-search/sqrt-x/) - binary search on answer
- [x] [Search in Rotated Sorted Array](./binary-search/search-in-rotated-sorted-array/) - one binary search, check sorted half
- [x] [Find Minimum in Rotated Sorted Array](./binary-search/find-minimum-in-rotated-sorted-array/) - compare mid with high
- [x] [Koko Eating Bananas](./binary-search/koko-eating-bananas/) - binary search on answer, ceil division
- [x] [First Bad Version](./binary-search/first-bad-version/) - binary search, find first true
- [x] [Time Based Key-Value Store](./binary-search/time-based-key-value-store/) - hash map + binary search, last ≤ timestamp

## Linked Lists (8)

- [x] [Reverse Linked List](./linked-lists/reverse-linked-list/) - iterative, recursive
- [x] [Merge Two Sorted Lists](./linked-lists/merge-two-sorted-lists/) - dummy node, tail pointer
- [x] [Merge K Sorted Lists](./linked-lists/merge-k-sorted-lists/) - min-heap of k heads, O(N log k)
- [x] [Linked List Cycle](./linked-lists/linked-list-cycle/) - hash set, Floyd's algorithm
- [x] [Middle of the Linked List](./linked-lists/middle-of-the-linked-list/) - fast & slow pointers
- [x] [Intersection of Two Linked Lists](./linked-lists/intersection-of-two-linked-lists/) - hash set, two pointers equalizing distance
- [x] [Remove Nth Node From End of List](./linked-lists/remove-nth-node-from-end-of-list/) - counter, dummy + fast/slow
- [x] [Add Two Numbers](./linked-lists/add-two-numbers/) - digit-by-digit with carry

## Trees (14)

- [x] [Maximum Depth of Binary Tree](./trees/maximum-depth-of-binary-tree/) - DFS recursive
- [x] [Same Tree](./trees/same-tree/) - DFS recursive comparison
- [x] [Invert Binary Tree](./trees/invert-binary-tree/) - DFS swap left/right
- [x] [Path Sum](./trees/path-sum/) - DFS, leaf check
- [x] [Binary Tree Level Order Traversal](./trees/binary-tree-level-order-traversal/) - BFS with level tracking
- [x] [Lowest Common Ancestor of a BST](./trees/lowers-common-ancestor-of-a-binary-search-tree/) - BST property shortcut
- [x] [Binary Tree Right Side View](./trees/binary-tree-right-side-view/) - DFS right-first, depth tracking
- [x] [Validate Binary Search Tree](./trees/validate-binary-search-tree/) - range passing top-down
- [x] [Diameter of Binary Tree](./trees/diameter-of-binary-search/) - DFS post-order, max(left+right)
- [x] [Balanced Binary Tree](./trees/balanced-binary-tree/) - DFS post-order, height + flag
- [x] [Serialize and Deserialize Binary Tree](./trees/serialize-and-deserialize-binary-tree/) - BFS, nil markers, trim trailing #
- [x] [Construct Binary Tree from Preorder and Inorder Traversal](./trees/construct-binary-tree-from-preorder-and-inorder/) - recursion + index map
- [x] [Kth Smallest Element in a BST](./trees/kth-smallest-element-in-a-bst/) - in-order traversal
- [x] [Lowest Common Ancestor of a Binary Tree](./trees/lowest-common-ancestor-of-a-binary-tree/) - recursive DFS

## Graphs (9)

- [x] [Find if Path Exists in Graph](./graphs/find-if-path-exists-in-graph/) - DFS + adjacency list, Union Find
- [x] [Flood Fill](./graphs/flood-fill/) - DFS on matrix
- [x] [Number of Islands](./graphs/number-of-islands/) - DFS flood fill, component counting
- [x] [Clone Graph](./graphs/clone-graph/) - DFS + visited map (orig → clone)
- [x] [Course Schedule](./graphs/course-schedule/) - topological sort, DFS with 3 states
- [x] [Pacific Atlantic Water Flow](./graphs/pacific-atlantic-water-flow/) - two DFS from borders, reverse direction
- [x] [Rotting Oranges](./graphs/rotting-oranges/) - multi-source BFS by levels
- [x] [01 Matrix](./graphs/01-matrix/) - multi-source BFS from all zeros
- [x] [Accounts Merge](./graphs/accounts-merge/) - union-find (DSU), email components

## Dynamic Programming (10)

- [x] [Climbing Stairs](./dynamic-programming/climbing-stairs/) - Fibonacci, two variables
- [x] [Min Cost Climbing Stairs](./dynamic-programming/min-cost-climbing-stairs/) - in-place DP, min of two paths
- [x] [House Robber](./dynamic-programming/house-robber/) - in-place DP, max(skip, rob)
- [x] [Maximum Subarray](./dynamic-programming/maximum-subarray/) - Kadane's algorithm, in-place
- [x] [Coin Change](./dynamic-programming/coin-change/) - DP with full coin iteration, minimization
- [x] [Unique Paths](./dynamic-programming/unique-path/) - DP on matrix, top + left
- [x] [Longest Increasing Subsequence](./dynamic-programming/longest-increasing-subsequence/) - DP O(n²), ending at i
- [x] [Jump Game](./dynamic-programming/jump-game/) - greedy, track max reach
- [x] [Word Break](./dynamic-programming/word-break/) - DP on string prefixes, O(n²)
- [x] [Partition Equal Subset Sum](./dynamic-programming/partition-equal-subset-sum/) - 0/1 knapsack DP, target = total/2

## Greedy (3)

- [x] [Jump Game II](./greedy/jump-game-2/) - BFS-levels greedy, two boundaries
- [x] [Gas Station](./greedy/gas-station/) - greedy pass, tank reset
- [x] [Partition Labels](./greedy/partition-labels/) - last-index map, two passes

## Cache (2)

- [x] [LRU Cache](./cache/lru-cache/) - doubly linked list + hash map, sentinel
- [x] [LFU Cache](./cache/lfu-cache/) - frequency buckets + hash map, sentinel, minFreq

## Heap (6)

- [x] [Last Stone Weight](./heap/last-stone-weight/) - max-heap, smash two heaviest
- [x] [Kth Largest Element in a Stream](./heap/kth-largest-element-in-a-stream/) - min-heap of size k
- [x] [Kth Largest Element in Array](./heap/kth-largest-element-in-array/) - min-heap of size k, QuickSelect alternative
- [x] [K Closest Points to Origin](./heap/k-closest-points-to-origin/) - max-heap, squared distance
- [x] [Task Scheduler](./heap/task-scheduler/) - max-heap + cooldown queue, greedy formula
- [x] [Find Median from Data Stream](./heap/find-median-from-data-stream/) - two heaps (max + min), balanced

## Trie (3)

- [x] [Implement Trie](./trie/implement-trie/) - array of 26 children, isEnd
- [x] [Design Add and Search Words](./trie/design-add-and-search-words-data-structure/) - trie + DFS for wildcard
- [x] [Replace Words](./trie/replace-words/) - find shortest root prefix

## Backtracking (6)

- [x] [Subsets](./backtracking/subsets/) - include/exclude, start-index
- [x] [Permutations](./backtracking/permutations/) - swap in-place, remove element, visited map
- [x] [Combination Sum](./backtracking/combination-sum/) - start-index + reuse, sort + break
- [x] [Letter Combinations of a Phone Number](./backtracking/letter-combinations/) - backtracking over digit-to-letters map
- [x] [Generate Parentheses](./backtracking/generate-parentheses/) - prune invalid prefixes, open/close invariants
- [x] [Word Search](./backtracking/word-search/) - DFS on matrix, mark + backtrack

## Intervals (6)

- [x] [Meeting Rooms](./intervals/meeting-rooms/) - sort by start, check adjacent overlaps
- [x] [Summary Ranges](./intervals/summary-ranges/) - two pointers, build ranges
- [x] [Merge Intervals](./intervals/merge-intervals/) - sort by start, greedy merge
- [x] [Insert Interval](./intervals/insert-interval/) - linear 3-phase merge
- [x] [Non-overlapping Intervals](./intervals/non-overlapping-intervals/) - sort by start, keep min end
- [x] [Meeting Rooms II](./intervals/meeting-rooms-2/) - sweep-line, sort starts/ends