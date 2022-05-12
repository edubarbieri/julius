package main

import (
	"log"
	"regexp"
	"strings"

	"github.com/gocolly/colly/v2"
)

func main() {
	c := colly.NewCollector()

	c.OnHTML("iframe[src]", func(h *colly.HTMLElement) {
		src := h.Attr("src")
		c.Visit(src)
	})

	//Find store name
	c.OnHTML("table.NFCCabecalho td.NFCCabecalho_SubTitulo", func(el *colly.HTMLElement) {
		if el.Attr("align") == "left" {
			log.Printf("store is %v", el.Text)
		}
	})
	//find CNPJ
	c.OnHTML("table.NFCCabecalho td.NFCCabecalho_SubTitulo1", func(el *colly.HTMLElement) {
		if strings.Contains(el.Text, "CNPJ:") {
			re := regexp.MustCompile(`([0-9\./-]{18})`)
			log.Printf("cnpj is %v", re.FindString(strings.ReplaceAll(el.Text, "\n", "")))
		}
	})

	//Find date
	c.OnHTML("table td.NFCCabecalho_SubTitulo", func(el *colly.HTMLElement) {
		if strings.Contains(el.Text, "Emiss") {
			re := regexp.MustCompile(`(\d{2})/(\d{2})/(\d{4}) (\d{2}):(\d{2}):(\d{2})`)
			log.Printf("date is %v", re.FindString(strings.ReplaceAll(el.Text, "\n", "")))
		}
	})

	//Find items
	c.OnHTML("table.NFCCabecalho tr", func(el *colly.HTMLElement) {
		if !strings.HasPrefix(el.Attr("id"), "Item +") {
			return
		}
		log.Println("---------------------------------------")
		log.Printf("item description: %s", el.ChildText(".NFCDetalhe_Item:nth-of-type(2)"))
		log.Printf("item quantity: %s", el.ChildText(".NFCDetalhe_Item:nth-of-type(3)"))
		log.Printf("item unity: %s", el.ChildText(".NFCDetalhe_Item:nth-of-type(4)"))
		log.Printf("item unity price: %s", el.ChildText(".NFCDetalhe_Item:nth-of-type(5)"))
		log.Printf("item total price: %s", el.ChildText(".NFCDetalhe_Item:nth-of-type(6)"))
		log.Println("---------------------------------------")
	})

	c.OnHTML("table td.NFCCabecalho_SubTitulo", func(el *colly.HTMLElement) {
		re := regexp.MustCompile(`([\d\s]{54})`)
		acccessKey := re.FindString(el.Text)
		if acccessKey != "" {
			log.Printf("chave de acesso is %v", acccessKey)
		}
	})

	c.OnHTML("table.NFCCabecalho tr", func(el *colly.HTMLElement) {
		firstTrText := el.ChildText(".NFCDetalhe_Item:nth-of-type(1)")
		if strings.Contains(firstTrText, "Valor total") {
			log.Printf("valor total: %v", el.ChildText(".NFCDetalhe_Item:nth-of-type(2)"))
		} else if strings.Contains(firstTrText, "Valor descontos") {
			log.Printf("valor descontos: %v", el.ChildText(".NFCDetalhe_Item:nth-of-type(2)"))
		}
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
	})

	c.Visit("https://www.sefaz.rs.gov.br/NFCE/NFCE-COM.aspx?p=43220490919200000569650080000895589133243560%7C2%7C1%7C28%7C146.62%7C65798DBD06DB772568CBBFBADC16EC2F9EB78145%7C1%7CB83112B3D6582EC49DF1837606A226254B1A0F83")

}
