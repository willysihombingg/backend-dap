// Package handler
package handler

import (
	"context"

	"gitlab.com/willysihombing/task-c3/internal/appctx"
	"gitlab.com/willysihombing/task-c3/internal/consts"
	uContract "gitlab.com/willysihombing/task-c3/internal/ucase/contract"
	"gitlab.com/willysihombing/task-c3/pkg/awssqs"
)

// SQSConsumerHandler sqs consumer message processor handler
func SQSConsumerHandler(msgHandler uContract.MessageProcessor) awssqs.MessageProcessorFunc {
	return func(decoder *awssqs.MessageDecoder) error {
		return msgHandler.Serve(context.Background(), &appctx.ConsumerData{
			Body:        []byte(*decoder.Body),
			Key:         []byte(*decoder.MessageId),
			ServiceType: consts.ServiceTypeConsumer,
		})
	}
}
