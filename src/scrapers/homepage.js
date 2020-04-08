(function() {
    var cards = document.getElementsByClassName("ic-DashboardCard");
    var courses = [];
    for (var i = 0; i < cards.length; ++i) {
        var card = cards.item(i);
        var course = {
            "name": card.getAttribute("aria-label")
        };
        var title = card.getElementsByClassName("ic-DashboardCard__header-title");
        if (title.length === 1) {
            title = title.item(0);
            var span = title.getElementsByTagName("span");
            if (span.length === 1) {
                span = span.item(0);
                course.color = span.style.color;
            }
        }
        var subtitle = card.getElementsByClassName("ic-DashboardCard__header-subtitle");
        if (subtitle.length === 1) {
            subtitle = subtitle.item(0);
            course.code = subtitle.innerHTML.trim();
        }
        var term = card.getElementsByClassName("ic-DashboardCard__header-term");
        if (term.length === 1) {
            term = term.item(0);
            course.term = term.innerHTML.trim();
        }
        var as = card.getElementsByTagName("a");
        for (var j = 0; j < as.length; ++j) {
            var a = as.item(j);
            if (!a.href.endsWith("s")) {
                if (course.id) {
                    course.id = null;
                    break;
                } else {
                    var p = a.href.split("/");
                    course.id = p[p.length - 1];
                }
            }
        }
        courses.push(course);
    }
    fetch("scraped://homepage/", {
        "method": "POST",
        "body": JSON.stringify(courses)
    });
})();
