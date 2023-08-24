package main

import (
	"flag"
	"fmt"
	"strings"
	"sync"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/acm"
)

func scanRegion(region, profile string, wgRegion *sync.WaitGroup) {
	defer wgRegion.Done()

	sess, err := session.NewSessionWithOptions(session.Options{
		Profile: profile,
		Config:  aws.Config{Region: aws.String(region)},
	})
	if err != nil {
		fmt.Println("Error creating session:", err)
		return
	}

	svc := acm.New(sess)
	input := &acm.ListCertificatesInput{}
	result, err := svc.ListCertificates(input)
	if err != nil {
		fmt.Printf("Error listing certificates in region %s: %v\n", region, err)
		return
	}

	var wgCerts sync.WaitGroup
	for _, summary := range result.CertificateSummaryList {
		wgCerts.Add(1)
		go func(summary *acm.CertificateSummary) {
			defer wgCerts.Done()

			describeInput := &acm.DescribeCertificateInput{
				CertificateArn: summary.CertificateArn,
			}
			descResult, err := svc.DescribeCertificate(describeInput)
			if err != nil {
				fmt.Printf("Error describing certificate in region %s: %v\n", region, err)
				return
			}
			for _, validationOption := range descResult.Certificate.DomainValidationOptions {
				if validationOption.DomainName != nil && validationOption.ValidationMethod != nil {
					fmt.Printf("[Region: %s] %s uses %s for validation\n", region, *validationOption.DomainName, *validationOption.ValidationMethod)
				}
			}
		}(summary)
	}
	wgCerts.Wait()
}

func main() {
	var profile string
	var regions string

	flag.StringVar(&profile, "profile", "default", "AWS CLI profile to use.")
	flag.StringVar(&regions, "regions", "us-west-2", "Comma-separated list of AWS regions to scan.")
	flag.Parse()

	regionList := strings.Split(regions, ",")
	var wg sync.WaitGroup

	for _, region := range regionList {
		wg.Add(1)
		go scanRegion(region, profile, &wg)
	}

	wg.Wait() // Wait for all goroutines to finish
}
