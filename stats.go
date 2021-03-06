package main

import "time"

type Snapshot struct {
	Timestamp   time.Time `json:""`
	Hour        int       `json:"hour"`
	SixHour     int       `json:"SixHour"`
	Day         int       `json:"day"`
	ThreeDay    int       `json:"threeDay"`
	SevenDay    int       `json:"sevenDay"`
	FourteenDay int       `json:"fourteenDay"`
	ThirtyDay   int       `json:"thirtyDay"`
	NinetyDay   int       `json:"ninetyDay"`
	HalfYear    int       `json:"halfYear"`
	Year        int       `json:"year"`
	AllTime     int       `json:"allTime"`
}

type StatType int

const (
	StatAll StatType = iota
	StatClearnet
	StatTorOnly
	StatDualStack
	StatListings
	StatRatings
	StatVendors
)

type StatsLogger struct {
	db       Datastore
	getNodes func() []Node
}

func NewStatsLogger(db Datastore, getNodes func() []Node) *StatsLogger {
	sl := &StatsLogger{
		db:       db,
		getNodes: getNodes,
	}
	return sl
}

func (sl *StatsLogger) run() {
	t := time.NewTicker(time.Hour)
	for range t.C {
		nodes := sl.getNodes()

		all := Snapshot{Timestamp: time.Now()}
		clearnet := Snapshot{Timestamp: time.Now()}
		torOnly := Snapshot{Timestamp: time.Now()}
		dualStack := Snapshot{Timestamp: time.Now()}
		listings := Snapshot{Timestamp: time.Now()}
		ratings := Snapshot{Timestamp: time.Now()}
		vendors := Snapshot{Timestamp: time.Now()}


		for _, node := range nodes {
			var rec Snapshot
			all.AllTime++
			listings.AllTime += node.Listings
			ratings.AllTime += node.Ratings
			if node.Vendor {
				vendors.AllTime++
			}
			if node.LastConnect.Add(time.Hour).After(time.Now()) {
				rec.Hour = 1
				all.Hour++
				listings.Hour += node.Listings
				ratings.Hour += node.Ratings
				if node.Vendor {
					vendors.Hour++
				}
			}
			if node.LastConnect.Add(time.Hour * 6).After(time.Now()) {
				rec.SixHour = 1
				all.SixHour++
				listings.SixHour += node.Listings
				ratings.SixHour += node.Ratings
				if node.Vendor {
					vendors.SixHour++
				}
			}
			if node.LastConnect.Add(time.Hour * 24).After(time.Now()) {
				rec.Day = 1
				all.Day++
				listings.Day += node.Listings
				ratings.Day += node.Ratings
				if node.Vendor {
					vendors.Day++
				}
			}
			if node.LastConnect.Add(time.Hour * 24 * 3).After(time.Now()) {
				rec.ThreeDay = 1
				all.ThreeDay++
				listings.ThreeDay += node.Listings
				ratings.ThreeDay += node.Ratings
				if node.Vendor {
					vendors.ThreeDay++
				}
			}
			if node.LastConnect.Add(time.Hour * 24 * 7).After(time.Now()) {
				rec.SevenDay = 1
				all.SevenDay++
				listings.SevenDay += node.Listings
				ratings.SevenDay += node.Ratings
				if node.Vendor {
					vendors.SevenDay++
				}
			}
			if node.LastConnect.Add(time.Hour * 24 * 14).After(time.Now()) {
				rec.FourteenDay = 1
				all.FourteenDay++
				listings.FourteenDay += node.Listings
				ratings.FourteenDay += node.Ratings
				if node.Vendor {
					vendors.FourteenDay++
				}
			}
			if node.LastConnect.Add(time.Hour * 24 * 30).After(time.Now()) {
				rec.ThirtyDay = 1
				all.ThirtyDay++
				listings.ThirtyDay += node.Listings
				ratings.ThirtyDay += node.Ratings
				if node.Vendor {
					vendors.ThirtyDay++
				}
			}
			if node.LastConnect.Add(time.Hour * 24 * 90).After(time.Now()) {
				rec.NinetyDay = 1
				all.NinetyDay++
				listings.NinetyDay += node.Listings
				ratings.NinetyDay += node.Ratings
				if node.Vendor {
					vendors.NinetyDay++
				}
			}
			if node.LastConnect.Add(time.Hour * 24 * 182).After(time.Now()) {
				rec.HalfYear = 1
				all.HalfYear++
				listings.HalfYear += node.Listings
				ratings.HalfYear += node.Ratings
				if node.Vendor {
					vendors.HalfYear++
				}
			}
			if node.LastConnect.Add(time.Hour * 24 * 365).After(time.Now()) {
				rec.Year = 1
				all.Year++
				listings.Year += node.Listings
				ratings.Year += node.Ratings
				if node.Vendor {
					vendors.Year++
				}
			}
			nodeType := GetNodeType(node.PeerInfo.Addrs)
			switch nodeType {
			case Clearnet:
				clearnet.Hour += rec.Hour
				clearnet.SixHour += rec.SixHour
				clearnet.Day += rec.Day
				clearnet.ThreeDay += rec.ThreeDay
				clearnet.SevenDay += rec.SevenDay
				clearnet.FourteenDay += rec.FourteenDay
				clearnet.ThirtyDay += rec.ThirtyDay
				clearnet.NinetyDay += rec.NinetyDay
				clearnet.HalfYear += rec.HalfYear
				clearnet.Year += rec.Year
				clearnet.AllTime++
			case TorOnly:
				torOnly.Hour += rec.Hour
				torOnly.SixHour += rec.SixHour
				torOnly.Day += rec.Day
				torOnly.ThreeDay += rec.ThreeDay
				torOnly.SevenDay += rec.SevenDay
				torOnly.FourteenDay += rec.FourteenDay
				torOnly.ThirtyDay += rec.ThirtyDay
				torOnly.NinetyDay += rec.NinetyDay
				torOnly.HalfYear += rec.HalfYear
				torOnly.Year += rec.Year
				torOnly.AllTime++
			case DualStack:
				dualStack.Hour += rec.Hour
				dualStack.SixHour += rec.SixHour
				dualStack.Day += rec.Day
				dualStack.ThreeDay += rec.ThreeDay
				dualStack.SevenDay += rec.SevenDay
				dualStack.FourteenDay += rec.FourteenDay
				dualStack.ThirtyDay += rec.ThirtyDay
				dualStack.NinetyDay += rec.NinetyDay
				dualStack.HalfYear += rec.HalfYear
				dualStack.Year += rec.Year
				dualStack.AllTime++
			}
		}
		sl.db.PutStat(StatAll, all)
		sl.db.PutStat(StatClearnet, clearnet)
		sl.db.PutStat(StatTorOnly, torOnly)
		sl.db.PutStat(StatDualStack, dualStack)
		sl.db.PutStat(StatListings, listings)
		sl.db.PutStat(StatRatings, ratings)
		sl.db.PutStat(StatVendors, vendors)
	}
}
