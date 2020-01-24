$(function() {

    $.getJSON('http://163.221.29.46:13131/getItems', {})
    .done(function(data) {
        if(data) {
            console.log(data)
            for(let l in data) {
                var id = data[l]['ID']
                var name = data[l]['Name']
                var price = data[l]['Price']
                var cost = data[l]['Cost']
                var amount = data[l]['Amount']
                $('#list_body').append("\<tr\>\<td\>" +
                id +  "\</td\>\<td\>" +
                name + "\</td\>\<td\>" +
                price + "円" + "\</td\>\<td\>" +
                cost + "円" + "\</td\>\<td\>" +
                amount + "個" + "\</td\>\</tr\>")
            }
        } else {
            console.log("error");
        }   
    });

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
        var name = $('input[name="iname"]').val();
        var price = $('input[name="iprice"]').val();
        var cost = $('input[name="icost"]').val();
        var amount = $('input[name="iamount"]').val();
        console.log(name, price, cost, amount);
        
        $('#result').replaceWith('<div id="result"></div>')
        $('#result').append("\<table\>\<thead\>\<tr\>\<th\>name\</th\>\<th\>price\</th\>\<th\>cost\</th\>\<th\>amont\</th\>\</tr\>\</thead\>\<tbody\>\<tr\>\<td\>" + 
        name + "\</td\>\<td\>" +
        price + "円" + "\</td\>\<td\>" +
        cost + "円" + "\</td\>\<td\>" +
        amount + "個" + "\</td\>\</tr\>\</tbody\>\</table\>")
    });
});