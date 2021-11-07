function getPort() {
    return "/panel/"
}

function getQueryString(name) {
    const reg = new RegExp("(^|&)" + name + "=([^&]*)(&|$)");
    const r = window.location.search.substr(1).match(reg);
    if (r != null) return unescape(r[2]);
    else {
        const urls = window.location.href.split('/');
        let i = 0,
            iLoop = urls.length;
        for (; i < iLoop; i++) {
            if (urls[i] === name) {
                return urls[i + 1].split('.')[0];
            }
        }
    }
    return null;
}

function getNextDate(date, day) {
    var dd = new Date(date);
    dd.setDate(dd.getDate() + day);
    var y = dd.getFullYear();
    var m = dd.getMonth() + 1 < 10 ? "0" + (dd.getMonth() + 1) : dd.getMonth() + 1;
    var d = dd.getDate() < 10 ? "0" + dd.getDate() : dd.getDate();
    return y + "-" + m + "-" + d;
};