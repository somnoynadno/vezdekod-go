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

## Задание 40

Билдим и запускаем сервер
```bash
 $ cd server
 $ go build main.go && ./main -port 8000
```

Билдим и запускаем тестилку для него
```bash
 $ cd tester
 $ go build main.go && ./main
```

Тестилка проверяет три основных use-кейса и никого не дудосит, потому что зачем?)

*Тесты могли бы быть и лучше, но времени у нас не так много...*

### Запросы к серверу

#### Создание задачи

```bash
# создаем задачу асинхронно на 5 секунд
 $ curl -X POST -H "Content-Type: application/json" --data '{"duration": "5s"}' http://localhost:8000/add/async

# создаем задачу синхронно на 10 секунд (соединение будет открыто на протяжении этого времени)
 $ curl -X POST -H "Content-Type: application/json" --data '{"duration": "10s"}' http://localhost:8000/add/sync

```

P.S. Если не указывать поле name в JSONе, то оно будет случайным

P.P.S. Будет код ответа 400, если тело запроса некорректно

#### Получение очереди задач

```bash
 $ curl http://localhost:8000/schedule
 # получаем JSON вида [{"name":"2","duration":"1s"}]
 # (или null, если очередь пуста)
```

#### Получение временной оценки

```bash
 $ curl http://localhost:8000/time
 # получаем текстовый ответ вида 40s
```

