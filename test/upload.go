package main

import (
	"context"
	"log"
	"strings"

	"github.com/0gfoundation/0g-storage-client/common/blockchain"
	"github.com/0gfoundation/0g-storage-client/core"
	"github.com/0gfoundation/0g-storage-client/indexer"
	"github.com/0gfoundation/0g-storage-client/transfer"
)

func main() {
	w3client := blockchain.MustNewWeb3("https://evmrpc-testnet.0g.ai/", "")
	defer w3client.Close()

	indexerClient, err := indexer.NewClient("https://indexer-storage-testnet-turbo.0g.ai")
	if err != nil {
		log.Fatalf("create indexer client error: %v", err)
	}

	ctx := context.Background()
	nodes, err := indexerClient.SelectNodes(ctx, 1, []string{
		"http://34.174.223.105:5678",
		"http://104.196.238.89:5678",
		"http://34.57.99.219:5678",
		"http://34.55.197.204:5678",
		"http://34.133.200.179:5678",
	}, "max", true)
	if err != nil {
		log.Fatalf("select nodes error: %v", err)
	}

	uploader, err := transfer.NewUploader(ctx, w3client, nodes)
	if err != nil {
		log.Fatalf("create uploader error: %v", err)
	}

	file, err := core.Open("./main.go")
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	/*fragmentSizeStr :=
	fragmentSize, err := strconv.ParseInt(fragmentSizeStr, 10, 64)
	if err != nil {
		log.Fatalf("Error fragment size: %v", err)
	}*/

	log.Printf("Begin to upload file ...\n")
	_, roots, err := uploader.SplitableUpload(ctx, file, 429496730, transfer.UploadOption{
		FinalityRequired: transfer.FileFinalized,
		SkipTx:           true,
		NRetries:         10,
		TaskSize:         10,
		FullTrusted:      false,
		Method:           "10",
	})
	if err != nil {
		log.Fatalf("upload file error: %v", err)
	}
	log.Printf("Upload successful!\n")
	log.Printf("Roots size: %d\n", len(roots))
	s := make([]string, len(roots))
	for i, root := range roots {
		s[i] = root.String()
	}
	log.Printf("File uploaded in %v fragments, roots = %v", len(roots), strings.Join(s, ","))

}
