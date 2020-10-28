package main

import "testProject/learning/example/api_jwt_mongo/worker/queue2queue/runner"

func main() {
	m := runner.NewManager()
	m.Monitor()
	m.Listening()
}
