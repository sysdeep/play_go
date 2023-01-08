package LZTree

import (
	"fmt"
	"io/ioutil"
	"log"
)

/*

2021.04.29

вариант алгоритма LZ с сохранением данных в дереве и отдельной последовательностью
дерево растёт на произвольную глубину


				  aa
				/ |  \
			  10 22 02
				/


результат 1 - успешно упаковали данные в дерево и успешно их распаковали
результат 2 - на разных наборах входных данных - разные результаты... картинка жмётся очень плохо, а база данных - очень хорошо...

========================================================
total bytes count:  726
nodes count:  113
sequence count:  114
========================================================


Проблемы:

	- надо как то хранить дерево



2021.04.30 - по всем призракам - это алгоритм LZW
*/

func ExampleLZTree1() {
	// var data = []byte{33, 1, 2, 3, 12, 45, 33, 34, 33, 33, 33}
	var data = []byte{
		33, 1, 2, 3, 12, 45, 33, 34, 33, 33, 33, 33, 1, 2, 3, 12, 45, 33, 34, 33, 33, 33, 33, 1, 2, 3, 12, 45, 33, 34, 33, 33, 33,
		33, 1, 2, 3, 12, 45, 33, 34, 33, 33, 33, 33, 1, 2, 3, 12, 45, 33, 34, 33, 33, 33, 33, 1, 2, 3, 12, 45, 33, 34, 33, 33, 33,
		33, 1, 2, 3, 12, 45, 33, 34, 33, 33, 33, 33, 1, 2, 3, 12, 45, 33, 34, 33, 33, 33, 33, 1, 2, 3, 12, 45, 33, 34, 33, 33, 33,
		33, 1, 2, 3, 12, 45, 33, 34, 33, 33, 33, 33, 1, 2, 3, 12, 45, 33, 34, 33, 33, 33, 33, 1, 2, 3, 12, 45, 33, 34, 33, 33, 33,
		33, 1, 2, 3, 12, 45, 33, 34, 33, 33, 33, 33, 1, 2, 3, 12, 45, 33, 34, 33, 33, 33, 33, 1, 2, 3, 12, 45, 33, 34, 33, 33, 33,
		33, 1, 2, 3, 12, 45, 33, 34, 33, 33, 33, 33, 1, 2, 3, 12, 45, 33, 34, 33, 33, 33, 33, 1, 2, 3, 12, 45, 33, 34, 33, 33, 33,
		33, 1, 2, 3, 12, 45, 33, 34, 33, 33, 33, 33, 1, 2, 3, 12, 45, 33, 34, 33, 33, 33, 33, 1, 2, 3, 12, 45, 33, 34, 33, 33, 33,
		33, 1, 2, 3, 12, 45, 33, 34, 33, 33, 33, 33, 1, 2, 3, 12, 45, 33, 34, 33, 33, 33, 33, 1, 2, 3, 12, 45, 33, 34, 33, 33, 33,
		33, 1, 2, 3, 12, 45, 33, 34, 33, 33, 33, 33, 1, 2, 3, 12, 45, 33, 34, 33, 33, 33, 33, 1, 2, 3, 12, 45, 33, 34, 33, 33, 33,
		33, 1, 2, 3, 12, 45, 33, 34, 33, 33, 33, 33, 1, 2, 3, 12, 45, 33, 34, 33, 33, 33, 33, 1, 2, 3, 12, 45, 33, 34, 33, 33, 33,
		33, 1, 2, 3, 12, 45, 33, 34, 33, 33, 33, 33, 1, 2, 3, 12, 45, 33, 34, 33, 33, 33, 33, 1, 2, 3, 12, 45, 33, 34, 33, 33, 33,
		33, 1, 2, 3, 12, 45, 33, 34, 33, 33, 33, 33, 1, 2, 3, 12, 45, 33, 34, 33, 33, 33, 33, 1, 2, 3, 12, 45, 33, 34, 33, 33, 33,
		33, 1, 2, 3, 12, 45, 33, 34, 33, 33, 33, 33, 1, 2, 3, 12, 45, 33, 34, 33, 33, 33, 33, 1, 2, 3, 12, 45, 33, 34, 33, 33, 33,
		33, 1, 2, 3, 12, 45, 33, 34, 33, 33, 33, 33, 1, 2, 3, 12, 45, 33, 34, 33, 33, 33, 33, 1, 2, 3, 12, 45, 33, 34, 33, 33, 33,
		33, 1, 2, 3, 12, 45, 33, 34, 33, 33, 33, 33, 1, 2, 3, 12, 45, 33, 34, 33, 33, 33, 33, 1, 2, 3, 12, 45, 33, 34, 33, 33, 33,
		33, 1, 2, 3, 12, 45, 33, 34, 33, 33, 33, 33, 1, 2, 3, 12, 45, 33, 34, 33, 33, 33, 33, 1, 2, 3, 12, 45, 33, 34, 33, 33, 33,
		33, 1, 2, 3, 12, 45, 33, 34, 33, 33, 33, 33, 1, 2, 3, 12, 45, 33, 34, 33, 33, 33, 33, 1, 2, 3, 12, 45, 33, 34, 33, 33, 33,
		33, 1, 2, 3, 12, 45, 33, 34, 33, 33, 33, 33, 1, 2, 3, 12, 45, 33, 34, 33, 33, 33, 33, 1, 2, 3, 12, 45, 33, 34, 33, 33, 33,
		33, 1, 2, 3, 12, 45, 33, 34, 33, 33, 33, 33, 1, 2, 3, 12, 45, 33, 34, 33, 33, 33, 33, 1, 2, 3, 12, 45, 33, 34, 33, 33, 33,
		33, 1, 2, 3, 12, 45, 33, 34, 33, 33, 33, 33, 1, 2, 3, 12, 45, 33, 34, 33, 33, 33, 33, 1, 2, 3, 12, 45, 33, 34, 33, 33, 33,
		33, 1, 2, 3, 12, 45, 33, 34, 33, 33, 33, 33, 1, 2, 3, 12, 45, 33, 34, 33, 33, 33, 33, 1, 2, 3, 12, 45, 33, 34, 33, 33, 33,
		33, 1, 2, 3, 12, 45, 33, 34, 33, 33, 33, 33, 1, 2, 3, 12, 45, 33, 34, 33, 33, 33, 33, 1, 2, 3, 12, 45, 33, 34, 33, 33, 33,
	}
	// var data = []byte{0xAA, 0xAA, 0xAA, 0xAA, 0xAA}

	worker := NewWorker()
	worker.Start(data)
}

/*
========================================================
main.go
total bytes count:  902
nodes count:  257
sequence count:  258
========================================================

========================================================
/home/nia/21.dcat
total bytes count:  8682496
nodes count:  13406
sequence count:  13407

zip 8682496 -> 30770
========================================================

========================================================
/home/nia/about.png
total bytes count:  1 499 035
nodes count:  519 088
sequence count:  519 089

zip 1 499 035 -> 1 499 151
========================================================
*/
func ExampleLZTree2() {

	data, err := ioutil.ReadFile("./main.go")
	// data, err := ioutil.ReadFile("/home/nia/21.dcat")
	// data, err := ioutil.ReadFile("/home/nia/about.png")
	if err != nil {
		log.Fatal("unable read file")
	}

	// var data = []byte{0xAA, 0xAA, 0xAA, 0xAA, 0xAA}

	worker := NewWorker()
	worker.Start(data)
}

// ExampleLZTree3 -
func ExampleLZTree3() {
	const srcFile string = "./add/CalgaryCorpus/calgarycorpus/bib"
	fmt.Println(srcFile)
	data, err := ioutil.ReadFile(srcFile)

	if err != nil {
		log.Fatal("unable read file")
	}

	// var data = []byte{0xAA, 0xAA, 0xAA, 0xAA, 0xAA}

	worker := NewWorker()
	worker.Start(data)
}

// ExampleLZTree4 -
func ExampleLZTree4() {

	var srcFiles = []string{
		"./add/CalgaryCorpus/calgarycorpus/bib",
		"./add/CalgaryCorpus/calgarycorpus/book1",
		"./add/CalgaryCorpus/calgarycorpus/book2",
		"./add/CalgaryCorpus/calgarycorpus/geo",
		"./add/CalgaryCorpus/calgarycorpus/news",
		"./add/CalgaryCorpus/calgarycorpus/obj1",
		"./add/CalgaryCorpus/calgarycorpus/obj2",
		"./add/CalgaryCorpus/calgarycorpus/paper1",
		"./add/CalgaryCorpus/calgarycorpus/paper2",
		"./add/CalgaryCorpus/calgarycorpus/pic",
		"./add/CalgaryCorpus/calgarycorpus/progc",
		"./add/CalgaryCorpus/calgarycorpus/progl",
		"./add/CalgaryCorpus/calgarycorpus/progp",
		"./add/CalgaryCorpus/calgarycorpus/trans",
	}

	for _, filePath := range srcFiles {
		fmt.Println("**********************************************************")
		fmt.Println(filePath)
		data, err := ioutil.ReadFile(filePath)

		if err != nil {
			log.Print("unable read file")
			continue
		}

		// var data = []byte{0xAA, 0xAA, 0xAA, 0xAA, 0xAA}

		worker := NewWorker()
		worker.Start(data)
		fmt.Println("**********************************************************")
	}
}

/*
========================================================
total bytes count:  3_141_622
nodes count:  469379
sequence count:  469380

result: 938_758
========================================================
*/
// ExampleLZTree5 - все файлы в 1 дерево
func ExampleLZTree5() {

	var srcFiles = []string{
		"./add/CalgaryCorpus/calgarycorpus/bib",
		"./add/CalgaryCorpus/calgarycorpus/book1",
		"./add/CalgaryCorpus/calgarycorpus/book2",
		"./add/CalgaryCorpus/calgarycorpus/geo",
		"./add/CalgaryCorpus/calgarycorpus/news",
		"./add/CalgaryCorpus/calgarycorpus/obj1",
		"./add/CalgaryCorpus/calgarycorpus/obj2",
		"./add/CalgaryCorpus/calgarycorpus/paper1",
		"./add/CalgaryCorpus/calgarycorpus/paper2",
		"./add/CalgaryCorpus/calgarycorpus/pic",
		"./add/CalgaryCorpus/calgarycorpus/progc",
		"./add/CalgaryCorpus/calgarycorpus/progl",
		"./add/CalgaryCorpus/calgarycorpus/progp",
		"./add/CalgaryCorpus/calgarycorpus/trans",
	}

	worker := NewWorker()

	for _, filePath := range srcFiles {
		fmt.Println("**********************************************************")
		fmt.Println(filePath)
		data, err := ioutil.ReadFile(filePath)

		if err != nil {
			log.Print("unable read file")
			continue
		}

		// var data = []byte{0xAA, 0xAA, 0xAA, 0xAA, 0xAA}

		worker.StartPart(data)
		fmt.Println("**********************************************************")
	}

	worker.StopParts()
}
