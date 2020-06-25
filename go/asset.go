package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"database/sql"
	_ "github.com/lib/pq"
)

// Define the Smart Contract structure
type SmartContract struct {
	contractapi.Contract
}

// データ構造の定義
type Asset struct {
	Year     string `json:"year"`     // 初度登録年
	Month    string `json:"month"`	  // 初度登録月
	Mileage  int    `json:"mileage"`  // 走行距離(km)
	Battery  int    `json:"battery"`  // バッテリーライフ(%)
	Location string `jasn:"location"` // 位置
}
type AssetWithOwner struct {
	Name    string	// 名前
	Country string	// 国
	City    string	// 都道府県
	Addr    string	// 市区町村
	Record  *Asset
}
// クエリ結果（レーコード検索）
type QueryResult struct {
	Key     string		// レコードID(VINコード)
	Record  *Asset
}
// クエリ結果（履歴検索）
type GetHisResult struct {
	TxId      string	// トランザクションID
	Timestamp string	// タイムスタンプ
	IsDelete  bool		// 削除(廃車)フラグ
	Record    *Asset
}

func (s *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
	fmt.Println("InitLedger")

	return nil
}

func (s *SmartContract) InitDB(ctx contractapi.TransactionContextInterface) error {
	fmt.Println("InitDB")

        db, err := sql.Open("postgres", "host=pgsql port=5432 user=postgres password=secret dbname=asset sslmode=disable")
        defer db.Close()

        if err != nil {
                return fmt.Errorf("sql.Open: %s", err.Error())
        }

        rows, err := db.Query("SELECT * FROM asset;")

        if err != nil {
                return fmt.Errorf("sql.Query: %s", err.Error())
        }

        var id string
        for rows.Next() {
                var asset Asset
                rows.Scan(&id, &asset.Year, &asset.Month, &asset.Mileage, &asset.Battery, &asset.Location)
                assetAsBytes, _ := json.Marshal(asset)
                ctx.GetStub().PutState(id, assetAsBytes)
        }

	return nil
}

func (s *SmartContract) QueryAsset(ctx contractapi.TransactionContextInterface, key string) (*Asset, error) {
	fmt.Println("QueryAsset")

        assetAsBytes, err := ctx.GetStub().GetState(key)

        if err != nil {
                return nil, fmt.Errorf("Failed to read from world state. %s", err.Error())
        }

        if assetAsBytes == nil {
                return nil, fmt.Errorf("%s does not exist", key)
        }

        asset := new(Asset)
        _ = json.Unmarshal(assetAsBytes, asset)

        return asset, nil
}

func (s *SmartContract) QueryAssetWithOwner(ctx contractapi.TransactionContextInterface, key string) (*AssetWithOwner, error) {
	fmt.Println("QueryAssetWithOwner")

        assetAsBytes, err := ctx.GetStub().GetState(key)

        if err != nil {
                return nil, fmt.Errorf("Failed to read from world state. %s", err.Error())
        }

        if assetAsBytes == nil {
                return nil, fmt.Errorf("%s does not exist", key)
        }

        asset := new(Asset)
        _ = json.Unmarshal(assetAsBytes, asset)

	// tmp code start
	/*
        db, err := sql.Open("postgres", "host=pgsql port=5432 user=postgres password=secret dbname=asset sslmode=disable")
        defer db.Close()
        if err != nil {
                return nil, fmt.Errorf("sql.Open: %s", err.Error())
        }
	sql := "SELECT * FROM owner WHERE id = " + key + ";"
	fmt.Println(sql)
	awo := new(AssetWithOwner)
	awo.Record = asset
	*/
	// tmp code end

        db, err := sql.Open("postgres", "host=pgsql port=5432 user=postgres password=secret dbname=asset sslmode=disable")
        defer db.Close()
        if err != nil {
                return nil, fmt.Errorf("sql.Open: %s", err.Error())
        }

	sql := "SELECT * FROM owner WHERE id = '" + key + "';"
        rows, err := db.Query(sql)
        if err != nil {
                return nil, fmt.Errorf("db.Query: %s", err.Error())
        }

        var id string
	awo := new(AssetWithOwner)
        for rows.Next() {
                rows.Scan(&id, &awo.Name, &awo.Country, &awo.City, &awo.Addr)
		awo.Record = asset
        }

        return awo, nil
}

func (s *SmartContract) CreateAsset(ctx contractapi.TransactionContextInterface, key string, year string, month string, sMileage string, sBattery string, location string) error {
	fmt.Println("CreateAsset")

	mileage, _ :=strconv.Atoi(sMileage)
	battery, _ :=strconv.Atoi(sBattery)
	asset := Asset{
		Year:     year,
		Month:    month,
		Mileage:  mileage,
		Battery:  battery,
		Location: location,
	}
	assetAsBytes, _ := json.Marshal(asset)

	return ctx.GetStub().PutState(key, assetAsBytes)
}

func (s *SmartContract) UpdateAsset(ctx contractapi.TransactionContextInterface, key string, sMileage string, sBattery string, location string) error {
	fmt.Println("UpdateAsset")

	assetAsBytes, _ := ctx.GetStub().GetState(key)
	asset := new(Asset)
	json.Unmarshal(assetAsBytes, &asset)


	if sMileage != "" {
		mileage, _ := strconv.Atoi(sMileage) // inpu mileage (km)
		if mileage < asset.Mileage {
			return fmt.Errorf("Invalid argument.")
		}
		asset.Mileage = mileage
	}

	if sBattery != "" {
		battery, _ := strconv.Atoi(sBattery) // inpu battery (%)
		if battery > (asset.Battery + 5) {
			return fmt.Errorf("Invalid argument.")
		}
		asset.Battery = battery
	}

	if location != "" {
		asset.Location = location
	}

	assetAsBytes, _ = json.Marshal(asset)

	return ctx.GetStub().PutState(key, assetAsBytes)
}

func (s *SmartContract) ResetAsset(ctx contractapi.TransactionContextInterface, key string, sMileage string, sBattery string, location string) error {
	fmt.Println("ResetAsset")

	assetAsBytes, _ := ctx.GetStub().GetState(key)
	asset := new(Asset)
	json.Unmarshal(assetAsBytes, &asset)


	if sMileage != "" {
		mileage, _ := strconv.Atoi(sMileage) // inpu mileage (km)
		asset.Mileage = mileage
	}

	if sBattery != "" {
		battery, _ := strconv.Atoi(sBattery) // inpu battery (%)
		asset.Battery = battery
	}

	if location != "" {
		asset.Location = location
	}

	assetAsBytes, _ = json.Marshal(asset)

	return ctx.GetStub().PutState(key, assetAsBytes)
}

func (s *SmartContract) QueryAllAssets(ctx contractapi.TransactionContextInterface) ([]QueryResult, error) {
	fmt.Println("QueryAllAssets")

	startKey := ""
	endKey   := ""

	resultsIterator, err := ctx.GetStub().GetStateByRange(startKey, endKey)
	if err != nil {
		return nil, fmt.Errorf("ctx.GetStub().GetStateByRange: %s", err.Error())
	}
	defer resultsIterator.Close()

	results := []QueryResult{}

	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		asset := new(Asset)
		_ = json.Unmarshal(queryResponse.Value, asset)
		queryResult := QueryResult{Key: queryResponse.Key, Record: asset}
		results = append(results, queryResult)
	}

	return results, nil
}

func (s *SmartContract) QueryRangeAssets(ctx contractapi.TransactionContextInterface, sPageSize string, bookmark string) ([]QueryResult, error) {
	fmt.Println("QueryRangeAssets")

	startKey := ""
	endKey := ""
	var pageSize int32
	tmpSize, _ := strconv.Atoi(sPageSize)
        pageSize = int32(tmpSize)

	resultsIterator, _, err := ctx.GetStub().GetStateByRangeWithPagination(startKey, endKey, pageSize, bookmark)
	if err != nil {
		return nil, fmt.Errorf("ctx.GetStub().GetStateByRangeWithPagination: %s", err.Error())
	}
	defer resultsIterator.Close()

	results := []QueryResult{}

	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		asset := new(Asset)
		_ = json.Unmarshal(queryResponse.Value, asset)
		queryResult := QueryResult{Key: queryResponse.Key, Record: asset}
		results = append(results, queryResult)
	}

	return results, nil
}

func (s *SmartContract) DeleteAsset(ctx contractapi.TransactionContextInterface, key string) error {
	fmt.Println("DeleteAsset")

	return ctx.GetStub().DelState(key)
}


func (s *SmartContract) GetHistoryOfAsset(ctx contractapi.TransactionContextInterface, key string) ([]GetHisResult, error) {
	fmt.Println("GetHistoryOfAsset")

	resultsIterator, err := ctx.GetStub().GetHistoryForKey(key)
	if err != nil {
		return nil, fmt.Errorf("ctx.GetStub().GetHistoryForKey: %s", err.Error())
	}
	defer resultsIterator.Close()

	results := []GetHisResult{}

	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
                        return nil, err
                }
		asset := new(Asset)
		_ = json.Unmarshal(queryResponse.Value, asset)
		t := time.Unix(queryResponse.Timestamp.Seconds, 0).String()
		//t := time.Unix(queryResponse.Timestamp.Seconds, int64(queryResponse.Timestamp.Nanos)).String()
		queryResult := GetHisResult{TxId:queryResponse.TxId, Timestamp:t, IsDelete:queryResponse.IsDelete, Record: asset}
		results = append(results, queryResult)
	}

	return results, nil
}

func main() {

        chaincode, err := contractapi.NewChaincode(new(SmartContract))

        if err != nil {
                fmt.Printf("Error create asset chaincode: %s", err.Error())
                return
        }

        if err := chaincode.Start(); err != nil {
                fmt.Printf("Error starting asset chaincode: %s", err.Error())
        }
}
