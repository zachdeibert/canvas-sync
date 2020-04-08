var fields = [];

function submitting() {
    var payload = {};
    fields.forEach(function(el) {
        localStorage[el.name] = el.value;
        payload[el.name] = el.value;
    });
    fetch("scraped://startup/", {
        "method": "POST",
        "body": JSON.stringify(payload)
    });
}

addEventListener("load", function() {
    if (!localStorage.out) {
        localStorage.out = "./db/";
    }
    var inputs = document.getElementsByTagName("input");
    for (var i = 0; i < inputs.length; ++i) {
        var input = inputs.item(i);
        if (input.type !== "submit") {
            fields.push(input);
            var v = localStorage[input.name];
            if (v) {
                input.value = v;
            }
        }
    }
    M.updateTextFields();
});
