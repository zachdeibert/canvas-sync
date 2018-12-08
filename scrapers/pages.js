const { Condition } = require("selenium-webdriver");
const PromiseUtil = require("../util/promise");

const lut = {
    "announcements": require("./announcements")
};
let lutWarnings = {};

module.exports = (driver, db, config) => {
    const courses = db.collection("courses");
    return courses.find({
        "active": true
    }, {
        "projection": {
            "_id": 1,
            "pages.type": 1,
            "pages.slug": 1
        }
    }).toArray().then(courseList => {
        return PromiseUtil.syncAll(courseList.map(course => {
            return () => {
                if (course.pages) {
                    return PromiseUtil.syncAll(course.pages.map(page => {
                        return () => {
                            if (lut[page.type]) {
                                driver.get(`https://${config.domain}/courses/${course._id}${page.slug}`);
                                return driver.wait(new Condition("document loaded", d => d.executeScript("return document.readyState;").then(s => s === "complete"))).then(() => {
                                    return lut[page.type](db.collection(`${course._id}_${page.type}`), driver, db, course._id, config);
                                });
                            } else if (!lutWarnings[page.type]) {
                                console.warn(`Warning: unknown page type '${page.type}'`);
                                lutWarnings[page] = true;
                            }
                        };
                    }));
                }
            };
        }));
    });
};
