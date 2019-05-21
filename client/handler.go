package client

import (
	"context"

	//"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/apache/thrift/lib/go/thrift"
	"github.com/melkor/test-thrift/payment"
)

func (c *Client) Authorize(ctx context.Context) (bool, error) {

	defer c.Transport.Close()
	if err := c.Transport.Open(); err != nil {
		return false, err
	}
	iprot := c.ProtocolFactory.GetProtocol(c.Transport)
	oprot := c.ProtocolFactory.GetProtocol(c.Transport)

	return payment.NewPaymentClient(thrift.NewTStandardClient(iprot, oprot)).Authorize(ctx, &c.Card)
}
