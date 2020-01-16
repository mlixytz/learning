package leetcode355

import (
	"testing"
)

func TestSolution(t *testing.T) {
	twitter := Constructor()
	twitter.PostTweet(1, 5)
	if !isEqual(twitter.GetNewsFeed(1), []int{5}) {
		t.Error("test error")
	}
	twitter.Follow(1, 2)
	twitter.PostTweet(2, 6)
	if !isEqual(twitter.GetNewsFeed(1), []int{6, 5}) {
		t.Error("test error")
	}
	twitter.Unfollow(1, 2)
	if !isEqual(twitter.GetNewsFeed(1), []int{5}) {
		t.Error("test error")
	}

	twitter2 := Constructor()
	twitter2.PostTweet(1, 5)
	twitter2.Follow(1, 1)
	if !isEqual(twitter2.GetNewsFeed(1), []int{5}) {
		t.Error("test error")
	}

	twitter3 := Constructor()
	twitter3.PostTweet(1, 5)
	twitter3.PostTweet(1, 3)
	twitter3.PostTweet(1, 101)
	twitter3.PostTweet(1, 13)
	twitter3.PostTweet(1, 10)
	twitter3.PostTweet(1, 2)
	twitter3.PostTweet(1, 94)
	twitter3.PostTweet(1, 505)
	twitter3.PostTweet(1, 333)
	if !isEqual(twitter3.GetNewsFeed(1), []int{333, 505, 94, 2, 10, 13, 101, 3, 5}) {
		t.Error("test error")
	}
}

func isEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
