#!/usr/bin/env node

const { MongoClient } = require("mongodb");
const { Builder, By, Condition, until } = require("selenium-webdriver");
require("./config-safe")(config => {

    const mongo = new MongoClient(config.dbURL);
    mongo.connect(err => {
        if (err) {
            console.error(err);
            mongo.close();
        } else {
            const db = mongo.db(config.dbName);
            const courses = db.collection("courses");
            const driver = new Builder()
                .forBrowser("firefox")
                .build();
            driver.get(`https://${config.domain}/courses`);
            driver.wait(until.urlIs(config.loginURL), 10000).then(() => {
                driver.wait(new Condition("document loaded", d => d.executeScript("return document.readyState;").then(s => s === "complete"))).then(() => {
                    Promise.all([
                        driver.findElement(By.css(config.loginSel)),
                        driver.findElement(By.css(config.usernameSel)).sendKeys(config.username),
                        driver.findElement(By.css(config.passwordSel)).sendKeys(config.password)
                    ]).then(args => args[0].click());
                });
            }).catch(() => {});
            Promise.all([
                courses.find({}, {
                    "projection": { "_id": 1 }
                }).toArray(),
                driver.wait(until.titleIs("Courses")).then(() => {
                    driver.wait(new Condition("document loaded", d => d.executeScript("return document.readyState;").then(s => s === "complete")))
                }).then(() => {
                    return driver.findElements(By.className("course-list-table-row"))
                }).then(courses => {
                    return Promise.all(courses.map(course =>
                        Promise.all([
                            course.findElement(By.className("course-list-course-title-column")).getText(),
                            course.findElement(By.className("course-list-term-column")).getText(),
                            new Promise(resolve => course.findElement(By.tagName("a")).getAttribute("href").then(resolve).catch(() => resolve()))
                        ]).then(args => {
                            return {
                                "name": args[0],
                                "term": args[1],
                                "_id": args[2] && parseInt(args[2].substr(args[2].lastIndexOf("/") + 1))
                            };
                        })));
                })
            ]).then(args => {
                let currentSet = [];
                args[0].forEach(obj => {
                    currentSet.push(obj._id);
                });
                const newCourses = args[1].filter(course => course._id && currentSet.indexOf(course._id) < 0);
                if (newCourses.length > 0) {
                    return courses.insertMany(newCourses);
                }
            }).then(() => {
                driver.quit();
                mongo.close();
            });
        }
    });
});
