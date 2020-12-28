// const path = require('path')
const express = require('express')
// const hbs = require('express-handlebars')


const stockRouter = require('./routes/stock')
const IndexRouter = require('./routes')


const logger = require('./middleware/logger')

const app = express()

// template engines
// app.engine('hbs', hbs({extname : 'hbs'}))
// app.set('view engine', 'hbs')

// middlesware
// app.use(express.static(path.join(__dirname, 'public')))
app.use(express.json())
app.use(express.urlencoded({extended : false}))

// custom middleware 
app.use(logger)


// routes
app.use('/stock', stockRouter)
app.use('/', IndexRouter)

app.use(function(req, res, next) {
    res.status(404);
    res.send({'message' : 'Not Found'});
});

app.listen(8080, () => {
    console.log('listen port 8080')
})

module.exports = app