# TCalendar

Простой консольный календать.

## Установка
```bash
go install ./cmd/tcalendar/
```


## Запуск в отдельном окне

### ST

```
       -c class
              defines the window class (default $TERM).

        
```


```bash
# попытка запустить с отдельным классом приводит к падению приложения
st -class kitty-float ./tcalendar


# попытка запустить с отдельным id приводит к падению приложения
st -windowid term_float ./tcalendar 

# ok
st -t tcalendar ./tcalendar 
```


### XTerm

```bash
# привело к смене BG и карокозябрам для Русского
xterm -class kitty_float ./tcalendar

# ok
xterm -title tcalendar ./tcalendar

# размер шрифта
xterm -title tcalendar -fa "Ubuntu Mono:size=14" ./tcalendar

# размер
xterm -title tcalendar -fa "Ubuntu Mono:size=14" -geometry 37x16 ./tcalendar

# размер и позиция справа внизу
xterm -title tcalendar -fa "Ubuntu Mono:size=14" -geometry 37x16-4-36 ./tcalendar
```
