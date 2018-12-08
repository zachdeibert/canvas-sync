const { By, Condition } = require("selenium-webdriver");
const PromiseUtil = require("../util/promise");

module.exports = (driver, db, config) => {
    const courses = db.collection("courses");
    return courses.find({
        "active": true
    }, {
        "projection": {
            "_id": 1,
            "pages.type": 1
        }
    }).toArray().then(courseList => {
        return PromiseUtil.syncAll(courseList.map(course => {
            return () => {
                driver.get(`https://${config.domain}/courses/${course._id}`);
                return driver.wait(new Condition("document loaded", d => d.executeScript("return document.readyState;").then(s => s === "complete"))).then(() => {
                    return Promise.all([
                        driver.findElement(By.id("content")).getAttribute("innerHTML"),
                        driver.findElement(By.id("section-tabs")).then(el => {
                            return el.findElements(By.tagName("a"));
                        }).then(els => {
                            return Promise.all(els.map(el => {
                                return Promise.all([
                                    el.getAttribute("title"),
                                    el.getAttribute("class")
                                ]).then(args => {
                                    return {
                                        "name": args[0],
                                        "type": args[1].split(" ", 2)[0]
                                    };
                                });
                            }));
                        })
                    ]);
                }).then(args => {
                    let currentSet = [];
                    if (course.pages) {
                        course.pages.forEach(page => {
                            currentSet.push(page.type);
                        });
                    }
                    const newSet = args[1].filter(page => page.type !== "home" && currentSet.indexOf(page) < 0);
                    if (newSet.length > 0) {
                        return courses.updateOne({
                            "_id": course._id
                        }, {
                            "$push": {
                                "pages": {
                                    "$each": newSet
                                }
                            },
                            "$set": {
                                "homepage": args[0]
                            }
                        });
                    }
                });
            };
        }));
    });
};
