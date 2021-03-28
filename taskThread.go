package main

import (
	"bufio"
	"database/sql"
	err "github.com/qeetell/2103-0918-58"
	"github.com/qamarian-lib/str"
	"fmt"
	"golang.org/x/net/html"
	"io"
	"net/http"
	_ "modernc.org/sqlite"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"os"
	"runtime"
)

func taskThread (c chan [2]string) {
	// ---- //
	c <- [2]string{"l2", "Startup phase"}
	_ = <- c

	c <- [2]string{"l3", `Listing-codes-collection source path:
Enter it here > `}
	_ = <- c

	softwareInputSrc := bufio.NewReader (os.Stdin)
	input, _, errX := softwareInputSrc.ReadLine ()
	if errX != nil {
		c <- [2]string{"l4", "An error occured [Ref 200]: Unable to read from input source: " +
			errX.Error ()}
		_ = <- c

		c <- [2]string{"hl", ""}
		_ = <- c

		return
	}

	// ---- //
	codesCllctnSrc, errY := sql.Open ("sqlite", string (input))
	if errY != nil {
		c <- [2]string{"l4", "An error occured [Ref 220]: Unable to read in the codes " +
		"collection: " + errY.Error ()}
		_ = <- c

		c <- [2]string{"hl", ""}
		_ = <- c

		return
	}
	codesCollection, errZ := codesCllctnSrc.Query ("select * from car_model")
	if errZ != nil {
		c <- [2]string{"l4", "An error occured [Ref 240]: Unable to read in the codes " +
		"collection: " + errZ.Error ()}
		_ = <- c

		c <- [2]string{"hl", ""}
		_ = <- c

		return
	}
		//
		procssdCdsCllctn := make ([][3]string, 0, 5)
	for {
		newRwAvlbltyStts := codesCollection.Next ()
		var (
			id         string
			brand_name string
			model_name string
			discard    string
		)
		codesCollection.Scan (&discard, &brand_name, &model_name, &id)
		_ = discard
		if newRwAvlbltyStts == true {
			procssdCdsCllctn = append (procssdCdsCllctn, [3]string{id, brand_name,
				model_name})
		}
		
		if newRwAvlbltyStts == false {
			break
		}

		/*
		if errA != nil {
			c <- [2]string{"l4", "An error occured [Ref 260]: Unable to read in a " +
				"listing code: " + errA.Error ()}
			_ = <- c
	
			c <- [2]string{"hl", ""}
			_ = <- c
	
			return
		}
		*/
	}

	// ---- //
	c <- [2]string{"l3", `Export-destination path:
Enter it here > `}
	_ = <- c

	softwareInputSrc200 := bufio.NewReader (os.Stdin)
	input200, _, err200 := softwareInputSrc200.ReadLine ()
	if err200 != nil {
		c <- [2]string{"l4", "An error occured [Ref 280]: Unable to read from input source: " +
			err200.Error ()}
		_ = <- c

		c <- [2]string{"hl", ""}
		_ = <- c

		return
	}

	_, err220 := os.Stat (string (input200))
	if err220 == os.ErrNotExist {
		c <- [2]string{"l4", "Path is invalid"}
		_ = <- c

		c <- [2]string{"hl", ""}
		_ = <- c

		return
	}

	// ---- //
	c <- [2]string{"l2", "Phase 1/3"}
	_ = <- c
	
	lstngDcmntsPrtsI := [][3]string {}

	fetchAFromCache := true

if fetchAFromCache == true {
	/* ---- */
	
	dataA, errA23 := os.ReadFile ("/home/qeetell/Temp/a.txt")
	if errA23 != nil {
		outputA23 := fmt.Sprintf ("Listing documents parts info: Fetching from cache " +
			"failed: " + errA23.Error ())
		c <- [2]string{"l3", outputA23}
		_ = <- c

		c <- [2]string{"hl", ""}
		_ = <- c

		return
	}
	
	dataA2 := strings.Split (string (dataA), "!")
	for _, currentA3 := range dataA2 {
		vimA4 := strings.Split (currentA3, "~")
		lstngDcmntsPrtsI = append (lstngDcmntsPrtsI, [3]string {vimA4 [0], vimA4 [1],
			vimA4 [2]})
	}	

} else {
	
	for i, documentsInfo := range procssdCdsCllctn {
		output := fmt.Sprintf ("Listing document %d/%d: No of parts being determined",
			i + 1, len (procssdCdsCllctn))
		c <- [2]string{"l3", output}
		_ = <- c
		
		lastPage := 1
		
		for {
			getLastPgnOnPg := func (listingDcmntId string, page int) (int,
				error) {
				
				lastPagenoOnPage := 1

				url := fmt.Sprintf ("https://buy.cars45.com/cars" +
					"?filter=%s&page=%d", listingDcmntId, page)
				response, err240 := http.Get (url)

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
						if document.Err () == io.EOF {
							return lastPage, nil
						}
						
						return 0, err.Create ("Listing document " +
							"processing failed 1: " +
							document.Err ().Error ())
					}

					token := document.Token ()

					targetAttrbtFnd := false
					for _, attribute := range token.Attr {
						if attribute.Key == "class" &&
							attribute.Val == "pagination" {
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
								"document processing failed 2: " +
								document.Err ().Error ())
						}

						token := document.Token ()

						if token.String () == "</ul>" {
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
					"parts determination failed: %s", i + 1,
					len (procssdCdsCllctn), err320.Error ())
				c <- [2]string{"l3", output320}
				_ = <- c
				c <- [2]string{"hl", ""}
				_ = <- c
				return
			}

			if lastPage == crrntKnwnLstPg {
				break
			}
			
			lastPage = crrntKnwnLstPg
		}

		output65X := fmt.Sprintf ("Listing document %d/%d: Parts info being stored",
			i + 1, len (procssdCdsCllctn))
		c <- [2]string{"l3", output65X}
		_ = <- c
		
		for x := 1; x <= lastPage; x ++ {
			lstngDcumntsPRC := fmt.Sprintf ("https://" +
				"buy.cars45.com/cars?filter=%s&page=%d",
				strings.TrimLeft (documentsInfo [0], "0"), x)
			lstngDcmntsPrtsI = append (lstngDcmntsPrtsI, [3]string{documentsInfo [1],
				documentsInfo [2], lstngDcumntsPRC})
		}
	}
}

	serialA := ""
	for _, x1234 := range lstngDcmntsPrtsI {
		z123 := fmt.Sprintf ("%s~%s~%s!", x1234 [0], x1234 [1], x1234 [2])
		serialA = serialA + z123
	}
	serialA = serialA + "!"
	serialA = strings.Replace (serialA, "!!", "", -1)
	os.WriteFile ("/home/qeetell/Temp/a.txt", []byte (serialA), 0666)

	// ---- //
	c <- [2]string{"l2", "Phase 2/3"}
	_ = <- c
	
	lstngGttngI := [][3]string {}

	fetchBFromCache := true

if fetchBFromCache == true {
	/* ---- */
	
	dataB, errB23 := os.ReadFile ("/home/qeetell/Temp/b.txt")
	if errB23 != nil {
		outputB23 := fmt.Sprintf ("Listing document info [B]: Fetching from cache " +
			"failed: " + errB23.Error ())
		c <- [2]string{"l3", outputB23}
		_ = <- c

		c <- [2]string{"hl", ""}
		_ = <- c

		return
	}
	
	dataB2 := strings.Split (string (dataB), "!")
	for _, currentB3 := range dataB2 {
		vimB4 := strings.Split (currentB3, "~")
		lstngGttngI = append (lstngGttngI, [3]string {vimB4 [0], vimB4 [1],
			vimB4 [2]})
	}	

} else {
	
	for j, dcmntPrtInfrmtn := range lstngDcmntsPrtsI {
		output := fmt.Sprintf ("Listing document page %d/%d: Listings on it being " +
			"extracted", j + 1, len (lstngDcmntsPrtsI))
		c <- [2]string{"l3", output}
		_ = <- c
		
		response360, err360 := http.Get (dcmntPrtInfrmtn [2])
		fmt.Println (dcmntPrtInfrmtn [2])
		
		if err360 != nil {
			output360 := fmt.Sprintf ("Listing document page %d/%d: Listings on " +
				"it extraction failed: %s", j + 1, len (lstngDcmntsPrtsI),
				err360.Error ())
			c <- [2]string{"l3", output360}
			_ = <- c

			c <- [2]string{"hl", ""}
			_ = <- c

			return
		}
		
		if response360.StatusCode != http.StatusOK {
			output380 := fmt.Sprintf ("Listing document page %d/%d: Listings on " +
				"it extraction failed: HTTP response code not 200 OK ... %d",
				j + 1, len (lstngDcmntsPrtsI), response360.StatusCode)
			c <- [2]string{"l3", output380}
			_ = <- c

			c <- [2]string{"hl", ""}
			_ = <- c

			return
		}
				
		documentPart := html.NewTokenizer (response360.Body)
		
		for {
			tokenType := documentPart.Next ()
			if tokenType == html.ErrorToken {
				if documentPart.Err () == io.EOF {
					break
				}
				output400 := fmt.Sprintf ("Listing document page %d/%d: " +
					"Listings on it extraction failed: Unable to get next " +
					"token: %s", j + 1, len (lstngDcmntsPrtsI),
					documentPart.Err ().Error ())
				c <- [2]string{"l3", output400}
				_ = <- c
				c <- [2]string{"hl", ""}
				_ = <- c
				return
			}
			
			token := documentPart.Token ()
			done := false

			if token.String () == "<!-- LISTING CONTENT -->" {
				lastListingURL := ""
				for {
					tokenType420 := documentPart.Next ()
					if tokenType420 == html.ErrorToken {
						if documentPart.Err () == io.EOF {
							break
						}
						output420 := fmt.Sprintf ("Listing document " +
							"page %d/%d: Listings on it " +
							"extraction failed: Unable to get next " +
							"token: %s", j + 1, len (lstngDcmntsPrtsI),
							documentPart.Err ().Error ())
						
						c <- [2]string{"l3", output420}
						_ = <- c
						c <- [2]string{"hl", ""}
						_ = <- c
						return
					}

					token420 := documentPart.Token ()
			
					if token420.String () == "<!-- row -->" {
						done = true
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
							[3]string{dcmntPrtInfrmtn [0],
							dcmntPrtInfrmtn [1], listingURL})
						
						lastListingURL = listingURL
					}
				}
			}
			
			if done == true {
				break
			}
		}
		response360.Body.Close ()
	}
}

	serialB := ""
	for _, x2345 := range lstngGttngI {
		z234 := fmt.Sprintf ("%s~%s~%s!", x2345 [0], x2345 [1], x2345 [2])
		serialB = serialB + z234
	}
	serialB = serialB + "!"
	serialB = strings.Replace (serialB, "!!", "", -1)
	os.WriteFile ("/home/qeetell/Temp/b.txt", []byte (serialB), 0666)
	
	// ---- //
	c <- [2]string{"l2", "Phase 3/3"}
	_ = <- c

	var (
		listingInformation [][9]string = make ([][9]string, 0, 2)
		// record_id, brand_name, model_name, year, transmission, condition, mileage,
		// location, price, and listing url
		listingImages map[int][]string = map[int][]string {}
		
		yrTrnsmssnMlgDS string
		cndtnPrcDtSrc string
		locationDataSrc string
		imageURLDataSrc string
	)

	fetchCFromCache := true

if fetchCFromCache == true {
	/* ---- */
	
	dataC, errC23 := os.ReadFile ("/home/qeetell/Temp/c.txt")
	if errC23 != nil {
		outputC23 := fmt.Sprintf ("Listing document: Fetching from cache " +
			"failed: " + errC23.Error ())
		c <- [2]string{"l3", outputC23}
		_ = <- c

		c <- [2]string{"hl", ""}
		_ = <- c

		return
	}
	
	dataC2 := strings.Split (string (dataC), "!")
	for _, currentC3 := range dataC2 {
		vimC4 := strings.Split (currentC3, "~")
		listingInformation = append (listingInformation, [9]string {vimC4 [0], vimC4 [1], vimC4 [2],
			vimC4 [3], vimC4 [4], vimC4 [5], vimC4 [6], vimC4 [7], vimC4 [8]})
	}	
	
	dataC23BB, errC23BB := os.ReadFile ("/home/qeetell/Temp/c2.txt")
	if errC23BB != nil {
		outputC23BB := fmt.Sprintf ("Listing document: Fetching from cache " +
			"failed: " + errC23BB.Error ())
		c <- [2]string{"l3", outputC23BB}
		_ = <- c

		c <- [2]string{"hl", ""}
		_ = <- c

		return
	}

	dataC2BB := strings.Split (string (dataC23BB), "!")
	for tv, currentC3BB := range dataC2BB {
		huey := strings.Split (currentC3BB, "~")
		listingImages [tv + 1] = huey
	}

} else {
	start := 4401
	complete := 400


	completed := 0
	for k, onLstngGttngI := range lstngGttngI {
		if (k + 1) < start {
			runtime.Gosched ()
			continue
		}

		if completed == complete {
			break
		}

		completed ++

		yrTrnsmssnMlgDS = ""
		cndtnPrcDtSrc = ""
		locationDataSrc = ""
		imageURLDataSrc = ""

		output510 := fmt.Sprintf ("Listing %d/%d: Data from source being fetched", k + 1,
			complete)
		c <- [2]string{"l3", output510}
		_ = <- c

		response540, err540 := http.Get (onLstngGttngI [2])
		fmt.Println (onLstngGttngI [2])
		if err540 != nil {
			output540 := fmt.Sprintf ("Listing %d/%d: Data from source fetching " +
				"failed 0: %s", k + 1, complete, err540.Error ())
			c <- [2]string{"l3", output540}
			_ = <- c

			c <- [2]string{"hl", ""}
			_ = <- c

			return
		}
		if response540.StatusCode == http.StatusNotFound ||
			response540.StatusCode == http.StatusInternalServerError {
			continue
		}
		if response540.StatusCode != http.StatusOK {
			output560 := fmt.Sprintf ("Listing %d/%d: Data from source fetching " +
				"failed: HTTP response code not 200 OK: %d", k + 1,
				complete, response540.StatusCode)
			c <- [2]string{"l3", output560}
			_ = <- c

			c <- [2]string{"hl", ""}
			_ = <- c

			return
		}
				
		listing := html.NewTokenizer (response540.Body)

		for {
			tokenType580 := listing.Next ()
			if tokenType580 == html.ErrorToken {
				output580 := fmt.Sprintf ("Listing %d/%d: Data from source " +
					"fetching failed 1: Unable to get next token: %s", k + 1,
					complete, listing.Err ().Error ())
				
				c <- [2]string{"l3", output580}
				_ = <- c
				
				c <- [2]string{"hl", ""}
				_ = <- c

				return
			}
			
			token600 := listing.Token ()

/*
			fmt.Print (token600.String ())
			
			targetAttrbtFnd := false
			for _, attribute := range token600.Attr {
				if attribute.Key == "class" && attribute.Val == "slick-track" {
					targetAttrbtFnd = true
					break
				}
			}
*/			
			if token600.String () == "<!-- Detail slider -->" {
				for {
					tokenType620 := listing.Next ()
					if tokenType620 == html.ErrorToken {
						output620 := fmt.Sprintf ("Listing %d/%d: " +
							"Data from source fetching failed 2: " +
							"Unable to get next token: %s", k + 1,
							complete, listing.Err ().Error ())
						
						c <- [2]string{"l3", output620}
						_ = <- c
						
						c <- [2]string{"hl", ""}
						_ = <- c

						return
					}
					token640 := listing.Token ()

					// <!-- detail_slider_wrap -->"
					if token640.String () == "</button>" {
						break
					}

					imageURLDataSrc = imageURLDataSrc + token640.String ()
				}
				break
			}
		}

/*		for {
			tokenType660 := listing.Next ()
			if tokenType660 == html.ErrorToken {
				output660 := fmt.Sprintf ("Listing %d/%d: Data from source " +
					"fetching failed 3: Unable to get next token: %s", k + 1,
					len (lstngGttngI), listing.Err ().Error ())
				
				c <- [2]string{"l3", output660}
				_ = <- c
				
				c <- [2]string{"hl", ""}
				_ = <- c

				return
			}
			
			token660 := listing.Token ()

			if token660.String () == "<!-- THIS PART HIDES ON MOBILE VERSION " +
				"(ANOTHER COPY IS UNDER DETAIL PAGE SLIDER) -->" {
				break
			}
		}
*/

		rcrdYrTMDS := false
		for {
			tokenType680 := listing.Next ()
			if tokenType680 == html.ErrorToken {
				output680 := fmt.Sprintf ("Listing %d/%d: Data from source " +
					"fetching failed 6: Unable to get next token: %s", k + 1,
					complete, listing.Err ().Error ())
				
				c <- [2]string{"l3", output680}
				_ = <- c
				
				c <- [2]string{"hl", ""}
				_ = <- c

				return
			}
			
			token680 := listing.Token ()

			if token680.String () == "<!-- /car_views_label_wrap -->" {
				break
			}
			
			cndtnPrcDtSrc = cndtnPrcDtSrc + token680.String ()
			
			if rcrdYrTMDS == true {
				yrTrnsmssnMlgDS = token680.String ()
				rcrdYrTMDS = false
			}
			
			for _, attribute700 := range token680.Attr {
				if attribute700.Key == "class" &&
					attribute700.Val == "text-fourteen text-muted mb-2" {
					rcrdYrTMDS = true
					break
				}
			}
		}

		for {
			/*
			tokenType720 := listing.Next ()
			if tokenType720 == html.ErrorToken {
				if listing.Err () == io.EOF {
					break
				}
				output720 := fmt.Sprintf ("Listing %d/%d: Data from source " +
					"fetching failed 4: Unable to get next token: %s", k + 1,
					len (lstngGttngI), listing.Err ().Error ())
				
				c <- [2]string{"l3", output720}
				_ = <- c
				
				c <- [2]string{"hl", ""}
				_ = <- c

				return
			}
			
			token720 := listing.Token ()

			/*
			targetAttrbtFnd740 := false
			for _, attribute740 := range token720.Attr {
				if attribute740.Key == "class" &&
					attribute740.Val == "information_list" {
					fmt.Println (9, token720.String ())
					targetAttrbtFnd740 = true
					break
				}
			}
			*/
			if true {
				for {
					tokenType760 := listing.Next ()
					if tokenType760 == html.ErrorToken {
						output760 := fmt.Sprintf ("Listing %d/%d: Data " +
							"from source fetching failed 5: Unable " +
							"to get next token: %s", k + 1,
							complete,
							listing.Err ().Error ())
						
						c <- [2]string{"l3", output760}
						_ = <- c
						
						c <- [2]string{"hl", ""}
						_ = <- c

						return
					}
					token760 := listing.Token ()

					if token760.String () == "<!-- tab-pane -->" {
						break
					}

					locationDataSrc = locationDataSrc + token760.String ()
				}
				break
			}
		}

				
		response540.Body.Close ()

		output820 := fmt.Sprintf ("Listing %d/%d: Data from source being saved", k + 1,
			len (lstngGttngI))
		c <- [2]string{"l3", output820}
		_ = <- c

		var (
			year string
			transmission string
			condition string
			mileage string
			location string
			price string
		)

		/*
		fmt.Println (1, yrTrnsmssnMlgDS)
		fmt.Println (2, cndtnPrcDtSrc)
		fmt.Println (3, locationDataSrc)
		fmt.Println (4, imageURLDataSrc)*/
		
		
		r := regexp.MustCompile (`\d{4,4}\s*•\s*\w+\s*•\s*[\d,]+`)
		importantPart := r.FindString (yrTrnsmssnMlgDS)
		importantPart = strings.ReplaceAll (importantPart, " ", "")
		importantPart = strings.ReplaceAll (importantPart, "	", "")
		importantPart = strings.ReplaceAll (importantPart, "\r\n", "")
		importantPart = strings.ReplaceAll (importantPart, "\n", "")
		importantPart = strings.ReplaceAll (importantPart, ",", "")
		threeData := strings.Split (importantPart, "•")

		if len (threeData) != 3 {
			continue
		}
		
		year = threeData [0]
		transmission = strings.Title (threeData [1])
		mileage = threeData [2]

		r780 := regexp.MustCompile (`((Foreign|Nigerian) used|New)`)
		condition = r780.FindString (cndtnPrcDtSrc)

		r790 := regexp.MustCompile (`₦ [\d,]+`)
		price = r790.FindString (cndtnPrcDtSrc)
		price = strings.ReplaceAll (price, "₦", "")
		price = strings.ReplaceAll (price, " ", "")
		price = strings.ReplaceAll (price, ",", "")
		fmt.Println (price)

		r800 := regexp.MustCompile (`C45( ,)?[\s\w]+`)
		location = r800.FindString (locationDataSrc)
		location = strings.ReplaceAll (location, "C45", "")
		location = strings.ReplaceAll (location, " ", "")
		location = strings.ReplaceAll (location, ",", "")
		location = strings.ReplaceAll (location, "	", "")
		location = strings.ReplaceAll (location, "\r\n", "")
		location = strings.ReplaceAll (location, "\n", "")

		r810 := regexp.MustCompile (` src="https:\/\/buy\.cars45\.com[\d\-_\/\w]+\.` +
			`(jpg|jpeg|png)`)
		images := r810.FindAllString (imageURLDataSrc, -1)
		lastImage := ""
		newImages := []string {}
		for _, image := range images {
			if len (newImages) == 12 {
				break
			}
			newImage := strings.ReplaceAll (image, ` src="`, "")
			if lastImage == newImage {
				continue
			}
			newImages = append (newImages, newImage)
			lastImage = newImage
		}
		images = newImages
		
		listingInformation = append (listingInformation, [9]string{
			onLstngGttngI [0],
			onLstngGttngI [1],
			year,
			transmission,
			condition,
			mileage,
			location,
			price,
			onLstngGttngI [2],
		})
		
		listingImages [len (listingInformation)] = images
	}
}

	serialC := ""
	for i943, x3456 := range listingInformation {
		outputV29 := fmt.Sprintf ("Listing %d/%d: Processing for backup", i943 + 1,
			len (listingInformation))	
		c <- [2]string{"l3", outputV29}
		_ = <- c
		
		z345 := fmt.Sprintf ("%s~%s~%s~%s~%s~%s~%s~%s~%s!", x3456 [0], x3456 [1], x3456 [2],
			x3456 [3], x3456 [4], x3456 [5], x3456 [6], x3456 [7], x3456 [8])
		serialC = serialC + z345
	}

	outputV30 := fmt.Sprintf ("Listings: Backing up")	
	c <- [2]string{"l3", outputV30}
	_ = <- c

	serialC = serialC + "!"
	serialC = strings.Replace (serialC, "!!", "", -1)
	os.WriteFile ("/home/qeetell/Temp/c.txt", []byte (serialC), 0666)

	serialC2 := ""
	for m3 := 1; m3 <= len (listingImages); m3 ++ {
		outputV31 := fmt.Sprintf ("Listing pictures set %d/%d: Processing for backup",
			m3, len (listingImages))	
		c <- [2]string{"l3", outputV31}
		_ = <- c
		
		for _, x7890 := range listingImages [m3] {
			serialC2 = serialC2 + x7890 + "~"
		}
		serialC2 = serialC2 + "~"
		serialC2 = strings.Replace (serialC2, "~~", "", -1)
		serialC2 = serialC2 + "!"
	}

	outputV32 := fmt.Sprintf ("Listing pictures sets: Backing up")	
	c <- [2]string{"l3", outputV32}
	_ = <- c

	serialC2 = serialC2 + "!"
	serialC2 = strings.Replace (serialC2, "!!", "", -1)
	os.WriteFile ("/home/qeetell/Temp/c2.txt", []byte (serialC2), 0666)
	
	// ---- //
	c <- [2]string{"l2", "Result exporting"}
	_ = <- c
	
	resultDatabase, err820 := sql.Open ("sqlite", filepath.Join (string (input200) + "/",
		"result-v2.db"))
	if err820 != nil {
		c <- [2]string{"l3", "Result exporting: Failed: Database creation failed: " +
			err820.Error ()}
		_ = <- c
		
		c <- [2]string{"hl", ""}
		_ = <- c

		return
	}

	_, err830 := resultDatabase.Exec (`create table if not exists listing (
		record_id nchar (32) primary key,
		brand_name varchar (32),
		model_name varchar (32),
		year nchar (4),
		transmission varchar (16),
		usage_condition varchar (16),
		mileage varchar (8),
		location varchar (16),
		price varchar(16),
		picture_url text,
		source_url text
	);
`)
/*
	create table if not exists listing_picture (
		record_id nchar (32) primary key,
		listing_record_id nchar (32),
		picture_url text,
		foreign key (listing_record_id) references listing (record_id)
	);
*/
	if err830 != nil {
		c <- [2]string{"l3", "Result exporting: Failed: Database creation failed: " +
			err830.Error ()}
		_ = <- c
		
		c <- [2]string{"hl", ""}
		_ = <- c

		return
	}

	for v, aListing := range listingInformation {		
		output832 := fmt.Sprintf ("Listing %d/%d: Details being exported", v + 1,
			len (listingInformation))
		c <- [2]string{"l3", output832}
		_ = <- c
		
		record_id, err835 := str.UniquePredsafeStr (32)
		if err835 != nil {
			output835 := fmt.Sprintf ("Listing %d/%d: Exporting failed: %s", v + 1,
				len (listingInformation), err835.Error ())
			c <- [2]string{"l3", output835}
			_ = <- c
			
			c <- [2]string{"hl", ""}
			_ = <- c
			
			return
		}

		if len (listingImages [v + 1]) < 3 {
			runtime.Gosched ()
			continue
		}
		
		_, err840 := resultDatabase.Exec (`insert into listing (
			record_id, brand_name, model_name, year, transmission,
			usage_condition, mileage, location, price, picture_url, source_url)
			values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
			record_id,
			aListing [0],
			aListing [1],
			aListing [2],
			aListing [3],
			aListing [4],
			aListing [5],
			aListing [6],
			aListing [7],
			listingImages [v + 1][2],
			aListing [8],
		)
		
		if err840 != nil {
			output840 := fmt.Sprintf ("Listing %d/%d: Exporting failed: %s", v + 1,
				len (listingInformation), err840.Error ())
			c <- [2]string{"l3", output840}
			_ = <- c
			
			c <- [2]string{"hl", ""}
			_ = <- c
			
			return
		}

		/*
		for _, picture := range listingImages [v + 1] {

			record_id2, err888 := str.UniquePredsafeStr (32)
			if err888 != nil {
				output888 := fmt.Sprintf ("Listing %d/%d: Exporting failed: %s", v + 1,
					len (listingInformation), err888.Error ())
				c <- [2]string{"l3", output888}
				_ = <- c
			
				c <- [2]string{"hl", ""}
				_ = <- c
			
				return
			}
			_, err860 := resultDatabase.Exec (`insert into listing_picture (
				record_id, listing_record_id, picture_url)
				values (?, ?, ?)`,
				record_id2,
				record_id,
				picture,
			)

			if err860 != nil {
				output860 := fmt.Sprintf ("Listing %d/%d: Exporting failed: %s",
					v + 1, len (listingInformation), err860.Error ())
				c <- [2]string{"l3", output860}
				_ = <- c
				
				c <- [2]string{"hl", ""}
				_ = <- c
				
				return
			}
		}*/
	}

	// ---- //
	c <- [2]string{"l2", "Task completed"}
	_ = <- c

	c <- [2]string{"hl", ""}
	_ = <- c
}
