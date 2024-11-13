package main

import (
	"fmt"
)

// Contract is a general interface for different types of share contracts
type Contract interface {
	Describe() string
	KeyTerms() string
}

// SharePurchaseAgreement represents a Share Purchase Agreement (SPA)
type SharePurchaseAgreement struct {
	Buyer       string
	Seller      string
	ShareAmount int
}

func (spa SharePurchaseAgreement) Describe() string {
	return "Share Purchase Agreement (SPA): A contract where the buyer purchases shares from the seller."
}

func (spa SharePurchaseAgreement) KeyTerms() string {
	return fmt.Sprintf("Buyer: %s, Seller: %s, Shares: %d", spa.Buyer, spa.Seller, spa.ShareAmount)
}

// ShareholdersAgreement represents a Shareholders' Agreement (SHA)
type ShareholdersAgreement struct {
	Shareholders []string
	CompanyName  string
	VotingRights bool
}

func (sha ShareholdersAgreement) Describe() string {
	return "Shareholders' Agreement (SHA): Regulates the relationship between the shareholders and the management of the company."
}

func (sha ShareholdersAgreement) KeyTerms() string {
	voting := "not provided"
	if sha.VotingRights {
		voting = "provided"
	}
	return fmt.Sprintf("Company: %s, Shareholders: %v, Voting Rights: %s", sha.CompanyName, sha.Shareholders, voting)
}

// AssetPurchaseAgreement represents an Asset Purchase Agreement (APA)
type AssetPurchaseAgreement struct {
	Buyer  string
	Seller string
	Asset  string
}

func (apa AssetPurchaseAgreement) Describe() string {
	return "Asset Purchase Agreement (APA): A contract where specific assets of a company are purchased rather than shares."
}

func (apa AssetPurchaseAgreement) KeyTerms() string {
	return fmt.Sprintf("Buyer: %s, Seller: %s, Asset: %s", apa.Buyer, apa.Seller, apa.Asset)
}

// ContractEducation educates about the contract type by printing details
func ContractEducation(c Contract) {
	fmt.Println(c.Describe())
	fmt.Println("Key Terms:", c.KeyTerms())
	fmt.Println()
}

func main() {
	spa := SharePurchaseAgreement{
		Buyer:       "Alice",
		Seller:      "Bob",
		ShareAmount: 100,
	}

	sha := ShareholdersAgreement{
		Shareholders: []string{"Alice", "Bob", "Charlie"},
		CompanyName:  "Tech Corp",
		VotingRights: true,
	}

	apa := AssetPurchaseAgreement{
		Buyer:  "Alice",
		Seller: "Bob",
		Asset:  "Intellectual Property",
	}

	fmt.Println("Educating on Different Types of Share Contracts:\n")
	ContractEducation(spa)
	ContractEducation(sha)
	ContractEducation(apa)
}
