const fs = require("fs");
const configGen = require("./config-gen");

module.exports = callback => {
    fs.access("config.js", err => {
        if (err) {
            configGen((err, config) => {
                if (err) {
                    console.error(err);
                } else {
                    callback(config);
                }
            });
        } else {
            const config = require("./config");
            if (config.version === configGen.version) {
                callback(config);
            } else {
                configGen((err, config) => {
                    if (err) {
                        console.error(err);
                    } else {
                        callback(config);
                    }
                });
            }
        }
    });
};
