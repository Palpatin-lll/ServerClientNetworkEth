# gRPC
Делать сервер можно как из под Windows, так и из под Linux. Далее будет описаны все два метода. Если у вас не установлен Linux на виртуальную машину, или как основаня ОС, или как вторая ОС, можно воспользоваться [WSL (Windows Subsystem for Linux)](https://docs.microsoft.com/ru-ru/windows/wsl/install), как её установить описано в [WSL.md](https://github.com/Net2Fox/PM.12/blob/master/WSL.md)

Вся разработка проводилась на Ubuntu 22.04 с версией Go 1.18.1 и Protobuf Compiler 3.12.4

## Подготовка Windows для разработки
1. Устанавливаем [Go](https://go.dev/dl/)
2. После установки вводим в командной строке 
    ```
    go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
    ```
3. Windows готова

## Подготовка Ubuntu 22.04 для разработки
1. Выполняем следующие команды:
    ```
    $ sudo apt update
    $ sudo apt full-upgrade
    ```
2. Теперь переходим к установке Go и Protobuf Compiler. Вводим следующие команды:
    ```
    $ sudo apt install golang
    $ sudo apt install -y protobuf-compiler
    $ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
    $ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
    ```
3. Далее надо прописать путь к Go в PATH. Для этого необходимо открыть файл .bashrc:
    ```
    $ nano ~/.bashrc
    ```
    И вставить в самый конец (по файлу можно перемещаться стрелочками) следующую строку:
    ~~~
    export PATH="$PATH:$(go env GOPATH)/bin"
    ~~~
    После чего нажать Ctrl+O -> Enter -> Ctrl+X и ввести 
    ```
    source ~/.bashrc
    ```
4. Linux настроен и готов.

## gRPC-сервер
### Сборка сервера в Windows
1. Заходим в папку gRPC\server в командной строке
    ```
    (Буква диска, на котором у вас находится проект, у меня это D:)
    D:
    cd (папка, где находится проект)\gRPC\server
    ```
2. После чего вводим
    ```
    go build -v -o bin\server.exe
    ```
3. Просто запускаем EXE-файл

### Сборка сервера в Linux
1. Заходим в папку gRPC/server через терминал Linux
    ```
    $ cd (папка, где находится проект)/gRPC/server
    ```
2. В терминале вводим
    ```
    $ go build -v -o bin/server
    ```
3. Переходим в папку bin и запускаем клиент
    ```
    $ ./server
    ```

## gRPC-клиент
### Сборка клиента в Windows
1. Заходим в папку client в cmd, для этого пишем:
    ```
    (Буква диска, на котором у вас находится проект, у меня это D:)
    D:
    cd (папка, где находится проект)\gRPC\client
    ```
2. После чего в cmd пишем:
    ```
    go build -v -o bin\client.exe
    ```
3. Просто запускаем EXE-файл

    P.S. Перед запуском клиента Вы должны запустить сервер

### Сборка клиента в Linux
1. Заходим в папку gRPC/client через терминал Linux
    ```
    $ cd (папка, где находится проект)/gRPC/client
    ```
2. В терминале вводим
    ```
    $ go build -v -o bin/client
    ```
3. Переходим в папку bin и запускаем клиент
    ```
    $ ./client
    ```
    P.S. Перед запуском клиента Вы должны запустить сервер

## Как это написать самому?
А фиг знает, не пофиг ли вам?