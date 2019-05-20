#
# pre-requisite:
#
#   export TEST_BASIC_PASSWORD=<put password here>
#

run-baseline1:
	TEST_WORKER=1 go run baseline.go | tee 1.run

run-baseline2:
	TEST_WORKER=2 go run baseline.go | tee 2.run

run-baseline4:
	TEST_WORKER=4 go run baseline.go | tee 4.run

run-baseline8:
	TEST_WORKER=8 go run baseline.go | tee 8.run

run-baseline16:
	TEST_WORKER=16 go run baseline.go | tee 16.run
