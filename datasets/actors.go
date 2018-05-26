package datasets

// import (
// 	bolt "github.com/coreos/bbolt"
// 	"github.com/vmihailenco/msgpack"
// 	"regexp"
// )

// var (
// 	actorLineRE = regexp.MustCompile(`^([^\t]+)?\t+(.*)`)
// )

// *
// -----------------------------------------------------------------------------
// RULES:
// 1       Movies and recurring TV roles only, no TV guest appearances
// 2       Please submit entries in the format outlined at the end of the list
// 3       Feel free to submit new actors

// "xxxxx"        = a television series
// "xxxxx" (mini) = a television mini-series
// [xxxxx]        = character name
// <xx>           = number to indicate billing position in credits
// (TV)           = TV movie, or made for cable movie
// (V)            = made for video movie (this category does NOT include TV
//                  episodes repackaged for video, guest appearances in
//                  variety/comedy specials released on video, or
//                  self-help/physical fitness videos)

// THE ACTORS LIST
// ===============

// Name                    Titles
// ----                    ------
// #1, Buffy               Closet Monster (2015)  [Buffy 4]  <31>

// $, Claw                 "OnCreativity" (2012)  [Himself]

// $, Homo                 Nykytaiteen museo (1986)  [Himself]  <25>
//                         Suuri illusioni (1985)  [Guests]  <22>

// $, Steve                E.R. Sluts (2003) (V)  <12>

// $hutter                 Battle of the Sexes (2017)  (as $hutter Boy)  [Bobby Riggs Fan]  <10>
//                         NVTION: The Star Nation Rapumentary (2017)  (as $hutter Boy)  [Himself]  <1>
//                         Secret in Their Eyes (2015)  (uncredited)  [2002 Dodger Fan]
//                         Steve Jobs (2015)  (uncredited)  [1988 Opera House Patron]
//                         Straight Outta Compton (2015)  (uncredited)  [Club Patron/Dopeman]

// $lim, Bee Moe           Fatherhood 101 (2013)  (as Brandon Moore)  [Himself - President, Passages]
//                         For Thy Love 2 (2009)  [Thug 1]
//                         Night of the Jackals (2009) (V)  [Trooth]
//                         "Idle Talk" (2013)  (as Brandon Moore)  [Himself]
//                         "Idle Times" (2012) {(#1.1)}  (as Brandon Moore)  [Detective Ryan Turner]
//                         "Idle Times" (2012) {(#1.2)}  (as Brandon Moore)  [Detective Ryan Turner]
//                         "Idle Times" (2012) {(#1.3)}  (as Brandon Moore)  [Detective Ryan Turner]
//                         "Idle Times" (2012) {(#1.4)}  (as Brandon Moore)  [Detective Ryan Turner]
//                         "Idle Times" (2012) {(#1.5)}  (as Brandon Moore)  [Detective Ryan Turner]
//                         "Money Power & Respect: The Series" (2010) {Bottom Dollar (#2.7)}  [Bee Moe Slim]
//                         "Money Power & Respect: The Series" (2010) {Flip Side (#2.5)}  [Bee Moe Slim]
//                         "Red River" (2015)  [Victor Charles]
//                         "Red River" (2015) {Under Control (#1.10)}  [Victor Charles]

// $ly, Yung               Town Bizzness Pt 3 (2014) (V)  [Yung $ly]
//                         "From Tha Bottom 2 Tha Top" (2016)  [Yung $ly]
//                         "From Tha Bottom 2 Tha Top" (2016) {T-Pain (#1.2)}  [Yung $ly]

// *

// var actorsBucketName = "actors"

// func ExportActorsDataset(path string, db *bolt.DB) error {
// 	scanner, err := NewScanner(path)
// 	if err != nil {
// 		return err
// 	}
// 	defer scanner.Close()

// 	scanner.Skip(239)

// 	var (
// 		currentActor *string
// 		inserts      int
// 	)

// 	tx, err := db.Begin(true)
// 	if err != nil {
// 		return err
// 	}

// 	b, err := tx.CreateBucketIfNotExists(bucketName)
// 	if err != nil {
// 		return err
// 	}

// 	for scanner.Scan() {

// 		if tx == nil {
// 			tx, err = db.Begin(true)
// 			if err != nil {
// 				return err
// 			}

// 			b, err = tx.CreateBucketIfNotExists(bucketName)
// 			if err != nil {
// 				return err
// 			}
// 		}

// 		actorTitle := actorLineRE.FindStringSubmatch(scanner.Text())

// 		var titleExpr string
// 		if len(actorTitle) == 3 {
// 			if len(actorTitle[1]) != 0 {
// 				currentActor = &actorTitle[1]
// 			}
// 			titleExpr = actorTitle[2]
// 		} else {
// 			continue
// 		}

// 		if currentActor == nil {
// 			continue
// 		}

// 		movieTitle, movieYear, err := parseTitleAndYear(titleExpr)
// 		if err != nil {
// 			return err
// 		}

// 		movieId := makeId(movieTitle, movieYear)

// 		key := []byte(movieId)

// 		entryBytes := b.Get(key)

// 		var movie Movie

// 		if entryBytes == nil {
// 			movie.Title = movieTitle
// 			movie.Year = movieYear
// 			movie.MovieId = movieId
// 		} else {
// 			if err := msgpack.Unmarshal(entryBytes, &movie); err != nil {
// 				return err
// 			}
// 		}

// 		movie.Actors = append(movie.Actors, *currentActor)

// 		if entryBytes, err = msgpack.Marshal(&movie); err != nil {
// 			return err
// 		}

// 		if err := b.Put(key, entryBytes); err != nil {
// 			return err
// 		}

// 		inserts += 1

// 		if inserts%10000 == 0 {
// 			if err := tx.Commit(); err != nil {
// 				return err
// 			}
// 			tx = nil
// 		}
// 	}

// 	if tx != nil {
// 		if err := tx.Commit(); err != nil {
// 			return err
// 		}
// 	}

// 	return scanner.Err()
// }
