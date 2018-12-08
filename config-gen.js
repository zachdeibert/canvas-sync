#!/usr/bin/env node

const fs = require("fs");
const readline = require("readline");

module.exports = callback => {
    const rl = readline.createInterface({
        "input": process.stdin,
        "output": process.stdout,
        "terminal": false
    });
    let options = [
        [ "Canvas Domain: ", "domain" ],
        [ "Login Page URL: ", "loginURL" ],
        [ "Username: ", "username" ],
        [ "Username Selector: ", "usernameSel" ],
        [ "Password: ", "password" ],
        [ "Password Selector: ", "passwordSel" ],
        [ "Login Button Selector: ", "loginSel" ]
    ];
    let config = {
        "version": module.exports.version
    };
    fs.access("config.js", err => {
        let oldConfig = {};
        if (!err) {
            oldConfig = require("./config");
        }
        const loop = i => {
            if (i < options.length) {
                if (oldConfig[options[i][1]]) {
                    config[options[i][1]] = oldConfig[options[i][1]];
                    loop(i + 1);
                } else {
                    rl.question(options[i][0], res => {
                        config[options[i][1]] = res;
                        loop(i + 1);
                    });
                }
            } else {
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
            }
        };
        loop(0);
    });
};

module.exports.version = 1;

if (require.main === module) {
    module.exports(err => {
        if (err) {
            console.error(err);
        }
    });
}
