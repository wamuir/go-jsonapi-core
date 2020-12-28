package core

type Included []Resource

func (i Included) Merge(c Included) Included {

	for _, r := range c {
		i = i.MergeResource(r)
	}

	return i

}

// Merge a resource into the collection, such as for `Included`
func (i Included) MergeResource(r Resource) Included {

	idx, found := i.FindIndex(r)
	if !found {
		i = append(i, r)
		return i
	}

	for key, rel := range r.Relationships {

		if _, ok := i[idx].Relationships[key]; ok {
			// Relationship already exists
			if i[idx].Relationships[key].Data != nil {
				// And data already exists
				continue
			}
		}

		i[idx].Relationships[key] = rel
	}

	return i
}

func (i Included) FindIndex(r Resource) (int, bool) {

	var idx int
	for j := range i {
		if i[j].Type == r.Type && i[j].Identifier == r.Identifier {
			return j, true
		}
	}
	return idx, false
}

func (i Included) Find(r Resource) (*Resource, bool) {

	idx, found := i.FindIndex(r)
	if !found {
		return nil, false
	}

	return &i[idx], true
}
