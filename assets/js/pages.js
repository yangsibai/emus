$(function () {
    var $iframe = $('#iframe');
    $('li').click(function (){
        $(this).addClass("current").siblings().removeClass("current");
        var pageID = $(this).data('page-id');
        $iframe.attr("src", "/page/" + pageID);
    });
});
