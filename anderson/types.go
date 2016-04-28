package anderson

type (
	// RVal string
	// Maybe the RVal should be a ssa.Value ?

	Constraint interface {
	}

	MaybePointsTo struct {
		Receiver string
		RVal     ssa.Value
	}

	MaybeContains struct {
		Receiver string
		RVal     ssa.Value
	}

	PointsTo struct {
		Receiver  string
		MaybeVals []ssa.Value
	}

	Contains struct {
		Receiver  string
		MaybeVals []ssa.Value
	}

	PointsToSet map[string]*PointsTo
)

func (pt *PointsTo) Copy() *PointsTo {
	vals := make([]ssa.Value)
	copy(vals, pt.MaybeVals)
	return &PointsTo{
		Receiver:  pt.Receiver,
		MaybeVals: vals,
	}
}

func (pts *PointsToSet) Copy() *PointsToSet {
	result := make(PointsToSet, 0)

	for key, val := range pts {
		result[key] = val.Copy()
	}

	return result
}

func Union(first, second *PointsToSet) PointsToSet {
	result = make(PointsToSet)

	// Need to remove duplicates
	union := func(a, b *PointsTo) {
		if a.Receiver != b.Receiver {
			log.Fatal("Inapproperate receiver types")
		}
		unionSet := &PointsTo{
			Receiver:  a.Receiver,
			MaybeVals: make([]ssa.Value, 0),
		}

		for _, ssaVal := range a.MaybeVals {
			append(unionSet, ssaVal)
		}

		for _, ssaVal := range b.MaybeVals {
			append(unionSet, ssaVal)
		}
		return unionSet
	}

	// TODO implement the union of these two
	return result
}
