package main

/*

	323 -> 	41
	728 ->  67
	2627506 	-> 	256			// witout chunks


	dSaize		cSize	chunks		tCount		total	elapsed			ratio
	449			4		112			48			160		4.821251ms		2
	519			8		64			49			113		3.957190695s	4
	2627506		4		656876		256			657132	29.509242452s	3			2.5mb
	2627506		8		328438		???			???		не дождался..	?			2.5mb



*/
import (
	"fmt"
	"gotree/internal/chunk_tree"
	"log"
	"os"
	"time"
)

func exFile(filePath string, chunkSize int) {

	// chunkSize := 4

	fd, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer fd.Close()

	fileInfo, _ := fd.Stat()

	chunks := fileInfo.Size() / int64(chunkSize)
	log.Println("data count: ", fileInfo.Size())
	log.Println("chunks count: ", chunks)
	log.Println("chunk size: ", chunkSize)

	tree := chunk_tree.NewTree()

	t1 := time.Now()
	chunk := make([]byte, chunkSize)
	for {
		n, err := fd.Read(chunk)
		if err != nil {
			log.Println("Error:")
			log.Println(err)
			break
		}

		if n == chunkSize {
			tree.AppendChunk(chunk)
		} else {
			log.Println("read less then 4: ", n)
		}
	}

	elapsed := time.Since(t1)
	log.Println("time elapsed: ", elapsed)

	// log.Println("data count: ", len(PAYLOADEX1))

	// for i := 0; i < (len(PAYLOADEX1) / 4); i++ {
	// 	chunk := PAYLOADEX1[i*4 : i*4+4]
	// 	tree.AppendChunk(chunk)
	// }

	treeCount := tree.GetCount()
	totalCount := treeCount + int(chunks)
	ratio := fileInfo.Size() / int64(totalCount)
	log.Println("tree count: ", treeCount)
	log.Println("total count: ", totalCount)
	log.Println("ration: ", ratio)

	fmt.Println("dSaize\tcSize\tchunks\ttCount\ttotal\telapsed\t\tratio")
	fmt.Printf("%d\t%d\t%d\t%d\t%d\t%s\t%d\n",
		fileInfo.Size(), chunkSize, chunks, treeCount, totalCount, elapsed, ratio)
}
