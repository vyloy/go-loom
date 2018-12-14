// +build js,wasm

package plugin

import (
	"bytes"
	"context"
	"encoding/binary"
	"github.com/gogo/protobuf/proto"
	"github.com/loomnetwork/go-loom"
	"github.com/loomnetwork/go-loom/plugin/types"
	"github.com/perlin-network/life/gowasm"
	"google.golang.org/grpc"
	"syscall/js"
)

type WASMContext struct {
	GRPCContext
}

func MakeWASMContext(object js.Value, req *types.ContractCallRequest) *WASMContext {
	return &WASMContext{
		GRPCContext: GRPCContext{
			GRPCAPIClient: &GRPCAPIClient{
				client: &WASMAPIClient{
					ccr: object,
				},
			},
			message:      req.Message,
			block:        req.Block,
			contractAddr: loom.UnmarshalAddressPB(req.ContractAddress),
		},
	}
}

type WASMAPIClient struct {
	ccr js.Value
}

func (c *WASMAPIClient) run(name string, resp proto.Message, args ...proto.Message) error {
	buffer := &bytes.Buffer{}
	err := binary.Write(buffer, binary.BigEndian, int16(len(name)))
	if err != nil {
		return err
	}
	_, err = buffer.WriteString(name)
	if err != nil {
		return err
	}
	err = binary.Write(buffer, binary.BigEndian, int16(len(args)))
	if err != nil {
		return err
	}
	if len(args) > 0 {
		data := make([][]byte, int16(len(args)))
		for i, arg := range args {
			data[i], err = proto.Marshal(arg)
			if err != nil {
				return err
			}
		}
		for _, d := range data {
			err = binary.Write(buffer, binary.BigEndian, int16(len(d)))
			if err != nil {
				return err
			}
			_, err = buffer.Write(d)
			if err != nil {
				return err
			}
		}
	}

	bs := buffer.Bytes()
	outData := c.ccr.Call("APIRequest", gowasm.ByteSlice2JSValue(bs))
	return proto.Unmarshal(gowasm.JSValue2ByteSlice(outData), resp)
}

func (c *WASMAPIClient) Get(ctx context.Context, in *types.GetRequest, opts ...grpc.CallOption) (*types.GetResponse, error) {
	out := new(types.GetResponse)
	err := c.run("Get", out, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *WASMAPIClient) Has(ctx context.Context, in *types.HasRequest, opts ...grpc.CallOption) (*types.HasResponse, error) {
	out := new(types.HasResponse)
	err := c.run("Has", out, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *WASMAPIClient) Range(ctx context.Context, in *types.RangeRequest, opts ...grpc.CallOption) (*types.RangeResponse, error) {
	out := new(types.RangeResponse)
	err := c.run("Range", out, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *WASMAPIClient) ValidatorPower(ctx context.Context, in *types.ValidatorPowerRequest, opts ...grpc.CallOption) (*types.ValidatorPowerResponse, error) {
	out := new(types.ValidatorPowerResponse)
	err := c.run("ValidatorPower", out, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *WASMAPIClient) StaticCall(ctx context.Context, in *types.CallRequest, opts ...grpc.CallOption) (*types.CallResponse, error) {
	out := new(types.CallResponse)
	err := c.run("StaticCall", out, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *WASMAPIClient) Resolve(ctx context.Context, in *types.ResolveRequest, opts ...grpc.CallOption) (*types.ResolveResponse, error) {
	out := new(types.ResolveResponse)
	err := c.run("Resolve", out, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *WASMAPIClient) Emit(ctx context.Context, in *types.EmitRequest, opts ...grpc.CallOption) (*types.EmitResponse, error) {
	out := new(types.EmitResponse)
	err := c.run("Emit", out, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *WASMAPIClient) GetEvmTxReceipt(ctx context.Context, in *types.EvmTxReceiptRequest, opts ...grpc.CallOption) (*types.EvmTxReceipt, error) {
	out := new(types.EvmTxReceipt)
	err := c.run("GetEvmTxReceipt", out, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *WASMAPIClient) Set(ctx context.Context, in *types.SetRequest, opts ...grpc.CallOption) (*types.SetResponse, error) {
	out := new(types.SetResponse)
	err := c.run("Set", out, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *WASMAPIClient) Delete(ctx context.Context, in *types.DeleteRequest, opts ...grpc.CallOption) (*types.DeleteResponse, error) {
	out := new(types.DeleteResponse)
	err := c.run("Delete", out, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *WASMAPIClient) Call(ctx context.Context, in *types.CallRequest, opts ...grpc.CallOption) (*types.CallResponse, error) {
	out := new(types.CallResponse)
	err := c.run("Call", out, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *WASMAPIClient) SetValidatorPower(ctx context.Context, in *types.SetValidatorPowerRequest, opts ...grpc.CallOption) (*types.SetValidatorPowerResponse, error) {
	out := new(types.SetValidatorPowerResponse)
	err := c.run("SetValidatorPower", out, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *WASMAPIClient) ContractRecord(ctx context.Context, in *types.ContractRecordRequest, opts ...grpc.CallOption) (*types.ContractRecordResponse, error) {
	out := new(types.ContractRecordResponse)
	err := c.run("ContractRecord", out, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _ types.APIClient = &WASMAPIClient{}

func ServeWASMContract(contract Contract) (err error) {
	ccr := js.Global().Get("ContractCallHandler")
	method := ccr.Call("GetMethod")

	defer func() {
		if err != nil {
			ccr.Call("SetError", err.Error())
		}
	}()

	switch method.String() {
	case "Meta":
		meta, err := contract.Meta()
		if err != nil {
			return err
		}
		data, err := proto.Marshal(&meta)
		if err != nil {
			return err
		}
		ccr.Call("SetResponse", gowasm.ByteSlice2JSValue(data))
	case "Init":
		req := &types.ContractCallRequest{}
		bs := ccr.Call("GetRequest")
		err := proto.Unmarshal(gowasm.JSValue2ByteSlice(bs), req)
		if err != nil {
			return err
		}
		err = contract.Init(MakeWASMContext(ccr, req), req.Request)
		if err != nil {
			return err
		}
	case "Call":
		req := &types.ContractCallRequest{}
		bs := ccr.Call("GetRequest")
		err := proto.Unmarshal(gowasm.JSValue2ByteSlice(bs), req)
		if err != nil {
			return err
		}
		resp, err := contract.Call(MakeWASMContext(ccr, req), req.Request)
		if err != nil {
			return err
		}
		data, err := proto.Marshal(resp)
		if err != nil {
			return err
		}
		ccr.Call("SetResponse", gowasm.ByteSlice2JSValue(data))
	case "StaticCall":
		req := &types.ContractCallRequest{}
		bs := ccr.Call("GetRequest")
		err := proto.Unmarshal(gowasm.JSValue2ByteSlice(bs), req)
		if err != nil {
			return err
		}
		resp, err := contract.StaticCall(MakeWASMContext(ccr, req), req.Request)
		if err != nil {
			return err
		}
		data, err := proto.Marshal(resp)
		if err != nil {
			return err
		}
		ccr.Call("SetResponse", gowasm.ByteSlice2JSValue(data))
	}
	return nil
}
