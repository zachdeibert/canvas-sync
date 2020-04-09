(function() {
    var announcement = {
        "replies": []
    };
    var message = document.querySelectorAll(".announcement .message_wrapper");
    if (message.length === 1) {
        announcement.html = message.item(0).innerHTML;
    }
    var author = document.querySelectorAll(".announcement .author");
    if (author.length === 1) {
        announcement.author = author.item(0).innerText;
    }
    var threads = document.querySelectorAll(".discussion_subentries:not(.hidden) > ul > li.entry");
    for (var i = 0; i < threads.length; ++i) {
        var thread = threads.item(i);
        var messages = thread.getElementsByTagName("article");
        if (messages.length > 0) {
            var messageGroup = [];
            for (var j = 0; j < messages.length; ++j) {
                var article = messages.item(j);
                var message = {};
                author = article.getElementsByClassName("author");
                if (author.length === 1) {
                    message.author = author.item(0).innerText;
                }
                var time = article.getElementsByTagName("time");
                if (time.length === 1) {
                    message.time = new Date(time.item(0).getAttribute("datetime")).getTime();
                }
                var html = article.getElementsByClassName("message");
                if (html.length === 1) {
                    message.html = html.item(0).innerHTML;
                    message.text = html.item(0).innerText;
                }
                messageGroup.push(message);
            }
            announcement.replies.push(messageGroup);
        }
    }
    fetch("scraped://announcement/", {
        "method": "POST",
        "body": JSON.stringify(announcement)
    });
})();
