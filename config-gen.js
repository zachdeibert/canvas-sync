#!/usr/bin/env node

const fs = require("fs");
const readline = require("readline");

module.exports = callback => {
    const rl = readline.createInterface({
        "input": process.stdin,
        "output": process.stdout,
        "terminal": false
    });
    let config = {};
    rl.question("Canvas Domain: ", domain => {
        config.domain = domain;
        const stream = fs.createWriteStream("config.js", {
            "mode": 0644
        });
        stream.on("ready", () => {
            stream.write(`module.exports = ${JSON.stringify(config)};\n`);
            stream.close();
            rl.close();
            callback(null, config);
        });
        stream.on("error", err => {
            rl.close();
            callback(err);
        });
    });
};

if (require.main === module) {
    module.exports(err => {
        if (err) {
            console.error(err);
        }
    });
}
