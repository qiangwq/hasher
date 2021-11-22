# hasher
A hashing web API

To run the service
1. clon
To start the

qiang@qiang-VirtualBox:~/go/src/qiangwq/hasher$ go run .
op# 2
op# 1
op# 3
op# 4
2021/11/21 16:28:01 http: Server closed
exit status 1

curl -X POST --data "password=angryMonkey" http://localhost:8080/shutdown



qiang@qiang-VirtualBox:~/go/src/qiangwq/hasher/client$ go run client.go 
Info: tester: 1, http://localhost:8080/hash returns hashID 1
Info: tester: 0, http://localhost:8080/hash returns hashID 2
Info: tester: 1, hashID 1 not ready yet
Info: tester: 0, hashID 2 not ready yet
Info: tester: 1, hashID 1 not ready yet
Info: tester: 0, hashID 2 not ready yet
Info: tester: 1, hashID 1 not ready yet
Info: tester: 0, hashID 2 not ready yet
Info: tester: 0, hashID: 2, hashValue: 3a81oZNherrMQXNJriBBMRLm-k6JqX6iCp7u5ktV05ohkpkqJ0_BqDa6PCOj_uu9RU1EI2Q86A4qmslPpUyknw==
Info: tester: 1, hashID: 1, hashValue: 3a81oZNherrMQXNJriBBMRLm-k6JqX6iCp7u5ktV05ohkpkqJ0_BqDa6PCOj_uu9RU1EI2Q86A4qmslPpUyknw==
Info: stats: {"total":2,"average":5003}
Info: stats: {"total":2,"average":5003}
Info: tester: 2, http://localhost:8080/hash returns hashID 3
Info: tester: 3, http://localhost:8080/hash returns hashID 4
Info: tester: 2, hashID 3 not ready yet
Info: tester: 3, hashID 4 not ready yet
Info: tester: 2, hashID 3 not ready yet
Info: tester: 3, hashID 4 not ready yet
Info: tester: 3, hashID 4 not ready yet
Info: tester: 2, hashID 3 not ready yet
Info: tester: 3, hashID: 4, hashValue: 3a81oZNherrMQXNJriBBMRLm-k6JqX6iCp7u5ktV05ohkpkqJ0_BqDa6PCOj_uu9RU1EI2Q86A4qmslPpUyknw==
Info: tester: 2, hashID: 3, hashValue: 3a81oZNherrMQXNJriBBMRLm-k6JqX6iCp7u5ktV05ohkpkqJ0_BqDa6PCOj_uu9RU1EI2Q86A4qmslPpUyknw==
Info: stats: {"total":4,"average":5002}
Info: stats: {"total":4,"average":5002}
test failed with, error Post "http://localhost:8080/hash": dial tcp 127.0.0.1:8080: connect: connection refused
test failed with, error Post "http://localhost:8080/hash": dial tcp 127.0.0.1:8080: connect: connection refused
test failed with, error Post "http://localhost:8080/hash": dial tcp 127.0.0.1:8080: connect: connection refused
test failed with, error Post "http://localhost:8080/hash": dial tcp 127.0.0.1:8080: connect: connection refused
test failed with, error Post "http://localhost:8080/hash": dial tcp 127.0.0.1:8080: connect: connection refused
test failed with, error Post "http://localhost:8080/hash": dial tcp 127.0.0.1:8080: connect: connection refused
