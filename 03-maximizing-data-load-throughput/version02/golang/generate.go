/*
Copyright 2018 Google Inc. All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"log"
	"math/rand"
	"sync"
)

func generateExceptionLogs(gcs *GCSclient, wg *sync.WaitGroup) {
	log.Printf("Generating %v exception logs ...", *numExceptionLogs)
	var accs []*dbExceptionLogs
	var recordsPerFile = *numRecordsPerFile
	for i := 0; i < *numExceptionLogs; i++ {
		accs = append(accs, NewGeneratedExceptionLogs(int64(i)))
		if len(accs)%recordsPerFile == 0 {
			go gcs.writeCSV(wg, *bucketName, *folderPath, "exceptionlogs", i/recordsPerFile, accs)
			accs = []*dbExceptionLogs{}
			log.Printf("Generated %v Exception logs of %v ...", i, *numExceptionLogs)
		}
	}
	go gcs.writeCSV(wg, *bucketName, *folderPath, "exceptionlogs", *numExceptionLogs/recordsPerFile, accs)
	log.Printf("Generated %v Exception logs of %v ...", *numExceptionLogs, *numExceptionLogs)
}

func generateAccounts(gcs *GCSclient, wg *sync.WaitGroup) {
	log.Printf("Generating %v accounts...", *numAccounts)
	var accs []*dbAccount
	var recordsPerFile = *numRecordsPerFile
	for i := 0; i < *numAccounts; i++ {
		accs = append(accs, NewGeneratedAccount(int64(i)))
		if len(accs)%recordsPerFile == 0 {
			go gcs.writeCSV(wg, *bucketName, *folderPath, "accounts", i/recordsPerFile, accs)
			accs = []*dbAccount{}
			log.Printf("Generated %v Accounts of %v ...", i, *numAccounts)
		}
	}
	go gcs.writeCSV(wg, *bucketName, *folderPath, "accounts", *numAccounts/recordsPerFile, accs)
	log.Printf("Generated %v Accounts of %v ...", *numAccounts, *numAccounts)
}

func generateOrders(gcs *GCSclient, wg *sync.WaitGroup) {
	log.Printf("Generating %v orders, order lineitems and payments...", *numOrders)
	var recordsPerFile = *numRecordsPerFile
	var os []*dbOrder
	var ols []*dbOrderLineItem
	var ps []*dbPayment
	for i := 0; i < *numOrders; i++ {
		o := NewGeneratedOrder(int64(i), int64(rand.Intn(*numAccounts)))
		os = append(os, o)

		for j := 0; j < rand.Intn(10); j++ {
			ol := NewGeneratedOrderLineItem(o.OldID, int64(rand.Intn(*numProducts)))
			ols = append(ols, ol)
		}

		p := NewGeneratedPayment(o.OldID)
		ps = append(ps, p)

		if len(os)%recordsPerFile == 0 {
			go gcs.writeCSV(wg, *bucketName, *folderPath, "orders", i/recordsPerFile, os)
			log.Printf("Generated %v Orders of %v ...", i, *numOrders)
			os = []*dbOrder{}
			go gcs.writeCSV(wg, *bucketName, *folderPath, "orderlineitems", i/recordsPerFile, ols)
			log.Printf("Generated %v OrderLineItems ...", len(ols))
			ols = []*dbOrderLineItem{}
			go gcs.writeCSV(wg, *bucketName, *folderPath, "payments", i/recordsPerFile, ps)
			log.Printf("Generated %v Payments of %v ...", i, *numOrders)
			ps = []*dbPayment{}

		}
	}
	go gcs.writeCSV(wg, *bucketName, *folderPath, "orders", *numOrders/recordsPerFile, os)
	log.Printf("Generated %v Orders of %v ...", *numOrders, *numOrders)
	go gcs.writeCSV(wg, *bucketName, *folderPath, "orderlineitems", *numOrders/recordsPerFile, ols)
	log.Printf("Generated %v OrderLineItems ...", len(ols))
	go gcs.writeCSV(wg, *bucketName, *folderPath, "payments", *numOrders/recordsPerFile, ps)
	log.Printf("Generated %v Payments of %v ...", *numOrders, *numOrders)
}

func generateProducts(gcs *GCSclient, wg *sync.WaitGroup) {
	log.Printf("Generating %v products...", *numProducts)
	var recordsPerFile = *numRecordsPerFile
	var ps []*dbProduct
	var pis []*dbProductInfo
	var pps []*dbProductPrice
	var pss []*dbProductStock
	for i := 0; i < *numProducts; i++ {
		p := NewGeneratedProduct(int64(i))
		ps = append(ps, p)

		pi := NewGeneratedProductInfo(p.OldID, *infoDataSize)
		pis = append(pis, pi)

		for j := 0; j < rand.Intn(10)+1; j++ {
			pp := NewGeneratedProductPrice(p.OldID)
			pps = append(pps, pp)
		}

		for j := 0; j < rand.Intn(10)+1; j++ {
			ps := NewGeneratedProductStock(p.OldID, int64(rand.Intn(*numWarehouses)))
			pss = append(pss, ps)
		}
		if len(ps)%recordsPerFile == 0 {
			go gcs.writeCSV(wg, *bucketName, *folderPath, "products", i/recordsPerFile, ps)
			log.Printf("Generated %v Products of %v ...", i, *numProducts)
			ps = []*dbProduct{}
			go gcs.writeCSV(wg, *bucketName, *folderPath, "productinfos", i/recordsPerFile, pis)
			log.Printf("Generated %v ProductInfos of %v ...", i, *numProducts)
			pis = []*dbProductInfo{}
			go gcs.writeCSV(wg, *bucketName, *folderPath, "productprices", i/recordsPerFile, pps)
			log.Printf("Generated %v ProductPrices", len(pps))
			pps = []*dbProductPrice{}
			go gcs.writeCSV(wg, *bucketName, *folderPath, "productstocks", i/recordsPerFile, pss)
			log.Printf("Generated %v ProductStocks", len(pps))
			pss = []*dbProductStock{}
		}

	}
	go gcs.writeCSV(wg, *bucketName, *folderPath, "products", *numProducts/recordsPerFile, ps)
	log.Printf("Generated %v Products of %v ...", *numProducts, *numProducts)
	go gcs.writeCSV(wg, *bucketName, *folderPath, "productinfos", *numProducts/recordsPerFile, pis)
	log.Printf("Generated %v ProductInfos of %v ...", *numProducts, *numProducts)
	go gcs.writeCSV(wg, *bucketName, *folderPath, "productprices", *numProducts/recordsPerFile, pps)
	log.Printf("Generated %v ProductPrices", len(pps))
	go gcs.writeCSV(wg, *bucketName, *folderPath, "productstocks", *numProducts/recordsPerFile, pss)
	log.Printf("Generated %v ProductStocks", len(pps))
}

func generateWarehouses(gcs *GCSclient, wg *sync.WaitGroup) {
	log.Printf("Generating %v warehouses...", *numWarehouses)
	var recordsPerFile = *numRecordsPerFile
	var whs []*dbWarehouse
	for i := 0; i < *numWarehouses; i++ {
		wh := NewGeneratedWarehouse(int64(i))
		whs = append(whs, wh)
		if len(whs)%recordsPerFile == 0 {
			go gcs.writeCSV(wg, *bucketName, *folderPath, "warehouses", i/recordsPerFile, whs)
			whs = []*dbWarehouse{}
			log.Printf("Generated %v Warehouses of %v ...", i, *numWarehouses)
		}
	}
	go gcs.writeCSV(wg, *bucketName, *folderPath, "warehouses", *numWarehouses/recordsPerFile, whs)
	log.Printf("Generated %v Warehouses of %v ...", *numWarehouses, *numWarehouses)
}
