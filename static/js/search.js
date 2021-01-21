const search = (key, element, itemSelector) => {
    const val = $(element).children(itemSelector).text().toUpperCase();

    if (val.indexOf(key) > -1) {
        $(element).css("display", "");
    } else {
        $(element).css("display", "none");
    }
}