$(function() {
    jQuery.support.cors = true;

    $('#submit').click(function() {
        var uid = $('input[name="uid"]').val();
        var pwd = $('input[name="password"]').val();
        console.log(uid, pwd)

        $.post("http://163.221.29.46:13131/login", {
            crossDomain: true,
            xhrFields: {
               withCredentials: true
            },
            data: { "Username": uid, "Password": pwd }
        })
        .done(function (data, textStatus) {
            console.log(data, textStatus, "hoge")
            window.location.href = "http://163.221.29.46:13131/buyItemHTML"
        })
        .fail(function (data, textStatus) {
            console.log(data, textStatus)
        })
    })

});
