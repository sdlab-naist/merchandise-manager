$(function() {
    jQuery.support.cors = true;

    $('#submit').click(function() {
        var uid = $('input[name="uid"]').val();
        var pwd = $('input[name="password"]').val();
        console.log(uid, pwd)

        $.ajax({
            type: "POST",
            crossDomain: true,
            xhrFields: {
               withCredentials: true
            },
            url: "http://163.221.29.46:13131/login",
            data: { "Username": uid, "Password": pwd },
            success: function (jsondata) {
                console.log(jsondata)
                window.location.href = "../AddItem/add_item_view.html"
            },
            error: function (jsondata) {
                console.log(jsondata)
            }

         }) 
        })
    });
