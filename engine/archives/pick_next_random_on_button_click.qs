quando.key.handleKey id=0, key="a", ctrl=false, alt=false, callback = {
    quando.pick.random type="random", callback={
        quando.pick.one reset="", inverted=true, callback={
            quando.control.key ch="A", upDown="down", onOff=val
            quando.control.key ch="a", upDown="down", onOff=val
            quando.control.key ch="b", upDown="down", onOff=val
        }
        quando.pick.one reset="", inverted=true, callback={
            quando.control.key ch="1", upDown="down", onOff=val
            quando.control.key ch="2", upDown="down", onOff=val
            quando.control.key ch="3", upDown="down", onOff=val
        }
        quando.pick.one reset="", inverted=true, callback={
            quando.control.key ch="w", upDown="down", onOff=val
            quando.control.key ch="s", upDown="down", onOff=val
            quando.control.key ch="d", upDown="down", onOff=val
        }
    }
}