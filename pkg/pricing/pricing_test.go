/*
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

// TODO: use _test module
// cannot use a _test module here because tests are calling unexported methods
package pricing

import (
	"reflect"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/pricing"
)

func TestOnDemandPage(t *testing.T) {
	expected := map[string]float64{
		"m5ad.8xlarge": 1.648,
	}
	priceList := []string{
		`{"product":{"productFamily":"Compute Instance","attributes":{"enhancedNetworkingSupported":"Yes","intelTurboAvailable":"No","memory":"128 GiB","dedicatedEbsThroughput":"3500 Mbps","vcpu":"32","classicnetworkingsupport":"false","capacitystatus":"Used","locationType":"AWS Region","storage":"2 x 600 NVMe SSD","instanceFamily":"General purpose","operatingSystem":"Linux","intelAvx2Available":"No","regionCode":"us-west-2","physicalProcessor":"AMD EPYC 7571","clockSpeed":"2.5 GHz","ecu":"NA","networkPerformance":"Up to 10 Gigabit","servicename":"Amazon Elastic Compute Cloud","gpuMemory":"NA","vpcnetworkingsupport":"true","instanceType":"m5ad.8xlarge","tenancy":"Shared","usagetype":"USW2-BoxUsage:m5ad.8xlarge","normalizationSizeFactor":"64","intelAvxAvailable":"No","processorFeatures":"AVX; AVX2; AMD Turbo","servicecode":"AmazonEC2","licenseModel":"No License required","currentGeneration":"Yes","preInstalledSw":"NA","location":"US West (Oregon)","processorArchitecture":"64-bit","marketoption":"OnDemand","operation":"RunInstances","availabilityzone":"NA"},"sku":"2A7A5TKMHVKD2KNQ"},"serviceCode":"AmazonEC2","terms":{"OnDemand":{"2A7A5TKMHVKD2KNQ.JRTCKXETXF":{"priceDimensions":{"2A7A5TKMHVKD2KNQ.JRTCKXETXF.6YS6EN2CT7":{"unit":"Hrs","endRange":"Inf","description":"$1.648 per On Demand Linux m5ad.8xlarge Instance Hour","appliesTo":[],"rateCode":"2A7A5TKMHVKD2KNQ.JRTCKXETXF.6YS6EN2CT7","beginRange":"0","pricePerUnit":{"USD":"1.6480000000"}}},"sku":"2A7A5TKMHVKD2KNQ","effectiveDate":"2023-04-01T00:00:00Z","offerTermCode":"JRTCKXETXF","termAttributes":{}}},"Reserved":{"2A7A5TKMHVKD2KNQ.CUZHX8X6JH":{"priceDimensions":{"2A7A5TKMHVKD2KNQ.CUZHX8X6JH.6YS6EN2CT7":{"unit":"Hrs","endRange":"Inf","description":"Linux/UNIX (Amazon VPC), m5ad.8xlarge reserved instance applied","appliesTo":[],"rateCode":"2A7A5TKMHVKD2KNQ.CUZHX8X6JH.6YS6EN2CT7","beginRange":"0","pricePerUnit":{"USD":"0.5770000000"}},"2A7A5TKMHVKD2KNQ.CUZHX8X6JH.2TG2D8R56U":{"unit":"Quantity","description":"Upfront Fee","appliesTo":[],"rateCode":"2A7A5TKMHVKD2KNQ.CUZHX8X6JH.2TG2D8R56U","pricePerUnit":{"USD":"5053"}}},"sku":"2A7A5TKMHVKD2KNQ","effectiveDate":"2021-04-01T00:00:00Z","offerTermCode":"CUZHX8X6JH","termAttributes":{"LeaseContractLength":"1yr","OfferingClass":"convertible","PurchaseOption":"Partial Upfront"}},"2A7A5TKMHVKD2KNQ.7NE97W5U4E":{"priceDimensions":{"2A7A5TKMHVKD2KNQ.7NE97W5U4E.6YS6EN2CT7":{"unit":"Hrs","endRange":"Inf","description":"Linux/UNIX (Amazon VPC), m5ad.8xlarge reserved instance applied","appliesTo":[],"rateCode":"2A7A5TKMHVKD2KNQ.7NE97W5U4E.6YS6EN2CT7","beginRange":"0","pricePerUnit":{"USD":"1.2110000000"}}},"sku":"2A7A5TKMHVKD2KNQ","effectiveDate":"2021-04-01T00:00:00Z","offerTermCode":"7NE97W5U4E","termAttributes":{"LeaseContractLength":"1yr","OfferingClass":"convertible","PurchaseOption":"No Upfront"}},"2A7A5TKMHVKD2KNQ.6QCMYABX3D":{"priceDimensions":{"2A7A5TKMHVKD2KNQ.6QCMYABX3D.6YS6EN2CT7":{"unit":"Hrs","endRange":"Inf","description":"USD 0.0 per Linux/UNIX (Amazon VPC), m5ad.8xlarge reserved instance applied","appliesTo":[],"rateCode":"2A7A5TKMHVKD2KNQ.6QCMYABX3D.6YS6EN2CT7","beginRange":"0","pricePerUnit":{"USD":"0.0000000000"}},"2A7A5TKMHVKD2KNQ.6QCMYABX3D.2TG2D8R56U":{"unit":"Quantity","description":"Upfront Fee","appliesTo":[],"rateCode":"2A7A5TKMHVKD2KNQ.6QCMYABX3D.2TG2D8R56U","pricePerUnit":{"USD":"8489"}}},"sku":"2A7A5TKMHVKD2KNQ","effectiveDate":"2021-04-01T00:00:00Z","offerTermCode":"6QCMYABX3D","termAttributes":{"LeaseContractLength":"1yr","OfferingClass":"standard","PurchaseOption":"All Upfront"}},"2A7A5TKMHVKD2KNQ.38NPMPTW36":{"priceDimensions":{"2A7A5TKMHVKD2KNQ.38NPMPTW36.2TG2D8R56U":{"unit":"Quantity","description":"Upfront Fee","appliesTo":[],"rateCode":"2A7A5TKMHVKD2KNQ.38NPMPTW36.2TG2D8R56U","pricePerUnit":{"USD":"8662"}},"2A7A5TKMHVKD2KNQ.38NPMPTW36.6YS6EN2CT7":{"unit":"Hrs","endRange":"Inf","description":"Linux/UNIX (Amazon VPC), m5ad.8xlarge reserved instance applied","appliesTo":[],"rateCode":"2A7A5TKMHVKD2KNQ.38NPMPTW36.6YS6EN2CT7","beginRange":"0","pricePerUnit":{"USD":"0.3300000000"}}},"sku":"2A7A5TKMHVKD2KNQ","effectiveDate":"2021-04-01T00:00:00Z","offerTermCode":"38NPMPTW36","termAttributes":{"LeaseContractLength":"3yr","OfferingClass":"standard","PurchaseOption":"Partial Upfront"}},"2A7A5TKMHVKD2KNQ.R5XV2EPZQZ":{"priceDimensions":{"2A7A5TKMHVKD2KNQ.R5XV2EPZQZ.2TG2D8R56U":{"unit":"Quantity","description":"Upfront Fee","appliesTo":[],"rateCode":"2A7A5TKMHVKD2KNQ.R5XV2EPZQZ.2TG2D8R56U","pricePerUnit":{"USD":"10113"}},"2A7A5TKMHVKD2KNQ.R5XV2EPZQZ.6YS6EN2CT7":{"unit":"Hrs","endRange":"Inf","description":"Linux/UNIX (Amazon VPC), m5ad.8xlarge reserved instance applied","appliesTo":[],"rateCode":"2A7A5TKMHVKD2KNQ.R5XV2EPZQZ.6YS6EN2CT7","beginRange":"0","pricePerUnit":{"USD":"0.3850000000"}}},"sku":"2A7A5TKMHVKD2KNQ","effectiveDate":"2021-04-01T00:00:00Z","offerTermCode":"R5XV2EPZQZ","termAttributes":{"LeaseContractLength":"3yr","OfferingClass":"convertible","PurchaseOption":"Partial Upfront"}},"2A7A5TKMHVKD2KNQ.BPH4J8HBKS":{"priceDimensions":{"2A7A5TKMHVKD2KNQ.BPH4J8HBKS.6YS6EN2CT7":{"unit":"Hrs","endRange":"Inf","description":"Linux/UNIX (Amazon VPC), m5ad.8xlarge reserved instance applied","appliesTo":[],"rateCode":"2A7A5TKMHVKD2KNQ.BPH4J8HBKS.6YS6EN2CT7","beginRange":"0","pricePerUnit":{"USD":"0.7120000000"}}},"sku":"2A7A5TKMHVKD2KNQ","effectiveDate":"2021-04-01T00:00:00Z","offerTermCode":"BPH4J8HBKS","termAttributes":{"LeaseContractLength":"3yr","OfferingClass":"standard","PurchaseOption":"No Upfront"}},"2A7A5TKMHVKD2KNQ.NQ3QZPMQV9":{"priceDimensions":{"2A7A5TKMHVKD2KNQ.NQ3QZPMQV9.2TG2D8R56U":{"unit":"Quantity","description":"Upfront Fee","appliesTo":[],"rateCode":"2A7A5TKMHVKD2KNQ.NQ3QZPMQV9.2TG2D8R56U","pricePerUnit":{"USD":"16284"}},"2A7A5TKMHVKD2KNQ.NQ3QZPMQV9.6YS6EN2CT7":{"unit":"Hrs","endRange":"Inf","description":"USD 0.0 per Linux/UNIX (Amazon VPC), m5ad.8xlarge reserved instance applied","appliesTo":[],"rateCode":"2A7A5TKMHVKD2KNQ.NQ3QZPMQV9.6YS6EN2CT7","beginRange":"0","pricePerUnit":{"USD":"0.0000000000"}}},"sku":"2A7A5TKMHVKD2KNQ","effectiveDate":"2021-04-01T00:00:00Z","offerTermCode":"NQ3QZPMQV9","termAttributes":{"LeaseContractLength":"3yr","OfferingClass":"standard","PurchaseOption":"All Upfront"}},"2A7A5TKMHVKD2KNQ.Z2E3P23VKM":{"priceDimensions":{"2A7A5TKMHVKD2KNQ.Z2E3P23VKM.6YS6EN2CT7":{"unit":"Hrs","endRange":"Inf","description":"Linux/UNIX (Amazon VPC), m5ad.8xlarge reserved instance applied","appliesTo":[],"rateCode":"2A7A5TKMHVKD2KNQ.Z2E3P23VKM.6YS6EN2CT7","beginRange":"0","pricePerUnit":{"USD":"0.8310000000"}}},"sku":"2A7A5TKMHVKD2KNQ","effectiveDate":"2021-04-01T00:00:00Z","offerTermCode":"Z2E3P23VKM","termAttributes":{"LeaseContractLength":"3yr","OfferingClass":"convertible","PurchaseOption":"No Upfront"}},"2A7A5TKMHVKD2KNQ.HU7G6KETJZ":{"priceDimensions":{"2A7A5TKMHVKD2KNQ.HU7G6KETJZ.2TG2D8R56U":{"unit":"Quantity","description":"Upfront Fee","appliesTo":[],"rateCode":"2A7A5TKMHVKD2KNQ.HU7G6KETJZ.2TG2D8R56U","pricePerUnit":{"USD":"4331"}},"2A7A5TKMHVKD2KNQ.HU7G6KETJZ.6YS6EN2CT7":{"unit":"Hrs","endRange":"Inf","description":"Linux/UNIX (Amazon VPC), m5ad.8xlarge reserved instance applied","appliesTo":[],"rateCode":"2A7A5TKMHVKD2KNQ.HU7G6KETJZ.6YS6EN2CT7","beginRange":"0","pricePerUnit":{"USD":"0.4940000000"}}},"sku":"2A7A5TKMHVKD2KNQ","effectiveDate":"2021-04-01T00:00:00Z","offerTermCode":"HU7G6KETJZ","termAttributes":{"LeaseContractLength":"1yr","OfferingClass":"standard","PurchaseOption":"Partial Upfront"}},"2A7A5TKMHVKD2KNQ.MZU6U2429S":{"priceDimensions":{"2A7A5TKMHVKD2KNQ.MZU6U2429S.2TG2D8R56U":{"unit":"Quantity","description":"Upfront Fee","appliesTo":[],"rateCode":"2A7A5TKMHVKD2KNQ.MZU6U2429S.2TG2D8R56U","pricePerUnit":{"USD":"19821"}},"2A7A5TKMHVKD2KNQ.MZU6U2429S.6YS6EN2CT7":{"unit":"Hrs","endRange":"Inf","description":"USD 0.0 per Linux/UNIX (Amazon VPC), m5ad.8xlarge reserved instance applied","appliesTo":[],"rateCode":"2A7A5TKMHVKD2KNQ.MZU6U2429S.6YS6EN2CT7","beginRange":"0","pricePerUnit":{"USD":"0.0000000000"}}},"sku":"2A7A5TKMHVKD2KNQ","effectiveDate":"2021-04-01T00:00:00Z","offerTermCode":"MZU6U2429S","termAttributes":{"LeaseContractLength":"3yr","OfferingClass":"convertible","PurchaseOption":"All Upfront"}},"2A7A5TKMHVKD2KNQ.4NA7Y494T4":{"priceDimensions":{"2A7A5TKMHVKD2KNQ.4NA7Y494T4.6YS6EN2CT7":{"unit":"Hrs","endRange":"Inf","description":"Linux/UNIX (Amazon VPC), m5ad.8xlarge reserved instance applied","appliesTo":[],"rateCode":"2A7A5TKMHVKD2KNQ.4NA7Y494T4.6YS6EN2CT7","beginRange":"0","pricePerUnit":{"USD":"1.0380000000"}}},"sku":"2A7A5TKMHVKD2KNQ","effectiveDate":"2021-04-01T00:00:00Z","offerTermCode":"4NA7Y494T4","termAttributes":{"LeaseContractLength":"1yr","OfferingClass":"standard","PurchaseOption":"No Upfront"}},"2A7A5TKMHVKD2KNQ.VJWZNREJX2":{"priceDimensions":{"2A7A5TKMHVKD2KNQ.VJWZNREJX2.2TG2D8R56U":{"unit":"Quantity","description":"Upfront Fee","appliesTo":[],"rateCode":"2A7A5TKMHVKD2KNQ.VJWZNREJX2.2TG2D8R56U","pricePerUnit":{"USD":"9903"}},"2A7A5TKMHVKD2KNQ.VJWZNREJX2.6YS6EN2CT7":{"unit":"Hrs","endRange":"Inf","description":"USD 0.0 per Linux/UNIX (Amazon VPC), m5ad.8xlarge reserved instance applied","appliesTo":[],"rateCode":"2A7A5TKMHVKD2KNQ.VJWZNREJX2.6YS6EN2CT7","beginRange":"0","pricePerUnit":{"USD":"0.0000000000"}}},"sku":"2A7A5TKMHVKD2KNQ","effectiveDate":"2021-04-01T00:00:00Z","offerTermCode":"VJWZNREJX2","termAttributes":{"LeaseContractLength":"1yr","OfferingClass":"convertible","PurchaseOption":"All Upfront"}}}},"version":"20230407152936","publicationDate":"2023-04-07T15:29:36Z"}`,
	}
	p := NewStaticProvider()
	prices := make(map[string]float64)
	err := p.onDemandPage(prices, &pricing.GetProductsOutput{
		PriceList: priceList,
	})
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	if !reflect.DeepEqual(expected, prices) {
		t.Fatalf("unexpected prices; expected %v, got %v", expected, prices)
	}
}

func TestFargatePage(t *testing.T) {
	expectedGBPrice := 0.004445
	expectedVCPUPrice := 0.04048
	priceList := []string{
		`{"product":{"productFamily":"Compute","attributes":{"tiertype":"HAStandard","regionCode":"us-west-2","servicecode":"AmazonEKS","usagetype":"USW2-AmazonEKS-Local-Outposts-Hours:perCluster","locationType":"AWS Outposts","location":"US West (Oregon)","servicename":"Amazon Elastic Container Service for Kubernetes","operation":"CreateOperation"},"sku":"DNPPFW9XJKCJA2KM"},"serviceCode":"AmazonEKS","terms":{"OnDemand":{"DNPPFW9XJKCJA2KM.JRTCKXETXF":{"priceDimensions":{"DNPPFW9XJKCJA2KM.JRTCKXETXF.6YS6EN2CT7":{"unit":"hours","endRange":"Inf","description":"Amazon EKS local cluster usage on AWS Outposts","appliesTo":[],"rateCode":"DNPPFW9XJKCJA2KM.JRTCKXETXF.6YS6EN2CT7","beginRange":"0","pricePerUnit":{"USD":"0.1000000000"}}},"sku":"DNPPFW9XJKCJA2KM","effectiveDate":"2023-03-01T00:00:00Z","offerTermCode":"JRTCKXETXF","termAttributes":{}}}},"version":"20230306232536","publicationDate":"2023-03-06T23:25:36Z"}`,
		`{"product":{"productFamily":"Compute","attributes":{"tiertype":"HAStandard","regionCode":"us-west-2","servicecode":"AmazonEKS","usagetype":"USW2-AmazonEKS-Hours:perCluster","locationType":"AWS Region","location":"US West (Oregon)","servicename":"Amazon Elastic Container Service for Kubernetes","operation":"CreateOperation"},"sku":"DYQZA4NFAQKBU6H4"},"serviceCode":"AmazonEKS","terms":{"OnDemand":{"DYQZA4NFAQKBU6H4.JRTCKXETXF":{"priceDimensions":{"DYQZA4NFAQKBU6H4.JRTCKXETXF.6YS6EN2CT7":{"unit":"Hours","endRange":"Inf","description":"Amazon EKS cluster usage in US West (Oregon)","appliesTo":[],"rateCode":"DYQZA4NFAQKBU6H4.JRTCKXETXF.6YS6EN2CT7","beginRange":"0","pricePerUnit":{"USD":"0.1000000000"}}},"sku":"DYQZA4NFAQKBU6H4","effectiveDate":"2023-03-01T00:00:00Z","offerTermCode":"JRTCKXETXF","termAttributes":{}}}},"version":"20230306232536","publicationDate":"2023-03-06T23:25:36Z"}`,
		`{"product":{"productFamily":"Compute","attributes":{"regionCode":"us-west-2","servicecode":"AmazonEKS","tenancy":"Shared","usagetype":"USW2-Fargate-GB-Hours","locationType":"AWS Region","location":"US West (Oregon)","servicename":"Amazon Elastic Container Service for Kubernetes","memorytype":"perGB","operation":""},"sku":"GU5BE86QZK55S7P4"},"serviceCode":"AmazonEKS","terms":{"OnDemand":{"GU5BE86QZK55S7P4.JRTCKXETXF":{"priceDimensions":{"GU5BE86QZK55S7P4.JRTCKXETXF.6YS6EN2CT7":{"unit":"hours","endRange":"Inf","description":"AWS Fargate - Memory  - US West (Oregon)","appliesTo":[],"rateCode":"GU5BE86QZK55S7P4.JRTCKXETXF.6YS6EN2CT7","beginRange":"0","pricePerUnit":{"USD":"0.0044450000"}}},"sku":"GU5BE86QZK55S7P4","effectiveDate":"2023-03-01T00:00:00Z","offerTermCode":"JRTCKXETXF","termAttributes":{}}}},"version":"20230306232536","publicationDate":"2023-03-06T23:25:36Z"}`,
		`{"product":{"productFamily":"Compute","attributes":{"regionCode":"us-west-2","servicecode":"AmazonEKS","tenancy":"Shared","usagetype":"USW2-Fargate-vCPU-Hours:perCPU","locationType":"AWS Region","location":"US West (Oregon)","servicename":"Amazon Elastic Container Service for Kubernetes","cputype":"perCPU","operation":""},"sku":"JRYUZKUX5CVYXGTE"},"serviceCode":"AmazonEKS","terms":{"OnDemand":{"JRYUZKUX5CVYXGTE.JRTCKXETXF":{"priceDimensions":{"JRYUZKUX5CVYXGTE.JRTCKXETXF.6YS6EN2CT7":{"unit":"hours","endRange":"Inf","description":"AWS Fargate - vCPU - US West (Oregon)","appliesTo":[],"rateCode":"JRYUZKUX5CVYXGTE.JRTCKXETXF.6YS6EN2CT7","beginRange":"0","pricePerUnit":{"USD":"0.0404800000"}}},"sku":"JRYUZKUX5CVYXGTE","effectiveDate":"2023-03-01T00:00:00Z","offerTermCode":"JRTCKXETXF","termAttributes":{}}}},"version":"20230306232536","publicationDate":"2023-03-06T23:25:36Z"}`,
	}
	p := NewStaticProvider()
	err := p.fargatePage(&pricing.GetProductsOutput{
		PriceList: priceList,
	})
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	if p.fargateGBPricePerHour != expectedGBPrice {
		t.Errorf("unexpected fargateGBPricePerHour; expected %f, got %f", expectedGBPrice, p.fargateGBPricePerHour)
	}
	if p.fargateVCPUPricePerHour != expectedVCPUPrice {
		t.Errorf("unexpected fargateVCPUPricePerHour; expected %f, got %f", expectedVCPUPrice, p.fargateVCPUPricePerHour)
	}
}
