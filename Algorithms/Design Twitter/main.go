package main

import (
	"fmt"
)

func main() {
	twitter := Constructor()

	// User 1 posts a new tweet (id = 5).
	twitter.PostTweet(1, 5)

	// User 1's news feed should return a list with 1 tweet id -> [5].
	fmt.Println(twitter.GetNewsFeed(1))

	// User 1 follows user 2.
	twitter.Follow(1, 2)

	// User 2 posts a new tweet (id = 6).
	twitter.PostTweet(2, 6)

	// User 1's news feed should return a list with 2 tweet ids -> [6, 5].
	// Tweet id 6 should precede tweet id 5 because it is posted after tweet id 5.
	fmt.Println(twitter.GetNewsFeed(1))

	// User 1 unfollows user 2.
	twitter.Unfollow(1, 2)

	// User 1's news feed should return a list with 1 tweet id -> [5],
	// since user 1 is no longer following user 2.
	fmt.Println(twitter.GetNewsFeed(1))
}

type Twitter struct {
	tweet [][]int
	// 用户关注的人
	users map[int]map[int]struct{}
}

type Tweet struct {
	// 推特id
	tweet int
	// 哪个用户发表的
	user int
}

/** Initialize your data structure here. */
func Constructor() Twitter {
	return Twitter{
		tweet: make([][]int, 0),
		users: make(map[int]map[int]struct{}),
	}
}

/** Compose a new tweet. */
func (this *Twitter) PostTweet(userId int, tweetId int) {
	this.tweet = append(this.tweet, []int{tweetId, userId})
}

/** Retrieve the 10 most recent tweet ids in the user's news feed. Each item in the news feed must be posted by users who the user followed or by the user herself. Tweets must be ordered from most recent to least recent. */
func (this *Twitter) GetNewsFeed(userId int) []int {
	res := make([]int, 0, 10)
	for i := len(this.tweet) - 1; i >= 0; i-- {
		if this.tweet[i][1] == userId {
			res = append(res, this.tweet[i][0])
		}
		if followee, ok := this.users[userId]; ok {
			if _, ok := followee[this.tweet[i][1]]; ok {
				res = append(res, this.tweet[i][0])
			}
		}
		if len(res) >= 10 {
			break
		}
	}
	return res
}

/** Follower follows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Follow(followerId int, followeeId int) {
	if followeeId == followerId {
		return
	}
	if _, ok := this.users[followerId]; !ok {
		this.users[followerId] = make(map[int]struct{})
	}
	this.users[followerId][followeeId] = struct{}{}
}

/** Follower unfollows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Unfollow(followerId int, followeeId int) {
	if _, ok := this.users[followerId]; ok {
		delete(this.users[followerId], followeeId)
	}
}
