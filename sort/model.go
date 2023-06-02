package sort

type Vendor struct {
	ID       string
	Distance float32
}

type VendorList []Vendor

func (v VendorList) Len() int {
	return len(v)
}

func (v VendorList) Less(i, j int) bool {
	return v[i].Distance < v[j].Distance
}

func (v VendorList) Swap(i, j int) {
	v[i], v[j] = v[j], v[i]
}
