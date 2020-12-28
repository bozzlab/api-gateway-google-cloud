const express = require('express')
const router = express.Router()
const axios = require('axios')


router.get('/info', (req, res) => {
    console.log(JSON.stringify(req.headers));
    res.json({
        "serviceName" : "Stock API",
        "language" : "Node.js",
        "framework" : "Express"
    })
})


module.exports = router