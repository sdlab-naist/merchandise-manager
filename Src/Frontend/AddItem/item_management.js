$(function() {
    $.getJSON('http://163.221.29.46:13131/getItems', {})
    .done(function(data) {
        if(data) {
            console.log(result);
        } else {
            console.log("error");
        }   
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