let scrapers = [];

module.exports = {
    "register": (name, callback) => {
        scrapers.push({
            "name": name,
            "callback": callback
        });
    },
    "getNames": () => {
        return scrapers.map(s => s.name);
    },
    "execute": (courseId, config, db) => {
        scrapers.forEach((scraper, idx) => {
            if (config[idx].update || config[idx].delete) {
                return scraper.callback(courseId, db.forScraper(scraper.name), config[idx].update, config[idx].delete);
            } else {
                return Promise.resolve();
            }
        });
    }
};
