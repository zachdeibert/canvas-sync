#!/usr/bin/env node

const { Builder, By, Condition, until } = require("selenium-webdriver");
require("./config-safe")(config => {

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
        driver.wait(new Condition("document loaded", d => d.executeScript("return document.readyState;").then(s => s === "complete"))).then(() => {
            driver.findElements(By.className("course-list-table-row")).then(courses => {
                Promise.all(courses.map(course =>
                    Promise.all([
                        course.findElement(By.className("course-list-course-title-column")).getText(),
                        course.findElement(By.className("course-list-term-column")).getText(),
                        new Promise(resolve => course.findElement(By.tagName("a")).getAttribute("href").then(resolve).catch(() => resolve()))
                    ]).then(args => {
                        return {
                            "name": args[0],
                            "term": args[1],
                            "courseId": args[2] && parseInt(args[2].substr(args[2].lastIndexOf("/") + 1))
                        };
                    }))).then(courses => {
                        courses.filter(course => course.courseId).forEach(course => {
                            console.log(course);
                        });
                    }).then(() => {
                        driver.quit();
                    });
            });
        });
    });
});
