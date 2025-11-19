package main

import (
	"context"
	"log"
	"strings"

	"github.com/0gfoundation/0g-storage-client/common/blockchain"
	"github.com/0gfoundation/0g-storage-client/indexer"
)

func main() {
	w3client := blockchain.MustNewWeb3("https://evmrpc-testnet.0g.ai/", "")
	defer w3client.Close()

	indexerClient, err := indexer.NewClient("https://indexer-storage-testnet-turbo.0g.ai")
	if err != nil {
		log.Fatalf("create indexer client error: %v", err)
	}

	ctx := context.Background()
	roots := "0x1f26cb89c25f3c9427234461abffb54b59a775fc3b535619e01d40571f72d8ac"
	if err := indexerClient.DownloadFragments(ctx, strings.Split(roots, ","), "downloaded_main.go", false); err != nil {
		log.Fatalf("Download file error: %v", err)
	}
}
