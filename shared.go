package app

import "os"

const TransferMoneyTaskQueue = "TRANSFER_MONEY_TASK_QUEUE"

var HostPort = os.Getenv("TEMPORAL_CLUSTER_HOSTPORT")

type TransferDetails struct {
	Amount      float32
	FromAccount string
	ToAccount   string
	ReferenceID string
}
