/* 
# #!/data/data/com.termux/files/user/bin/bash
# 
# DB = "db.txt"
# GPSOUT = "gpsout.txt"
# LOCATION = "termux-location -p gps"
# AWKPROG = "gps.awk"
# AWK = "awk -F: -f $AWKPROG"
# SMS = "termux-sms-send -n xxx-yyy-zzzz"
# 
# cat <<-! > $AWKPROG
# { print \$2; }
# !
# 
# if [ -w $DB ]
# then
#     mv $DB ${DB}O
# fi
# 
# while true
# do
#     $LOCATION > $GPSOUT
#     LAT=`cat $GPSOUT  | grep latitude  | $AWK`
#     LONG=`cat $GPSOUT | grep longitude | $AWK`
#     D=`date +%H:%M:%S`
#     echo $D $LAT $LONG | tee -a $DB
#     sleep 2   # 60
# done

*/

package main

import (  
    "fmt"
    "encoding/json"
    "math"
    "os/exec"
    "time"
)

const (
    ITERATIONS       = 20
    POLL_INTERVAL    = 2 * time.Second
    LOCATION_COMMAND = "termux-location"
)

type Location struct {
    Latitude  float64
    Longitude float64
    Altitude  float64
    Accuracy  float64
}

func hsin(theta float64) float64 {
    return math.Pow(math.Sin(theta/2), 2)
}

// From https://gist.github.com/cdipaolo/d3f8db3848278b49db68
//
// Distance function returns the distance (in meters) between two points of
// a given longitude and latitude relatively accurately (using a spherical
// approximation of the Earth) through the Haversin Distance Formula for
// great arc distance on a sphere with accuracy for small distances
// point coordinates are supplied in degrees and converted into rad. in the func
// distance returned is METERS!!!!!!
// http://en.wikipedia.org/wiki/Haversine_formula
func distance(lat1, lon1, lat2, lon2 float64) float64 {
    const (
        rad = math.Pi/180   // convert to radians
        r   = 6378100       // Earth radius in meters
    )
    
    la1 := lat1*rad
    lo1 := lon1*rad
    la2 := lat2*rad
    lo2 := lon2*rad

    h := hsin(la2 - la1) + math.Cos(la1)*math.Cos(la2)*hsin(lo2 - lo1)
    return 2*r*math.Asin(math.Sqrt(h))
}

func main() {  
    var loc, priorloc Location
    var dist, totaldist, seconds, totalseconds, avgspeed float64
    var priortime time.Time

    for i := 0; i < 20; i++ {
        out, err := exec.Command(LOCATION_COMMAND).Output()
        if err != nil {
            fmt.Println(err)
        }
        now := time.Now()

        err = json.Unmarshal([]byte(out), &loc)
	if err != nil {
		fmt.Println("error parsing location:", err)
	}

//	fmt.Printf("%+v, %+v\n", priorloc, loc)

        if priorloc.Latitude != 0 {
            dist = distance(priorloc.Latitude, priorloc.Longitude,
                            loc.Latitude, loc.Longitude)
            totaldist += dist
            seconds = time.Since(priortime).Seconds()
            totalseconds +=  seconds
//            if seconds != 0 {
//                speed = dist*3.6/seconds
//            } else {
//                speed = 0
//            }
            if totalseconds != 0 {
                avgspeed = totaldist*3.6/totalseconds
            } else {
                avgspeed = 0
            }
        }
        fmt.Printf("%s %f, %f: %.0fm\t%.0fkm/h\n",
                   now.Format("01/02/2006 15:04:05"),
                   loc.Latitude, loc.Longitude,
                   totaldist, avgspeed)
        priorloc = loc
        priortime = now

        time.Sleep(POLL_INTERVAL)
    }
}
