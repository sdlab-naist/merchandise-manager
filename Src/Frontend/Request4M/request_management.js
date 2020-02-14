$(function() {
    var requestlist
    $.getJSON('http://163.221.29.46:13131/api/getRequests', {})
    .done(function(data) {
        if(data) {
            console.log(data)
            requestlist = data
            for(let l in data) {
                var id = data[l]['ID']
                var uid = data[l]['Username']
                var iname = data[l]['Itemname']
                var amount = data[l]['Amount']
                var status = data[l]['Status']
                $('#list_body').append("\<tr>\<td\>" +
                id +  "\</td\>\<td\>" +
                uid + "\</td\>\<td\>" +
                iname + "円" + "\</td\>\<td\>" +
                amount + "個" + "\</td\>\<td\>" +
                status + "\</td\>" + 
                "\<td>\<input id=" + id + " type=\"button\" class=\"btn btn-default select\" value=\"Select\"\>\</td>\</tr\>")
            }
        } else {
            console.log("error");
        }   
    });

    var over_flg

    $(function () {
        $('#navbar-button').click(function() {
            console.log(over_flg)
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

    var id
    $(document).on('click','.select', function() {
        console.log($(this).attr('id'))
        for (l in requestlist) {
            if (requestlist[l]['ID'] == $(this).attr('id')){
                id = requestlist[l]['ID']
                var name = requestlist[l]['Name']
                var amount = requestlist[l]['Amount']
                console.log(id, name)
                $('#id').val(id)
                $('#item_name').val(name)
                $('#item_amount').val(amount)
                break
            }
        }
    });

    $('#submit').click(function() {
        var id = $('input[name="id"]').val();
        var name = $('input[name="iname"]').val();
        console.log(name)
        var amount = parseInt($('input[name="iamount"]').val());
        console.log(typeof(amount))
        
        $.post('http://163.221.29.46:13131/api/deleteRequest', {
            ID: id
        })
        .done(function(data) {
            console.log(data)
            $('#result').replaceWith('<div id="result"></div>')
            $('#result').append("\<div\>Deleted Request\</div\>\<table\>\<thead\>\<tr\>\<th\>id\</th\>\<th\>item name\</th\>\<th\>amount\</th\>\</tr\>\</thead\>\<tbody\>\<tr\>\<td\>" + 
            id + "\</td\>\<td\>" +
            name + "\</td\>\<td\>" +
            amount + "個" + "\</td\>\</tr\>\</tbody\>\</table\>")
            })
    });
});