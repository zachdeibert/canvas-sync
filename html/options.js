var results = [];

function init(courses, scrapers) {
    if (!sessionStorage.isDev) {
        sessionStorage.courses = JSON.stringify(courses);
        sessionStorage.scrapers = JSON.stringify(scrapers);
        sessionStorage.isDev = true;
    }
    var container = document.getElementById("container");
    courses.forEach(function(course) {
        var col = document.createElement("div");
        col.classList.add("col");
        col.classList.add("s12");
        col.classList.add("m6");
        col.classList.add("l4");
        var card = document.createElement("div");
        card.classList.add("card");
        card.classList.add("darken-2");
        card.style.backgroundColor = course.color;
        var content = document.createElement("div");
        content.classList.add("card-content");
        content.classList.add("white-text");
        var title = document.createElement("span");
        title.classList.add("card-title");
        title.innerText = course.name;
        content.appendChild(title);
        var code = document.createElement("h6");
        code.innerText = course.code;
        content.appendChild(code);
        var term = document.createElement("p");
        term.innerText = course.term;
        content.appendChild(term);
        var table = document.createElement("table");
        var thead = document.createElement("thead");
        var tr = document.createElement("tr");
        var th = document.createElement("th");
        th.innerText = "Content Type";
        tr.appendChild(th);
        th = document.createElement("th");
        th.innerText = "Update";
        tr.appendChild(th);
        th = document.createElement("th");
        th.innerText = "Delete";
        tr.appendChild(th);
        th = document.createElement("th");
        th.innerText = "Sync";
        tr.appendChild(th);
        thead.appendChild(tr);
        table.appendChild(thead);
        var tbody = document.createElement("tbody");
        var updateBoxes = [];
        var deleteBoxes = [];
        var syncBoxes = [];
        var result = JSON.parse(JSON.stringify(course));
        results.push(result);
        result.scrapers = [];
        scrapers.forEach(function(scraper) {
            var tr = document.createElement("tr");
            var td = document.createElement("td");
            td.innerText = scraper;
            tr.appendChild(td);
            td = document.createElement("td");
            var label = document.createElement("label");
            label.classList.add("checkbox");
            var updateBox = document.createElement("input");
            var deleteBox = document.createElement("input");
            var syncBox = document.createElement("input");
            var scraperConfig = {
                "update": false,
                "delete": false
            };
            result.scrapers.push(scraperConfig);
            updateBoxes.push(updateBox);
            deleteBoxes.push(deleteBox);
            syncBoxes.push(syncBox);
            function updateSyncBox() {
                if (updateBox.checked) {
                    if (deleteBox.checked) {
                        syncBox.indeterminate = false;
                        syncBox.checked = true;
                    } else {
                        syncBox.indeterminate = true;
                    }
                } else if (deleteBox.checked) {
                    syncBox.indeterminate = true;
                } else {
                    syncBox.indeterminate = false;
                    syncBox.checked = false;
                }
                syncBox.dispatchEvent(new Event("change"));
            }
            updateBox.type = "checkbox";
            updateBox.addEventListener("change", function() {
                scraperConfig.update = updateBox.checked;
                updateSyncBox();
            });
            label.appendChild(updateBox);
            var span = document.createElement("span");
            label.appendChild(span);
            td.appendChild(label);
            tr.appendChild(td);
            td = document.createElement("td");
            label = document.createElement("label");
            label.classList.add("checkbox");
            deleteBox.type = "checkbox";
            deleteBox.addEventListener("change", function() {
                scraperConfig.delete = deleteBox.checked;
                updateSyncBox();
            });
            label.appendChild(deleteBox);
            span = document.createElement("span");
            label.appendChild(span);
            td.appendChild(label);
            tr.appendChild(td);
            td = document.createElement("td");
            label = document.createElement("label");
            label.classList.add("checkbox");
            syncBox.type = "checkbox";
            syncBox.addEventListener("change", function() {
                if (!syncBox.indeterminate) {
                    scraperConfig.update = scraperConfig.delete = updateBox.checked = deleteBox.checked = syncBox.checked;
                }
            });
            label.appendChild(syncBox);
            span = document.createElement("span");
            label.appendChild(span);
            td.appendChild(label);
            tr.appendChild(td);
            tbody.appendChild(tr);
        });
        tr = document.createElement("tr");
        var td = document.createElement("td");
        td.innerText = "All";
        tr.appendChild(td);
        td = document.createElement("td");
        var label = document.createElement("label");
        label.classList.add("checkbox");
        var updateBox = document.createElement("input");
        var deleteBox = document.createElement("input");
        var syncBox = document.createElement("input");
        function updateSyncBox() {
            if (updateBox.checked) {
                if (deleteBox.checked) {
                    syncBox.indeterminate = false;
                    syncBox.checked = true;
                } else {
                    syncBox.indeterminate = true;
                }
            } else if (deleteBox.checked) {
                syncBox.indeterminate = true;
            } else {
                syncBox.indeterminate = false;
                syncBox.checked = false;
            }
        }
        function checkAll(list, checked) {
            list.forEach(function(el) {
                el.checked = checked;
                el.indeterminate = false;
                var ev = new Event("change");
                ev.synthetic = true;
                el.dispatchEvent(ev);
            });
        }
        function updateCol(list, checkbox) {
            var anyChecked = false;
            var anyIndeterminate = false;
            var allChecked = true;
            list.forEach(function(el) {
                if (el.indeterminate) {
                    anyIndeterminate = true;
                } else if (el.checked) {
                    anyChecked = true;
                } else {
                    allChecked = false;
                }
            });
            if (allChecked && !anyIndeterminate) {
                checkbox.checked = true;
            } else if (!anyChecked && !anyIndeterminate) {
                checkbox.checked = false;
            } else {
                checkbox.indeterminate = true;
                return;
            }
            checkbox.indeterminate = false;
        }
        function updateAll(ev) {
            if (!ev.synthetic) {
                updateCol(updateBoxes, updateBox);
                updateCol(deleteBoxes, deleteBox);
                updateCol(syncBoxes, syncBox);
            }
        }
        updateBoxes.forEach(function(el) {
            el.addEventListener("change", updateAll);
        });
        deleteBoxes.forEach(function(el) {
            el.addEventListener("change", updateAll);
        });
        syncBoxes.forEach(function(el) {
            el.addEventListener("change", updateAll);
        });
        updateBox.type = "checkbox";
        updateBox.addEventListener("change", function() {
            updateSyncBox();
            checkAll(updateBoxes, updateBox.checked);
        });
        label.appendChild(updateBox);
        var span = document.createElement("span");
        label.appendChild(span);
        td.appendChild(label);
        tr.appendChild(td);
        td = document.createElement("td");
        label = document.createElement("label");
        label.classList.add("checkbox");
        deleteBox.type = "checkbox";
        deleteBox.addEventListener("change", function() {
            updateSyncBox();
            checkAll(deleteBoxes, deleteBox.checked);
        });
        label.appendChild(deleteBox);
        span = document.createElement("span");
        label.appendChild(span);
        td.appendChild(label);
        tr.appendChild(td);
        td = document.createElement("td");
        label = document.createElement("label");
        label.classList.add("checkbox");
        syncBox.type = "checkbox";
        syncBox.addEventListener("change", function() {
            updateBox.checked = deleteBox.checked = syncBox.checked;
            updateBox.indeterminate = deleteBox.indeterminate = false;
            checkAll(syncBoxes, syncBox.checked);
        });
        label.appendChild(syncBox);
        span = document.createElement("span");
        label.appendChild(span);
        td.appendChild(label);
        tr.appendChild(td);
        tbody.appendChild(tr);
        table.appendChild(tbody);
        content.appendChild(table);
        card.appendChild(content);
        col.appendChild(card);
        container.appendChild(col);
    });
}

function submit() {
    fetch("scraped://options/", {
        "method": "POST",
        "body": JSON.stringify(results)
    });
}

addEventListener("load", function() {
    if (sessionStorage.isDev) {
        init(JSON.parse(sessionStorage.courses), JSON.parse(sessionStorage.scrapers));
    }
});
