package googlestorage

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"

	"cloud.google.com/go/storage"
)

const project_ID = "14099599407"

func TestXxx(t *testing.T) {

	fmt.Println(os.Environ())
	os.Setenv("GOLANG_SAMPLES_PROJECT_ID", project_ID)
	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", `C:\Users\SonyPC\AppData\Roaming\gcloud\application_default_credentials.json`)
	fmt.Println(os.Environ())
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	// Close the client when finished.
	if err := client.Close(); err != nil {
		// TODO: handle error.
	}
}
