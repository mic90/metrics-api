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
	name          string
	_type         string
	shards        []*Shard
	shardsLookup  map[string][]int
	shardDuration time.Duration
	index         int
}

func NewBucket(name, _type string, dur time.Duration) (*Bucket, error) {
	var (
		m   metrics.Metric
		err error
	)
	if m, err = metrics.FromType(name, _type); err != nil {
		return nil, err
	}
	shard := NewShard(m, dur)
	return &Bucket{
		name,
		_type,
		[]*Shard{shard},
		map[string][]int{
			timeHash(shard.MinT()): {0},
		},
		dur,
		0,
	}, nil
}

func (b *Bucket) AddData(dataPoint data.Point) error {
	err := b.shards[b.index].AddData(dataPoint)

	if err != nil && !errors.Is(err, ErrShardMaxReached) {
		return err
	}

	var m metrics.Metric

	// we have reached max timestamp for the current shard, make a new one
	if m, err = metrics.FromType(b.name, b._type); err != nil {
		return err
	}

	shard := NewShard(m, b.shardDuration)
	if err := shard.AddData(dataPoint); err != nil {
		return err
	}

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

func (b *Bucket) Data(from, to time.Time) ([]data.Point, error) {
	var (
		startIndexes []int
		endIndexes   []int
		err          error
		ret          []data.Point
	)

	if startIndexes, err = b.findIndexes(from); err != nil {
		return ret, err
	} else if endIndexes, err = b.findIndexes(to); err != nil {
		return ret, err
	}
	// indexes in lookup table are stored in ascending order
	// use first one for start, and last one for end indexes range
	ranged := b.shards[startIndexes[0] : endIndexes[len(endIndexes)-1]+1]

	// to finish range retrieval, cut first and last shards to desired time range
	// copy other shards as they are
	for index, shard := range ranged {
		if index == 1 {
			ret = append(ret, shard.DataFrom(from)...)
		} else if index == len(ranged)-1 {
			ret = append(ret, shard.DataTo(to)...)
		}
		ret = append(ret, shard.Data()...)
	}

	return ret, nil
}

func (b *Bucket) Size() int {
	return len(b.shards)
}

func (b *Bucket) findIndexes(t time.Time) ([]int, error) {
	hash := timeHash(t)
	lastTime := t

	for {
		// we heave reached end of the available shard timestamps
		if lastTime.After(b.shards[b.index].MaxT()) {
			break
		}

		if indexes, ok := b.shardsLookup[hash]; !ok {
			// if no shard for this time hash exists, increase lookup by shard duration
			lastTime = lastTime.Add(b.shardDuration)
			hash = timeHash(lastTime)
		} else {
			return indexes, nil
		}
	}

	return []int{}, ErrIndexNotFound
}

func timeHash(t time.Time) string {
	return fmt.Sprintf("%d%d%d%d", t.Day(), t.Month(), t.Year(), t.Hour())
}
