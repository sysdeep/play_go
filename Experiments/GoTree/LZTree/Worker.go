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

	//--- pack
	for _, b := range data {
		// log.Println("input byte:", i, b)

		w.tree.Append(b)
	}
	w.tree.Finish()
	w.tree.PrinfInfo()

	// w.tree.Unpack()

	// //--- pack data
	// nodes := w.tree.GetNodes()
	// sequence := w.tree.GetSequence()

	// // fmt.Println(sequence)

	// //--- unpack
	// ut := NewTree()
	// ut.SetSequence(sequence)
	// ut.SetNodes(nodes)

}

// StartData - добавление в дерево несколько частей, в конце - StopParts
func (w *Worker) StartPart(data []byte) {

	//--- pack
	for _, b := range data {
		// log.Println("input byte:", i, b)

		w.tree.Append(b)
	}
}

// StopParts - отчёт о добавленных частях - после StartPart
func (w *Worker) StopParts() {

	w.tree.Finish()
	w.tree.PrinfInfo()

}
