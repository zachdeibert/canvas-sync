const app = require("./app");
const proto = require("./proto");
const scraper = require("./scrapers/scraper");

module.exports = {
    "startCollection": db => {
        proto.once("homepage", res => {
            app.window.webContents.loadURL("static://options/options.html").then(() => {
                app.window.webContents.executeJavaScript(`init(${JSON.stringify(res)}, ${JSON.stringify(scraper.getNames())});`);
            });
            proto.once("options", res => {
                let resolve;
                let queue = new Promise((_resolve, reject) => {
                    resolve = _resolve;
                });
                res.forEach(course => {
                    let needed = false;
                    course.scrapers.forEach(s => {
                        needed = needed || s.update || s.delete;
                    });
                    if (needed) {
                        let subDb = db.forCourse(course.id, course.name, course.code, course.term);
                        queue = queue.then(() => scraper.execute(course.id, course.scrapers, subDb));
                    }
                });
                queue.then(() => {
                    console.log("Done!");
                });
                resolve();
            });
        });
        app.runScript("homepage");
    }
};
