$(function() {
    $('#submit').click(function() {
        var uid = $('input[name="uid"]').val();
        var pwd = $('input[name="password"]').val();
        console.log(uid, pwd)

        $.post('http://163.221.29.46:13131/login' ,{
            Username: uid,
            Password: pwd
        })
        .done(function(data) {
            console.log(data)
        })
    })

});