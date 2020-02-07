$(function() {
    var orderlist
    var itemlist

    $.getJSON('http://163.221.29.46:13131/getItems', {})
    .done(function(data) {
        if(data) {
            itemlist = data
            console.log(itemlist)
        } else {
            console.log("error");
        } 
    });

    $.getJSON('http://163.221.29.46:13131/getOrders', {})
    .done(function(data) {
        if(data) {
            orderlist = data
            console.log(data)
            for(let l in data) {
                var id = data[l]['ID']
                var oid = data[l]['OrderID']
                var iid = data[l]['ItemID']
                var amount = parseInt(data[l]['Amount'])
                var iname
                var price 
                for(i in itemlist) {
                    if(itemlist[i]['ID'] == iid){
                        console.log(itemlist[i], iid)
                        iname = itemlist[i]['Name']
                        price = parseInt(itemlist[i]['Price'])
                        break;
                    }
                }
                if(typeof iname !== 'undefined') {
                    var total = price * amount
                    $('#list_body').append("\<tr\>\<td\>" +
                    id + "\</td\>\<td\>" +
                    iname +  "\</td\>\<td\>" +
                    price + "円" + "\</td\>\<td\>" +
                    amount + "個" + "\</td\>\<td\>" +
                    total + "円" + "\</td\>" +
                    "\<td>\<input id=" + id + " type=\"button\" class=\"btn btn-default select\" value=\"Select\"\>\</td>\</tr\>")
                }
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

    var orderid
    $(document).on('click','.select', function() {
        for (l in orderlist) {
            if (orderlist[l]['ID'] == $(this).attr('id')){
                orderid = orderlist[l]['ID']
                var iname
                var price
                for(i in itemlist) {
                    if(itemlist[i]['ID'] == orderlist[l]['ItemID']){
                        iname = itemlist[i]['Name']
                        price = parseInt(itemlist[i]['Price'])
                        break;
                    }
                }
                var amount = parseInt(orderlist[l]['Amount'])
                var total = price * amount
                $('#order_id').val(orderid)
                $('#order_name').val(iname)
                $('#order_price').val(price)
                $('#order_amount').val(amount)
                $('#order_total').val(total)
                break
            }
        }
    });

    var orderId = ""
    $('#submit').click(function() {
        
        if(orderId == "") {
            console.log(orderId)
            $.post('http://163.221.29.46:13131/makeOrder', {
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
            console.log(orderId, itemid)
            $.post('http://163.221.29.46:13131/registerOrder', {
                OrderID: orderId,
                ItemID: itemid,
                Amount: amount
            }
        )
        .done(function(data) { 
            console.log(data, itemid)
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