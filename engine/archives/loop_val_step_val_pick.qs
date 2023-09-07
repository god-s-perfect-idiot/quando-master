quando.time.vary count=2, units="seconds", mode="seesaw", times=4, timesUnits="seconds", inverted=false, callback={
    quando.pick.val, callback={
        quando.log text="0 -> 20"
        quando.log text="20 -> 40"
        quando.log text="40 -> 60"
        quando.log text="60 -> 80"
        quando.log text="80 -> 100"
    }
}
