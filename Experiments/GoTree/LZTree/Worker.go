package LZTree

/*



 */

// Worker
type Worker struct {
	tree *Tree
}

func NewWorker() *Worker {
	worker := &Worker{
		tree: NewTree(),
	}
	return worker
}

func (w *Worker) Start(data []byte) {
	for _, b := range data {
		// log.Println("input byte:", i, b)

		w.tree.Append(b)
	}
	w.tree.Finish()
	w.tree.PrinfInfo()
}
