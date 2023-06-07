# Ethereum

## Развертывание тестовой сети блокчейна Ethereum и создание нескольких кошельков для криптовалюты
### Лучше всего делать это на Linux, я делал всё в Ubuntu Ubuntu 22.04 с версией Go 1.18.1 и Geth 1.10.18-unstable.

### За основу всего ниже бралась [статья с Хабр](https://habr.com/ru/post/654835/).
1. Выполняем следующие команды
    ```
    $ sudo apt update
    $ sudo apt full-upgrade
    ```
2. Устанавливаем Go (если ещё не установлен)
    ```
    $ sudo apt install golang
    ```
3. Добавляем PPA Ethereum
    ```
    $ sudo add-apt-repository -y ppa:ethereum/ethereum
    $ sudo apt-get update
    $ sudo apt-get install ethereum-unstable
    ```
4. Создаём файл genesis.json и вставляем в него это
    ```json
    {
        "config": {
            "chainId": 98760,
            "homesteadBlock": 0,
            "eip150Block": 0,
            "eip155Block": 0,
            "eip158Block": 0,
            "byzantiumBlock": 0,
            "constantinopleBlock": 0,
            "istanbulBlock": 0,
            "petersburgBlock": 0
        },
        "difficulty": "10",
        "gasLimit": "5100000",
        "alloc": {}
    }
    ```
5. Создаём каталог для работы с блокчейном
    ```
    $ mkdir node1
    ```
6. Создаём аккаунт в блокчейне
    ```
    $ geth --datadir ~/node1 account new
    ```
    Команда account new выведет на консоль параметр Public address of the key — так называемый адрес узла. Мы будем указывать его в различных командах. Обязательно сохраните его.

    Также обратите внимание на путь к файлу с секретным ключом Path of the secret key file. Этот файл необходим для выполнения транзакций. Его можно скопировать и хранить в безопасном месте.
7. Запускаем инициализацию узла блокчейна
    ```
    $ geth --datadir ~/node1 init ~/genesis.json
    ```
    Команда выполнит инициализацию и выведет на консоль результаты своей работы.
8. Для запуска узла, вводим в консоли
    ```
    $ geth --datadir ~/node1 --ipcpath "~/node1/geth.ipc" --nodiscover --mine --miner.threads 1 --maxpeers 0 --verbosity 3 --networkid 98760 console
    ```
    При первом запуске узла нужно дождаться завершения процесса генерации DAG.
9. Откройте ещё одну консоль (не закрывая прошлую) и подключитесь к нашему узлу
    ```
    $ geth --datadir ~/node1 --networkid 98760 attach
    ```
    Эта команда откроет консоль geth и подключится к вашему приватному узлу. Для выхода можно нажать Ctrl+D.
10. Узнаем существующие адреса аккаунтов в нашем узле
    ```
    > web3.eth.accounts
    ```
11. Проверим баланс
    ```
    > web3.fromWei(eth.getBalance(eth.coinbase))
    ```
12. Для остановки майнинга в первой консоли нужно ввести
    ```
    > miner.stop()
    ```
    Для запуска майнинга снова нужно ввести
    ```
    > miner.start(1)
    ```
    В скобочках указываем количество потоков процессора, которые будут выделенны под майнинг. Чем больше, тем быстрее происходить процесс.
13. Чтобы полностью остановить узел, необходимо выйти из коносли Geth. Для этого можно нажать Ctrl+D, либо ввести
    ```
    > exit
    ```

## Разворачивание смарт-контракта в тестовой сети
1. Для начала, необходимо установить node и npm
    ```
    $ sudo apt install curl
    $ curl -fsSL https://deb.nodesource.com/setup_18.x | sudo -E bash -
    $ sudo apt-get install -y nodejs
    ```
2. Теперь надо запустить тестовую сеть и начать майнинг
    ```
    $ geth --datadir ~/node1 --ipcpath "~/node1/geth.ipc" --nodiscover --mine --miner.threads 1 --maxpeers 0 --verbosity 3 --networkid 98760 console
    ```
3. Перемещаем файлы HelloSol.sol, deploy_contract_promise.js, call_contract_set.js, call_contract_get.js в home/имя_пользователя
4. Открываем второй терминал, для самого разворачивания смарт-контракта мы будем использовать уже готовый скрипт, deploy_contract_promise.js, чтобы было удобнее. Для запуска скрипта надо написать "node deploy_contract_promise.js Название_вашего_пользователя Название_контракта Пароль_от_кошелька_тестовой_сети ipc"
    ```
    $ node deploy_contract_promise.js net2fox HelloSol ********** ipc
    ```
5. После этого нужно дождаться, пока контракт опубликуется, вам выведеться сообщение по типу этого
    ```
    Contract Deploy script: HelloSol
    Unlocked: true
    Gas estimation: 361588
    transactionHash:        0x7ef97cb7a59b80f37da4876b0c7d9544e4abf98f72d600d88ef65781fab03629
    on transactionHash:     0x7ef97cb7a59b80f37da4876b0c7d9544e4abf98f72d600d88ef65781fab03629
    on confirmation: 0
    contractAddress: 0xB632B0a196ecd77d578B42e09bAA74207B90c2B8
    Contract Deployed
    coinbase: 0x49536b7a06b0fa1a4c41bd11e3ca63ea3ae3389b
    New Contract address: 0xB632B0a196ecd77d578B42e09bAA74207B90c2B8
    ```
6. Смарт-контракт HelloSol успешно опубликовался. В данном контракте есть две функции, set и get. Для их удобного вызова воспользуемся готовыми скриптами, call_contract_set.js и call_contract_get.js
    ```
    $ node call_contract_set.js net2fox HelloSol ********** 111 hello
    $ node call_contract_get.js net2fox HelloSol
    ```
    В результате call_contract_set.js мы присовили в смарт-контракте число 111 и строку "hello", после чего получили эти данные через call_contract_get.js