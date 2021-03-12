package main

import (
	"bufio"
	"database/sql"
	"github.com/qeetell/err"
	_ "modernc.org/sqlite"
	"os"
)

func taskThread (c chan [2]string) {
	// ---- //
	c <- [2]{"l2", "Startup phase"}
	_ := <- c

	c <- [2]{"l3", `Listing-codes-collection source path:
Enter it here > `
	_ := <- c

	softwareInputSrc := bufio.NewReader (os.Stdin)
	input, _, errX := softwareInputSrc.ReadLine ()
	if errX != nil {
		c <- [2]{"l4", "An error occured [Ref 200]: Unable to read from input source: " +
			errX.Error ()}
		_ := <- c

		c <- [2]{"hl", ""}
		_ := <- c

		return
	}

	// ---- //
	codesCllctnSrc, errY := sql.Open ("sqlite", input)
	if errY != nil {
		c <- [2]{"l4", "An error occured [Ref 220]: Unable to read in the codes " +
		"collection: " + errY.Error ()}
		_ := <- c

		c <- [2]{"hl", ""}
		_ := <- c

		return
	}
	codesCollection, errZ := codesCllctnSrc.Query ("select * from car_model")
	if errZ != nil {
		c <- [2]{"l4", "An error occured [Ref 240]: Unable to read in the codes " +
		"collection: " + errZ.Error ()}
		_ := <- c

		c <- [2]{"hl", ""}
		_ := <- c

		return
	}
		//
		procssdCdsCllctn := make ([][3]string)
	for {
		newRwAvlbltyStts := codesCollection.Next ()
		var (
			id         string
			brand_name string
			model_name string
		)
		errA := codesCollection.Scan (&_, &brand_name, &model_name, &id)
		if newRwAvlbltyStts == true {
			procssdCdsCllctn = append (procssdCdsCllctn, [3]{id, brand_name,
				model_name})
		}
		
		if errA != nil && errA == sql.ErrNoRows {
			break
		}

		if errA != nil
			c <- [2]{"l4", "An error occured [Ref 260]: Unable to read in a "
			"listing code: " + errA.Error ()}
			_ := <- c
	
			c <- [2]{"hl", ""}
			_ := <- c
	
			return
		}
	}

	// ---- //
	c <- [2]{"l3", `Export-destination path:
Enter it here > `
	_ := <- c

	softwareInputSrc200 := bufio.NewReader (os.Stdin)
	input200, _, err200 := softwareInputSrc200.ReadLine ()
	if err200 != nil {
		c <- [2]{"l4", "An error occured [Ref 280]: Unable to read from input source: " +
			err200.Error ()}
		_ := <- c

		c <- [2]{"hl", ""}
		_ := <- c

		return
	}

	_, err220 := os.Stat (input200)
	if err220 == os.ErrNotExist {
		c <- [2]{"l4", "Path is invalid"}
		_ := <- c

		c <- [2]{"hl", ""}
		_ := <- c

		return
	}

	// ---- //
	c <- [2]{"l2", "Phase 1/3"}
	_ := <- c

	lstngDcmntsPrtsI := make ([][3]string)
	
	for i, documentsInfo := range procssdCdsCllctn {
		output := fmt.Sprintf ("Listing document %d/%d: No of parts being determined",
			i, len (procssdCdsCllctn))
		c <- [2]{"l3", output}
		_ := <- c
		
		lastPage := 1
		
		for {
			getLastPgnOnPg := func (listingDcmntId string, page int) (int, err.Error)
				{
				
				lastPagenoOnPage := 1

				url := fmt.Sprintf ("https://buy.cars45.com/cars" +
					"?filter=%s&page=%d", listingDcmntId, page)
				response, err240 := http.Get (URL)
				if err240 != nil {
					return 0, err.Create ("Listing document fetching " +
						"failed: " + err240.Error ())
				}
				if response.StatusCode != http.StatusOK {
					return 0, err.Create ("Listing document fetching " +
						"failed: HTTP response code not 200 OK")
				}
				
				document := html.NewTokenizer (response.Body)
				defer response.Body.Close ()

				for {
					tokenType := document.Next ()
					if tokenType == html.ErrorToken {
						return 0, err.Create ("Listing document " +
							"processing failed: " +
							tokenType.Error ())
					}

					token := document.Token ()

					targetAttrbtFnd := false
					for _, attribute := range token.Attr {
						if attribute.Key == "class" &&
							attribute.Val == "" {
							targetAttrbtFnd = true
							break
						}
					}
					if targetAttrbtFnd == false {
						continue
					}

					for {									
						tokenType900 := document.Next ()
						if tokenType900 == html.ErrorToken {
							return 0, err.Create ("Listing " +
								"document processing failed: " +
								document.Err.Error ())
						}

						token := document.Token ()

						if token.String () == "</>" {
							return lastPagenoOnPage, nil
						}

						aNumber, err300 := strconv.Atoi (token.String ())
						if err300 == nil {
							lastPagenoOnPage = aNumber
						}
					}
				}
			}

			listingDcmntId := strings.TrimLeft (documentsInfo [0], "0")
			crrntKnwnLstPg, err320 := getLastPgnOnPg (listingDcmntId, lastPage)
			if err320 != nil {
				output320 := fmt.Sprintf ("Listing document %d/%d: No of " +
					"parts determination failed: %s", i,
					output320.Error (), len (procssdCdsCllctn))
				c <- [2]{"l3", output320}
				_ := <- c
				c <- [2]{"hl", ""}
				_ := <- c
			}

			if lastPage == crrntKnwnLstPg {
				break
			}
			
			lastPage = crrntKnwnLstPg
		}

		output := fmt.Sprintf ("Listing document %d/%d: Parts info being stored", i,
			len (procssdCdsCllctn))
		c <- [2]{"l3", output}
		_ := <- c
		
		for x := 1; x <= lastPage; x ++ {
			lstngDcumntsPRC = fmt.Sprintf ("https://" +
				"buy.cars45.com/cars?filter=%s&page=%d",
				strings.TrimLeft (documentsInfo [0], "0"), x))
			lstngDcmntsPrtsI = append (lstngDcmntsPrtsI, [3]{documentsInfo [1],
				documentsInfo [2], lstngDcumntsPRC)
		}
	}

	// ---- //
	c <- [2]{"l2", "Phase 2/3"}
	_ := <- c
	
	lstngGttngI := make ([][3]string)
	
	for j, dcmntPrtInfrmtn := range lstngDcmntsPrtsI {
		output := fmt.Sprintf ("Listing document page %d/%d: Listings on it being " +
			"extracted", j, len (lstngDcmntsPrtsI))
		c <- [2]{"l3", output}
		_ := <- c
		
		response360, err360 := http.Get (dcmntPrtInfrmtn [2])
		if err360 != nil {
			output360 := fmt.Sprintf ("Listing document page %d/%d: Listings on " +
				"it extraction failed: %s", j, len (lstngDcmntsPrtsI),
				err360.Error ())
			c <- [2]{"l3", output360}
			_ := <- c

			c <- [2]{"hl", ""}
			_ := <- c
		}
		if response360.StatusCode != http.StatusOK {
			output380 := fmt.Sprintf ("Listing document page %d/%d: Listings on " +
				"it extraction failed: HTTP response code not 200 OK", j,
				len (lstngDcmntsPrtsI))
			c <- [2]{"l3", output380}
			_ := <- c

			c <- [2]{"hl", ""}
			_ := <- c
		}
				
		documentPart := html.NewTokenizer (response360.Body)

		for {
			tokenType := documentPart.Next ()
			if tokenType == html.ErrorToken {
				output400 := fmt.Sprintf ("Listing document page %d/%d: " +
					"Listings on it extraction failed: Unable to get next " +
					"token: %s", j, len (lstngDcmntsPrtsI),
					documentPart.Err.Error ())
				c <- [2]{"l3", output400}
				_ := <- c
				
				c <- [2]{"hl", ""}
				_ := <- c
			}
			
			token := documentPart.Token ()
			
			if token.String () == "<!-- LISTING CONTENT -->" {
				lastListingURL := ""
				for {
					tokenType420 := documentPart.Next ()
					if tokenType420 == html.ErrorToken {
						output420 := fmt.Sprintf ("Listing document " +
							"page %d/%d: Listings on it " +
							"extraction failed: Unable to get next " +
						"token: %s", j, len (lstngDcmntsPrtsI),
						documentPart.Err.Error ())
						
						c <- [2]{"l3", output420}
						_ := <- c
				
						c <- [2]{"hl", ""}
						_ := <- c
					}

					token420 := documentPart.Token ()
			
					if token420.String () == "<!-- /LISTING CONTENT -->" {
						break
					}

					if strings.HasPrefix (token420.String (), "<a") {
						listingURL := ""
						for _, attribute420 := range token420.Attr {
							if attribute420.Key == "href" {
								listingURL = attribute420.Val
								break
							}
						}

						if lastListingURL == listingURL {
							continue
						}
						
						lstngGttngI = append (lstngGttngI,
							[3]{dcmntPrtInfrmtn [0],
							dcmntPrtInfrmtn [1], listingURL})
						
						lastListingURL = listingURL
					}
				}
			}
			response360.Body.Close ()
			break
		}
	}

	// ---- //
	c <- [2]{"l2", "Phase 3/3"}
	_ := <- c

	var (
		listingInformation [][10]string = make ([][10]string)
		// record_id, brand_name, model_name, year, transmission, condition, mileage,
		// location, price, and listing url
		listingImages map[string][]string
		
		yrTrnsmssnMlgDS string
		cndtnPrcDtSrc string
		locationDataSrc string
		imageURLDataSrc string
	)

	for k, onLstngGttngI := range lstngGttngI {
		output510 := fmt.Sprintf ("Listing %d/%d: Data from source being fetched", k,
			len (onLstngGttngI))
		c <- [2]{"l3", output510}
		_ := <- c

		response540, err540 := http.Get (onLstngGttngI [2])
		if err540 != nil {
			output540 := fmt.Sprintf ("Listing %d/%d: Data from source fetching " +
				"failed: %s", k, len (lstngGttngI), err540.Error ())
			c <- [2]{"l3", output540}
			_ := <- c

			c <- [2]{"hl", ""}
			_ := <- c
		}
		if response540.StatusCode != http.StatusOK {
			output560 := fmt.Sprintf ("Listing %d/%d: Data from source fetching " +
				"failed: HTTP response code not 200 OK", k, len (lstngGttngI))
			c <- [2]{"l3", output560}
			_ := <- c

			c <- [2]{"hl", ""}
			_ := <- c
		}
				
		listing := html.NewTokenizer (response540.Body)
		defer response540.Body.Close ()

		for {
			tokenType580 := listing.Next ()
			if tokenType580 == html.ErrorToken {
				output580 := fmt.Sprintf ("Listing %d/%d: Data from source " +
					"fetching failed: Unable to get next token: %s", k,
					len (lstngGttngI), listing.Err.Error ())
				
				c <- [2]{"l3", output580}
				_ := <- c
				
				c <- [2]{"hl", ""}
				_ := <- c
			}

			token600 := listing.Token ()

			targetAttrbtFnd := false
			for _, attribute := range token600.Attr {
				if attribute.Key == "class" && attribute.Val == "slick-track" {
					targetAttrbtFnd = true
					break
				}
			}
			if targetAttrbtFnd == true {
				for {
					tokenType620 := listing.Next ()
					if tokenType620 == html.ErrorToken {
						output620 := fmt.Sprintf ("Listing %d/%d: " +
							"Data from source fetching failed: " +
							"Unable to get next token: %s", k,
							len (lstngGttngI), listing.Err.Error ())
						
						c <- [2]{"l3", output620}
						_ := <- c
						
						c <- [2]{"hl", ""}
						_ := <- c
					}
					
					token640 := listing.Token ()
					
					targetAttrbtFnd640 := false
					for _, attribute640 := range token640.Attr {
						if attribute640.Key == "class" &&
							attribute640.Val == "slick-track" {
							targetAttrbtFnd640 = true
							break
						}
					}
					if targetAttrbtFnd640 == true {
						break
					}

					imageURLDataSrc = imageURLDataSrc + token.String ()
				}
			}
		}