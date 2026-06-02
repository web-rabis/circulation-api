package model

func NewListFromProto[P any, D any](items []P, fn func(P) D) []D {
	if items == nil {
		return nil
	}
	result := make([]D, len(items))
	for i, item := range items {
		result[i] = fn(item)
	}
	return result
}
