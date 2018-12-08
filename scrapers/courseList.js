const { By } = require("selenium-webdriver");

module.exports = (driver, db) => {
    const courses = db.collection("courses");
    return Promise.all([
        courses.find({}, {
            "projection": { "_id": 1 }
        }).toArray(),
        driver.findElements(By.className("course-list-table-row")).then(courses => {
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
        const currentCourses = args[1].filter(course => course._id);
        const newCourses = currentCourses.filter(course => currentSet.indexOf(course._id) < 0);
        if (newCourses.length > 0) {
            return courses.insertMany(newCourses);
        }
        return currentCourses;
    });
};
