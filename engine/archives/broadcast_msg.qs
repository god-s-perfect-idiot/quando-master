quando.message.send message="hi"

# receiver
quando.message.add_handler('hi', (txt) => {
quando.text(' received hi ',true,false)
})
