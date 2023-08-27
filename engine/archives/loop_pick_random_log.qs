quando.time.every count=1, units="seconds", callback={
    quando.pick.random type="random", callback={
        quando.log text="a"
        quando.log text="b"
        quando.log text="c"
    }
}
quando.time.every count=1, units="seconds", callback={
    quando.pick.random type="sequence", callback={
        quando.log text="a"
        quando.log text="b"
        quando.log text="c"
    }
}
quando.time.every count=1, units="seconds", callback={
    quando.pick.random type="unique", callback={
        quando.log text="a"
        quando.log text="b"
        quando.log text="c"
    }
}
