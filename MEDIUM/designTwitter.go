package medium

import "container/heap"

type MaxPq []Post

func (p *MaxPq) Push(e any) {
	*p = append(*p, e.(Post))
}

func (p *MaxPq) Pop() any {
	old := *p
	res := old[len(*p)-1]
	*p = old[0 : len(*p)-1]
	return res
}

func (p *MaxPq) Peak() any {
	old := *p
	return old[0]
}

func (p MaxPq) Len() int      { return len(p) }
func (h MaxPq) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h MaxPq) Less(i, j int) bool {
	return h[i].timestamp > h[j].timestamp
}

type Twitter struct {
	followMap map[int]map[int]int
	posts     map[int][]Post
	timestamp int
}

type Post struct {
	tweetId   int
	timestamp int
}

func Constructor() Twitter {
	return Twitter{
		followMap: map[int]map[int]int{},
		posts:     map[int][]Post{},
		timestamp: 0,
	}
}

func (this *Twitter) PostTweet(userId int, tweetId int) {
	this.timestamp++
	posts, ok := this.posts[userId]
	if !ok {
		posts = []Post{}
	}
	newPost := Post{
		tweetId:   tweetId,
		timestamp: this.timestamp,
	}
	posts = append(posts, newPost)
	this.posts[userId] = posts
}

func (this *Twitter) GetNewsFeed(userId int) []int {

	allFollowees, ok := this.followMap[userId]
	if !ok {
		allFollowees = map[int]int{
			userId: 1,
		}
	}
	feed := []int{}

	feadQueue := MaxPq{}
	heap.Init(&feadQueue)

	for fid := range allFollowees {
		count := 0
		if allFollowees[fid] == 0 {
			continue
		}
		for _, post := range this.posts[allFollowees[fid]] {
			heap.Push(&feadQueue, post)
			count++
			if count == 10 {
				break
			}
		}
	}

	for i := 0; i < 10; i++ {
		if feadQueue.Len() == 0 {
			break
		}
		post := heap.Pop(&feadQueue).(Post).tweetId
		feed = append(feed, post)
		i++
	}
	return feed

}

func (this *Twitter) Follow(followerId int, followeeId int) {
	user, ok := this.followMap[followerId]
	if !ok {
		user = map[int]int{}
		user[followerId] = 1
	}
	user[followeeId] = 1
	this.followMap[followerId] = user
}

func (this *Twitter) Unfollow(followerId int, followeeId int) {
	user, ok := this.followMap[followerId]
	if !ok {
		user = map[int]int{}
		user[followerId] = 1
	}
	user[followeeId] = 0
	this.followMap[followerId] = user
}
