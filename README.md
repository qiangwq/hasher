# hasher
A hashing web API

To run the service:
1. clone qiangwq/hasher repo
2. cd hasher, run: go build
3. run service with: go run .
4. to test the service: at another console, cd hasher/client, run: go run client.go 
5. to shutdown the service gracefully, send shutdown request: curl -X POST --data "password=angryMonkey" http://localhost:8080/shutdown
6. below is the test result from step 4:
    each tester is a goruntine executing steps: 1, post password, query hash by hashID every 2 seconds until hashValue is fetched.
    then the stats is printed and tester/goruntine stops.
    there are 2 tester/goruntines executing at the same time.
    then shutdown is sent to the service at the end, you can see all requests being processed, and port is closed.

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

