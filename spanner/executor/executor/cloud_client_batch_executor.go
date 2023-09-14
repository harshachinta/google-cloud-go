package executor

import (
	"context"
	"errors"
	"fmt"
	"log"

	"cloud.google.com/go/spanner"
	executorpb "cloud.google.com/go/spanner/executor/proto"
	"google.golang.org/api/option"
	spb "google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type startBatchTxnHandler struct {
	action        *executorpb.StartBatchTransactionAction
	txnContext    context.Context
	flowContext   *executionFlowContext
	outcomeSender *outcomeSender
	options       []option.ClientOption
}

func (h *startBatchTxnHandler) executeAction(ctx context.Context) error {
	/*h.flowContext.mu.Lock()
	defer h.flowContext.mu.Unlock()
	if h.flowContext.isTransactionActive() {
		return errors.New("already in a transaction")
	}

	if h.flowContext.database == "" {
		return fmt.Errorf("database path must be set for this action")
	}

	client, err := spanner.NewClient(h.txnContext, h.flowContext.database, h.options...)
	if err != nil {
		return err
	}
	var txn *spanner.BatchReadOnlyTransaction
	if h.action.GetBatchTxnTime() != nil {
		txn, err = client.BatchReadOnlyTransaction(h.txnContext, spanner.ReadTimestamp(h.action.GetBatchTxnTime()))
		if err != nil {
			return err
		}
	} else if h.action.GetTid() != nil {
		h.action.get
		client.BatchReadOnlyTransactionFromID()
	}
	*/
	return nil
}

type partitionReadActionHandler struct {
	action        *executorpb.GenerateDbPartitionsForReadAction
	flowContext   *executionFlowContext
	outcomeSender *outcomeSender
}

func (h *partitionReadActionHandler) executeAction(ctx context.Context) error {
	metadata := &tableMetadataHelper{}
	metadata.initFromTableMetadatas(h.action.GetTable())

	h.flowContext.mu.Lock()
	defer h.flowContext.mu.Unlock()

	h.flowContext.tableMetadata = metadata
	readAction := h.action.GetRead()

	typeList, err := h.flowContext.tableMetadata.getKeyColumnTypes(readAction.GetTable())
	if err != nil {
		return status.Error(codes.InvalidArgument, fmt.Sprintf("Can't extract types from metadata: %s", err))
	}

	keySet, err := keySetProtoToCloudKeySet(readAction.GetKeys(), typeList)
	if err != nil {
		return status.Error(codes.InvalidArgument, fmt.Sprintf("Can't convert rowSet: %s", err))
	}

	batchtxn, err := h.flowContext.getBatchTransaction()
	if err != nil {
		return fmt.Errorf("can't get batch transaction: %s", err)
	}

	partitionOptions := spanner.PartitionOptions{PartitionBytes: h.action.GetDesiredBytesPerPartition(), MaxPartitions: h.action.GetMaxPartitionCount()}
	var partitions []*spanner.Partition
	if readAction.Index != nil {
		partitions, err = batchtxn.PartitionReadUsingIndexWithOptions(h.flowContext.txnContext, readAction.GetTable(), readAction.GetIndex(), keySet, readAction.GetColumn(), partitionOptions, spanner.ReadOptions{})
	} else {
		partitions, err = batchtxn.PartitionReadWithOptions(h.flowContext.txnContext, readAction.GetTable(), keySet, readAction.GetColumn(), partitionOptions, spanner.ReadOptions{})
	}
	if err != nil {
		return err
	}
	var batchPartitions []*executorpb.BatchPartition
	for _, part := range partitions {
		partitionInstance, _ := part.MarshalBinary()
		batchPartition := &executorpb.BatchPartition{
			Partition: partitionInstance,
			///PartitionToken: part,
			Table: &readAction.Table,
			Index: readAction.Index,
		}
		batchPartitions = append(batchPartitions, batchPartition)
	}
	spannerActionOutcome := &executorpb.SpannerActionOutcome{
		Status:      &spb.Status{Code: int32(codes.OK)},
		DbPartition: batchPartitions,
	}
	err = h.outcomeSender.sendOutcome(spannerActionOutcome)
	if err != nil {
		return h.outcomeSender.finishWithError(err)
	}
	return h.outcomeSender.finishSuccessfully()
}

type partitionQueryActionHandler struct {
	action        *executorpb.GenerateDbPartitionsForQueryAction
	flowContext   *executionFlowContext
	outcomeSender *outcomeSender
}

func (h *partitionQueryActionHandler) executeAction(ctx context.Context) error {
	h.flowContext.mu.Lock()
	defer h.flowContext.mu.Unlock()

	batchTxn, err := h.flowContext.getBatchTransaction()
	if err != nil {
		return fmt.Errorf("can't get batch transaction: %s", err)
	}
	stmt, err := buildQuery(h.action.GetQuery())
	if err != nil {
		return err
	}
	partitionOptions := spanner.PartitionOptions{PartitionBytes: h.action.GetDesiredBytesPerPartition()}
	partitions, err := batchTxn.PartitionQuery(h.flowContext.txnContext, stmt, partitionOptions)
	if err != nil {
		return err
	}
	var batchPartitions []*executorpb.BatchPartition
	for _, partition := range partitions {
		partitionInstance, err := partition.MarshalBinary()
		if err != nil {
			return err
		}
		batchPartition := &executorpb.BatchPartition{
			Partition: partitionInstance,
			///PartitionToken: spanner.Partition{},
		}
		batchPartitions = append(batchPartitions, batchPartition)
	}

	spannerActionOutcome := &executorpb.SpannerActionOutcome{
		Status:      &spb.Status{Code: int32(codes.OK)},
		DbPartition: batchPartitions,
	}
	err = h.outcomeSender.sendOutcome(spannerActionOutcome)
	if err != nil {
		return h.outcomeSender.finishWithError(err)
	}
	return h.outcomeSender.finishSuccessfully()
}

type executePartition struct {
	action        *executorpb.ExecutePartitionAction
	flowContext   *executionFlowContext
	outcomeSender *outcomeSender
}

func (h *executePartition) executeAction(ctx context.Context) error {
	batchTxn, err := h.flowContext.getBatchTransaction()
	partitionBinary := h.action.GetPartition().GetPartition()
	if partitionBinary == nil || len(partitionBinary) == 0 {
		return spanner.ToSpannerError(status.Errorf(codes.InvalidArgument, "Invalid batchPartition %s", h.action))
	}
	if h.action.GetPartition().Table != nil {
		h.outcomeSender.hasReadResult = true
		h.outcomeSender.table = h.action.GetPartition().GetTable()
		if h.action.GetPartition().Index != nil {
			h.outcomeSender.index = h.action.GetPartition().Index
		}
	} else {
		h.outcomeSender.hasQueryResult = true
	}
	var partition *spanner.Partition
	if err = partition.UnmarshalBinary(partitionBinary); err != nil {
		return fmt.Errorf("deserializing Partition failed %v", err)
	}
	iter := batchTxn.Execute(h.flowContext.txnContext, partition)
	defer iter.Stop()
	err = processResults(iter, 0, h.outcomeSender, h.flowContext)
	if err != nil {
		return h.outcomeSender.finishWithError(err)
	}
	return h.outcomeSender.finishSuccessfully()
}

type partitionedUpdate struct {
	action        *executorpb.PartitionedUpdateAction
	flowContext   *executionFlowContext
	outcomeSender *outcomeSender
}

func (h *partitionedUpdate) executeAction(ctx context.Context) error {
	opts := h.action.GetOptions()
	stmt := spanner.Statement{SQL: h.action.GetUpdate().GetSql()}
	count, err := h.flowContext.dbClient.PartitionedUpdateWithOptions(h.flowContext.txnContext, stmt, spanner.QueryOptions{
		Priority:   opts.GetRpcPriority(),
		RequestTag: opts.GetTag(),
	})
	if err != nil {
		return err
	}
	spannerActionOutcome := &executorpb.SpannerActionOutcome{
		Status:          &spb.Status{Code: int32(codes.OK)},
		DmlRowsModified: []int64{count},
	}
	err = h.outcomeSender.sendOutcome(spannerActionOutcome)
	if err != nil {
		return h.outcomeSender.finishWithError(err)
	}
	return h.outcomeSender.finishSuccessfully()
}

type closeBatchTxnHandler struct {
	action        *executorpb.CloseBatchTransactionAction
	flowContext   *executionFlowContext
	outcomeSender *outcomeSender
}

func (h *closeBatchTxnHandler) executeAction(ctx context.Context) error {
	log.Printf("closing batch transaction %v", h.action)
	if h.action.GetCleanup() {
		if h.flowContext.batchTxn == nil {
			return h.outcomeSender.finishWithError(errors.New("not in a batch transaction"))
		}
		h.flowContext.batchTxn.Close()
	}
	return h.outcomeSender.finishSuccessfully()
}
