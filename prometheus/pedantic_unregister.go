package prometheus

// Unregister implements Registerer.
func (r *pedanticRegistry) Unregister(c Collector) bool {
	ch := make(chan *Desc)
	go func() {
		c.Describe(ch)
		close(ch)
	}()

	r.mtx.Lock()
	defer r.mtx.Unlock()

	var ids []uint64
	for d := range ch {
		ids = append(ids, d.id)
	}

	ok := r.Registerer.Unregister(c)
	if ok {
		for _, id := range ids {
			delete(r.checkedDescIDs, id)
		}
	}
	return ok
}
