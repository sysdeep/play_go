# Calgary corpus

- https://ru.wikipedia.org/wiki/Calgary_Corpus
- http://www.data-compression.info/Corpora/CalgaryCorpus/

Calgary corpus - набор текстовых и двоичных файлов, часто использовавшийся в качестве стандартного теста алгоритмов сжатия данных и сравнения их эффективности. Набор был собран в Университете Калгари в 1987 году и широко применялся в 1990-х. В 1997 был предложен новый тестовый набор Canterbury corpus, в котором были учтены некоторые замечания к репрезентативности корпуса Калгари.

## Состав корпуса

В наиболее часто используемой форме корпус Калгари состоит из 14 файлов общим объёмом 3141622 байт:

Размер (байт)	Имя файла	Описание
111,261	        BIB	        Текст ASCII в формате утилиты UNIX "refer" с 725 библиографическими записями.
768,771	        BOOK1	    Неформатированный текст ASCII новеллы Томаса Харди Far from the Madding Crowd.
610,856	        BOOK2	    Текст ASCII в формате "troff" – Ian H. Witten: Principles of Computer Speech.
102,400	        GEO	        Сейсмические данные в виде 32 битных чисел с плавающей запятой в формате IBM.
377,109	        NEWS	    Текст ASCII – набор сообщений из групп USENET.
21,504	        OBJ1	    Исполняемый файл для VAX, полученный компиляцией PROGP.
246,814	        OBJ2	    Исполняемый файл для Macintosh, программа "Knowledge Support System".
53,161	        PAPER1	    Статья в формате "troff" – Witten, Neal, Cleary: Arithmetic Coding for Data Compression.
82,199	        PAPER2	    Статья в формате "troff" – Witten: Computer (in)security.
513,216	        PIC	        Изображение размером 1728 x 2376: текст на французском и линейные диаграммы.
39,611	        PROGC	    Исходный код на языке C – программа UNIX compress v4.0.
71,646	        PROGL	    Исходный код на языке Lisp – системное приложение.
49,379	        PROGP	    Исходный код на языке Pascal – программа для оценки сжатия PPM.
93,695	        TRANS	    Текст ASCII и управляющие последовательности - запись терминальной сессии.

Реже используется набор из 18 файлов, в который дополнительно включены 4 текстовых файла в формате "troff" - PAPER3-PAPER6.

## Тестирование

Корпус Calgary часто использовался для сравнения эффективности сжатия в 1990-е годы. Результаты часто указывались в виде коэффициента бит на байт (среднее количество бит в сжатом файле, требуемое для кодирования 1 байта исходного файла) для каждого файла из набора, затем они усреднялись. Затем чаще стали указывать суммарный размер всех сжатых файлов.

Некоторые архиваторы допускали более эффективное сжатие при одновременной обработке всего корпуса (например, при их помещении в несжатый контейнер тира tar), за счет использования взаимной информации. Другие архиваторы, наоборот, хуже сжимали такой вариант из-за медленной реакции компрессора на изменение характеристик данных. Одновременное сжатие всего корпуса использовалось Matt Mahoney в его книге Data Compression Explained.

В таблице указаны размеры сжатого корпуса для нескольких популярных архиваторов.

Архиватор	    Опции	        Сжатие 14 отдельных файлов	    Объединенный tar архив
Без сжатия		                3,141,622	                    3,152,896
compress		                1,272,772	                    1,319,521
Info-ZIP 2.32	-9	            1,020,781	                    1,023,042
gzip 1.3.5	    -9	            1,017,624	                    1,022,810
bzip2 1.0.3	    -9	            828,347	                        860,097
7-zip 9.12b		                848,687	                        824,573
ppmd Jr1	    -m256 -o16	    740,737	                        754,243
ppmonstr J		                675,485	                        669,497


## Ссылки на наборы




[The Large Calgary Corpus](http://www.data-compression.info/files/corpora/largecalgarycorpus.zip) - ZIP-file with: bib, book1, book2, geo, news, obj1, obj2, paper1, paper2, paper3, paper4, paper5, paper6, pic, progc, progl, progp and trans
 

[The Standard Calgary Corpus](http://www.data-compression.info/files/corpora/calgarycorpus.zip) - ZIP-file with: bib, book1, book2, geo, news, obj1, obj2, paper1, paper2, pic, progc, progl, progp and trans