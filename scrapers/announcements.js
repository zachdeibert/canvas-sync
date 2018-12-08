const request = require("request");
const { By } = require("selenium-webdriver");
const DOMParser = require("xmldom").DOMParser;

module.exports = (collection, driver, db, courseId, config) => {
    return Promise.all([
        collection.find({}, {
            "projection": {
                "_id": 1
            }
        }),
        driver.findElement(By.id("external_feed")).then(el => {
            el.click();
            return new Promise(resolve => {
                setTimeout(resolve, 1000);
            });
        }).then(() => {
            return driver.findElement(By.id("rss-feed-link"));
        }).then(el => {
            return el.getAttribute("href");
        }).then(feedURL => {
            return new Promise((resolve, reject) => {
                request(feedURL, (err, res, body) => {
                    if (err) {
                        reject(err);
                    } else {
                        const parser = new DOMParser();
                        const document = parser.parseFromString(body, "text/xml");
                        const entries = document.getElementsByTagNameNS("http://www.w3.org/2005/Atom", "entry");
                        let parsedEntries = [];
                        for (let i = 0; i < entries.length; ++i) {
                            const entry = entries.item(i);
                            parsedEntries.push({
                                "title": entry.getElementsByTagNameNS("http://www.w3.org/2005/Atom", "title").item(0).textContent,
                                "_id": entry.getElementsByTagNameNS("http://www.w3.org/2005/Atom", "id").item(0).textContent,
                                "date": new Date(entry.getElementsByTagNameNS("http://www.w3.org/2005/Atom", "published").item(0).textContent),
                                "author": entry.getElementsByTagNameNS("http://www.w3.org/2005/Atom", "name").item(0).textContent,
                                "content": entry.getElementsByTagNameNS("http://www.w3.org/2005/Atom", "content").item(0).textContent
                            });
                        }
                        resolve(parsedEntries);
                    }
                });
            });
        })
    ]).then(args => {
        let currentSet = [];
        args[0].forEach(entry => {
            currentSet.push(entry._id);
        });
        const newSet = args[1].filter(entry => currentSet.indexOf(entry._id) < 0);
        if (newSet.length > 0) {
            return collection.insertMany(newSet);
        }
    });
};
