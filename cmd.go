package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"

	asset "cloud.google.com/go/asset/apiv1"
	"cloud.google.com/go/storage"
	assetpb "google.golang.org/genproto/googleapis/cloud/asset/v1"
)

func main() {
	ctx := context.Background()
	if len(os.Args) < 2 {
		log.Fatalf("Usage: ccat PROJECT_ID")
	}
	projectID := os.Args[1]
	client, err := asset.NewClient(ctx)
	if err != nil {
		log.Fatal(err)
	}

	bucketName := fmt.Sprintf("%s-for-assets", projectID)
	assetDumpFile := fmt.Sprintf("gs://%s/my-assets.txt", bucketName)

	req := &assetpb.ExportAssetsRequest{
		Parent:      fmt.Sprintf("projects/%s", projectID),
		ContentType: 1,
		OutputConfig: &assetpb.OutputConfig{
			Destination: &assetpb.OutputConfig_GcsDestination{
				GcsDestination: &assetpb.GcsDestination{
					ObjectUri: &assetpb.GcsDestination_Uri{
						Uri: string(assetDumpFile),
					},
				},
			},
		},
	}

	operation, err := client.ExportAssets(ctx, req)
	if err != nil {
		log.Fatal(err)
	}

	if _, err := operation.Wait(ctx); err != nil {
		log.Fatal(err)
	}

	storageClient, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	bkt := storageClient.Bucket(bucketName)
	obj := bkt.Object("my-assets.txt")
	r, err := obj.NewReader(ctx)
	if err != nil {
		log.Fatalf("Failed to create reader: %v", err)
	}

	defer r.Close()

	if _, err := io.Copy(os.Stdout, r); err != nil {
		log.Fatalf("Failed to copy assets: %v", err)
	}
}
