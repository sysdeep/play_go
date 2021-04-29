package LZTree

/*

2021.04.29

вариант алгоритма LZ с сохранением данных в дереве и отдельной последовательностью



*/

func ExampleLZTree1() {
	var data = []byte{33, 1, 2, 3, 12, 45, 33, 34, 33, 33, 33}
	// var data = []byte{0xAA, 0xAA, 0xAA, 0xAA, 0xAA}

	worker := NewWorker()
	worker.Start(data)
}
