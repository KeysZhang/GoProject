$(document).ready(function() {
    $.ajax({
        url: "/static/error"
    }).then(function(data) {
        if(data.iserror){
            for (var i=0;i<data.errorslice.length;i++){
                $('#'+ data.errorslice[i].Id).text(data.errorslice[i].Message);
            }
        }
    });
});