package main

import (
	"Didux-blackbox/src/data"
	"Didux-blackbox/src/data/types"
	"encoding/base64"
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGetAll(t *testing.T) {
	LoadConfig("/Users/Elkan/Projects/Didux.io/Didux-blackbox/test/test.conf")
	data.Start()
	peers := make([]types.Peer, 0)
	err := types.GetAll(&peers)
	if err != nil {
		require.Fail(t, "Unexpected error retrieving peers")
	}
	fmt.Println("Peers: ", len(peers))
	fmt.Println("")

	for _, peer := range peers {

		fmt.Println("Url: ", peer.URL)
		encoded := base64.StdEncoding.EncodeToString(peer.PublicKeys[0])
		fmt.Println("Key: ", encoded)
		fmt.Println("Failures: ", peer.Failures)
		fmt.Println("LastFailure: ", peer.LastFailure)
		fmt.Println("")
		require.NoError(t, err)
	}


	fmt.Println("--------------------------------------")
	fmt.Println("")


	encRaw := make([]types.EncryptedRawTransaction, 0)
	err = types.GetAll(&encRaw)
	if err != nil {
		require.Fail(t, "Unexpected error retrieving encryptedRawTransactions")
	}

	fmt.Println("encryptedRawTransactions: ", len(encRaw))
	fmt.Println("")

	for _, rawTx := range encRaw {
		fmt.Println("Hash: ", string(base64.StdEncoding.EncodeToString(rawTx.Hash)))
		fmt.Println("Sender: ", string(base64.StdEncoding.EncodeToString(rawTx.Sender)))
		fmt.Println("Timestamp: ", rawTx.Timestamp)
		fmt.Println("EncodedPayload: ", string(base64.StdEncoding.EncodeToString(rawTx.EncodedPayload)))
		fmt.Println("")
		require.NoError(t, err)
	}


	fmt.Println("--------------------------------------")
	fmt.Println("")

	encTx := make([]types.EncryptedTransaction, 0)
	err = types.GetAll(&encTx)
	if err != nil {
		require.Fail(t, "Unexpected error retrieving encryptedTransactions")
	}
	fmt.Println("encryptedTransactions: ", len(encTx))
	fmt.Println("")

	for _, tx := range encTx {
		fmt.Println("Hash: ", string(base64.StdEncoding.EncodeToString(tx.Hash)))
		fmt.Println("Timestamp: ", tx.Timestamp)
		fmt.Println("EncodedPayload: ", string(base64.StdEncoding.EncodeToString(tx.EncodedPayload)))
		fmt.Println("")
		require.NoError(t, err)
	}



}
