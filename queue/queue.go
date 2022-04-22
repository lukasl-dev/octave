package queue

import (
	"github.com/lukasl-dev/waterlink/v2/track"
	"sync"
)

// Queue is a thread-safe queue for audio tracks. It is used to store tracks
// that are ready to be played.
type Queue struct {
	// mu is the mutex protecting the following fields from concurrent access.
	mu sync.RWMutex

	// tracks is a slice of enqueued tracks.
	tracks []track.Track
}

// NewQueue returns a new Queue with the given tracks as the initial contents.
func NewQueue(tr ...track.Track) *Queue {
	return &Queue{tracks: tr}
}

// Range iterates over the tracks in the queue and calls fn() for each track.
func (q *Queue) Range(fn func(track.Track)) {
	q.mu.RLock()
	for _, t := range q.tracks {
		fn(t)
	}
	q.mu.RUnlock()
}

// Len returns the number of tracks in the queue.
func (q *Queue) Len() int {
	q.mu.RLock()
	defer q.mu.RUnlock()
	return len(q.tracks)
}

// Push adds the given tracks to the queue.
func (q *Queue) Push(tr ...track.Track) {
	q.mu.Lock()
	q.tracks = append(q.tracks, tr...)
	q.mu.Unlock()
}

// Remove deletes the next track from the queue and returns true if one was
// removed.
func (q *Queue) Remove() bool {
	if q.Len() == 0 {
		return false
	}
	q.mu.Lock()
	q.tracks = q.tracks[1:]
	q.mu.Unlock()
	return true
}

// Peek returns the next track in the queue without removing it. If the queue
// is empty, nil is returned.
func (q *Queue) Peek() *track.Track {
	if q.Len() == 0 {
		return nil
	}
	q.mu.RLock()
	defer q.mu.RUnlock()
	return &q.tracks[0]
}

// Pop returns the next track in the queue and removes it. If the queue is empty,
// nil is returned.
func (q *Queue) Pop() *track.Track {
	if q.Len() == 0 {
		return nil
	}
	q.mu.Lock()
	tr := q.tracks[0]
	q.tracks = q.tracks[1:]
	q.mu.Unlock()
	return &tr
}
