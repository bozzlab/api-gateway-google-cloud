const express = require('express')
const router = express.Router()
const stockData = require('../data')

router.get('/', (req, res) => {
    res.json(stockData)
})

router.get('/:id', (req, res) => {
    let stockId = Number.parseInt(req.params.id, 10)
    let stock = stockData.find((stock)=> stockId == stock.id)
    res.json(stock)
})

router.post('/', (req, res) => {
    let newStock = {
        ...req.body
    }
    stockData.push(req.body)
    res.json(newStock)
})

router.put('/:id', (req, res) => {
    let stockId = Number.parseInt(req.params.id, 10)
    let stockIndex = stockData.findIndex((stock)=> stockId == stock.id)

    let updateStock = {
        id: stockId,
        ...req.body
    }
        
    stockData[stockIndex] = updateStock
    res.json(updateStock)
})

router.delete('/:id', (req, res) => {
    let stockId = Number.parseInt(req.params.id, 10)
    let stockIndex = stockData.findIndex((stock)=> stockId == stock.id)
    stockData.splice(stockIndex, 1)
    res.sendStatus(204)
})


module.exports = router