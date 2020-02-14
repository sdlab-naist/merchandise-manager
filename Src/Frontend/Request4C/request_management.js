$(function() {
    var over_flg

    $(function () {
        $('#navbar-button').click(function() {
            console.log(over_flg)
           //$('#navcontents').slideToggle('fast');
            $('#navbarSupportedContent').collapse('toggle');
        });

        $('header, ul').hover(function() {
            over_flg = true;
        }, function() {
            over_flg = false;
        });

        $('body').click(function() {
            if(over_flg == false) {
                $('#navbarSupportedContent').collapse('hide')
            }
        });
    });

    $('#submit').click(function() {
        var uid = $('input[name="uid"]').val();
        var iname = $('input[name="iname"]').val();
        var amount = $('input[name="iamount"]').val();
        
        $.post('http://163.221.29.46:13131/api/requestItem', {
            Username: uid,
            Itemname: iname,
            Amount: amount
        })
        .done(function(data) {
            console.log(data)
            $('#result').replaceWith('<div id="result"></div>')
            $('#result').append("\<p class=title\>Requested Item\</p\>\<table\>\<thead\>\<tr\>\<th\>User Name\</th\>\<th\>Item Name\</th\>\<th\>amount\</th\>\</tr\>\</thead\>\<tbody\>\<tr\>\<td\>" + 
            uid + "\</td\>\<td\>" +
            iname + "\</td\>\<td\>" +
            amount + "個" + "\</td\>\</tr\>\</tbody\>\</table\>")
            })
    });
});