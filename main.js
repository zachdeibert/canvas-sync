#!/usr/bin/env node

const { MongoClient } = require("mongodb");
const { Builder, By, Condition, until } = require("selenium-webdriver");
const courseList = require("./scrapers/courseList");
require("./config-safe")(config => {

    const mongo = new MongoClient(config.dbURL);
    mongo.connect(err => {
        if (err) {
            console.error(err);
            mongo.close();
        } else {
            const db = mongo.db(config.dbName);
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
            driver.wait(until.titleIs("Courses")).then(() => {
                driver.wait(new Condition("document loaded", d => d.executeScript("return document.readyState;").then(s => s === "complete")))
            }).then(() => courseList(driver, db))
            .then(courses => {
                console.log(courses);
                driver.quit();
                mongo.close();
            });
        }
    });
});
