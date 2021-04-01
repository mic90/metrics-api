package persistance

import (
	"errors"
	"github.com/mic90/metrics-api/metrics"
	"github.com/mic90/metrics-api/metrics/data"
	"sort"
	"time"
)

var ErrNoIndexFound = errors.New("no index found")

// Bucket stores single metric data, sharded in smaller chunks
type Bucket struct {
	descriptor    metrics.Descriptor
	shards        []*Shard
	shardDuration time.Duration
	index         int
}

// NewBucket creates new bucket for given metric
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
		descriptor:    desc,
		shards:        []*Shard{shard},
		shardDuration: dur,
	}, nil
}

// Descriptor returns metric descriptor for this bucket
func (b *Bucket) Descriptor() metrics.Descriptor {
	return b.descriptor
}

// AddData adds new data point to the bucket.
// If given data point exceeds current chunk time range
// new chunk will be created internally
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

	b.shards = append(b.shards, shard)
	b.index++

	return nil
}

// Data returns data merged from all chunks based on provided time range
func (b *Bucket) Data(from, to time.Time) []data.Point {
	var (
		fromIndex int
		toIndex   int
		err       error
		ret       []data.Point
	)

	if fromIndex, toIndex, err = b.findIndexRange(from, to); err != nil {
		// no data was found that matches given time range
		return []data.Point{}
	}
	// indexes in lookup table are stored in ascending order
	// use first one for start, and last one for end indexes range
	ranged := b.shards[fromIndex : toIndex+1]

	// if all data is contained in one shard
	if len(ranged) == 1 {
		return ranged[0].DataRange(from, to)
	}

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

// Size returns number of data chunks
func (b *Bucket) Size() int {
	return len(b.shards)
}

// Hash returns underlying metric unique hash
func (b *Bucket) Hash() string {
	return b.descriptor.Hash()
}

func (b *Bucket) findIndex(t time.Time) (int, error) {
	index := sort.Search(len(b.shards), func(i int) bool {
		return b.shards[i].EndT().Equal(t) || b.shards[i].EndT().After(t)
	})

	// no index was found, data probably ends before time t
	if index == len(b.shards) {
		return -1, ErrNoIndexFound
	}

	return index, nil
}

func (b *Bucket) findLast(t time.Time) int {
	var (
		index int
		err   error
	)

	// no index was found - data probably ends before time t, return last available shard index
	if index, err = b.findIndex(t); err != nil {
		return len(b.shards) - 1
	}

	// make sure the shard fits given time range or is before
	for {
		if (b.shards[index].MinT().Before(t) || b.shards[index].MinT().Equal(t)) &&
			(b.shards[index].EndT().After(t) || b.shards[index].EndT().Equal(t)) {
			return index
		}
		index--
	}
}

func (b *Bucket) findIndexRange(from, to time.Time) (int, int, error) {
	var (
		fromIndex int
		toIndex   int
		err       error
	)

	if fromIndex, err = b.findIndex(from); err != nil {
		return -1, -1, err
	}

	// when there is already a start index do not report errors
	// in case we missed the time range, just return last available index
	toIndex = b.findLast(to)

	return fromIndex, toIndex, nil
}
