$(function() {
    var itemlist
    $.getJSON('http://163.221.29.46:13131/getItems', {})
    .done(function(data) {
        if(data) {
            console.log(data)
            itemlist = data
            for(let l in data) {
                var id = data[l]['ID']
                var name = data[l]['Name']
                var price = data[l]['Price']
                var cost = data[l]['Cost']
                var amount = data[l]['Amount']
                $('#list_body').append("\<tr>\<td\>" +
                id +  "\</td\>\<td\>" +
                name + "\</td\>\<td\>" +
                price + "円" + "\</td\>\<td\>" +
                cost + "円" + "\</td\>\<td\>" +
                amount + "個" + "\</td\>" + 
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
        for (l in itemlist) {
            if (itemlist[l]['ID'] == $(this).attr('id')){
                id = itemlist[l]['ID']
                var name = itemlist[l]['Name']
                $('#item_name').val(name)
                break
            }
        }
    });

    $('#submit').click(function() {
        var name = $('input[name="iname"]').val();
        console.log(name)
        var amount = parseInt($('input[name="iamount"]').val());
        console.log(typeof(amount))
        
        $.post('http://163.221.29.46:13131/deleteItem', {
            ID: id,
            Amount: amount
        })
        .done(function(data) {
            console.log(data)
            $('#result').replaceWith('<div id="result"></div>')
            $('#result').append("\<div\>Deleted Item\</div\>\<table\>\<thead\>\<tr\>\<th\>name\</th\>\<th\>amount\</th\>\</tr\>\</thead\>\<tbody\>\<tr\>\<td\>" + 
            name + "\</td\>\<td\>" +
            amount + "個" + "\</td\>\</tr\>\</tbody\>\</table\>")
            })
    });
});