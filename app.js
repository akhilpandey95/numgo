/*
 * numgo http server
 * MPL License
 * Rex Labs
 */

const f = require('fs');
const h = require('http');
const c = require('cluster');
const e = require('express');
const bp = require('body-parser');
const cp = require('cookie-parser');
const n = require('os').cpus().length;

if (c.isMaster) {
    process.stdout.write(`master ${process.pid} is running\n`);

for (let i = 0; i < n; i++) {
    c.fork();
}

c.on('exit', (worker, code, signal) => {
    process.stdout.write(`worker ${worker.process.pid} is dead\n`);
    c.fork();
});

} else {
        var port = process.env.PORT || 4567;
        var app = e();
        var router = e.Router();

        app.use(e.static('assets'));
        app.set('title', "home | numgo");
        app.use(cp());
        app.use(bp.json());
        app.use(bp.urlencoded({extended: true}));

        router.get('/', (req, res) => {
                res.sendFile(__dirname + '/index.html');
        });

        router.get('/ufunc', (req, res) => {
                res.sendFile(__dirname + '/ufunc.html');
        });

        router.get('/linalg', (req, res) => {
                res.sendFile(__dirname + '/linalg.html');
        });

        router.get('/routines', (req, res) => {
                res.sendFile(__dirname + '/routines.html');
        });
        app.use('/', router);

        h.createServer(app).listen(port);
    process.stdout.write(`worker ${process.pid} has started\n`);
}
