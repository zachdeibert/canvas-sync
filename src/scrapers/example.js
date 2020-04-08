const scraper = require("./scraper");

scraper.register("Example", (course, db, update, del) => {
    console.log(`course = ${course}, db = ${db}, update = ${update}, delete = ${del}`);
});
