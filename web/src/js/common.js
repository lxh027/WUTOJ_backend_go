function getPort() {
    return "http://localhost:5000/panel/"
}

function getQueryString(name)
{
    const reg = new RegExp("(^|&)" + name + "=([^&]*)(&|$)");
    const r = window.location.search.substr(1).match(reg);
    if(r!=null)return unescape(r[2]);
    else {
        const urls = window.location.href.split('/');
        let i = 0, iLoop = urls.length;
        for (; i<iLoop; i++) {
            if (urls[i] === name) {
                return urls[i+1].split('.')[0];
            }
        }
    }
    return null;
}