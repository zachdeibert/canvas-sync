const electron = require("electron");
const path = require("path");
const app = require("./app");

electron.protocol.registerSchemesAsPrivileged([
    {
        "scheme": "static",
        "privileges": {
            "standard": true,
            "secure": true
        }
    }, {
        "scheme": "scraped",
        "privileges": {
            "standard": true,
            "secure": true,
            "bypassCSP": true,
            "corsEnabled": true,
            "supportFetchAPI": true
        }
    }
]);

let scraperMap = {};

app.onLoad(() => {
    electron.protocol.registerFileProtocol("static", (req, callback) => {
        let url = req.url.split("/").slice(3).join("/");
        if (url.startsWith("node_modules")) {
            callback({
                path: path.normalize(`${__dirname}/../${url}`)
            });
        } else {
            callback({
                path: path.normalize(`${__dirname}/../html/${url}`)
            });
        }
    });
    electron.protocol.registerStringProtocol("scraped", (req, callback) => {
        if (req.uploadData && req.uploadData.length === 1) {
            let prefix = req.url.split("/", 4)[2];
            if (scraperMap[prefix]) {
                let obj = JSON.parse(req.uploadData[0].bytes.toString());
                if (obj) {
                    callback("ok");
                    setTimeout(() => {
                        scraperMap[prefix](obj);
                    });
                } else {
                    console.error("could not deserialize obejct");
                    callback("error: could not deserialize object");
                }
            } else {
                console.error("URI not mapped");
                callback("error: URI not mapped");
            }
        } else {
            console.error("could not understand request");
            callback("error: could not understand request");
        }
    });
});

module.exports = {
    "on": (prefix, callback) => {
        if (scraperMap[prefix]) {
            return Promise.reject("Scraped URI already bound");
        }
        scraperMap[prefix] = callback;
        return Promise.resolve();
    },
    "once": (prefix, callback) => {
        const listener = data => {
            module.exports.remove(prefix, listener);
            callback(data);
        };
        return module.exports.on(prefix, listener);
    },
    "remove": (prefix, callback) => {
        if (!scraperMap[prefix]) {
            return Promise.reject("Scraped URI not bound");
        }
        if (scraperMap[prefix] != callback) {
            return Promise.reject("Wrong scraper callback");
        }
        scraperMap[prefix] = null;
        return Promise.resolve();
    }
};
