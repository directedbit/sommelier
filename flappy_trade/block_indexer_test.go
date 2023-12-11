package flappy_trade

import (
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestGetPoolPrice(t *testing.T) {
	price := GetPoolPrice(SOMM_OSMO_POOL, "SOMM", "OSMO")
	assert.Greater(t, 0.5, float64(price))
	assert.Less(t, 0.2, float64(price))
}

func TestParquetWriter(t *testing.T) {
	filename := "/Users/richard/workspace/training-data/test.parquet"
	pw := initParquetWriter(filename)
	data := []Transaction{
		{BlockNumber: 1,
			Time:            2,
			MessageType:     "3",
			Source:          "4",
			Destination:     "5",
			Amount:          6,
			OsmoToEthPrice:  -1.0,
			SommToOsmoPrice: -1.0},
	}

	for _, testTransaction := range data {
		if err := pw.Write(testTransaction); err != nil {
			log.Fatal("Failed writing to parquet file:", err)
		}
	}
	if err := pw.Flush(true); err != nil {
		log.Fatal("Failed flushing to parquet file:", err)
	}
	if err := pw.WriteStop(); err != nil {
		log.Fatal("Failed writing to stop parquet file:", err)
	}
	pw.PFile.Close()

	pr := initParquetReader(filename)
	rows := int(pr.GetNumRows())
	for i := 0; i < rows; i++ {
		log.Println("reading row", zap.Int("row", i))
		testTransaction := Transaction{}
		transactionRow := make([]Transaction, 1)
		if err := pr.Read(&transactionRow); err != nil {
			log.Fatal("Failed reading:", err)
		}
		for _, testTransaction = range transactionRow {
			log.Print("read row")
			testTransaction.OsmoToEthPrice = 12.34
			//testTransaction.OsmoToEthPrice = 12.34
			testTransaction.SommToOsmoPrice = 23.78
		}
	}
	pr.ReadStop()
	if err := pr.PFile.Close(); err != nil {
		log.Fatal("Failed closing file:", err)
	}
	//now write the price data

	if err := os.Remove(filename); err != nil {
		log.Fatal("Failed removing parquet file:", err)
	}
}
