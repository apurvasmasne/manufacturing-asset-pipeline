# Manufacturing Asset Ingestion Pipeline & Queue

A high-throughput, concurrent manufacturing ingestion service built in Go. This service is designed to ingest multi-megabyte graphic files (e.g., sticker designs), track atomic asset states, and safely stream payloads to downstream physical hardware printing queues without blocking system memory.

## 🚀 System Architecture & Features

This project addresses the real-world scale challenges of a high-volume manufacturing platform by utilizing native Go concurrency primitives instead of relying on heavy, abstract frameworks:

*   **Bounded Worker Pool:** Utilizes a configurable pool of concurrent workers (simulating a scalable GCP Cloud Run architecture) to process variable payloads asynchronously.
*   **Memory Safeguards:** Implements bounded Go channels to throttle sudden ingestion spikes and prevent system memory exhaustion under intense parallel load.
*   **Deterministic State Simulation:** Uses cryptographically secure random processing ranges to accurately evaluate pipeline performance metrics, simulating actual vector rasterization times.
*   **Move-Fast Philosophy:** Architected to prioritize core processing throughput and high availability while intentionally deprioritizing non-critical edge cases to optimize shipping speed.

## 🛠️ Tech Stack

*   **Language:** Go (Golang) 1.21
*   **Concurreny Models:** Goroutines, Bounded Channels, `sync.WaitGroup`
*   **Deployment Target:** Optimized for containerized microservices (Docker / GCP Cloud Run)

## 📦 Local Setup & Execution

### Prerequisites
Make sure you have Go installed locally (version 1.21 or higher).

### Running the Pipeline
Clone the repository and execute the binary using the native Go toolchain:

```bash
git clone https://github.com/xxx/manufacturing-asset-pipeline.git
cd manufacturing-asset-pipeline
go run main.go
```

## 📊 Sample Pipeline Telemetry
When executed, the system initializes the worker matrix and streams parallel transactions across the channel network:

```text
🚀 Initializing High-Throughput Manufacturing Ingestion Pipeline...
✓ [MULE-QUEUE] Job JOB-TXN-0001 successfully processed by Worker 1 in 84ms
✓ [MULE-QUEUE] Job JOB-TXN-0002 successfully processed by Worker 3 in 65.65ms
✓ [MULE-QUEUE] Job JOB-TXN-0003 successfully processed by Worker 2 in 93.28ms
...
🏁 Pipeline Finished. Successfully dispatched 20/20 assets to printer queues.
```
