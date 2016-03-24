$(function () {
    var $iframe = $('#iframe');
    $('li').click(function (){
        var pageID = $(this).data('page-id');
        $iframe.attr("src", "/page/" + pageID);
    });
});
