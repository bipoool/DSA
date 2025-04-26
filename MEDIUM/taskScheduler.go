package medium

import "github.com/emirpasic/gods/queues/priorityqueue"

// Sort the array from frequency
// Add all tasks frequency in priority queue
// Once you schedule a task just reduce the frequency and add it to the queue
// Each time check the queue the queue and then schedule the tasks
func leastInterval(tasks []byte, n int) int {
	ar := [26]int{}
	for i := range tasks {
		ar[tasks[i]-'A']++
	}

	pq := priorityqueue.NewWith(func(a, b any) int {
		return b.(int) - a.(int)
	})

	for i := range ar {
		if ar[i] == 0 {
			continue
		}
		pq.Enqueue(ar[i])
	}

	count := 0
	q := [][]int{}

	for !(len(q) == 0 && pq.Size() == 0) {
		for len(q) > 0 && q[0][1] <= count {
			pq.Enqueue(q[0][0])
			q = q[1:]
		}
		if pq.Size() != 0 {
			pqele, _ := pq.Dequeue()
			if pqele.(int)-1 != 0 {
				q = append(q, []int{pqele.(int) - 1, count + n + 1})
			}
		}
		count++
	}

	return count

}
