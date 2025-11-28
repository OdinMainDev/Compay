package rates

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"strconv"
	"strings"

	"golang.org/x/text/encoding/charmap"
)

func getUsdToRubRateFromBestChange() (float64, error) {
	// 1) грузим info.zip
	body, err := doRequest(parserCfg["bestchange"])
	if err != nil {
		return 0, fmt.Errorf("bestchange 1: request error: %w", err)
	}

	zr, err := zip.NewReader(bytes.NewReader(body), int64(len(body)))
	if err != nil {
		return 0, fmt.Errorf("bestchange 2: zip open error: %w", err)
	}

	// 2) читаем нужные файлы из архива
	var cyRaw, ratesRaw []byte
	for _, f := range zr.File {
		switch strings.ToLower(f.Name) {
		case "bm_cy.dat":
			rc, err := f.Open()
			if err != nil {
				return 0, fmt.Errorf("bestchange 3: open bm_cy.dat: %w", err)
			}
			cyRaw, err = io.ReadAll(rc)
			rc.Close()
			if err != nil {
				return 0, fmt.Errorf("bestchange 4: read bm_cy.dat: %w", err)
			}
		case "bm_rates.dat":
			rc, err := f.Open()
			if err != nil {
				return 0, fmt.Errorf("bestchange 5: open bm_rates.dat: %w", err)
			}
			ratesRaw, err = io.ReadAll(rc)
			rc.Close()
			if err != nil {
				return 0, fmt.Errorf("bestchange 6: read bm_rates.dat: %w", err)
			}
		}
	}
	if len(cyRaw) == 0 || len(ratesRaw) == 0 {
		return 0, fmt.Errorf("bestchange 7: required files not found in zip")
	}

	// 3) оба файла в CP1251 — декодируем в UTF-8 (как у тебя в PHP)
	// см. твои классы BestChange/Currencies/Rates с iconv('CP1251','UTF-8',...) :contentReference[oaicite:3]{index=3} :contentReference[oaicite:4]{index=4}
	decoder := charmap.Windows1251.NewDecoder()
	cyUtf8, err := decoder.Bytes(cyRaw)
	if err != nil {
		return 0, fmt.Errorf("bestchange 8: decode bm_cy.dat: %w", err)
	}
	ratesUtf8, err := decoder.Bytes(ratesRaw)
	if err != nil {
		return 0, fmt.Errorf("bestchange 9: decode bm_rates.dat: %w", err)
	}

	// 4) построим карты currencyID -> name и выделим наборы id для RUB и USD/USDT
	// Формат bm_cy.dat: id;parent;name;...
	lines := strings.Split(string(cyUtf8), "\n")
	id2name := make(map[int]string)
	for _, ln := range lines {
		ln = strings.TrimSpace(ln)
		if ln == "" {
			continue
		}
		parts := strings.Split(ln, ";")
		if len(parts) < 3 {
			continue
		}
		id, err := strconv.Atoi(strings.TrimSpace(parts[0]))
		if err != nil {
			continue
		}
		name := strings.TrimSpace(parts[2])
		id2name[id] = name
	}

	usdLike := make(map[int]struct{})
	rubLike := make(map[int]struct{})

	for id, name := range id2name {
		up := strings.ToUpper(name)
		if up == "Т-БАНК RUB" {
			rubLike[id] = struct{}{}
			continue
		}
		if up == "TETHER TRC20 (USDT)" {
			usdLike[id] = struct{}{}
			continue
		}
	}

	if len(usdLike) == 0 || len(rubLike) == 0 {
		return 0, fmt.Errorf("bestchange 10: can't pick USD/USDT or RUB IDs")
	}

	var best float64
	best = 0
	for _, ln := range strings.Split(string(ratesUtf8), "\n") {
		ln = strings.TrimSpace(ln)
		if ln == "" {
			continue
		}
		parts := strings.Split(ln, ";")
		if len(parts) < 5 {
			continue
		}
		recvID, _ := strconv.Atoi(parts[0])
		giveID, _ := strconv.Atoi(parts[1])

		if _, ok := rubLike[recvID]; !ok {
			continue
		}
		if _, ok := usdLike[giveID]; !ok {
			continue
		}

		rg, _ := strconv.ParseFloat(strings.TrimSpace(parts[3]), 64)
		rr, _ := strconv.ParseFloat(strings.TrimSpace(parts[4]), 64)
		if rg <= 0 || rr <= 0 {
			continue
		}
		rate := rg / rr
		if rate <= 0 {
			continue
		}

		if best == 0 || rate < best {
			best = rate
		}
	}

	if best == 0 {
		return 0, fmt.Errorf("bestchange 11: no USD/USDT→RUB rates found")
	}
	return best, nil
}
