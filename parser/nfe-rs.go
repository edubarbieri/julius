package parser

import (
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/edubarbieri/julius/entity"

	"github.com/gocolly/colly"
)

type NfeRs struct {
	nfe       entity.Nfe
	collector *colly.Collector
}

func NewNfeRs() NfeRs {
	return NfeRs{}
}

func (n NfeRs) Parse(url string) entity.Nfe {
	n.nfe = entity.Nfe{
		Url: url,
	}
	n.collector = colly.NewCollector()

	n.onIframe()
	n.storeName()
	n.cnpj()
	n.date()
	n.items()
	n.accessKey()
	n.nfeTotals()

	n.collector.Visit(url)

	return n.nfe
}
func (n *NfeRs) onIframe() {
	n.collector.OnHTML("iframe[src]", func(h *colly.HTMLElement) {
		src := h.Attr("src")
		n.collector.Visit(src)
	})
}

func (n *NfeRs) storeName() {
	n.collector.OnHTML("table.NFCCabecalho td.NFCCabecalho_SubTitulo", func(el *colly.HTMLElement) {
		if el.Attr("align") == "left" {
			n.nfe.StoreName = el.Text
		}
	})
}

func (n *NfeRs) cnpj() {
	n.collector.OnHTML("table.NFCCabecalho td.NFCCabecalho_SubTitulo1", func(el *colly.HTMLElement) {
		if strings.Contains(el.Text, "CNPJ:") {
			re := regexp.MustCompile(`([0-9\./-]{18})`)
			n.nfe.StoreCnpj = re.FindString(strings.ReplaceAll(el.Text, "\n", ""))
		}
	})
}

func (n *NfeRs) date() {
	n.collector.OnHTML("table td.NFCCabecalho_SubTitulo", func(el *colly.HTMLElement) {
		if strings.Contains(el.Text, "Emiss") {
			re := regexp.MustCompile(`(\d{2})/(\d{2})/(\d{4}) (\d{2}):(\d{2}):(\d{2})`)
			tz, _ := time.LoadLocation("America/Sao_Paulo")
			date, err := time.ParseInLocation("02/01/2006 15:04:05", re.FindString(strings.ReplaceAll(el.Text, "\n", "")), tz)
			if err != nil {
				log.Printf("error parsing nfe date %v", err)
				return
			}
			n.nfe.Date = date
		}
	})
}

func (n *NfeRs) items() {
	n.collector.OnHTML("table.NFCCabecalho tr", func(el *colly.HTMLElement) {
		if !strings.HasPrefix(el.Attr("id"), "Item +") {
			return
		}
		item := entity.NfeItem{}
		item.Description = el.ChildText(".NFCDetalhe_Item:nth-of-type(2)")
		item.UnitOfMeasure = strings.ToUpper(el.ChildText(".NFCDetalhe_Item:nth-of-type(4)"))

		quantity, err := parseFloat(el.ChildText(".NFCDetalhe_Item:nth-of-type(3)"))
		if err != nil {
			log.Printf("error parsing quantity %v", err)
		} else {
			item.Quantity = quantity
		}

		unityPrice, err := parseFloat(el.ChildText(".NFCDetalhe_Item:nth-of-type(5)"))
		if err != nil {
			log.Printf("error parsing unityPrice %v", err)
		} else {
			item.UnityPrice = unityPrice
		}

		totalPrice, err := parseFloat(el.ChildText(".NFCDetalhe_Item:nth-of-type(6)"))
		if err != nil {
			log.Printf("error parsing totalPrice %v", err)
		} else {
			item.TotalPrice = totalPrice
		}
		n.nfe.Items = append(n.nfe.Items, item)
	})
}

func (n *NfeRs) accessKey() {
	n.collector.OnHTML("table td.NFCCabecalho_SubTitulo", func(el *colly.HTMLElement) {
		re := regexp.MustCompile(`([\d\s]{54})`)
		acccessKey := re.FindString(el.Text)
		if acccessKey != "" {
			n.nfe.AccessKey = acccessKey
		}
	})
}

func (n *NfeRs) nfeTotals() {
	n.collector.OnHTML("table.NFCCabecalho tr", func(el *colly.HTMLElement) {
		firstTrText := el.ChildText(".NFCDetalhe_Item:nth-of-type(1)")
		if strings.Contains(firstTrText, "Valor total") {
			total, err := parseFloat(el.ChildText(".NFCDetalhe_Item:nth-of-type(2)"))
			if err != nil {
				log.Printf("error parsing totalPrice %v", err)
			} else {
				n.nfe.Total = total
			}
		} else if strings.Contains(firstTrText, "Valor descontos") {
			discounts, err := parseFloat(el.ChildText(".NFCDetalhe_Item:nth-of-type(2)"))
			if err != nil {
				log.Printf("error parsing discoumts %v", err)
			} else {
				n.nfe.Discount = discounts
			}
		}
	})
}

func parseFloat(rawValue string) (float64, error) {
	strNumber := strings.ReplaceAll(strings.ReplaceAll(rawValue, ".", ""), ",", ".")
	return strconv.ParseFloat(strNumber, 64)
}
