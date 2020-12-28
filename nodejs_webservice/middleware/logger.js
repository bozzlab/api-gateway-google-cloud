function logger(req, res, next) {
    console.log(`[Log] : Requesting ${req.method} to ${req.url} `);
    console.log(`[Log] : Header : ${JSON.stringify(req.headers)}`);
    console.log(`[Log] : x-endpoint-api-userinfo : ${JSON.stringify(req.headers["x-endpoint-api-userinfo"])}`);
    next()
} 

module.exports = logger
