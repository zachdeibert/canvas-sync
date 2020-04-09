(function() {
    var atRe = / at /;
    var amRe = /[aA][mM]$/;
    var pmRe = /[pP][mM]$/;
    var rows = document.getElementsByClassName("ic-announcement-row");
    var announcements = [];
    for (var i = 0; i < rows.length; ++i) {
        var row = rows.item(i);
        var announcement = {};
        var title = row.querySelectorAll(".ic-item-row__content-link-container > h3");
        if (title.length === 1) {
            announcement.title = title.item(0).lastChild.textContent;
        }
        var content = row.getElementsByClassName("ic-announcement-row__content");
        if (content.length === 1) {
            announcement.text = content.item(0).innerText;
        }
        var postDate = row.getElementsByClassName("ic-item-row__meta-content-timestamp");
        if (postDate.length === 1) {
            var str = postDate.item(0).innerText;
            str = str.replace(atRe, " ");
            str = str.replace(amRe, " am");
            str = str.replace(pmRe, " pm");
            announcement.postDate = new Date(str).getTime();
        }
        var replies = row.getElementsByClassName("ic-unread-badge__total-count");
        if (replies.length == 1) {
            announcement.replies = parseInt(replies.item(0).innerText);
        }
        var link = row.getElementsByClassName("ic-item-row__content-link");
        if (link.length > 0) {
            var parts = link.item(0).href.split("/");
            announcement.id = parts[parts.length - 1];
        }
        announcements.push(announcement);
    }
    fetch("scraped://announcements/", {
        "method": "POST",
        "body": JSON.stringify(announcements)
    });
})();
