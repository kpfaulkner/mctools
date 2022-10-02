// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package mcra

import "github.com/kpfaulkner/mctools/anvil/item"

// Query defines a generic search query.
type Query interface {
	// IsTarget returns true if the given block should be
	// included in the result set.
	IsTarget(Block) bool
}

// RadiusQuery finds all listed items within a given radius from a
// specific world location.
type RadiusQuery struct {
	items  []item.Id
	origin Location
	radius uint
}

// NewRadiusQuery creates a new distance query for the given values,
// Origin is the point from which
func NewRadiusQuery(origin Location, radius uint, items ...item.Id) Query {
	return &RadiusQuery{
		origin: origin,
		radius: radius,
		items:  items,
	}
}

// IsTarget returns true if the given block should be
// included in the result set.
func (q *RadiusQuery) IsTarget(b Block) bool {
	for _, v := range q.items {
		if v == b.Id {
			dist := q.origin.DistanceTo(b.Location)
			return dist <= q.radius
		}
	}

	return false
}

// InclusionQuery defines an inclusion search.
// This means that the result set will contain only item types included
// in this query.
type InclusionQuery []item.Id

// NewInclusionQuery creates a new inclusion query.
func NewInclusionQuery(items ...item.Id) Query {
	return InclusionQuery(items)
}

// IsTarget returns true if the given block should be
// included in the result set.
func (q InclusionQuery) IsTarget(b Block) bool {
	for _, v := range q {
		if v == b.Id {
			return true
		}
	}

	return false
}

// ExclusionQuery defines an exclusion search.
// This means that the result set will contain only item types NOT included
// in this query.
type ExclusionQuery []item.Id

// NewExclusionQuery creates a new exclusion query.
func NewExclusionQuery(items ...item.Id) Query {
	return ExclusionQuery(items)
}

// IsTarget returns true if the given block should be
// included in the result set.
func (q ExclusionQuery) IsTarget(b Block) bool {
	for _, v := range q {
		if v == b.Id {
			return false
		}
	}

	return true
}
