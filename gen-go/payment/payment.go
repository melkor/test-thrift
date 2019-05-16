// Autogenerated by Thrift Compiler (0.13.0)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package payment

import (
	"bytes"
	"context"
	"reflect"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = context.Background
var _ = reflect.DeepEqual
var _ = bytes.Equal

// The first thing to know about are types. The available types in Thrift are:
// 
//  bool        Boolean, one byte
//  i8 (byte)   Signed 8-bit integer
//  i16         Signed 16-bit integer
//  i32         Signed 32-bit integer
//  i64         Signed 64-bit integer
//  double      64-bit floating point value
//  string      String
//  binary      Blob (byte array)
//  map<t1,t2>  Map from one type to another
//  list<t1>    Ordered list of one type
//  set<t1>     Set of unique elements of one type
// 
// Did you also notice that Thrift supports C style comments?
// 
// Attributes:
//  - Number
//  - Cryptogram
//  - Holder
//  - Date
type Credit struct {
  Number int32 `thrift:"number,1" db:"number" json:"number"`
  Cryptogram int32 `thrift:"cryptogram,2" db:"cryptogram" json:"cryptogram"`
  Holder string `thrift:"holder,3" db:"holder" json:"holder"`
  Date int64 `thrift:"date,4" db:"date" json:"date"`
}

func NewCredit() *Credit {
  return &Credit{}
}


func (p *Credit) GetNumber() int32 {
  return p.Number
}

func (p *Credit) GetCryptogram() int32 {
  return p.Cryptogram
}

func (p *Credit) GetHolder() string {
  return p.Holder
}

func (p *Credit) GetDate() int64 {
  return p.Date
}
func (p *Credit) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if fieldTypeId == thrift.I32 {
        if err := p.ReadField1(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    case 2:
      if fieldTypeId == thrift.I32 {
        if err := p.ReadField2(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    case 3:
      if fieldTypeId == thrift.STRING {
        if err := p.ReadField3(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    case 4:
      if fieldTypeId == thrift.I64 {
        if err := p.ReadField4(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    default:
      if err := iprot.Skip(fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *Credit)  ReadField1(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadI32(); err != nil {
  return thrift.PrependError("error reading field 1: ", err)
} else {
  p.Number = v
}
  return nil
}

func (p *Credit)  ReadField2(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadI32(); err != nil {
  return thrift.PrependError("error reading field 2: ", err)
} else {
  p.Cryptogram = v
}
  return nil
}

func (p *Credit)  ReadField3(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 3: ", err)
} else {
  p.Holder = v
}
  return nil
}

func (p *Credit)  ReadField4(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadI64(); err != nil {
  return thrift.PrependError("error reading field 4: ", err)
} else {
  p.Date = v
}
  return nil
}

func (p *Credit) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("Credit"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField1(oprot); err != nil { return err }
    if err := p.writeField2(oprot); err != nil { return err }
    if err := p.writeField3(oprot); err != nil { return err }
    if err := p.writeField4(oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *Credit) writeField1(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("number", thrift.I32, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:number: ", p), err) }
  if err := oprot.WriteI32(int32(p.Number)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.number (1) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:number: ", p), err) }
  return err
}

func (p *Credit) writeField2(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("cryptogram", thrift.I32, 2); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:cryptogram: ", p), err) }
  if err := oprot.WriteI32(int32(p.Cryptogram)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.cryptogram (2) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 2:cryptogram: ", p), err) }
  return err
}

func (p *Credit) writeField3(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("holder", thrift.STRING, 3); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:holder: ", p), err) }
  if err := oprot.WriteString(string(p.Holder)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.holder (3) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 3:holder: ", p), err) }
  return err
}

func (p *Credit) writeField4(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("date", thrift.I64, 4); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 4:date: ", p), err) }
  if err := oprot.WriteI64(int64(p.Date)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.date (4) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 4:date: ", p), err) }
  return err
}

func (p *Credit) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("Credit(%+v)", *p)
}

type Payment interface {
  Authorize(ctx context.Context) (r bool, err error)
  // Parameters:
  //  - Amount
  Pay(ctx context.Context, amount int32) (r bool, err error)
}

type PaymentClient struct {
  c thrift.TClient
}

func NewPaymentClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *PaymentClient {
  return &PaymentClient{
    c: thrift.NewTStandardClient(f.GetProtocol(t), f.GetProtocol(t)),
  }
}

func NewPaymentClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *PaymentClient {
  return &PaymentClient{
    c: thrift.NewTStandardClient(iprot, oprot),
  }
}

func NewPaymentClient(c thrift.TClient) *PaymentClient {
  return &PaymentClient{
    c: c,
  }
}

func (p *PaymentClient) Client_() thrift.TClient {
  return p.c
}
func (p *PaymentClient) Authorize(ctx context.Context) (r bool, err error) {
  var _args0 PaymentAuthorizeArgs
  var _result1 PaymentAuthorizeResult
  if err = p.Client_().Call(ctx, "Authorize", &_args0, &_result1); err != nil {
    return
  }
  return _result1.GetSuccess(), nil
}

// Parameters:
//  - Amount
func (p *PaymentClient) Pay(ctx context.Context, amount int32) (r bool, err error) {
  var _args2 PaymentPayArgs
  _args2.Amount = amount
  var _result3 PaymentPayResult
  if err = p.Client_().Call(ctx, "Pay", &_args2, &_result3); err != nil {
    return
  }
  return _result3.GetSuccess(), nil
}

type PaymentProcessor struct {
  processorMap map[string]thrift.TProcessorFunction
  handler Payment
}

func (p *PaymentProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
  p.processorMap[key] = processor
}

func (p *PaymentProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
  processor, ok = p.processorMap[key]
  return processor, ok
}

func (p *PaymentProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
  return p.processorMap
}

func NewPaymentProcessor(handler Payment) *PaymentProcessor {

  self4 := &PaymentProcessor{handler:handler, processorMap:make(map[string]thrift.TProcessorFunction)}
  self4.processorMap["Authorize"] = &paymentProcessorAuthorize{handler:handler}
  self4.processorMap["Pay"] = &paymentProcessorPay{handler:handler}
return self4
}

func (p *PaymentProcessor) Process(ctx context.Context, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
  name, _, seqId, err := iprot.ReadMessageBegin()
  if err != nil { return false, err }
  if processor, ok := p.GetProcessorFunction(name); ok {
    return processor.Process(ctx, seqId, iprot, oprot)
  }
  iprot.Skip(thrift.STRUCT)
  iprot.ReadMessageEnd()
  x5 := thrift.NewTApplicationException(thrift.UNKNOWN_METHOD, "Unknown function " + name)
  oprot.WriteMessageBegin(name, thrift.EXCEPTION, seqId)
  x5.Write(oprot)
  oprot.WriteMessageEnd()
  oprot.Flush(ctx)
  return false, x5

}

type paymentProcessorAuthorize struct {
  handler Payment
}

func (p *paymentProcessorAuthorize) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
  args := PaymentAuthorizeArgs{}
  if err = args.Read(iprot); err != nil {
    iprot.ReadMessageEnd()
    x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
    oprot.WriteMessageBegin("Authorize", thrift.EXCEPTION, seqId)
    x.Write(oprot)
    oprot.WriteMessageEnd()
    oprot.Flush(ctx)
    return false, err
  }

  iprot.ReadMessageEnd()
  result := PaymentAuthorizeResult{}
var retval bool
  var err2 error
  if retval, err2 = p.handler.Authorize(ctx); err2 != nil {
    x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing Authorize: " + err2.Error())
    oprot.WriteMessageBegin("Authorize", thrift.EXCEPTION, seqId)
    x.Write(oprot)
    oprot.WriteMessageEnd()
    oprot.Flush(ctx)
    return true, err2
  } else {
    result.Success = &retval
}
  if err2 = oprot.WriteMessageBegin("Authorize", thrift.REPLY, seqId); err2 != nil {
    err = err2
  }
  if err2 = result.Write(oprot); err == nil && err2 != nil {
    err = err2
  }
  if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
    err = err2
  }
  if err2 = oprot.Flush(ctx); err == nil && err2 != nil {
    err = err2
  }
  if err != nil {
    return
  }
  return true, err
}

type paymentProcessorPay struct {
  handler Payment
}

func (p *paymentProcessorPay) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
  args := PaymentPayArgs{}
  if err = args.Read(iprot); err != nil {
    iprot.ReadMessageEnd()
    x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
    oprot.WriteMessageBegin("Pay", thrift.EXCEPTION, seqId)
    x.Write(oprot)
    oprot.WriteMessageEnd()
    oprot.Flush(ctx)
    return false, err
  }

  iprot.ReadMessageEnd()
  result := PaymentPayResult{}
var retval bool
  var err2 error
  if retval, err2 = p.handler.Pay(ctx, args.Amount); err2 != nil {
    x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing Pay: " + err2.Error())
    oprot.WriteMessageBegin("Pay", thrift.EXCEPTION, seqId)
    x.Write(oprot)
    oprot.WriteMessageEnd()
    oprot.Flush(ctx)
    return true, err2
  } else {
    result.Success = &retval
}
  if err2 = oprot.WriteMessageBegin("Pay", thrift.REPLY, seqId); err2 != nil {
    err = err2
  }
  if err2 = result.Write(oprot); err == nil && err2 != nil {
    err = err2
  }
  if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
    err = err2
  }
  if err2 = oprot.Flush(ctx); err == nil && err2 != nil {
    err = err2
  }
  if err != nil {
    return
  }
  return true, err
}


// HELPER FUNCTIONS AND STRUCTURES

type PaymentAuthorizeArgs struct {
}

func NewPaymentAuthorizeArgs() *PaymentAuthorizeArgs {
  return &PaymentAuthorizeArgs{}
}

func (p *PaymentAuthorizeArgs) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    if err := iprot.Skip(fieldTypeId); err != nil {
      return err
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *PaymentAuthorizeArgs) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("Authorize_args"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *PaymentAuthorizeArgs) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("PaymentAuthorizeArgs(%+v)", *p)
}

// Attributes:
//  - Success
type PaymentAuthorizeResult struct {
  Success *bool `thrift:"success,0" db:"success" json:"success,omitempty"`
}

func NewPaymentAuthorizeResult() *PaymentAuthorizeResult {
  return &PaymentAuthorizeResult{}
}

var PaymentAuthorizeResult_Success_DEFAULT bool
func (p *PaymentAuthorizeResult) GetSuccess() bool {
  if !p.IsSetSuccess() {
    return PaymentAuthorizeResult_Success_DEFAULT
  }
return *p.Success
}
func (p *PaymentAuthorizeResult) IsSetSuccess() bool {
  return p.Success != nil
}

func (p *PaymentAuthorizeResult) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 0:
      if fieldTypeId == thrift.BOOL {
        if err := p.ReadField0(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    default:
      if err := iprot.Skip(fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *PaymentAuthorizeResult)  ReadField0(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadBool(); err != nil {
  return thrift.PrependError("error reading field 0: ", err)
} else {
  p.Success = &v
}
  return nil
}

func (p *PaymentAuthorizeResult) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("Authorize_result"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField0(oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *PaymentAuthorizeResult) writeField0(oprot thrift.TProtocol) (err error) {
  if p.IsSetSuccess() {
    if err := oprot.WriteFieldBegin("success", thrift.BOOL, 0); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err) }
    if err := oprot.WriteBool(bool(*p.Success)); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T.success (0) field write error: ", p), err) }
    if err := oprot.WriteFieldEnd(); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err) }
  }
  return err
}

func (p *PaymentAuthorizeResult) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("PaymentAuthorizeResult(%+v)", *p)
}

// Attributes:
//  - Amount
type PaymentPayArgs struct {
  Amount int32 `thrift:"amount,1" db:"amount" json:"amount"`
}

func NewPaymentPayArgs() *PaymentPayArgs {
  return &PaymentPayArgs{}
}


func (p *PaymentPayArgs) GetAmount() int32 {
  return p.Amount
}
func (p *PaymentPayArgs) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if fieldTypeId == thrift.I32 {
        if err := p.ReadField1(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    default:
      if err := iprot.Skip(fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *PaymentPayArgs)  ReadField1(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadI32(); err != nil {
  return thrift.PrependError("error reading field 1: ", err)
} else {
  p.Amount = v
}
  return nil
}

func (p *PaymentPayArgs) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("Pay_args"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField1(oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *PaymentPayArgs) writeField1(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("amount", thrift.I32, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:amount: ", p), err) }
  if err := oprot.WriteI32(int32(p.Amount)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.amount (1) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:amount: ", p), err) }
  return err
}

func (p *PaymentPayArgs) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("PaymentPayArgs(%+v)", *p)
}

// Attributes:
//  - Success
type PaymentPayResult struct {
  Success *bool `thrift:"success,0" db:"success" json:"success,omitempty"`
}

func NewPaymentPayResult() *PaymentPayResult {
  return &PaymentPayResult{}
}

var PaymentPayResult_Success_DEFAULT bool
func (p *PaymentPayResult) GetSuccess() bool {
  if !p.IsSetSuccess() {
    return PaymentPayResult_Success_DEFAULT
  }
return *p.Success
}
func (p *PaymentPayResult) IsSetSuccess() bool {
  return p.Success != nil
}

func (p *PaymentPayResult) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 0:
      if fieldTypeId == thrift.BOOL {
        if err := p.ReadField0(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    default:
      if err := iprot.Skip(fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *PaymentPayResult)  ReadField0(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadBool(); err != nil {
  return thrift.PrependError("error reading field 0: ", err)
} else {
  p.Success = &v
}
  return nil
}

func (p *PaymentPayResult) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("Pay_result"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField0(oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *PaymentPayResult) writeField0(oprot thrift.TProtocol) (err error) {
  if p.IsSetSuccess() {
    if err := oprot.WriteFieldBegin("success", thrift.BOOL, 0); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err) }
    if err := oprot.WriteBool(bool(*p.Success)); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T.success (0) field write error: ", p), err) }
    if err := oprot.WriteFieldEnd(); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err) }
  }
  return err
}

func (p *PaymentPayResult) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("PaymentPayResult(%+v)", *p)
}


