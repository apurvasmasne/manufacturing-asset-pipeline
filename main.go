package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"sync"
	"time"
)

// AssetJob represents a print-ready file upload (e.g., a sticker graphic)
type AssetJob struct {
	ID        string
	FileName  string
	SizeMB    float64
	Timestamp time.Time
}

// JobResult represents the outcome of the asset processing
type JobResult struct {
	JobID       string
	Success     bool
	ProcessedBy int
	Duration    time.Duration
}

// Worker simulates an isolated processing thread on a manufacturing stack
func Worker(id int, jobs <-chan AssetJob, results chan<- JobResult, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		startTime := time.Now()

		// Simulate processing / rasterizing heavy graphics files
		// High-throughput pipelines prioritize speed and deprioritize heavy edge cases
		fakeProcessingTime := cryptoRandDuration(50, 150)
		time.Sleep(fakeProcessingTime)

		results <- JobResult{
			JobID:       job.ID,
			Success:     true,
			ProcessedBy: id,
			Duration:    time.Since(startTime),
		}
	}
}

func main() {
	fmt.Println("🚀 Initializing High-Throughput Manufacturing Ingestion Pipeline...")

	numWorkers := 5
	numJobs := 20

	// Bounded channels to prevent memory flooding under high concurrent loads
	jobsChannel := make(chan AssetJob, numJobs)
	resultsChannel := make(chan JobResult, numJobs)

	var wg sync.WaitGroup

	// 1. Initialize Worker Pool (Simulating GCP Cloud Run scalable environment)
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go Worker(w, jobsChannel, resultsChannel, &wg)
	}

	// 2. Ingest stream of incoming simulated manufacturing files
	for i := 1; i <= numJobs; i++ {
		jobsChannel <- AssetJob{
			ID:        fmt.Sprintf("JOB-TXN-%04d", i),
			FileName:  fmt.Sprintf("custom_sticker_artwork_%d.png", i),
			SizeMB:    float64(cryptoRandInt(5, 50)),
			Timestamp: time.Now(),
		}
	}
	close(jobsChannel) // Signal workers that ingestion phase is complete

	// Wait for workers to finish in an isolated thread
	go func() {
		wg.Wait()
		close(resultsChannel)
	}()

	// 3. Process and log results (Simulating database metric logging)
	successCount := 0
	for result := range resultsChannel {
		if result.Success {
			successCount++
			fmt.Printf("✓ [MULE-QUEUE] Job %s successfully processed by Worker %d in %v\n",
				result.JobID, result.ProcessedBy, result.Duration)
		}
	}

	fmt.Printf("\n🏁 Pipeline Finished. Successfully dispatched %d/%d assets to printer queues.\n", successCount, numJobs)
}

// Secure helper utilities for pseudo-random processing simulation
func cryptoRandDuration(minMs, maxMs int64) time.Duration {
	n, _ := rand.Int(rand.Reader, big.NewInt(maxMs-minMs))
	return time.Duration(minMs+n.Int64()) * time.Millisecond
}

func cryptoRandInt(min, max int64) int64 {
	n, _ := rand.Int(rand.Reader, big.NewInt(max-min))
	return min + n.Int64()
}
