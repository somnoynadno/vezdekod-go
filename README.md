# vezdekod-go

## Задания 10 и 20

Входные данные подаются через файл `input.txt`, либо же входной файл можно изменить в коде:
```go
const FILE = "../input.txt"
```

Сборка и запуск: 
```bash
 $ go build main.go && ./main 
```

P.S. Строки нумеруются с нуля)

## Задание 30

К команде запуска выше добавляются два флага командной строки:

```bash
 $ go build main.go && ./main --help
Usage of ./main:
  -max_queue_size int
        The size of job queue (default 10)
  -max_workers int
        The number of workers to start (default 2)
```
