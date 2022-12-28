package go_eth_multicall

import (
	"context"
	"encoding/json"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	MultiCall2 "github.com/truongpx396/go-eth-multicall/contracts/MultiCall"
)

type Call struct {
	Name     string         `json:"name"`
	Target   common.Address `json:"target"`
	CallData []byte         `json:"call_data"`
}

type CallResponse struct {
	Success    bool   `json:"success"`
	ReturnData []byte `json:"returnData"`
}

func (call Call) GetMultiCall() MultiCall2.Multicall2Call {
	return MultiCall2.Multicall2Call{Target: call.Target, CallData: call.CallData}
}

func (call Call) GetCustomMultiCall() MultiCall2.CustomMulticall2Call {
	return MultiCall2.CustomMulticall2Call{Target: call.Target, CallData: call.CallData}
}

func randomSigner() *bind.TransactOpts {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		panic(err)
	}

	signer, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(1))
	if err != nil {
		panic(err)
	}

	signer.NoSend = true
	signer.Context = context.Background()
	signer.GasPrice = big.NewInt(0)

	return signer
}

type EthMultiCaller struct {
	Signer          *bind.TransactOpts
	Client          *ethclient.Client
	Abi             abi.ABI
	ContractAddress common.Address
}

func New(rawurl, multilcalContractAddress string) EthMultiCaller {
	client, err := ethclient.Dial(rawurl)
	if err != nil {
		panic(err)
	}

	// Load Multicall abi for later use
	mcAbi, err := abi.JSON(strings.NewReader(MultiCall2.MultiCallABI))
	if err != nil {
		panic(err)
	}

	contractAddress := common.HexToAddress(multilcalContractAddress)

	return EthMultiCaller{
		Signer:          randomSigner(),
		Client:          client,
		Abi:             mcAbi,
		ContractAddress: contractAddress,
	}
}

func (caller *EthMultiCaller) Execute(calls []Call) map[string]CallResponse {
	var responses []CallResponse

	var multiCalls = make([]MultiCall2.Multicall2Call, 0, len(calls))

	// Add calls to multicall structure for the contract
	for _, call := range calls {
		multiCalls = append(multiCalls, call.GetMultiCall())
	}

	// Prepare calldata for multicall
	callData, err := caller.Abi.Pack("tryAggregate", false, multiCalls)
	if err != nil {
		panic(err)
	}

	// Perform multicall
	resp, err := caller.Client.CallContract(context.Background(), ethereum.CallMsg{To: &caller.ContractAddress, Data: callData}, nil)
	if err != nil {
		panic(err)
	}

	// Unpack results
	unpackedResp, _ := caller.Abi.Unpack("tryAggregate", resp)

	a, err := json.Marshal(unpackedResp[0])
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(a, &responses)
	if err != nil {
		panic(err)
	}

	// Create mapping for results. Be aware that we sometimes get two empty results initially, unsure why
	results := make(map[string]CallResponse)
	for i, response := range responses {
		results[calls[i].Name] = response
	}

	return results
}

func (caller *EthMultiCaller) ExecuteBalances(calls []Call, userAddress string) map[string]CallResponse {
	var responses []CallResponse

	var multiCalls = make([]MultiCall2.CustomMulticall2Call, 0, len(calls))

	// Add calls to multicall structure for the contract
	for _, call := range calls {
		multiCalls = append(multiCalls, call.GetCustomMultiCall())
	}

	// Prepare calldata for multicall
	callData, err := caller.Abi.Pack("tryAggregateBalances", false, multiCalls, common.HexToAddress(userAddress))
	if err != nil {
		panic(err)
	}

	// Perform multicall
	resp, err := caller.Client.CallContract(context.Background(), ethereum.CallMsg{To: &caller.ContractAddress, Data: callData}, nil)
	if err != nil {
		panic(err)
	}

	// Unpack results
	unpackedResp, _ := caller.Abi.Unpack("tryAggregateBalances", resp)

	// nativeBalance := new(big.Int).SetBytes(unpackedResp[1])

	a, err := json.Marshal(unpackedResp[0])
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(a, &responses)
	if err != nil {
		panic(err)
	}

	// Create mapping for results. Be aware that we sometimes get two empty results initially, unsure why
	results := make(map[string]CallResponse)
	for i, response := range responses {
		results[calls[i].Name] = response
	}

	return results
}
