package main

import (
    "encoding/json"
    "fmt"

    "github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type LedgerController struct {
    contractapi.Contract
}

type Record struct {
    DealerCode        string `json:"dealerCode"`
    PhoneNumber       string `json:"phoneNumber"`
    PinCode           string `json:"pinCode"`
    Amount            int    `json:"amount"`
    State             string `json:"state"`
    TransactionValue  int    `json:"transactionValue"`
    TransactionMode   string `json:"transactionMode"`
    Note              string `json:"note"`
}

func (l *LedgerController) RegisterRecord(ctx contractapi.TransactionContextInterface, recordId string, dealerCode, phoneNumber, pinCode string, amount int, state string, transactionValue int, transactionMode, note string) error {
    record := Record{dealerCode, phoneNumber, pinCode, amount, state, transactionValue, transactionMode, note}
    recordJSON, err := json.Marshal(record)
    if err != nil {
        return err
    }
    return ctx.GetStub().PutState(recordId, recordJSON)
}

func (l *LedgerController) FetchRecord(ctx contractapi.TransactionContextInterface, recordId string) (*Record, error) {
    recordJSON, err := ctx.GetStub().GetState(recordId)
    if err != nil {
        return nil, fmt.Errorf("failed to read from ledger: %v", err)
    }
    if recordJSON == nil {
        return nil, fmt.Errorf("the record %s does not exist", recordId)
    }
    var record Record
    err = json.Unmarshal(recordJSON, &record)
    if err != nil {
        return nil, err
    }
    return &record, nil
}

func (l *LedgerController) RetrieveRecordTimeline(ctx contractapi.TransactionContextInterface, recordId string) ([]*Record, error) {
    resultsIterator, err := ctx.GetStub().GetHistoryForKey(recordId)
    if err != nil {
        return nil, err
    }
    defer resultsIterator.Close()

    var history []*Record
    for resultsIterator.HasNext() {
        response, err := resultsIterator.Next()
        if err != nil {
            return nil, err
        }
        var record Record
        if len(response.Value) > 0 {
            err = json.Unmarshal(response.Value, &record)
            if err != nil {
                return nil, err
            }
            history = append(history, &record)
        }
    }
    return history, nil
}
