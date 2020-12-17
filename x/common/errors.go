package common

import (
	"encoding/json"
	"fmt"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// const uint32
const (
	DefaultCodespace = "common"

	CodeInternalError              uint32 = 60101
	CodeInvalidPaginateParam       uint32 = 60102
	CodeCreateAddrFromBech32Failed uint32 = 60103
	CodeMarshalJSONFailed          uint32 = 60104
	CodeUnMarshalJSONFailed        uint32 = 60105 //"incorrectly formatted request data", err.Error()
	CodeStrconvFailed              uint32 = 60106
	CodeParseDecCoinFailed         uint32 = 60107
	CodeUnknownProposalType        uint32 = 60108
)

type SDKError struct {
	Codespace string `json:"codespace"`
	Code      uint32 `json:"code"`
	Message   string `json:"message"`
}

func ParseSDKError(errMsg string) SDKError {
	var sdkErr SDKError
	err := json.Unmarshal([]byte(errMsg), &sdkErr)
	if err != nil {
		sdkErr = SDKError{
			Codespace: DefaultCodespace,
			Code:      CodeInternalError,
			Message:   "internal error",
		}
		return sdkErr
	}
	return sdkErr
}

// invalid paginate param
func ErrInvalidPaginateParam(page int, perPage int) sdk.Error {
	return sdkerrors.New(DefaultCodespace, CodeInvalidPaginateParam, fmt.Sprintf("invalid params: page=%d or per_page=%d", page, perPage))
}

// invalid address
func ErrCreateAddrFromBech32Failed(addr string) sdk.Error {
	return sdkerrors.New(DefaultCodespace, CodeCreateAddrFromBech32Failed, fmt.Sprintf("invalid address：%s", addr))
}

// could not marshal result to JSON
func ErrMarshalJSONFailed(msg string) sdk.Error {
	return sdkerrors.New(DefaultCodespace, CodeMarshalJSONFailed, fmt.Sprintf("could not marshal result to JSON, %s", msg))
}

// could not unmarshal result to origin
func ErrUnMarshalJSONFailed(msg string) sdk.Error {
	return sdkerrors.New(DefaultCodespace, CodeUnMarshalJSONFailed, fmt.Sprintf("incorrectly formatted request data, %s", msg))
}

func ErrStrconvFailed(msg string) sdk.Error {
	return sdkerrors.New(DefaultCodespace, CodeStrconvFailed, fmt.Sprintf("incorrectly string conversion "))
}

func ErrUnknownProposalType(codespace string, msg string) sdk.Error {
	return sdkerrors.New(codespace, CodeUnknownProposalType, fmt.Sprintf("unknown proposal content type: %s", msg))
}