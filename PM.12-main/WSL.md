# Установка WSL2 и настройка
У вас обязательно должна быть Windows 10, желательно версии 2004 или более новая (проверить можно через Win+R -> winver), либо Windows 11.
1. Заходите в PowerShell (просто пишете в поиске Windows 'PowerShell' и запускаете от имени администратора) и вводите следующую команду:
    ```
    wsl --install
    ```
    Перезагружаем ПК.
2. По умолчанию ставиться Ubuntu (20.04), однако он нам не нужен. Заходим в Microsoft Store и пишем в поиске 'Ubuntu 22.04 LTS', устанавливаем, запускаем.
3. Производите установку Ubuntu 22.04 LTS, там всё крайне просто и понятно. Если вы даже с этим не можете справиться, то советую задуматься.
4. После успешной установки Ubuntu у нас будет обычный терминал Linux. Выполняем следующие команды:
    ```console
    $ sudo apt update
    $ sudo apt full-upgrade
    ```
5. Теперь переходим к установке Go и Protobuf Compiler. Вводим следующие команды:
    ```
    $ sudo apt install golang
    $ sudo apt install -y protobuf-compiler
    $ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
    $ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
    ```
6. Далее надо прописать путь к Go в PATH. Для этого необходимо открыть файл .bashrc:
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
7. Linux настроен и готов.