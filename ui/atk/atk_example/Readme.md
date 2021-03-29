# ATK example

- https://github.com/visualfc/atk
- https://github.com/visualfc/atk_demo



## Run 

	go run ./example/main


## Windows x_64

2020.03.24

компилируется и для x86 и для x64

при запуске ругается - Failed to load tcl86t.dll

попытка подсунуть разные версии либ приводят к - Tcl_Init failed

2020.03.25 - установленна свежая версия ActiveTcl - ошибки те-же - Tcl_Init failed

2020.07.20 - библиотеки от проекта dtk - tcl86t.dll, tk86t.dll, а также необходимы скрипты инициализации TCL - папка library. 
Все помещено в архив for_win.zip. Для работы - извлечь в папку с бинарником.



set GOARCH=386
set GOARCH=amd64


