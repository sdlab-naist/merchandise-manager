$(function() {
    jQuery.support.cors = true;
    $('#submit').click(function() {
        var uid = $('input[name="uid"]').val();
        var pwd = $('input[name="password"]').val();
        console.log(uid, pwd)
        jQuery.support.cors = true;
        $.ajax({
            type: "POST",
            crossDomain: true,
            xhfFields: {
                withCredentials: true
            },
            url: "htt://163.221.29.46:13131/login",
            data: { "Username": uid, "password": pwd},
            success: function (jsondata) {
                console.log(jsondata)
                window.location.href = "http://163.221.29.46:13131/addItemHTML"
            }
        })
    })
});