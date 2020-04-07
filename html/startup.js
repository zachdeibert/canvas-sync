var fields = [];

function submitting() {
    fields.forEach(function(el) {
        localStorage[el.name] = el.value;
    });
}

addEventListener("load", function() {
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
