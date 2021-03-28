package persistance

import (
	"errors"
	"fmt"
	"github.com/mic90/metrics-api/metrics"
	"github.com/mic90/metrics-api/metrics/data"
	"time"
)

var ErrIndexNotFound = errors.New("time index not found")

type Bucket struct {
	descriptor    metrics.Descriptor
	shards        []*Shard
	shardsLookup  map[string][]int
	shardDuration time.Duration
	index         int
}

func NewBucket(desc metrics.Descriptor, dur time.Duration) (*Bucket, error) {
	var (
		m   metrics.Metric
		err error
	)
	if m, err = metrics.FromDescriptor(desc); err != nil {
		return nil, err
	}
	shard := NewShard(m, dur)
	return &Bucket{
		desc,
		[]*Shard{shard},
		map[string][]int{
			timeHash(shard.MinT()): {0},
		},
		dur,
		0,
	}, nil
}

func (b *Bucket) Descriptor() metrics.Descriptor {
	return b.descriptor
}

func (b *Bucket) AddData(dataPoint data.Point) error {
	err := b.shards[b.index].AddData(dataPoint)

	if err == nil {
		return nil
	}

	if !errors.Is(err, ErrShardMaxReached) {
		return err
	}

	var m metrics.Metric

	// we have reached max timestamp for the current shard, make a new one
	if m, err = metrics.FromDescriptor(b.descriptor); err != nil {
		return err
	}
	if err := m.AddData(dataPoint); err != nil {
		return err
	}

	shard := NewShard(m, b.shardDuration)
	hash := timeHash(shard.MinT())

	b.shards = append(b.shards, shard)
	b.index++

	// store shards lookup data
	if indexes, ok := b.shardsLookup[hash]; ok {
		b.shardsLookup[hash] = append(indexes, b.index)
	} else {
		b.shardsLookup[hash] = []int{b.index}
	}

	return nil
}

func (b *Bucket) Data(from, to time.Time) []data.Point {
	var (
		startIndexes []int
		endIndexes   []int
		err          error
		ret          []data.Point
	)

	if startIndexes, err = b.findIndexes(from, b.shardDuration); err != nil {
		return ret
	} else if endIndexes, err = b.findIndexes(to, -b.shardDuration); err != nil {
		return ret
	}
	// indexes in lookup table are stored in ascending order
	// use first one for start, and last one for end indexes range
	ranged := b.shards[startIndexes[0] : endIndexes[len(endIndexes)-1]+1]

	// to finish range retrieval, cut first and last shards to desired time range
	// copy other shards as they are
	for index, shard := range ranged {
		switch {
		case index == 0:
			ret = append(ret, shard.DataFrom(from)...)
		case index == len(ranged)-1:
			ret = append(ret, shard.DataTo(to)...)
		default:
			ret = append(ret, shard.Data()...)
		}
	}

	return ret
}

func (b *Bucket) Size() int {
	return len(b.shards)
}

func (b *Bucket) Hash() string {
	return b.descriptor.Hash()
}

func (b *Bucket) findIndexes(t time.Time, step time.Duration) ([]int, error) {
	hash := timeHash(t)
	lastTime := t

	for {
		if indexes, ok := b.shardsLookup[hash]; !ok {
			// if no shard for this time hash exists, move lookup window by step - left or right
			lastTime = lastTime.Add(step)
			hash = timeHash(lastTime)
		} else {
			return indexes, nil
		}

		if step > 0 && lastTime.After(b.shards[b.index].MaxT()) {
			// we heave reached end of the available shard timestamps
			return []int{}, ErrIndexNotFound
		} else if step < 0 && lastTime.Before(b.shards[0].MinT()) {
			// we moved before available timestamps window
			return []int{}, ErrIndexNotFound
		}
	}
}

func timeHash(t time.Time) string {
	return fmt.Sprintf("%d%d%d%d", t.Day(), t.Month(), t.Year(), t.Hour())
}
