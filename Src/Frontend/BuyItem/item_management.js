$(function() {
    var itemlist
    $.getJSON('http://163.221.29.46:13131/getItems', {})
    .done(function(data) {
        if(data) {
            itemlist = data
            for(let l in data) {
                var id = data[l]['ID']
                var name = data[l]['Name']
                var price = data[l]['Price']
                //var cost = data[l]['Cost']
                var amount = data[l]['Amount']
                $('#list_body').append("\<tr\>\<td\>" +
                id +  "\</td\>\<td\>" +
                name + "\</td\>\<td\>" +
                price + "円" + "\</td\>\<td\>" +
                /*cost + "円" + "\</td\>\<td\>" +*/
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

    var itemid
    $(document).on('click','.select', function() {
        for (l in itemlist) {
            if (itemlist[l]['ID'] == $(this).attr('id')){
                itemid = itemlist[l]['ID']
                var name = itemlist[l]['Name']
                var price = parseInt(itemlist[l]['Price'])
                //var amount = parseInt(itemlist[l]['Amount'])
                $('#item_name').val(name)
                $('#item_price').val(price)
                break
            }
        }
    });

    var orderId = ""
    $('#submit').click(function() {
        var name = $('input[name="iname"]').val();
        var price = parseInt($('input[name="iprice"]').val());
        //var cost = $('input[name="icost"]').val();
        var amount = parseInt($('input[name="iamount"]').val());
        
        if(orderId == "") {
            $.post('http://163.221.29.46:13131/registerOrder', {
                ItemID: itemid,
                Amount: amount
            }
        )
        .done(function(data) { 
            console.log(data)
            orderId = data;
            var total = price * amount
            $('#result').replaceWith('<div id="result"></div>')
            $('#result').append("\<div\>Added Item\</div\>\<table\>\<thead\>\<tr\>\<th\>name\</th\>\<th\>price\</th\>\<th\>amount\</th\>\<th\>total\</th\>\</tr\>\</thead\>\<tbody\>\<tr\>\<td\>" + 
            name + "\</td\>\<td\>" +
            price + "円" + "\</td\>\<td\>" +
            amount + "個" + "\</td\>\<td\>" +
            total + "円" + "\</td\>\</tr\>\</tbody\>\</table\>")
            })
        } else {
            $.post('http://163.221.29.46:13131/registerOrder', {
                OrderID: orderId,
                ItemID: itemid,
                Amount: amount
            }
        )
        .done(function(data) {
            orderId = data;
            var total = price * amount
            $('#result').replaceWith('<div id="result"></div>')
            $('#result').append("\<div\>Added Item\</div\>\<table\>\<thead\>\<tr\>\<th\>name\</th\>\<th\>price\</th\>\<th\>amount\</th\>\<th\>total\</th\>\</tr\>\</thead\>\<tbody\>\<tr\>\<td\>" + 
            name + "\</td\>\<td\>" +
            price + "円" + "\</td\>\<td\>" +
            amount + "個" + "\</td\>\<td\>" +
            total + "円" + "\</td\>\</tr\>\</tbody\>\</table\>")
            })
        }
    });
});