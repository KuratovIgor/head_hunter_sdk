package main

import (
	headhunter "github.com/KuratovIgor/head_hunter_sdk"
	"log"
)

func main() {
	client, _ := headhunter.NewClient("P3PD344RTN89887V9SQJB24QNA0U06SPP3JG6TSR59SKMQ1191C8VHRJLC17RO0D", "SMIR5S7AK9P7FMOE4MRLJ291N4939TTM5BO9RGM7NOFSEC9RSMF6CCD50TVFBQUS", "http://localhost")
	token := "VFD0TDCM0JNRMSKHPNRPTHRTOJFP07LDFK3S6E9PHPHJ5L3Q1961GG9051IP0B64"

	resumes := client.GetResumesIds(token)
	log.Println(resumes)

	client.SetToken(token)

	err := client.ApplyToJob("76004367", "e189bab6ff08a10b670039ed1f4b534f383461", "")
	if err != nil {
		log.Println(err)
	}
}
