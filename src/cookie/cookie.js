function setCookie(token, days) {
    // 设置过期时间
    let data = new Date(
        new Date().getTime() + days * 24 * 60 * 60 * 1000
    ).toUTCString();
    document.cookie = "token" + "=" + token + "; expires=" + data;
}

function getCookie(name) {
    var arr = document.cookie.match(new RegExp("(^| )" + name + "=([^;]*)(;|$)"));
    if (arr != null) {
        return unescape(arr[2])
    } else {
        return null
    }
}

function clearCookie(name) {
    let json = {};
    json[name] = '';
    setCookie(json, -1)
}

export default {
    setCookie,
    getCookie,
    clearCookie
}
