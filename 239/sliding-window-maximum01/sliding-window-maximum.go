package sliding_window_maximum01

func maxSlidingWindow(nums []int, k int) []int {
	res := []int{}
	mq := NewMQ()
	for i := 0; i < len(nums); i++ {
		if i < k-1 {
			mq.Push(nums[i])
			continue
		}

		mq.Push(nums[i])
		res = append(res, mq.Max())
		mq.Pop(nums[i-k+1])
	}

	return res
}

type monotonicQueue struct {
	queue []int
}

func NewMQ() *monotonicQueue {
	return &monotonicQueue{
		queue: make([]int, 0),
	}
}

func (mq *monotonicQueue) Push(n int) {
	for len(mq.queue) > 0 && mq.queue[len(mq.queue)-1] < n {
		mq.queue = mq.queue[:len(mq.queue)-1]
	}

	mq.queue = append(mq.queue, n)
}

func (mq *monotonicQueue) Pop(n int) {
	if len(mq.queue) > 0 && mq.queue[0] == n {
		mq.queue = mq.queue[1:]
	}
}

func (mq *monotonicQueue) Max() int {
	return mq.queue[0]
}
