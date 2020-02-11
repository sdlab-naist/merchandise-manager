$(function() {
    $('#submit').click(function() {
        var uid = $('input[name="uid"]').val();
        var pwd = $('input[name="password"]').val();
        var fname = $('input[name="fname"]').val();
        var lname = $('input[name="lname"]').val();
        var email = $('input[name="email"]').val();
        console.log(uid, pwd, fname, lname, email)

        $.post('http://163.221.29.46:13131/registerUser' ,{
            Username: uid,
            Password: pwd,
            Firstname: fname,
            Lastname: lname,
            Email: email
        })
        .done(function(data) {
            console.log(data)
            if(data['success']){
                $('#result').replaceWith('<div id="result"></div>')
                $('#result').append("\<a\>Check your E-mail\</a\>")
            } else {
                $('#result').replaceWith('<div id="result"></div>')
                $('#result').append("\<a\>error\</a\>")
            }
        })
    })

});