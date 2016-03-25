$(function() {
    var $iframe = $('#iframe');
    var currentPageID = null;

    function updateIframeSource(src) {
        $iframe.attr("src", src);
    }

    $('li').click(function() {
        $(this).addClass("current").siblings().removeClass("current");
        var pageID = $(this).data('page-id');
        currentPageID = pageID;
        updateIframeSource("/page/" + pageID);
    });

    $('li .delete').click(function() {
        var pageID = $(this).data('page-id');
        $.post("/page/delete/" + pageID, function(res) {
            if (res.code === 0) {
                $('#' + pageID).slideUp(function () {
                    $(this).remove();
                    if (currentPageID === pageID) {
                        updateIframeSource("");
                    }
                });
            } else {
                alert(res.error);
            }
        });
    });
});
